package pipes

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"

	openapiclient "github.com/turbot/pipes-sdk-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type pipesConfig struct {
	Token *string `hcl:"token"`
	Host  *string `hcl:"host"`
}

func ConfigInstance() interface{} {
	return &pipesConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) pipesConfig {
	if connection == nil || connection.Config == nil {
		return pipesConfig{}
	}
	config, _ := connection.Config.(pipesConfig)
	return config
}

func connect(_ context.Context, d *plugin.QueryData) (*openapiclient.APIClient, error) {
	config := GetConfig(d.Connection)

	token := os.Getenv("STEAMPIPE_CLOUD_TOKEN")
	// If `STEAMPIPE_CLOUD_TOKEN` is not set - we try to get the token from `PIPES_TOKEN`
	if token == "" {
		token = os.Getenv("PIPES_TOKEN")
	}
	// token value present in the config takes precedence over environment variable
	if config.Token != nil {
		token = *config.Token
	}
	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	configuration := openapiclient.NewConfiguration()
	configuration.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))

	host := os.Getenv("STEAMPIPE_CLOUD_HOST")
	// If `STEAMPIPE_CLOUD_HOST` is not set - we try to get the token from `PIPES_HOST`
	if host == "" {
		host = os.Getenv("PIPES_HOST")
	}
	// host value present in the config takes precedence over environment variable
	if config.Host != nil {
		host = *config.Host
	}

	parsedURL, parseErr := url.Parse(host)
	if parseErr != nil {
		return nil, fmt.Errorf(`invalid host: %v`, parseErr)
	}
	if parsedURL.Host == "" {
		return nil, errors.New(`missing protocol or host`)
	}

	if parsedURL.Host != "cloud.steampipe.io" && parsedURL.Host != "pipes.turbot.com" {
		// Parse and frame the Primary Servers
		var primaryServers []openapiclient.ServerConfiguration
		for _, server := range configuration.Servers {
			serverURL, parseErr := url.Parse(server.URL)
			if parseErr != nil {
				return nil, fmt.Errorf(`invalid host: %v`, parseErr)
			}
			primaryServers = append(primaryServers, openapiclient.ServerConfiguration{URL: fmt.Sprintf("%s://%s%s", serverURL.Scheme, parsedURL.Host, serverURL.Path), Description: server.Description})
		}
		configuration.Servers = primaryServers

		// Parse and frame the Operation Servers
		operationServers := make(map[string]openapiclient.ServerConfigurations)
		for service, servers := range configuration.OperationServers {
			var serviceServers []openapiclient.ServerConfiguration
			for _, server := range servers {
				serverURL, parseErr := url.Parse(server.URL)
				if parseErr != nil {
					return nil, fmt.Errorf(`invalid host: %v`, parseErr)
				}
				serviceServers = append(serviceServers, openapiclient.ServerConfiguration{URL: fmt.Sprintf("%s://%s%s", serverURL.Scheme, parsedURL.Host, serverURL.Path), Description: server.Description})
			}
			operationServers[service] = serviceServers
		}
		configuration.OperationServers = operationServers
	}

	apiClient := openapiclient.NewAPIClient(configuration)

	return apiClient, nil
}
