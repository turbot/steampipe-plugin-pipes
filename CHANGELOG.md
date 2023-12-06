## v0.12.2 [2023-12-06]

_Bug fixes_

- Fixed the invalid Go module path of the plugin. ([#20](https://github.com/turbot/steampipe-plugin-pipes/pull/20))

## v0.12.1 [2023-10-04]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#9](https://github.com/turbot/steampipe-plugin-pipes/pull/9))

## v0.12.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#7](https://github.com/turbot/steampipe-plugin-pipes/pull/7))
- Recompiled plugin with Go version `1.21`. ([#7](https://github.com/turbot/steampipe-plugin-pipes/pull/7))

## v0.11.1 [2023-08-31]

_Bug fixes_

- Fixed `data` column query errors in `pipes_workspace_snapshot` table. ([#5](https://github.com/turbot/steampipe-plugin-pipes/pull/5))

## v0.11.0 [2023-07-27]

_What's new?_

- The [Steampipe Cloud](https://hub.steampipe.io/plugins/turbot/steampipecloud/tables) tables have now been rebranded to use `Pipes` instead:
  - [pipes_audit_log](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_audit_log)
  - [pipes_connection](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_connection)
  - [pipes_organization](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_organization)
  - [pipes_organization_member](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_organization_member)
  - [pipes_organization_workspace_member](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_organization_workspace_member)
  - [pipes_process](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_process)
  - [pipes_token](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_token)
  - [pipes_user](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_user)
  - [pipes_user_email](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_user_email)
  - [pipes_user_preferences](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_user_preferences)
  - [pipes_workspace](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_workspace)
  - [pipes_workspace_aggregator](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_workspace_aggregator)
  - [pipes_workspace_connection](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_workspace_connection)
  - [pipes_workspace_db_log](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_workspace_db_log)
  - [pipes_workspace_mod](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_workspace_mod)
  - [pipes_workspace_mod_variable](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_workspace_mod_variable)
  - [pipes_workspace_pipeline](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_workspace_pipeline)
  - [pipes_workspace_process](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_workspace_process)
  - [pipes_workspace_snapshot](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_workspace_snapshot)
