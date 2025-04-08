---
title: "Steampipe Table: pipes_connection - Query Pipes Connections using SQL"
description: "Allows users to query Pipes Connections, providing insights into the connection details, status, and configuration."
folder: "Connection"
---

# Table: pipes_connection - Query Pipes Connections using SQL

Pipes is a service within Steampipe that allows you to create and manage connections between different plugins. It provides a unified interface to set up and manage connections for various plugins, including AWS, Azure, GCP, and more. Pipes help you stay informed about the health and status of your connections and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `pipes_connection` table provides insights into the connections within Steampipe's Pipes service. As a DevOps engineer, explore connection-specific details through this table, including connection status, configuration, and associated metadata. Utilize it to uncover information about connections, such as those with specific configurations, the health status of connections, and the verification of connection configurations.

## Examples

### Basic info
Explore which plugins are connected to your system and identify any associated handles. This can help you manage your system's connections more effectively.

```sql+postgres
select
  id,
  plugin,
  handle,
  identity_handle
from
  pipes_connection;
```

```sql+sqlite
select
  id,
  plugin,
  handle,
  identity_handle
from
  pipes_connection;
```

### List connections using AWS plugin
Determine the areas in which AWS plugin is being used for connections. This is useful to understand and manage the distribution and usage of plugins across your connections.

```sql+postgres
select
  id,
  plugin,
  handle,
  jsonb_pretty(config) as config,
  identity_handle
from
  pipes_connection
where
  plugin = 'aws';
```

```sql+sqlite
select
  id,
  plugin,
  handle,
  config,
  identity_handle
from
  pipes_connection
where
  plugin = 'aws';
```

### List user connections
Explore which user connections are currently active. This is useful for understanding user engagement and tracking resource usage.

```sql+postgres
select
  id,
  plugin,
  handle,
  jsonb_pretty(config) as config,
  identity_handle
from
  pipes_connection
where
  identity_type = 'user';
```

```sql+sqlite
select
  id,
  plugin,
  handle,
  config,
  identity_handle
from
  pipes_connection
where
  identity_type = 'user';
```

### List organization workspaces
Review the configuration of your organization's workspaces to understand the plugins in use and their respective settings. This can help in assessing the elements within your organization's setup and identify areas for optimization or troubleshooting.

```sql+postgres
select
  id,
  plugin,
  handle,
  jsonb_pretty(config) as config,
  identity_handle
from
  pipes_connection
where
  identity_type = 'org';
```

```sql+sqlite
select
  id,
  plugin,
  handle,
  config,
  identity_handle
from
  pipes_connection
where
  identity_type = 'org';
```