package pipes

import (
	"context"

	openapi "github.com/turbot/pipes-sdk-go"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablePipesTenant(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pipes_tenant",
		Description: "Tenants are logical groups of resources within Turbot Pipes.",
		List: &plugin.ListConfig{
			Hydrate: listTenants,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("handle"),
			Hydrate:    getTenant,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier of the tenant.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "handle",
				Description: "The handle name of the tenant.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "display_name",
				Description: "The display name of the tenant.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "avatar_url",
				Description: "The avatar URL of the tenant.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "state",
				Description: "The state of the tenant.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The time of creation in ISO 8601 UTC.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "created_by_id",
				Description: "The ID of the user that created this.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "updated_at",
				Description: "The time of the last update in ISO 8601 UTC.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "updated_by_id",
				Description: "The ID of the user that performed the last update.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "deleted_at",
				Description: "The time the item was deleted in ISO 8601 UTC.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "deleted_by_id",
				Description: "The ID of the user that performed the deletion.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "version_id",
				Description: "The version ID of this item.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "usage_compute_action",
				Description: "The action to take when compute usage exceeds the threshold.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "usage_compute_threshold",
				Description: "The threshold for compute usage.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "usage_storage_action",
				Description: "The action to take when storage usage exceeds the threshold.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "usage_storage_threshold",
				Description: "The threshold for storage usage.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "usage_user_action",
				Description: "The action to take when user usage exceeds the threshold.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "usage_user_threshold",
				Description: "The threshold for user usage.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromCamel(),
			},
		}),
	}
}

//// LIST FUNCTION

func listTenants(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create Session
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listTenants", "connection_error", err)
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
	var resp openapi.ListTenantsResponse
	var listDetails func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error)

	for pagesLeft {
		if resp.NextToken != nil {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.Tenants.List(ctx).NextToken(*resp.NextToken).Limit(maxResults).Execute()
				return resp, err
			}
		} else {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.Tenants.List(ctx).Limit(maxResults).Execute()
				return resp, err
			}
		}

		response, err := plugin.RetryHydrate(ctx, d, h, listDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

		if err != nil {
			plugin.Logger(ctx).Error("listTenants", "list", err)
			return nil, err
		}

		result := response.(openapi.ListTenantsResponse)

		if result.HasItems() {
			for _, tenant := range *result.Items {
				d.StreamListItem(ctx, tenant)

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
		}
		if result.NextToken == nil {
			pagesLeft = false
		} else {
			resp.NextToken = result.NextToken
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTION

func getTenant(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create Session
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getTenant", "connection_error", err)
		return nil, err
	}
	handle := d.EqualsQuals["handle"].GetStringValue()

	// check if id is empty
	if handle == "" {
		return nil, nil
	}

	var resp openapi.Tenant

	// execute get call
	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		resp, _, err = svc.Tenants.Get(ctx, handle).Execute()
		return resp, err
	}

	response, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

	if err != nil {
		plugin.Logger(ctx).Error("getTenant", "get", err)
		return nil, err
	}

	return response.(openapi.Tenant), nil
}
