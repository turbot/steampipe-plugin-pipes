---
title: "Steampipe Table: pipes_organization_member - Query Pipes Organization Members using SQL"
description: "Allows users to query Pipes Organization Members, specifically the membership details of users in an organization, providing insights into user roles, permissions, and participation in projects."
folder: "Organization"
---

# Table: pipes_organization_member - Query Pipes Organization Members using SQL

Pipes is a data integration service that allows users to connect, transform, and automate data processes across different platforms. An Organization Member in Pipes refers to a user who is part of a specific organization, with assigned roles and permissions to access and manage data processes. This includes information on the user's participation in projects, their roles, and the permissions granted to them.

## Table Usage Guide

The `pipes_organization_member` table provides insights into user memberships within Pipes Organizations. As a data analyst or DevOps engineer, explore user-specific details through this table, including roles, permissions, and project participation. Utilize it to uncover information about users, such as their roles in the organization, their permissions, and their involvement in various projects.

## Examples

### Basic info
Explore which members belong to your organization and their current status. This can help you manage your team effectively by identifying active and inactive members.

```sql+postgres
select
  id,
  org_id,
  user_handle,
  status
from
  pipes_organization_member;
```

```sql+sqlite
select
  id,
  org_id,
  user_handle,
  status
from
  pipes_organization_member;
```

### List invited members
Explore which members have been invited to join your organization, helping you keep track of pending invitations and manage your team more effectively.

```sql+postgres
select
  id,
  org_id,
  user_handle,
  status
from
  pipes_organization_member
where
  status = 'invited';
```

```sql+sqlite
select
  id,
  org_id,
  user_handle,
  status
from
  pipes_organization_member
where
  status = 'invited';
```

### List members with owner role
Explore which members hold the 'owner' role within your organization. This helps in managing access controls and understanding the distribution of administrative responsibilities.

```sql+postgres
select
  id,
  org_id,
  user_handle,
  status
from
  pipes_organization_member
where
  role = 'owner';
```

```sql+sqlite
select
  id,
  org_id,
  user_handle,
  status
from
  pipes_organization_member
where
  role = 'owner';
```