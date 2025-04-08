---
title: "Steampipe Table: pipes_workspace_connection - Query Pipes Workspace Connections using SQL"
description: "Allows users to query Pipes Workspace Connections, specifically information about each connection within a workspace, providing insights into connection details, associated workspace, and other metadata."
folder: "Connection"
---

# Table: pipes_workspace_connection - Query Pipes Workspace Connections using SQL

Pipes Workspace Connection is a resource within the Steampipe Pipes plugin that represents a connection within a workspace. It provides details about each individual connection, including its ID, name, and associated workspace. This resource is crucial for managing and understanding the connections within a workspace in Steampipe Pipes.

## Table Usage Guide

The `pipes_workspace_connection` table provides insights into the connections within a workspace in Steampipe Pipes. As a DevOps engineer, explore connection-specific details through this table, including connection ID, name, and associated workspace. Utilize it to uncover information about connections, such as their details, the workspaces they are associated with, and other metadata.

## Examples

### Basic info
Explore the connections established within your workspace by identifying the unique identifiers associated with each connection. This can be particularly useful to assess and manage the various connections in a larger workspace setup.

```sql+postgres
select
  id,
  connection_id,
  workspace_id,
  identity_id,
  jsonb_pretty(connection) as connection
from
  pipes_workspace_connection;
```

```sql+sqlite
select
  id,
  connection_id,
  workspace_id,
  identity_id,
  connection
from
  pipes_workspace_connection;
```

### List workspace connections using AWS plugin
Discover the segments that utilize the AWS plugin within your workspace connections. This is useful for understanding how many and which connections are specifically associated with AWS, thereby aiding in resource allocation and management.

```sql+postgres
select
  id,
  connection_id,
  workspace_id,
  identity_id,
  jsonb_pretty(connection) as connection
from
  pipes_workspace_connection
where
  connection ->> 'plugin' = 'aws';
```

```sql+sqlite
select
  id,
  connection_id,
  workspace_id,
  identity_id,
  connection
from
  pipes_workspace_connection
where
  json_extract(connection, '$.plugin') = 'aws';
```

### List user workspace connections
Explore which user workspace connections exist, focusing on those that have a specific identity ID pattern. This can be beneficial for managing and understanding the relationships between users, their workspaces, and their connections.

```sql+postgres
select
  id,
  connection_id,
  workspace_id,
  identity_id,
  jsonb_pretty(connection) as connection
from
  pipes_workspace_connection
where
  identity_id like 'u_%';
```

```sql+sqlite
select
  id,
  connection_id,
  workspace_id,
  identity_id,
  connection
from
  pipes_workspace_connection
where
  identity_id like 'u_%';
```

### List organization workspace connections
Explore which organization workspace connections are in use. This can help in understanding the interdependencies and managing resources more effectively.

```sql+postgres
select
  id,
  connection_id,
  workspace_id,
  identity_id,
  jsonb_pretty(connection) as connection
from
  pipes_workspace_connection
where
  identity_id like 'o_%';
```

```sql+sqlite
select
  id,
  connection_id,
  workspace_id,
  identity_id,
  connection
from
  pipes_workspace_connection
where
  identity_id like 'o_%';
```