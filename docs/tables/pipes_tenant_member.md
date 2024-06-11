---
title: "Steampipe Table: pipes_tenant_member - Query Pipes Tenant Members using SQL"
description: "Allows users to query Pipes Tenant Members, providing insights into member details such as creation time, role, status, and associated tenant, essential for tenant member management and monitoring."
---

# Table: pipes_tenant_member - Query Pipes Tenant Members using SQL

Pipes Tenant Members are individual users associated with a tenant in the Pipes service. They contain detailed information about each member, including their role, status, and timestamps for creation and updates. Managing these members is crucial for maintaining proper access control and monitoring user activity within the tenant.

## Table Usage Guide

The `pipes_tenant_member` table provides detailed insights into the members of tenants within the Pipes service. As a System Administrator or a DevOps Engineer, use this table to explore member-specific details such as their role, status, and the tenant they belong to. This information is vital for managing access controls, monitoring user activity, and ensuring proper role assignments.

## Examples

### Basic info
Retrieve basic information about tenant members, including their email, role, and status. This can help in understanding the overall composition of tenant members and monitoring their statuses.

```sql+postgres
select
  id,
  email,
  role,
  status,
  tenant_id
from
  pipes_tenant_member;
```

```sql+sqlite
select
  id,
  email,
  role,
  status,
  tenant_id
from
  pipes_tenant_member;
```

### List active members
Identify members who have accepted their invitation to join the tenant, ensuring that only active members are utilizing the services.

```sql+postgres
select
  id,
  email,
  role,
  status,
  tenant_id
from
  pipes_tenant_member
where
  status = 'accepted';
```

```sql+sqlite
select
  id,
  email,
  role,
  status,
  tenant_id
from
  pipes_tenant_member
where
  status = 'accepted';
```

### List members by role
Retrieve a list of tenant members filtered by their role. This can help in understanding the distribution of roles within a tenant.

```sql+postgres
select
  id,
  email,
  role,
  status,
  tenant_id
from
  pipes_tenant_member
where
  role = 'admin';
```

```sql+sqlite
select
  id,
  email,
  role,
  status,
  tenant_id
from
  pipes_tenant_member
where
  role = 'admin';
```

### Members created in the last 30 days
Find members who were added in the last 30 days to monitor new additions and ensure they are properly integrated into the tenant.

```sql+postgres
select
  id,
  email,
  role,
  status,
  tenant_id,
  created_at
from
  pipes_tenant_member
where
  created_at >= (current_date - interval '30' day);
```

```sql+sqlite
select
  id,
  email,
  role,
  status,
  tenant_id,
  created_at
from
  pipes_tenant_member
where
  created_at >= date('now','-30 day');
```