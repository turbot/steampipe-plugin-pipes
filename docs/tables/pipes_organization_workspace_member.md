---
title: "Steampipe Table: pipes_organization_workspace_member - Query Pipes Organization Workspace Members using SQL"
description: "Allows users to query Pipes Organization Workspace Members, specifically to fetch data about the members associated with a particular workspace within an organization."
---

# Table: pipes_organization_workspace_member - Query Pipes Organization Workspace Members using SQL

Pipes is a service that allows users to connect, transform, and observe data in motion. Within an organization's workspace in Pipes, members are the users who have access to the workspace and its resources. This table provides data about these members, including their roles, permissions, and other related details. Organization Workspace members can collaborate and share connections and dashboards.

## Table Usage Guide

The `pipes_organization_workspace_member` table provides insights into the members within a Pipes Organization Workspace. As a data analyst or a DevOps engineer, you can explore member-specific details through this table, including their roles, permissions, and other associated metadata. Utilize it to uncover information about members, such as those with administrative permissions, the roles assigned to various members, and the verification of their access rights.

## Examples

### Basic info
Explore which users are part of your organization's workspace in Steampipe and understand their roles and status. This can help in managing access and determining the level of participation of each member.

```sql+postgres
select
  id,
  org_handle,
  workspace_handle,
  user_handle,
  status,
  role
from
  pipes_organization_workspace_member;
```

```sql+sqlite
select
  id,
  org_handle,
  workspace_handle,
  user_handle,
  status,
  role
from
  pipes_organization_workspace_member;
```

### List invited members
Explore which members have been invited to join your organization's workspace. This can be useful in managing your team's access and roles, ensuring that all necessary invites have been sent, and tracking the status of these invitations.

```sql+postgres
select
  id,
  org_handle,
  workspace_handle,
  user_handle,
  status,
  role
from
  pipes_organization_workspace_member
where
  status = 'invited';
```

```sql+sqlite
select
  id,
  org_handle,
  workspace_handle,
  user_handle,
  status,
  role
from
  pipes_organization_workspace_member
where
  status = 'invited';
```

### List owners of an organization workspace
Explore which users have the 'owner' role in a specific organization workspace. This is beneficial in managing permissions and roles within the workspace, ensuring the right people have access to certain features and data.

```sql+postgres
select
  id,
  org_handle,
  workspace_handle,
  user_handle,
  status,
  role 
from
  pipes_organization_workspace_member 
where
  org_handle = 'testorg' 
  and workspace_handle = 'dev' 
  and role = 'owner';
```

```sql+sqlite
select
  id,
  org_handle,
  workspace_handle,
  user_handle,
  status,
  role 
from
  pipes_organization_workspace_member 
where
  org_handle = 'testorg' 
  and workspace_handle = 'dev' 
  and role = 'owner';
```

### Get details of a particular member in an organization workspace
Explore the status and role of a specific member within a particular workspace of an organization. This can help in understanding their access level and their contribution within the workspace.

```sql+postgres
select
  id,
  org_handle,
  workspace_handle,
  user_handle,
  status,
  role 
from
  pipes_organization_workspace_member 
where
  org_handle = 'testorg' 
  and workspace_handle = 'dev' 
  and user_handle = 'myuser';
```

```sql+sqlite
select
  id,
  org_handle,
  workspace_handle,
  user_handle,
  status,
  role 
from
  pipes_organization_workspace_member 
where
  org_handle = 'testorg' 
  and workspace_handle = 'dev' 
  and user_handle = 'myuser';
```