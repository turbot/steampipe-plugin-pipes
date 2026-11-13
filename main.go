package main

import (
	"github.com/turbot/steampipe-plugin-pipes/pipes"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: pipes.Plugin})
}
