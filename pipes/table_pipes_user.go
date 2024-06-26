package pipes

import (
	"context"

	openapi "github.com/turbot/pipes-sdk-go"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablePipesUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pipes_user",
		Description: "Users can manage connections, organizations, and workspaces.",
		List: &plugin.ListConfig{
			Hydrate: getUser,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "display_name",
				Description: "The display name for the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The user status.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "handle",
				Description: "The handle name of the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "url",
				Description: "The URL of the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "avatar_url",
				Description: "The avatar URL of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "created_at",
				Description: "The creation time of the user.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "preview_access_mode",
				Description: "The preview mode for the current user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "version_id",
				Description: "The version ID of the user.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "updated_at",
				Description: "The user's last updated time.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
		}),
	}
}

//// LIST FUNCTION

func getUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	getUserIdentityCached := plugin.HydrateFunc(getUserIdentity)
	commonData, err := getUserIdentityCached(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("getUser", "error", err)
		return nil, err
	}

	user := commonData.(openapi.User)

	d.StreamListItem(ctx, user)

	return nil, nil
}
