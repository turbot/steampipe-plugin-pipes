---
title: "Steampipe Table: pipes_workspace_process - Query Pipes Workspace Processes using SQL"
description: "Allows users to query Pipes Workspace Processes, providing detailed information about each workspace process including its status, workspace ID, workspace name, and more."
---

# Table: pipes_workspace_process - Query Pipes Workspace Processes using SQL

Pipes allows users to create, manage, and monitor workspace processes. It provides a centralized platform for managing all workspace processes, including their status, workspace ID, workspace name, and more. Pipes helps users stay informed about the health and performance of their workspace processes and take appropriate actions when needed.

## Table Usage Guide

The `pipes_workspace_process` table offers a comprehensive view of Pipes Workspace Processes. As a DevOps engineer, you can use this table to explore detailed information about each workspace process, including its status, workspace ID, workspace name, and more. Utilize it to monitor the health and performance of your workspace processes and take appropriate actions when needed.

**Important Notes**

- This table supports optional quals. Queries with optional quals in the `where` clause are optimised to use Turbot Pipes filters.

- Optional quals are supported for the following columns:

  - `created_at`
  - `id`
  - `identity_handle`
  - `identity_id`
  - `pipeline_id`
  - `query_where` - Allows use of [query filters](https://turbot.com/pipes/docs/reference/query-filter). For a list of supported columns for workspace proceses, please see [Supported APIs and Columns](https://turbot.com/pipes/docs/reference/query-filter#supported-apis--columns). Please note that any query filter passed into the `query_where` qual will be combined with other optional quals.
  - `state`
  - `type`
  - `updated_at`
  - `workspace_handle`
  - `workspace_id`

## Examples

### Basic info
Explore the processes within a workspace to understand their types, states, and creation dates. This can be useful for monitoring the health and activity of your pipelines over time.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_workspace_process;
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_workspace_process;
```

### List processes for a pipeline
Explore the various processes associated with a specific pipeline in your workspace. This can be beneficial in assessing the state and type of these processes, and understanding when they were created.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_workspace_process
where
  pipeline_id = 'pipe_cfcgiefm1tumv1dis7lg';
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_workspace_process
where
  pipeline_id = 'pipe_cfcgiefm1tumv1dis7lg';
```

### List running processes for a pipeline
Explore which processes are currently running for a specific pipeline to manage resources and troubleshoot potential issues efficiently. This helps in maintaining the smooth operation of workflows and identifying any bottlenecks swiftly.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_workspace_process
where
  pipeline_id = 'pipe_cfcgiefm1tumv1dis7lg'
  and state = 'running';
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_workspace_process
where
  pipeline_id = 'pipe_cfcgiefm1tumv1dis7lg'
  and state = 'running';
```

### List running processes for a pipeline using [query filter](https://turbot.com/pipes/docs/reference/query-filter)
Explore the status of active processes within a specific pipeline. This query is useful for monitoring the ongoing operations and tracking their progress in real-time.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_workspace_process
where
  query_where = 'pipeline_id = ''pipe_cfcgiefm1tumv1dis7lg'' and state = ''running''';
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_workspace_process
where
  pipeline_id = 'pipe_cfcgiefm1tumv1dis7lg' and state = 'running';
```