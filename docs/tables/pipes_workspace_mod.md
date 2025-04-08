---
title: "Steampipe Table: pipes_workspace_mod - Query Pipes Workspace Modules using SQL"
description: "Allows users to query Workspace in Pipes, specifically the details of each module in a workspace, providing insights into module configurations, versions, and dependencies."
folder: "Mod"
---

# Table: pipes_workspace_mod - Query Pipes Workspace Modules using SQL

Pipes Workspace Modules are components of the Pipes service that allow users to manage and configure their data pipelines. Each module in a workspace contains specific configurations, versions, and dependencies, which are crucial for the proper functioning and management of data pipelines.

A Steampipe mod is a portable, versioned collection of related Steampipe resources such as dashboards, benchmarks, queries, and controls. Steampipe mods and mod resources are defined in HCL, and distributed as simple text files. Modules can be found on the Steampipe Hub, and may be shared with others from any public git repository.

## Table Usage Guide

The `pipes_workspace_mod` table provides insights into the Workspace Modules within Pipes. As a Data Engineer, explore module-specific details through this table, including configurations, versions, and dependencies. Utilize it to uncover information about each module, such as its configuration details, the version it is running, and the dependencies it has on other modules.

## Examples

### Basic information about mods across all workspaces
Explore the status and details of modifications across all workspaces. This allows for efficient management and oversight of workspace modifications, including their current versions and state.

```sql+postgres
select
  id,
  path,
  alias,
  constraint,
  installed_version,
  state
from
  pipes_workspace_mod;
```

```sql+sqlite
select
  id,
  path,
  alias,
  constraint,
  installed_version,
  state
from
  pipes_workspace_mod;
```

### List mods for all workspaces of the user
Explore which modifications have been made across all of a user's workspaces. This can help in understanding the state of each workspace, including any constraints or version changes that have been applied.

```sql+postgres
select
  id,
  path,
  alias,
  constraint,
  installed_version,
  state
from
  pipes_workspace_mod
where
  identity_type = 'user';
```

```sql+sqlite
select
  id,
  path,
  alias,
  constraint,
  installed_version,
  state
from
  pipes_workspace_mod
where
  identity_type = 'user';
```

### List mods for all workspaces belonging to all organizations that the user is a member of
Explore which modifications are installed in all workspaces across organizations you're a member of. This can help assess the elements within each workspace and ensure they are up to date and functioning as expected.

```sql+postgres
select
  id,
  path,
  alias,
  constraint,
  installed_version,
  state
from
  pipes_workspace_mod
where
  identity_type = 'org';
```

```sql+sqlite
select
  id,
  path,
  alias,
  constraint,
  installed_version,
  state
from
  pipes_workspace_mod
where
  identity_type = 'org';
```

### List mods for a particular workspace belonging to an organization
Analyze the modifications made to a specific workspace within an organization to understand the current state, version, and constraints of those modifications. This is useful for managing and understanding the workspace's configuration and its impact on the organization's operations.

```sql+postgres
select 
  swm.id,
  swm.path,
  swm.alias,
  swm.constraint,
  swm.installed_version,
  swm.state
from 
  pipes_workspace_mod as swm 
  inner join pipes_organization as so on so.id = swm.identity_id
  inner join pipes_workspace as sw on sw.id = swm.workspace_id
where
  so.handle = 'testorg'
  and sw.handle = 'dev';
```

```sql+sqlite
select 
  swm.id,
  swm.path,
  swm.alias,
  swm.constraint,
  swm.installed_version,
  swm.state
from 
  pipes_workspace_mod as swm 
  join pipes_organization as so on so.id = swm.identity_id
  join pipes_workspace as sw on sw.id = swm.workspace_id
where
  so.handle = 'testorg'
  and sw.handle = 'dev';
```