---
title: "Steampipe Table: pipes_workspace_snapshot - Query Pipes Workspace Snapshots using SQL"
description: "Allows users to query Workspace Snapshots in Pipes, specifically snapshot details, providing insights into workspace data and potential changes."
folder: "Snapshot"
---

# Table: pipes_workspace_snapshot - Query Pipes Workspace Snapshots using SQL

Pipes Workspace Snapshots are a feature within the Pipes service that allow you to capture and store the state of your workspace at a specific point in time. This is useful for tracking changes, recovering data, and maintaining a historical record of your workspace. It provides a way to manage and monitor the evolution of your workspace over time.

## Table Usage Guide

The `pipes_workspace_snapshot` table provides insights into Workspace Snapshots within Pipes. As a data analyst or DevOps engineer, explore snapshot-specific details through this table, including snapshot creation time, workspace ID, and associated metadata. Utilize it to track changes over time, recover lost data, and maintain a historical record of your workspace.

**Important Notes**

- This table supports optional quals. Queries with optional quals in the `where` clause are optimised to use Turbot Pipes filters.

- Optional quals are supported for the following columns:

  - `created_at`
  - `dashboard_name`
  - `dashboard_title`
  - `id`
  - `query_where` - Allows use of [query filters](https://turbot.com/pipes/docs/reference/query-filter). For a list of supported columns for snapshots, please see [Supported APIs and Columns](https://turbot.com/pipes/docs/reference/query-filter#supported-apis--columns). Please note that any query filter passed into the `query_where` qual will be combined with other optional quals.
  - `visibility`

## Examples

### Basic info
Explore which workspaces are in use and their current state to manage resources effectively. This information can also provide insights into the visibility and the versioning of the schema used, along with any associated tags for better organization and tracking.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  state,
  visibility,
  dashboard_name,
  dashboard_title,
  schema_version,
  tags
from
  pipes_workspace_snapshot;
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  state,
  visibility,
  dashboard_name,
  dashboard_title,
  schema_version,
  tags
from
  pipes_workspace_snapshot;
```

### List snapshots for a specific workspace
Identify snapshots related to a specific workspace to gain insights into its state, visibility, and associated dashboard details. This can assist in managing and analyzing the workspace's configuration and evolution over time.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  state,
  visibility,
  dashboard_name,
  dashboard_title,
  schema_version,
  tags
from
  pipes_workspace_snapshot
where
  workspace_handle = 'dev';
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  state,
  visibility,
  dashboard_name,
  dashboard_title,
  schema_version,
  tags
from
  pipes_workspace_snapshot
where
  workspace_handle = 'dev';
```

### List public snapshots for the AWS Tags Limit benchmark dashboard across all workspaces
Explore which public snapshots are available for the AWS Tags Limit benchmark dashboard across all workspaces. This can be useful to assess the current configuration and visibility settings, providing insights into the overall state and benchmarking of AWS tags.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  state,
  visibility,
  dashboard_name,
  dashboard_title,
  schema_version,
  tags
from
  pipes_workspace_snapshot
where
  dashboard_name = 'aws_tags.benchmark.limit'
  and visibility = 'anyone_with_link';
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  state,
  visibility,
  dashboard_name,
  dashboard_title,
  schema_version,
  tags
from
  pipes_workspace_snapshot
where
  dashboard_name = 'aws_tags.benchmark.limit'
  and visibility = 'anyone_with_link';
```

### List snapshots for the AWS Compliance CIS v1.4.0 dashboard executed in the last 7 days
Explore the status and visibility of snapshots from the AWS Compliance CIS v1.4.0 dashboard taken in the past week. This can help in tracking compliance history and identifying changes or anomalies over time.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  state,
  visibility,
  dashboard_name,
  dashboard_title,
  schema_version,
  tags
from
  pipes_workspace_snapshot
where
  dashboard_name = 'aws_compliance.benchmark.cis_v140'
  and created_at >= now() - interval '7 days';
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  state,
  visibility,
  dashboard_name,
  dashboard_title,
  schema_version,
  tags
from
  pipes_workspace_snapshot
where
  dashboard_name = 'aws_compliance.benchmark.cis_v140'
  and created_at >= datetime('now', '-7 days');
```

### Get the raw data for a particular snapshot
Explore the raw snapshot data related to a specific user and workspace. This can be particularly useful for auditing changes or troubleshooting issues within a particular workspace environment.

```sql+postgres
select
  data
from
  pipes_workspace_snapshot
where
  identity_handle = 'myuser'
  and workspace_handle = 'dev'
  and id = 'snap_cc1ini7m1tujk0r0oqvg_12fie4ah78yl5rwadj7p6j63';
```

```sql+sqlite
select
  data
from
  pipes_workspace_snapshot
where
  identity_handle = 'myuser'
  and workspace_handle = 'dev'
  and id = 'snap_cc1ini7m1tujk0r0oqvg_12fie4ah78yl5rwadj7p6j63';
```

### List snapshots for the AWS Tags Limit benchmark dashboard executed in the last 7 days using [query filter](https://turbot.com/pipes/docs/reference/query-filter)
This example helps to identify the snapshots taken for the AWS Tags Limit benchmark dashboard in the last week. This can be useful for assessing recent performance metrics and understanding changes over time.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  state,
  visibility,
  dashboard_name,
  dashboard_title,
  schema_version,
  tags
from
  pipes_workspace_snapshot
where
  query_where = 'dashboard_name = ''aws_tags.benchmark.limit'' and created_at >= now() - interval ''7 days''';
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  state,
  visibility,
  dashboard_name,
  dashboard_title,
  schema_version,
  tags
from
  pipes_workspace_snapshot
where
  query_where = 'dashboard_name = ''aws_tags.benchmark.limit'' and created_at >= date('now','-7 days')';
```

### List all controls in alarm for a benchmark snapshot
This query helps you identify the controls that are in an alarm state for a specific benchmark snapshot. This is useful for pinpointing areas of concern within your system and addressing them proactively.

```sql+postgres
with unnest_panels as (
  select
    key as panel_name,
    value -> 'data' -> 'rows' as panel_data
  from
    pipes_workspace_snapshot as s
    cross join lateral jsonb_each(data -> 'panels')
  where
    -- modify with your snapshot ID
    s.id = 'snap_ckkeot9pveta5example_04czp8qosu9ir318e29example'
),
unnest_rows as (
  select
    panel_name,
    jsonb_array_elements(panel_data) as control
  from
    unnest_panels
)
select
  panel_name,
  control
from
  unnest_rows
where
  control ->> 'status' = 'alarm';
```

```sql+sqlite
with unnest_panels as (
  select
    key as panel_name,
    json_extract(value, '$.data.rows') as panel_data
  from
    pipes_workspace_snapshot as s,
    json_each(s.data, '$.panels')
  where
    -- modify with your snapshot ID
    s.id = 'snap_ckkeot9pveta5example_04czp8qosu9ir318e29example'
),
unnest_rows as (
  select
    panel_name,
    json_each(panel_data) as control
  from
    unnest_panels
)
select
  panel_name,
  control
from
  unnest_rows
where
  json_extract(control.value, '$.status') = 'alarm';
```

### List all controls for a benchmark snapshot
Determine the areas in which specific controls are set to 'alarm' status for a given snapshot. This is useful for identifying potential areas of concern or risk within your workspace.

```sql+postgres
with unnest_panels as (
  select
    key as panel_name,
    value -> 'data' -> 'rows' as panel_data
  from
    pipes_workspace_snapshot as s
    cross join lateral jsonb_each(data -> 'panels')
  where
    -- modify with your snapshot ID
    s.id = 'snap_ckkeot9pveta5example_04czp8qosu9ir318e29example'
),
unnest_rows as (
  select
    panel_name,
    jsonb_array_elements(panel_data) as control
  from
    unnest_panels
)
select
  panel_name,
  control
from
  unnest_rows
where
  control ->> 'status' = 'alarm';
```

```sql+sqlite
with unnest_panels as (
  select
    key as panel_name,
    json_extract(value, '$.data.rows') as panel_data
  from
    pipes_workspace_snapshot as s,
    json_each(s.data, '$.panels')
  where
    -- modify with your snapshot ID
    s.id = 'snap_ckkeot9pveta5example_04czp8qosu9ir318e29example'
),
unnest_rows as (
  select
    panel_name,
    json_each(panel_data) as control
  from
    unnest_panels
)
select
  panel_name,
  control.value
from
  unnest_rows
where
  json_extract(control.value, '$.status') = 'alarm';
```

### List all controls in alarm for multiple snapshots
This query is used to identify and analyze the controls that are in an alarm state across multiple snapshots. It's useful for monitoring and managing system health, allowing you to quickly pinpoint potential issues and take corrective action.

```sql+postgres
with unnest_panels as (
  select
    s.id as snapshot_id,
    s.created_at as snapshot_created_at,
    s.dashboard_name as dashboard_name,
    key as panel_name,
    value -> 'data' -> 'rows' as panel_data
  from
    pipes_workspace_snapshot as s
    cross join lateral jsonb_each(data -> 'panels')
  where
    -- Update or remove constraints for your workspace and goals
    s.identity_handle = 'example-community'
    and s.workspace_handle = 'tracker'
    and s.dashboard_name = 'github_tracker.benchmark.organization_checks'
    and s.created_at > current_timestamp - interval '3 days'
),
unnest_rows as (
  select
    up.snapshot_id,
    up.snapshot_created_at,
    up.dashboard_name,
    up.panel_name,
    jsonb_array_elements(panel_data) as control
  from
    unnest_panels as up
)
select
  *
from
  unnest_rows
where
  control ->> 'status' = 'alarm';
```

```sql+sqlite
with unnest_panels as (
  select
    s.id as snapshot_id,
    s.created_at as snapshot_created_at,
    s.dashboard_name as dashboard_name,
    key as panel_name,
    json_extract(value, '$.data.rows') as panel_data
  from
    pipes_workspace_snapshot as s,
    json_each(s.data, '$.panels')
  where
    -- Update or remove constraints for your workspace and goals
    s.identity_handle = 'example-community'
    and s.workspace_handle = 'tracker'
    and s.dashboard_name = 'github_tracker.benchmark.organization_checks'
    and s.created_at > datetime('now', '-3 days')
),
unnest_rows as (
  select
    up.snapshot_id,
    up.snapshot_created_at,
    up.dashboard_name,
    up.panel_name,
    json_each(up.panel_data) as control
  from
    unnest_panels as up
)
select
  *
from
  unnest_rows
where
  json_extract(control.value, '$.status') = 'alarm';
```