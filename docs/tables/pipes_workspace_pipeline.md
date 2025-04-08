---
title: "Steampipe Table: pipes_workspace_pipeline - Query Pipes Workspace Pipelines using SQL"
description: "Allows users to query Workspace Pipelines in Pipes, specifically the pipeline configuration, status, and related metadata, providing insights into the pipeline execution and its management."
folder: "Pipeline"
---

# Table: pipes_workspace_pipeline - Query Pipes Workspace Pipelines using SQL

A Pipes Workspace Pipeline is a crucial component within the Pipes service that allows for the execution of a sequence of data processing steps. These pipelines can be used to extract, transform, and load (ETL) data from various sources into a target destination. They are highly configurable and can be managed and monitored for efficient data processing and analytics.

## Table Usage Guide

The `pipes_workspace_pipeline` table provides insights into the Workspace Pipelines within the Pipes service. As a Data Engineer, explore pipeline-specific details through this table, including pipeline configuration, status, and related metadata. Utilize it to uncover information about pipelines, such as their current execution status, configuration details, and to manage and monitor them effectively.

**Important Notes**

- This table supports optional quals. Queries with optional quals in the `where` clause are optimised to use Turbot Pipes filters.

- Optional quals are supported for the following columns:

  - `created_at`
  - `id`
  - `identity_handle`
  - `identity_id`
  - `pipeline`
  - `query_where` - Allows use of [query filters](https://turbot.com/pipes/docs/reference/query-filter). For a list of supported columns for pipelines, please see [Supported APIs and Columns](https://turbot.com/pipes/docs/reference/query-filter#supported-apis--columns). Please note that any query filter passed into the `query_where` qual will be combined with other optional quals.
  - `title`
  - `updated_at`
  - `workspace_handle`
  - `workspace_id`

## Examples

### Basic info
Explore which workspace pipelines are active, how frequently they run, and any associated tags. This can be useful in understanding the flow of data and identifying any potential bottlenecks in the pipeline.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  title,
  frequency,
  pipeline,
  args,
  tags,
  last_process_id
from
  pipes_workspace_pipeline;
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  title,
  frequency,
  pipeline,
  args,
  tags,
  last_process_id
from
  pipes_workspace_pipeline;
```

### List pipelines for a specific workspace
Discover the segments that are part of a particular workspace in order to understand how frequently they run, their specific parameters, and their latest process. This is useful for assessing the operational efficiency and identifying any potential bottlenecks within a specific workspace.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  title,
  frequency,
  pipeline,
  args,
  tags,
  last_process_id
from
  pipes_workspace_pipeline
where
  workspace_handle = 'dev';
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  title,
  frequency,
  pipeline,
  args,
  tags,
  last_process_id
from
  pipes_workspace_pipeline
where
  workspace_handle = 'dev';
```

### List pipelines of frequency type `interval` for a specific workspace
Analyze the settings to understand the pipelines configured to run at regular intervals in a specific workspace. This can be useful in managing and optimizing scheduled tasks within your workspace.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  title,
  frequency,
  pipeline,
  args,
  tags,
  last_process_id
from
  pipes_workspace_pipeline
where
  workspace_handle = 'dev'
  and frequency->>'type' = 'interval';
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  title,
  frequency,
  pipeline,
  args,
  tags,
  last_process_id
from
  pipes_workspace_pipeline
where
  workspace_handle = 'dev'
  and json_extract(frequency, '$.type') = 'interval';
```

### List pipelines for the `AWS Compliance CIS v1.4.0` dashboard created in the last 7 days
Explore which pipelines have been created in the past week for the AWS Compliance CIS v1.4.0 dashboard. This is useful for auditing recent changes and ensuring compliance standards are being met.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  title,
  frequency,
  pipeline,
  args,
  tags,
  last_process_id
from
  pipes_workspace_pipeline
where
  args->>'resource' = 'aws_compliance.benchmark.cis_v140'
  and created_at >= now() - interval '7 days';
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  title,
  frequency,
  pipeline,
  args,
  tags,
  last_process_id
from
  pipes_workspace_pipeline
where
  json_extract(args, '$.resource') = 'aws_compliance.benchmark.cis_v140'
  and created_at >= datetime('now', '-7 days');
```

### List pipelines for the `AWS Compliance CIS v1.4.0` dashboard created in the last 7 days using [query filter](https://turbot.com/pipes/docs/reference/query-filter)
Explore recently created pipelines for a specific compliance dashboard. This is particularly useful for tracking the latest updates and changes made within the past week.

```sql+postgres
select
  id,
  identity_handle,
  workspace_handle,
  title,
  frequency,
  pipeline,
  args,
  tags,
  last_process_id
from
  pipes_workspace_pipeline
where
  query_where = 'title = ''Scheduled snapshot: CIS v1.4.0'' and created_at >= now() - interval ''7 days''';
```

```sql+sqlite
select
  id,
  identity_handle,
  workspace_handle,
  title,
  frequency,
  pipeline,
  args,
  tags,
  last_process_id
from
  pipes_workspace_pipeline
where
  query_where = 'title = ''Scheduled snapshot: CIS v1.4.0'' and created_at >= datetime('now', '-7 days')';
```