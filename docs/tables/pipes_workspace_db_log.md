---
title: "Steampipe Table: pipes_workspace_db_log - Query Pipes Workspace Database Logs using SQL"
description: "Allows users to query Pipes Workspace Database Logs. This table provides detailed information about all the logs related to the workspace database in the Pipes service."
folder: "Database Log"
---

# Table: pipes_workspace_db_log - Query Pipes Workspace Database Logs using SQL

Pipes is a service that provides a unified interface for querying, connecting, and managing data from different sources. The Workspace Database Logs in Pipes service offers a comprehensive log of all the activities and operations performed in the workspace database. It aids in tracking and monitoring the performance, errors, and other crucial metrics of the workspace database.

## Table Usage Guide

The `pipes_workspace_db_log` table provides insights into the logs of the workspace database in the Pipes service. As a database administrator or a DevOps engineer, explore log-specific details through this table, including error messages, timestamps, and associated metadata. Utilize it to monitor the performance, troubleshoot issues, and ensure the smooth operation of your workspace database.

## Examples

### List db logs for an actor by handle
Explore the database logs associated with a specific user to gain insights into their activities and the duration of their actions, which can be useful for auditing or troubleshooting purposes.

```sql+postgres
select
  id,
  workspace_id,
  workspace_handle,
  duration,
  query,
  log_timestamp
from
  pipes_workspace_db_log
where
  actor_handle = 'siddharthaturbot';
```

```sql+sqlite
select
  id,
  workspace_id,
  workspace_handle,
  duration,
  query,
  log_timestamp
from
  pipes_workspace_db_log
where
  actor_handle = 'siddharthaturbot';
```

### List db logs for an actor by handle in a particular workspace
Explore the activity history of a specific user within a particular workspace. This query is useful in monitoring user actions, identifying patterns, and troubleshooting issues.

```sql+postgres
select
  id,
  workspace_id,
  workspace_handle,
  duration,
  query,
  log_timestamp
from
  pipes_workspace_db_log
where
  actor_handle = 'siddharthaturbot'
  and workspace_handle = 'dev';
```

```sql+sqlite
select
  id,
  workspace_id,
  workspace_handle,
  duration,
  query,
  log_timestamp
from
  pipes_workspace_db_log
where
  actor_handle = 'siddharthaturbot'
  and workspace_handle = 'dev';
```

### List queries that took more than 30 seconds to execute
Identify instances where certain queries have taken longer than usual to execute. This can help in pinpointing inefficient queries, thus enabling optimization and improving overall system performance.

```sql+postgres
select
  id,
  workspace_id,
  workspace_handle,
  duration,
  query,
  log_timestamp
from
  pipes_workspace_db_log
where
  duration > 30000;
```

```sql+sqlite
select
  id,
  workspace_id,
  workspace_handle,
  duration,
  query,
  log_timestamp
from
  pipes_workspace_db_log
where
  duration > 30000;
```

### List all queries that ran in my workspace in the last hour
Explore which queries have been executed in your workspace within the past hour. This can be particularly useful for tracking recent activities and identifying any unusual or unexpected operations.

```sql+postgres
select
  id,
  workspace_id,
  workspace_handle,
  duration,
  query,
  log_timestamp
from
  pipes_workspace_db_log
where
  workspace_handle = 'dev'
  and log_timestamp > now() - interval '1 hr';
```

```sql+sqlite
select
  id,
  workspace_id,
  workspace_handle,
  duration,
  query,
  log_timestamp
from
  pipes_workspace_db_log
where
  workspace_handle = 'dev'
  and log_timestamp > datetime('now', '-1 hour');
```