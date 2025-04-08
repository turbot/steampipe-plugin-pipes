package pipes

import (
	"context"
	"fmt"

	openapi "github.com/turbot/pipes-sdk-go"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablePipesTenantIntegration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pipes_tenant_integration",
		Description: "Tenant integrations enable connections to services like AWS, GCP, Azure, etc.",
		List: &plugin.ListConfig{
			Hydrate: listTenantIntegrations,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getTenantIntegration,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the integration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "handle",
				Description: "The handle for the integration.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "tenant_id",
				Description: "The unique identifier for the tenant where this integration is created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "identity_id",
				Description: "The unique identifier for an identity where this integration is created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "type",
				Description: "The type of the integration.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "state",
				Description: "The state of the integration.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "state_reason",
				Description: "The reason for the state of the integration.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "config",
				Description: "The configuration for the integration.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "github_installation_id",
				Description: "The GitHub installation ID for this integration, only applicable when the integration type is github.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "pipeline_id",
				Description: "The pipeline ID for this integration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "pipeline",
				Description: "Pipeline information for this integration.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "created_at",
				Description: "The time when the integration was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "created_by_id",
				Description: "The unique identifier of the user who created the integration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "created_by",
				Description: "Information about the user who created the integration.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "updated_at",
				Description: "The time when the integration was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "updated_by_id",
				Description: "The unique identifier of the user who last updated the integration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "updated_by",
				Description: "Information about the user who last updated the integration.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "deleted_at",
				Description: "The time when the integration was deleted.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "deleted_by_id",
				Description: "The unique identifier of the user who deleted the integration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "deleted_by",
				Description: "Information about the user who deleted the integration.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "version_id",
				Description: "The version ID of the integration.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromCamel(),
			},
		}),
	}
}

//// LIST FUNCTION

func listTenantIntegrations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create Session
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listTenantIntegrations", "connection_error", err)
		return nil, err
	}

	// If the requested number of items is less than the paging max limit
	// set the limit to that instead
	maxResults := int32(100)
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < int64(maxResults) {
			if *limit < 1 {
				maxResults = int32(1)
			} else {
				maxResults = int32(*limit)
			}
		}
	}

	// execute list call
	pagesLeft := true
	var nextToken *string

	for pagesLeft {
		var listDetails func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error)

		// Create the request with or without nextToken
		if nextToken != nil {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err := svc.TenantIntegrations.List(ctx).NextToken(*nextToken).Limit(maxResults).Execute()
				if err != nil {
					plugin.Logger(ctx).Error("listTenantIntegrations", "api_error", err)
					return nil, err
				}
				plugin.Logger(ctx).Debug("listTenantIntegrations", "response_type", fmt.Sprintf("%T", resp))
				return resp, nil
			}
		} else {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err := svc.TenantIntegrations.List(ctx).Limit(maxResults).Execute()
				if err != nil {
					plugin.Logger(ctx).Error("listTenantIntegrations", "api_error", err)
					return nil, err
				}
				plugin.Logger(ctx).Debug("listTenantIntegrations", "response_type", fmt.Sprintf("%T", resp))
				return resp, nil
			}
		}

		// Execute the request with retry
		response, err := plugin.RetryHydrate(ctx, d, h, listDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
		if err != nil {
			plugin.Logger(ctx).Error("listTenantIntegrations", "list_error", err)
			return nil, err
		}

		// Handle the response as ListIntegrationsResponse
		resp, ok := response.(openapi.ListIntegrationsResponse)
		if !ok {
			plugin.Logger(ctx).Error("listTenantIntegrations", "type_error", fmt.Sprintf("Expected ListIntegrationsResponse, got %T", response))
			return nil, fmt.Errorf("expected ListIntegrationsResponse, got %T", response)
		}

		// Stream the results
		if resp.HasItems() {
			for _, integration := range *resp.Items {
				d.StreamListItem(ctx, integration)

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
		}

		// Check if there are more pages
		if resp.NextToken == nil {
			pagesLeft = false
		} else {
			nextToken = resp.NextToken
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getTenantIntegration(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	integrationId := d.EqualsQuals["id"].GetStringValue()

	// check if id is empty
	if integrationId == "" {
		return nil, nil
	}

	// Create Session
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getTenantIntegration", "connection_error", err)
		return nil, err
	}

	// execute get call
	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		resp, _, err := svc.TenantIntegrations.Get(ctx, integrationId).Execute()
		if err != nil {
			plugin.Logger(ctx).Error("getTenantIntegration", "api_error", err)
			return nil, err
		}
		plugin.Logger(ctx).Debug("getTenantIntegration", "response_type", fmt.Sprintf("%T", resp))
		return resp, nil
	}

	response, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
	if err != nil {
		plugin.Logger(ctx).Error("getTenantIntegration", "get_error", err)
		return nil, err
	}

	return response, nil
}
