package pipes

import (
	"context"

	openapi "github.com/turbot/pipes-sdk-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "user_id",
			Hydrate:     getUserIdentity,
			Type:        proto.ColumnType_STRING,
			Description: "The unique identifier for the user.",
			Transform:   transform.FromField("Id"),
		},
	}, c...)
}

func getUserIdForConnection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	user, err := getUserIdentityMemoize(ctx, d, h)
	if err != nil {
		return nil, err
	}

	return user.(openapi.User).Id, nil
}

var getUserIdentityMemoize = plugin.HydrateFunc(getUserIdentityUncached).Memoize(memoize.WithCacheKeyFunction(getUserIdentityCacheKey))

func getUserIdentityCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "GetUserIdentity"
	return key, nil
}

func getUserIdentity(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	user, err := getUserIdentityMemoize(ctx, d, h)
	if err != nil {
		return nil, err
	}

	return user.(openapi.User), nil
}

func getUserIdentityUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// get the service connection for the service
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("GetUserIdentity", "connection_error", err)
		return nil, err
	}

	var resp openapi.User

	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		resp, _, err = svc.Actors.Get(ctx).Execute()
		return resp, err
	}

	response, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{})

	if err != nil {
		plugin.Logger(ctx).Error("GetUserIdentity", "error", err)
		return nil, err
	}

	user := response.(openapi.User)

	return user, nil
}
