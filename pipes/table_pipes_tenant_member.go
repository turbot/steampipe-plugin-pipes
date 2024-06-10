package pipes

import (
	"context"

	openapi "github.com/turbot/pipes-sdk-go"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablePipesTenantMember(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pipes_tenant_member",
		Description: "Members of a Turbot Pipes tenant.",
		List: &plugin.ListConfig{
			ParentHydrate: listTenants,
			Hydrate:       listTenantMembers,
			KeyColumns:    plugin.OptionalColumns([]string{"tenant_id", "tenant_handle"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"tenant_handle", "id"}),
			Hydrate:    getTenantMember,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier of the tenant member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "tenant_id",
				Description: "The identifier of the tenant.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "tenant_handle",
				Description: "The handle name of the tenant.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "user_id",
				Description: "The identifier of the user that belongs to the tenant.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "email",
				Description: "The email of the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "role",
				Description: "The role of the tenant user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the tenant member, i.e., invited or accepted.",
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
				Name:        "version_id",
				Description: "The version ID of this item.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromCamel(),
			},
		}),
	}
}

type tenantMember struct {
	TenantHandle string
	openapi.TenantUser
}

//// LIST FUNCTION

func listTenantMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	tenant := h.Item.(openapi.Tenant)

	tenantId := d.EqualsQualString("tenant_id")
	tenantHandle := d.EqualsQualString("tenant_handle")

	// Minimize API call
	if tenantId != "" && tenantId != tenant.Id {
		return nil, nil
	}
	if tenantHandle != "" && tenantHandle != tenant.Handle {
		return nil, nil
	}

	// Create Session
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listTenantMembers", "connection_error", err)
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
	var resp openapi.ListTenantUsersResponse
	var listDetails func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error)

	for pagesLeft {
		if resp.NextToken != nil {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.TenantMembers.List(ctx, tenant.Handle).NextToken(*resp.NextToken).Limit(maxResults).Execute()
				return resp, err
			}
		} else {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.TenantMembers.List(ctx, tenant.Handle).Limit(maxResults).Execute()
				return resp, err
			}
		}

		response, err := plugin.RetryHydrate(ctx, d, h, listDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

		if err != nil {
			plugin.Logger(ctx).Error("listTenantMembers", "list", err)
			return nil, err
		}

		result := response.(openapi.ListTenantUsersResponse)

		if result.HasItems() {
			for _, member := range *result.Items {
				d.StreamListItem(ctx, tenantMember{tenant.Handle, member})

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

func getTenantMember(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create Session
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getTenantMember", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()
	tenantHandle := d.EqualsQuals["tenant_handle"].GetStringValue()

	// check if id is empty
	if id == "" || tenantHandle == "" {
		return nil, nil
	}

	var resp openapi.TenantUser

	// execute get call
	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		resp, _, err = svc.TenantMembers.Get(ctx, tenantHandle, id).Execute()
		return resp, err
	}

	response, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

	if err != nil {
		plugin.Logger(ctx).Error("getTenantMember", "get", err)
		return nil, err
	}

	res := response.(openapi.TenantUser)

	return tenantMember{tenantHandle, res}, nil
}
