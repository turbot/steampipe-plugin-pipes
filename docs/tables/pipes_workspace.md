---
title: "Steampipe Table: pipes_workspace - Query Pipes Workspaces using SQL"
description: "Allows users to query Pipes Workspaces, providing insights into workspace metadata such as creation time, last update time, and user information."
folder: "Workspace"
---

# Table: pipes_workspace - Query Pipes Workspaces using SQL

Workspaces provide a bounded context for managing, operating, and securing Steampipe resources. A workspace comprises a single Steampipe database instance as well as a directory of mod resources such as queries, benchmarks, and controls. Workspaces allow you to separate your Steampipe instances for security, operational, or organizational purposes.

## Table Usage Guide

The `pipes_workspace` table provides insights into workspaces within Pipes. As a data engineer, explore workspace-specific details through this table, including creation time, last update time, and associated user information. Utilize it to uncover information about workspaces, such as their current status, the user who last updated them, and the time of the last update.

## Examples

### Basic info
Explore which workspaces are active and who is managing them. This is useful for auditing purposes and ensuring proper workspace allocation.

```sql+postgres
select
  id,
  state,
  handle,
  identity_handle
from
  pipes_workspace;
```

```sql+sqlite
select
  id,
  state,
  handle,
  identity_handle
from
  pipes_workspace;
```

### List user workspaces
Explore which workspaces are specifically associated with user identities. This can help in managing user-specific resources and understanding user activities within these workspaces.

```sql+postgres
select
  id,
  state,
  handle,
  identity_handle
from
  pipes_workspace
where
  identity_type = 'user';
```

```sql+sqlite
select
  id,
  state,
  handle,
  identity_handle
from
  pipes_workspace
where
  identity_type = 'user';
```

### List organization workspaces
Explore which workspaces are linked to your organization. This query is useful in understanding the overall structure and dependencies within your organization.

```sql+postgres
select
  id,
  state,
  handle,
  identity_handle
from
  pipes_workspace
where
  identity_type = 'org';
```

```sql+sqlite
select
  id,
  state,
  handle,
  identity_handle
from
  pipes_workspace
where
  identity_type = 'org';
```

### List workspaces which are not running
Assess the elements within your system to pinpoint specific workspaces that are not currently active. This can help in managing resources and ensuring efficient workspace utilization.

```sql+postgres
select
  id,
  state,
  handle,
  identity_handle
from
  pipes_workspace
where
  state <> 'running';
```

```sql+sqlite
select
  id,
  state,
  handle,
  identity_handle
from
  pipes_workspace
where
  state <> 'running';
```