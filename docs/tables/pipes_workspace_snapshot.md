# Table: pipes_workspace_snapshot

A Steampipe snapshot is a point in time view of a benchmark. It can be shared across workspaces or made public.

**Important notes:**

This table supports optional quals. Queries with optional quals in the `where` clause are optimised to use Turbot Pipes filters.

Optional quals are supported for the following columns:

- `created_at`
- `dashboard_name`
- `dashboard_title`
- `id`
- `query_where` - Allows use of [query filters](https://turbot.com/pipes/docs/reference/query-filter). For a list of supported columns for snapshots, please see [Supported APIs and Columns](https://turbot.com/pipes/docs/reference/query-filter#supported-apis--columns). Please note that any query filter passed into the `query_where` qual will be combined with other optional quals.
- `visibility`

## Examples

### Basic info

```sql
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

```sql
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

```sql
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

```sql
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

### Get the raw data for a particular snapshot

```sql
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

```sql
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

### List all controls in alarm for a benchmark snapshot

```sql
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

### List all controls for a benchmark snapshot

```sql
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

### List all controls in alarm for multiple snapshots

```sql
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
