---
title: "Steampipe Table: pipes_audit_log - Query Pipes Audit Logs using SQL"
description: "Allows users to query Pipes Audit Logs, specifically the logs that capture the activity within the Pipes service, providing insights into user interactions and potential security incidents."
folder: "Audit Log"
---

# Table: pipes_audit_log - Query Pipes Audit Logs using SQL

Pipes Audit Logs is a feature within the Pipes service that records user activity. It provides a comprehensive log of actions taken within the service, including who performed the action, what the action was, and when it was done. Pipes Audit Logs is a critical tool for understanding user behavior, troubleshooting issues, and investigating potential security incidents.

## Table Usage Guide

The `pipes_audit_log` table provides insights into user activity within the Pipes service. As a security analyst or administrator, explore detailed logs through this table, including the actor, action, and timestamp. Utilize it to uncover information about user behavior, troubleshoot issues, and investigate potential security incidents.

**Important Notes**

- You must specify an organization or user ID, or an organization or user handle, in the where or join clause using the `identity_id` or `identity_handle` columns respectively.

## Examples

### List audit logs for a user handle
Discover the actions taken by a particular user by examining their audit logs. This can be useful for analyzing user behavior or investigating potential security issues.

```sql+postgres
select
  id,
  action_type,
  jsonb_pretty(data) as data
from
  pipes_audit_log
where
  identity_handle = 'myuser';
```

```sql+sqlite
select
  id,
  action_type,
  data
from
  pipes_audit_log
where
  identity_handle = 'myuser';
```

### List audit logs for a user ID
Explore the actions taken by a specific user by analyzing the audit logs. This can help in understanding user behavior or investigating suspicious activities.

```sql+postgres
select
  id,
  action_type,
  jsonb_pretty(data) as data
from
  pipes_audit_log
where
  identity_id = 'u_c6fdjke232example';
```

```sql+sqlite
select
  id,
  action_type,
  data
from
  pipes_audit_log
where
  identity_id = 'u_c6fdjke232example';
```

### List audit logs for an organization handle
Analyze the actions performed in your organization by exploring the audit logs. This can be useful to track changes, identify unusual activity, and maintain security within your organization.

```sql+postgres
select
  id,
  action_type,
  jsonb_pretty(data) as data
from
  pipes_audit_log
where
  identity_handle = 'myorg';
```

```sql+sqlite
select
  id,
  action_type,
  data
from
  pipes_audit_log
where
  identity_handle = 'myorg';
```

### List audit logs for an organization ID
Explore which actions have been taken within a specific organization by analyzing its audit logs. This allows you to monitor and understand the operational activities and changes within your organization.

```sql+postgres
select
  id,
  action_type,
  jsonb_pretty(data) as data
from
  pipes_audit_log
where
  identity_id = 'o_c6qjjsaa6guexample';
```

```sql+sqlite
select
  id,
  action_type,
  data
from
  pipes_audit_log
where
  identity_id = 'o_c6qjjsaa6guexample';
```