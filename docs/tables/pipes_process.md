---
title: "Steampipe Table: pipes_process - Query Pipes Processes using SQL"
description: "Allows users to query Pipes Processes, specifically the data related to the running processes in the Pipes plugin, providing insights into the process status, duration, and other details."
---

# Table: pipes_process - Query Pipes Processes using SQL

Pipes provides a way to query and manage processes for various Steampipe plugins, including retrieving process status, duration, and other details. It helps you stay informed about the state and performance of your Steampipe plugins and take appropriate actions based on the retrieved data.

## Table Usage Guide

The `pipes_process` table provides insights into running processes within Pipes. As a DevOps engineer, this table allows you to explore process-specific details, including status, duration, and associated metadata. Utilize it to monitor and manage processes, such as those with long durations, the state of different processes, and the verification of process details.

## Examples

### Basic info
Explore which processes are currently active and when they were created. This can help in understanding the operational flow and managing resources effectively.

```sql+postgres
select
  id,
  identity_handle,
  identity_type,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_process;
```

```sql+sqlite
select
  id,
  identity_handle,
  identity_type,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_process;
```

### List processes that are being run by an identity pipeline
Explore which processes are being managed by a specific pipeline. This can be useful to understand the pipeline's activity and monitor its performance.

```sql+postgres
select
  id,
  identity_handle,
  identity_type,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_process
where
  pipeline_id is not null;
```

```sql+sqlite
select
  id,
  identity_handle,
  identity_type,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_process
where
  pipeline_id is not null;
```

### List user processes
Explore which processes are initiated by users within a system. This can be particularly useful to monitor user activity and manage system resources effectively.

```sql+postgres
select
  id,
  identity_handle,
  identity_type,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_process
where
  identity_type = 'user';
```

```sql+sqlite
select
  id,
  identity_handle,
  identity_type,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_process
where
  identity_type = 'user';
```

### List running processes
Explore which processes are currently active. This is useful for monitoring ongoing operations and identifying any that may be causing issues or consuming excessive resources.

```sql+postgres
select
  id,
  identity_handle,
  identity_type,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_process
where
  state = 'running';
```

```sql+sqlite
select
  id,
  identity_handle,
  identity_type,
  pipeline_id,
  type,
  state,
  created_at
from
  pipes_process
where
  state = 'running';
```