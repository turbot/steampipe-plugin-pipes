package pipes

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-pipes"

// Plugin creates this (pipes) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromGo(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		},
		DefaultRetryConfig: &plugin.RetryConfig{
			ShouldRetryErrorFunc: shouldRetryErrorFunc,
			BackoffAlgorithm:     "Exponential",
			RetryInterval:        250,
			CappedDuration:       2000,
			MaxAttempts:          12,
			MaxDuration:          30,
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "user_id",
				Hydrate: getUserIdForConnection,
			},
		},
		TableMap: map[string]*plugin.Table{
			"pipes_audit_log":                     tablePipesAuditLog(ctx),
			"pipes_connection":                    tablePipesConnection(ctx),
			"pipes_organization_member":           tablePipesOrganizationMember(ctx),
			"pipes_organization":                  tablePipesOrganization(ctx),
			"pipes_process":                       tablePipesProcess(ctx),
			"pipes_organization_workspace_member": tablePipesOrganizationWorkspaceMember(ctx),
			"pipes_tenant":                        tablePipesTenant(ctx),
			"pipes_tenant_member":                 tablePipesTenantMember(ctx),
			"pipes_token":                         tablePipesToken(ctx),
			"pipes_user":                          tablePipesUser(ctx),
			"pipes_user_email":                    tablePipesUserEmail(ctx),
			"pipes_user_preferences":              tablePipesUserPreferences(ctx),
			"pipes_workspace":                     tablePipesWorkspace(ctx),
			"pipes_workspace_aggregator":          tablePipesWorkspaceAggregator(ctx),
			"pipes_workspace_connection":          tablePipesWorkspaceConnection(ctx),
			"pipes_workspace_mod":                 tablePipesWorkspaceMod(ctx),
			"pipes_workspace_mod_variable":        tablePipesWorkspaceModVariable(ctx),
			"pipes_workspace_db_log":              tablePipesWorkspaceDBLog(ctx),
			"pipes_workspace_pipeline":            tablePipesWorkspacePipeline(ctx),
			"pipes_workspace_process":             tablePipesWorkspaceProcess(ctx),
			"pipes_workspace_snapshot":            tablePipesWorkspaceSnapshot(ctx),
		},
	}

	return p
}
