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

func tablePipesConnection(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pipes_connection",
		Description: "Connections represent a set of tables for a single data source.",
		List: &plugin.ListConfig{
			Hydrate: listConnections,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "identity_handle",
					Require: plugin.Optional,
				},
				{
					Name:    "identity_id",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"handle", "identity_handle"}),
			Hydrate:    getConnection,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "handle",
				Description: "The handle name for the connection.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "identity_id",
				Description: "The unique identifier for an identity where the connection has been created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "identity_handle",
				Description: "The handle name for an identity where the connection has been created.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getIdentityDetailsForConnection,
			},
			{
				Name:        "identity_type",
				Description: "The type of identity, which can be 'user' or 'org'.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getIdentityDetailsForConnection,
			},
			{
				Name:        "type",
				Description: "The connection type.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "plugin",
				Description: "The plugin name for the connection.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "config",
				Description: "The connection config details.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "created_at",
				Description: "The connection created time.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "created_by_id",
				Description: "The unique identifier of the user who created the connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "created_by",
				Description: "Information about the user who created the connection.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "updated_at",
				Description: "The connection's updated time.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "updated_by_id",
				Description: "The unique identifier of the user who last updated the connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromCamel(),
			},
			{
				Name:        "updated_by",
				Description: "Information about the user who last updated the connection.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "version_id",
				Description: "The version ID for the connection.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromCamel(),
			},
		}),
	}
}

//// LIST FUNCTION

func listConnections(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create Session
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listConnections", "connection_error", err)
		return nil, err
	}

	getUserIdentityCached := plugin.HydrateFunc(getUserIdentity).WithCache()
	commonData, err := getUserIdentityCached(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("listConnections", "getUserIdentityCached", err)
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
		err = listActorConnections(ctx, d, h, svc, maxResults)
	} else if identityId != "" && strings.HasPrefix(identityId, "u_") {
		err = listUserConnections(ctx, d, h, identityId, svc, maxResults)
	} else if identityId != "" && strings.HasPrefix(identityId, "o_") {
		err = listOrgConnections(ctx, d, h, identityId, svc, maxResults)
	} else if identityHandle == user.Handle {
		err = listUserConnections(ctx, d, h, identityHandle, svc, maxResults)
	} else {
		err = listOrgConnections(ctx, d, h, identityHandle, svc, maxResults)
	}

	if err != nil {
		plugin.Logger(ctx).Error("listConnections", "list", err)
		return nil, err
	}
	return nil, nil
}

