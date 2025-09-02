---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/pipes.svg"
brand_color: "#FABF1B"
display_name: "Turbot Pipes"
short_name: "pipes"
description: "Steampipe plugin for querying workspaces, connections and more from Turbot Pipes."
og_description: "Query Turbot Pipes with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/pipes-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Turbot Pipes + Steampipe

[Turbot Pipes](https://turbot.com/pipes) is an intelligence, automation & security platform built specifically for DevOps.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

For example:

```sql
select
  user_handle,
  role,
  status
from
  pipes_organization_member
where
  status = 'accepted';
```

```
> select user_handle, role, status from pipes_organization_member where status = 'accepted';
+---------------------+--------+----------+
| user_handle         | role   | status   |
+---------------------+--------+----------+
| victor-ujkk         | owner  | accepted |
| graza-io            | owner  | accepted |
| mattystratton-volw  | owner  | accepted |
+---------------------+--------+----------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/pipes/tables)**

## Get started

### Install

Download and install the latest Pipes plugin:

```bash
steampipe plugin install pipes
```

### Configuration

Installing the latest Pipes plugin will create a config file (`~/.steampipe/config/pipes.spc`) with a single connection named `pipes`:

```hcl
connection "pipes" {
  plugin = "pipes"

  # Turbot Pipes API token. If `token` is not specified, it will be loaded
  # from the `STEAMPIPE_CLOUD_TOKEN` environment variable and if not found
  # there will fallback to the `PIPES_TOKEN` environment variable. If both
  # are set simultaneously, `STEAMPIPE_CLOUD_TOKEN` will take preference.
  # token = "tpt_thisisnotarealtoken_123"

  # Turbot Pipes host URL. This defaults to "https://pipes.turbot.com".
  # You only need to set this if connecting to a remote Turbot Pipes database
  # not hosted in "https://pipes.turbot.com".
  # If `host` is not specified, it will be loaded from the `STEAMPIPE_CLOUD_HOST`
  # environment variable and if not found there will fallback to the
  # `PIPES_HOST` environment variable. If both are set simultaneously,
  # `STEAMPIPE_CLOUD_HOST` will take preference.
  # host = "https://pipes.turbot.com"
}
```

- `token` (required) - [API tokens](https://turbot.com/pipes/docs/da-settings#tokens) can be used to access the Turbot Pipes API or to connect to Turbot Pipes workspaces from the Steampipe CLI. May alternatively be set via the `STEAMPIPE_CLOUD_TOKEN` or `PIPES_TOKEN`. Note that the value in `STEAMPIPE_CLOUD_TOKEN` will take preference if both are set.
- `host` (optional) The Turbot Pipes Host URL. This defaults to `https://pipes.turbot.com`. You only need to set this if you are connecting to a remote Turbot Pipes database that is NOT hosted in `https://pipes.turbot.com`. This can also be set via the `STEAMPIPE_CLOUD_HOST` or `PIPES_HOST`. Note that the value in `STEAMPIPE_CLOUD_HOST` will take preference if both are set.

## Get Involved

- Open source: https://github.com/turbot/steampipe-plugin-pipes
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
