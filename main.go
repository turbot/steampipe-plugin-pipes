package main

import (
	"github.com/turbot/steampipe-plugin-pipes/pipes"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: pipes.Plugin})
}
