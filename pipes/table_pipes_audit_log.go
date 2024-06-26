package pipes

import (
	"context"
	"strings"

	openapi "github.com/turbot/pipes-sdk-go"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablePipesAuditLog(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pipes_audit_log",
		Description: "Audit logs record a series of events performed on an identity.",
		List: &plugin.ListConfig{
			Hydrate:    listAuditLogs,
			KeyColumns: plugin.AnyColumn([]string{"identity_handle", "identity_id"}),
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for an audit log.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "identity_id",
				Description: "The unique identifier for an identity where the action has been performed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "identity_handle",
				Description: "The handle name for an identity where the action has been performed.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "action_type",
				Description: "The action performed on the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "actor_avatar_url",
				Description: "The avatar of an actor who has performed the action.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "actor_display_name",
				Description: "The display name of an actor.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "actor_handle",
				Description: "The handle name of an actor.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "actor_id",
				Description: "The unique identifier of an actor.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "actor_ip",
				Description: "The IP address of the actor.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The time when the action was performed.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "target_handle",
				Description: "The handle name of the entity where the action has been performed.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "target_id",
				Description: "The unique identifier of the entity where the action has been performed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "data",
				Description: "The data which has been modified on the entity.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "process_id",
				Description: "The process id which this entry is a part of.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
		}),
	}
}

//// LIST FUNCTION

func listAuditLogs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create Session
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listAuditLogs", "connection_error", err)
		return nil, err
	}

	getUserIdentityCached := plugin.HydrateFunc(getUserIdentity).WithCache()
	commonData, err := getUserIdentityCached(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("listAuditLogs", "getUserIdentityCached", err)
		return nil, err
	}

	user := commonData.(openapi.User)

	identityHandle := d.EqualsQuals["identity_handle"].GetStringValue()
	identityId := d.EqualsQuals["identity_id"].GetStringValue()

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

	if identityHandle == "" && identityId == "" {
		return nil, nil
	} else if identityId != "" && strings.HasPrefix(identityId, "u_") {
		err = listUserAuditLogs(ctx, d, h, identityId, svc, maxResults)
	} else if identityId != "" && strings.HasPrefix(identityId, "o_") {
		err = listOrgAuditLogs(ctx, d, h, identityId, svc, maxResults)
	} else if identityHandle == user.Handle {
		err = listUserAuditLogs(ctx, d, h, identityHandle, svc, maxResults)
	} else {
		err = listOrgAuditLogs(ctx, d, h, identityHandle, svc, maxResults)
	}

	if err != nil {
		plugin.Logger(ctx).Error("listAuditLogs", "list", err)
		return nil, err
	}
	return nil, nil
}

func listOrgAuditLogs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, handle string, svc *openapi.APIClient, maxResults int32) error {
	var err error

	// execute list call
	pagesLeft := true
	var resp openapi.ListAuditLogsResponse
	var listDetails func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error)

	for pagesLeft {
		if resp.NextToken != nil {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.Orgs.ListAuditLogs(ctx, handle).NextToken(*resp.NextToken).Limit(maxResults).Execute()
				return resp, err
			}
		} else {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.Orgs.ListAuditLogs(ctx, handle).Limit(maxResults).Execute()
				return resp, err
			}
		}

		response, err := plugin.RetryHydrate(ctx, d, h, listDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

		if err != nil {
			plugin.Logger(ctx).Error("listOrgAuditLogs", "list", err)
			return err
		}

		result := response.(openapi.ListAuditLogsResponse)

		if result.HasItems() {
			for _, log := range *result.Items {
				d.StreamListItem(ctx, log)

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil
				}
			}
		}
		if result.NextToken == nil {
			pagesLeft = false
		} else {
			resp.NextToken = result.NextToken
		}
	}

	return nil
}

func listUserAuditLogs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, handle string, svc *openapi.APIClient, maxResults int32) error {
	var err error

	// execute list call
	pagesLeft := true
	var resp openapi.ListAuditLogsResponse
	var listDetails func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error)

	for pagesLeft {
		if resp.NextToken != nil {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.Users.ListAuditLogs(ctx, handle).NextToken(*resp.NextToken).Limit(maxResults).Execute()
				return resp, err
			}
		} else {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.Users.ListAuditLogs(ctx, handle).Limit(maxResults).Execute()
				return resp, err
			}
		}

		response, err := plugin.RetryHydrate(ctx, d, h, listDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

		if err != nil {
			plugin.Logger(ctx).Error("listUserAuditLogs", "list", err)
			return err
		}

		result := response.(openapi.ListAuditLogsResponse)

		if result.HasItems() {
			for _, log := range *result.Items {
				d.StreamListItem(ctx, log)

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil
				}
			}
		}
		if resp.NextToken == nil {
			pagesLeft = false
		} else {
			resp.NextToken = result.NextToken
		}
	}

	return nil
}