func listOrgConnections(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, handle string, svc *openapi.APIClient, maxResults int32) error {
	var err error

	// execute list call
	pagesLeft := true
	var resp openapi.ListConnectionsResponse
	var listDetails func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error)

	for pagesLeft {
		if resp.NextToken != nil {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.OrgConnections.List(ctx, handle).NextToken(*resp.NextToken).Limit(maxResults).Execute()
				return resp, err
			}
		} else {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.OrgConnections.List(ctx, handle).Limit(maxResults).Execute()
				return resp, err
			}
		}

		response, err := plugin.RetryHydrate(ctx, d, h, listDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

		if err != nil {
			plugin.Logger(ctx).Error("listOrgConnections", "list", err)
			return err
		}

		result := response.(openapi.ListConnectionsResponse)

		for _, connection := range *result.Items {
			d.StreamListItem(ctx, connection)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil
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

func listUserConnections(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, handle string, svc *openapi.APIClient, maxResults int32) error {
	var err error

	// execute list call
	pagesLeft := true
	var resp openapi.ListConnectionsResponse
	var listDetails func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error)

	for pagesLeft {
		if resp.NextToken != nil {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.UserConnections.List(ctx, handle).NextToken(*resp.NextToken).Limit(maxResults).Execute()
				return resp, err
			}
		} else {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.UserConnections.List(ctx, handle).Limit(maxResults).Execute()
				return resp, err
			}
		}

		response, err := plugin.RetryHydrate(ctx, d, h, listDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

		if err != nil {
			plugin.Logger(ctx).Error("listUserConnections", "list", err)
			return err
		}

		result := response.(openapi.ListConnectionsResponse)

		for _, connection := range *result.Items {
			d.StreamListItem(ctx, connection)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil
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

func listActorConnections(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, svc *openapi.APIClient, maxResults int32) error {
	var err error

	// execute list call
	pagesLeft := true

	var resp openapi.ListConnectionsResponse
	var listDetails func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error)

	for pagesLeft {
		if resp.NextToken != nil {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.Actors.ListConnections(ctx).NextToken(*resp.NextToken).Limit(maxResults).Execute()
				return resp, err
			}
		} else {
			listDetails = func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
				resp, _, err = svc.Actors.ListConnections(ctx).Limit(maxResults).Execute()
				return resp, err
			}
		}

		response, err := plugin.RetryHydrate(ctx, d, h, listDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

		if err != nil {
			plugin.Logger(ctx).Error("listActorConnections", "list", err)
			return err
		}

		result := response.(openapi.ListConnectionsResponse)

		if result.HasItems() {
			for _, connection := range *result.Items {
				d.StreamListItem(ctx, connection)

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

//// HYDRATE FUNCTIONS

func getConnection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	identityHandle := d.EqualsQuals["identity_handle"].GetStringValue()
	handle := d.EqualsQuals["handle"].GetStringValue()

	// check if handle or identityHandle is empty
	if identityHandle == "" || handle == "" {
		return nil, nil
	}

	// Create Session
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getConnection", "connection_error", err)
		return nil, err
	}

	getUserIdentityCached := plugin.HydrateFunc(getUserIdentity).WithCache()
	commonData, err := getUserIdentityCached(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("getConnection", "getUserIdentityCached", err)
		return nil, err
	}

	user := commonData.(openapi.User)
	var resp interface{}

	if identityHandle == user.Handle {
		resp, err = getUserConnection(ctx, d, h, identityHandle, handle, svc)
	} else {
		resp, err = getOrgConnection(ctx, d, h, identityHandle, handle, svc)
	}

	if err != nil {
		plugin.Logger(ctx).Error("getConnection", "get", err)
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return resp.(openapi.Connection), nil
}

func getOrgConnection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, identityHandle string, handle string, svc *openapi.APIClient) (interface{}, error) {
	var err error

	// execute get call
	var resp openapi.Connection

	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		resp, _, err = svc.OrgConnections.Get(ctx, identityHandle, handle).Execute()
		return resp, err
	}

	response, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

	connection := response.(openapi.Connection)

	if err != nil {
		plugin.Logger(ctx).Error("getOrgConnection", "get", err)
		return nil, err
	}

	return connection, nil
}

func getUserConnection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, identityHandle string, handle string, svc *openapi.APIClient) (interface{}, error) {
	var err error

	// execute get call
	var resp openapi.Connection

	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		resp, _, err = svc.UserConnections.Get(ctx, identityHandle, handle).Execute()
		return resp, err
	}

	response, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

	connection := response.(openapi.Connection)

	if err != nil {
		plugin.Logger(ctx).Error("getUserConnection", "get", err)
		return nil, err
	}

	return connection, nil
}

func getIdentityDetailsForConnection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create Session
	svc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getIdentityDetailsForConnection", "connection_error", err)
		return nil, err
	}

	// Get the identity id from the connection hydrate object
	var identityId string
	switch w := h.Item.(type) {
	case openapi.Connection:
		identityId = *h.Item.(openapi.Connection).IdentityId
	case *openapi.Connection:
		identityId = *h.Item.(*openapi.Connection).IdentityId
	default:
		plugin.Logger(ctx).Debug("getIdentityDetailsForConnection", "Unknown Type", w)
	}

	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		resp, _, err := svc.Identities.Get(ctx, identityId).Execute()
		return resp, err
	}

	response, _ := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
	identity := response.(openapi.Identity)

	return &IdentityDetails{IdentityHandle: identity.Handle, IdentityType: identity.Type}, nil
}
