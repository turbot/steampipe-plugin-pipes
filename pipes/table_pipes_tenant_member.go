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
			Hydrate: listTenantMembers,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"tenant_id", "user_handle"}),
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
				Name:        "user_handle",
				Description: "The handle name of a user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Handle"),
			},
			{
				Name:        "member_id",
				Description: "The identifier of the user that belongs to the tenant.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("UserId"),
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
				Name:        "last_activity_at",
				Description: "The time of the last activity for the member.",
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
			// As user is a preserved keyword we can not use it as a column name. So renamed it to user_info.
			{
				Name:        "user_info",
				Description: "The user details.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("User"),
			},
		}),
	}
}

//// LIST FUNCTION

func listTenantMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// List the members for the tenant to which the calling user is in.
	// Only call list members for the tenant ID of the actor.
	callerIdentity, err := getUserIdentity(ctx, d, h)
	if err != nil {
		return nil, err
	}
	user := callerIdentity.(openapi.User)

	// Create Session
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pipes_tenant_member.listTenantMembers", "connection_error", err)
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
				resp, _, err = svc.TenantMembers.List(ctx, user.TenantId).NextToken(*resp.NextToken).Limit(maxResults).Execute()
				return resp, err
			}
		} else {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.TenantMembers.List(ctx, user.TenantId).Limit(maxResults).Execute()
				return resp, err
			}
		}

		response, err := plugin.RetryHydrate(ctx, d, h, listDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

		if err != nil {
			plugin.Logger(ctx).Error("pipes_tenant_member.listTenantMembers", "list", err)
			return nil, err
		}

		result := response.(openapi.ListTenantUsersResponse)

		if result.HasItems() {
			for _, member := range *result.Items {
				d.StreamListItem(ctx, member)

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
		plugin.Logger(ctx).Error("pipes_tenant_member.getTenantMember", "connection_error", err)
		return nil, err
	}
	userHandle := d.EqualsQuals["user_handle"].GetStringValue()
	tenantId := d.EqualsQuals["tenant_id"].GetStringValue()

	// check if id is empty
	if userHandle == "" || tenantId == "" {
		return nil, nil
	}

	var resp openapi.TenantUser

	// execute get call
	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		resp, _, err = svc.TenantMembers.Get(ctx, tenantId, userHandle).Execute()
		return resp, err
	}

	response, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

	if err != nil {
		plugin.Logger(ctx).Error("pipes_tenant_member.getTenantMember", "get", err)
		return nil, err
	}

	res := response.(openapi.TenantUser)

	return res, nil
}
