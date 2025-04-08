package pipes

import (
	"context"

	openapi "github.com/turbot/pipes-sdk-go"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablePipesUserIntegration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pipes_user_integration",
		Description: "User integrations enable connections to services like GitHub, Slack, etc.",
		List: &plugin.ListConfig{
			ParentHydrate: getUser,
			Hydrate:       listUserIntegrations,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "user_handle"}),
			Hydrate:    getUserIntegration,
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
				Name:        "user_handle",
				Description: "The handle name for the user where this integration is created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Identity.Handle").NullIfZero(),
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

func listUserIntegrations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	user := h.Item.(openapi.User)

	// Create Session
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listUserIntegrations", "connection_error", err)
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

	// Execute list call
	pagesLeft := true
	var nextToken *string

	for pagesLeft {
		var listDetails func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error)

		// Create the request with or without nextToken
		if nextToken != nil {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err := svc.UserIntegrations.List(ctx, user.Handle).NextToken(*nextToken).Limit(maxResults).Execute()
				if err != nil {
					plugin.Logger(ctx).Error("listUserIntegrations", "api_error", err)
					return nil, err
				}
				return resp, nil
			}
		} else {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err := svc.UserIntegrations.List(ctx, user.Handle).Limit(maxResults).Execute()
				if err != nil {
					plugin.Logger(ctx).Error("listUserIntegrations", "api_error", err)
					return nil, err
				}
				return resp, nil
			}
		}

		// Execute the request with retry
		response, err := plugin.RetryHydrate(ctx, d, h, listDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
		if err != nil {
			plugin.Logger(ctx).Error("listUserIntegrations", "list_error", err)
			return nil, err
		}

		// Handle the response
		result, ok := response.(openapi.ListIntegrationsResponse)
		if !ok {
			plugin.Logger(ctx).Error("listUserIntegrations", "type_error", "Expected ListIntegrationsResponse")
			return nil, err
		}

		// Stream the results
		if result.HasItems() {
			for _, integration := range *result.Items {
				d.StreamListItem(ctx, integration)

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
		}

		// Check for pagination
		if result.NextToken == nil {
			pagesLeft = false
		} else {
			nextToken = result.NextToken
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getUserIntegration(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	integrationId := d.EqualsQuals["id"].GetStringValue()
	userHandle := d.EqualsQuals["user_handle"].GetStringValue()

	// Check if required parameters are provided
	if integrationId == "" || userHandle == "" {
		return nil, nil
	}

	// Create Session
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getUserIntegration", "connection_error", err)
		return nil, err
	}

	// Execute get call
	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		resp, _, err := svc.UserIntegrations.Get(ctx, userHandle, integrationId).Execute()
		if err != nil {
			plugin.Logger(ctx).Error("getUserIntegration", "api_error", err)
			return nil, err
		}
		return resp, nil
	}

	response, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
	if err != nil {
		plugin.Logger(ctx).Error("getUserIntegration", "get_error", err)
		return nil, err
	}

	return response, nil
}
