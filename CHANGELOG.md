## v1.1.1 [2025-02-04]

_Bug fixes_

- Fixed the Github workflows to correctly build the Steampipe anywhere plugin builds.

## v1.1.0 [2025-01-31]

_Enhancements_

- Added the column `last_activity_at` to `pipes_organization_member`, `pipes_organization_workspace_member` and `pipes_tenant_member` tables. ([#47](https://github.com/turbot/steampipe-plugin-pipes/pull/47))

_Dependencies_

- Recompiled plugin with [pipes-go-sdk v0.12.0](https://github.com/turbot/pipes-sdk-go/blob/main/CHANGELOG.md#0120-2025-01-31). ([#47](https://github.com/turbot/steampipe-plugin-pipes/pull/47)) 

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. 
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. 

## v0.16.0 [2024-06-15]

_What's new?_

- New tables added
  - [pipes_tenant](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_tenant) ([#35](https://github.com/turbot/steampipe-plugin-pipes/pull/35))
  - [pipes_tenant_member](https://hub.steampipe.io/plugins/turbot/pipes/tables/pipes_tenant_member) ([#35](https://github.com/turbot/steampipe-plugin-pipes/pull/35))

## v0.15.0 [2024-05-22]

_Enhancements_

- The `user_id` column has now been assigned as a connection key column across all the tables which facilitates more precise and efficient querying across multiple Pipes connections. ([#27](https://github.com/turbot/steampipe-plugin-pipes/pull/27))
- The Plugin and the Steampipe Anywhere binaries are now built with the `netgo` package. ([#32](https://github.com/turbot/steampipe-plugin-pipes/pull/32))
- Added the `version` flag to the plugin's Export tool. ([#65](https://github.com/turbot/steampipe-export/pull/65))

_Bug fixes_

- Fixed the plugin to correctly authenticate against a custom tenant in `Pipes` instead of returning `401` errors. ([#30](https://github.com/turbot/steampipe-plugin-pipes/pull/30))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.10.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v5100-2024-04-10) that adds support for connection key columns. ([#27](https://github.com/turbot/steampipe-plugin-pipes/pull/27))

## v0.14.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#23](https://github.com/turbot/steampipe-plugin-pipes/pull/23))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#23](https://github.com/turbot/steampipe-plugin-pipes/pull/23))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-pipes/blob/main/docs/LICENSE). ([#23](https://github.com/turbot/steampipe-plugin-pipes/pull/23))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#22](https://github.com/turbot/steampipe-plugin-pipes/pull/22))

## v0.13.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#23](https://github.com/turbot/steampipe-plugin-pipes/pull/23))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#23](https://github.com/turbot/steampipe-plugin-pipes/pull/23))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-pipes/blob/main/docs/LICENSE). ([#23](https://github.com/turbot/steampipe-plugin-pipes/pull/23))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#22](https://github.com/turbot/steampipe-plugin-pipes/pull/22))

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
