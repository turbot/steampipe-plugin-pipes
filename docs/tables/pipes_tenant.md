---
title: "Steampipe Table: pipes_tenant - Query Pipes Tenants using SQL"
description: "Allows users to query Pipes Tenants, providing insights into tenant details such as creation time, update time, state, and usage thresholds, essential for tenant management and monitoring."
---

# Table: pipes_tenant - Query Pipes Tenants using SQL

Pipes Tenants are entities within the Pipes service representing groups or organizations that utilize the service. They contain comprehensive information about the tenant, including creation and update times, state, and various usage thresholds. Managing these tenants is crucial for ensuring proper allocation and monitoring of resources.

## Table Usage Guide

The `pipes_tenant` table provides detailed insights into the tenants within the Pipes service. As a System Administrator or a DevOps Engineer, use this table to explore tenant-specific details such as their state, creation and update times, and usage thresholds. This information is vital for tenant lifecycle management, resource allocation, and ensuring compliance with usage policies.

## Examples

### Basic info
Retrieve basic information about tenants, including their display name, handle, and current state. This can help in understanding the overall tenant landscape and monitoring their statuses.

```sql+postgres
select
  id,
  display_name,
  handle,
  state
from
  pipes_tenant;
```

```sql+sqlite
select
  id,
  display_name,
  handle,
  state
from
  pipes_tenant;
```

### List active tenants
Identify active tenants to manage and allocate resources efficiently. Active tenants are those that are currently utilizing the Pipes service.

```sql+postgres
select
  id,
  display_name,
  handle,
  state
from
  pipes_tenant
where
  state = 'active';
```

```sql+sqlite
select
  id,
  display_name,
  handle,
  state
from
  pipes_tenant
where
  state = 'active';
```

### List tenants created in the last 30 days
Find tenants that were created in the last 30 days to monitor new additions and ensure they are onboarded properly.

```sql+postgres
select
  id,
  display_name,
  handle,
  created_at
from
  pipes_tenant
where
  created_at >= (current_date - interval '30' day);
```

```sql+sqlite
select
  id,
  display_name,
  handle,
  created_at
from
  pipes_tenant
where
  created_at >= date('now','-30 day');
```

### Tenants exceeding compute usage threshold
Identify tenants whose compute usage exceeds the set threshold, which is essential for managing and preventing resource overuse.

```sql+postgres
select
  id,
  display_name,
  handle,
  usage_compute_threshold,
  usage_compute_action
from
  pipes_tenant
where
  usage_compute_threshold < (current_compute_usage);
```

```sql+sqlite
select
  id,
  display_name,
  handle,
  usage_compute_threshold,
  usage_compute_action
from
  pipes_tenant
where
  usage_compute_threshold < (current_compute_usage);
```