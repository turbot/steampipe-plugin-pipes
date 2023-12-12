---
title: "Steampipe Table: pipes_workspace_aggregator - Query Pipes Workspace Aggregators using SQL"
description: "Allows users to query Pipes Workspace Aggregators, specifically providing insights into the data aggregation and transformation processes within a workspace."
---

# Table: pipes_workspace_aggregator - Query Pipes Workspace Aggregators using SQL

Pipes is a service that allows for the aggregation and transformation of data from multiple sources. Workspace Aggregators in Pipes are responsible for executing these aggregation and transformation tasks within a specified workspace. They provide a unified view of data from disparate sources, enabling efficient data analysis and decision-making.

## Table Usage Guide

The `pipes_workspace_aggregator` table provides insights into the data aggregation processes within Pipes. As a data analyst, explore aggregator-specific details through this table, including the sources, transformations, and outputs of each aggregator. Utilize it to uncover information about the data pipelines, such as the data sources being aggregated, the transformations being applied, and the outputs being generated.

## Examples

### Basic info
Explore the connections and types of plugins used across different workspaces. This can help you manage and optimize the use of plugins in your system.

```sql+postgres
select
  id,
  handle,
  identity_handle,
  workspace_handle,
  type,
  plugin,
  connections
from
  pipes_workspace_aggregator;
```

```sql+sqlite
select
  id,
  handle,
  identity_handle,
  workspace_handle,
  type,
  plugin,
  connections
from
  pipes_workspace_aggregator;
```

### List aggregators for a specific workspace
Explore the different aggregators associated with a specific workspace to understand their types, plugins, and connections. This can be useful for managing and optimizing workspace resources.

```sql+postgres
select
  id,
  handle,
  identity_handle,
  workspace_handle,
  type,
  plugin,
  connections
from
  pipes_workspace_aggregator
where
  workspace_handle = 'dev';
```

```sql+sqlite
select
  id,
  handle,
  identity_handle,
  workspace_handle,
  type,
  plugin,
  connections
from
  pipes_workspace_aggregator
where
  workspace_handle = 'dev';
```

### List aggregators of plugin type `aws` for a specific workspace
Review the configuration for specific workspaces to pinpoint the specific locations where AWS plugins are used. This is especially useful for administrators who need to manage and monitor AWS resources across different workspaces.

```sql+postgres
select
  id,
  handle,
  identity_handle,
  workspace_handle,
  type,
  plugin,
  connections
from
  pipes_workspace_aggregator
where
  workspace_handle = 'dev'
  and plugin = 'aws';
```

```sql+sqlite
select
  id,
  handle,
  identity_handle,
  workspace_handle,
  type,
  plugin,
  connections
from
  pipes_workspace_aggregator
where
  workspace_handle = 'dev'
  and plugin = 'aws';
```

### List aggregators created in the last 7 days for a specific workspace
Discover the recently created aggregators within a specific workspace over the past week. This is beneficial for tracking recent changes and understanding the current state of your workspace.

```sql+postgres
select
  id,
  handle,
  identity_handle,
  workspace_handle,
  type,
  plugin,
  connections
from
  pipes_workspace_aggregator
where
  workspace_handle = 'dev'
  and created_at >= now() - interval '7 days';
```

```sql+sqlite
select
  id,
  handle,
  identity_handle,
  workspace_handle,
  type,
  plugin,
  connections
from
  pipes_workspace_aggregator
where
  workspace_handle = 'dev'
  and created_at >= datetime('now', '-7 days');
```