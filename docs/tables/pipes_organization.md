---
title: "Steampipe Table: pipes_organization - Query Pipes Organizations using SQL"
description: "Allows users to query Pipes Organizations. The table provides details such as the organization's name, description, and the date it was created."
folder: "Organization"
---

# Table: pipes_organization - Query Pipes Organizations using SQL

Users can create their own connections and workspaces, but they are not shared with other users. Organizations, on the other hand, include multiple users and are intended for organizations to collaborate and share workspaces and connections.

## Table Usage Guide

The `pipes_organization` table provides insights into Organizations within Pipes. As a data engineer, explore organization-specific details through this table, including the name, description, and creation date of the organization. Utilize it to uncover information about organizations, such as those with specific data flows, the relationships between organizations, and the verification of data pipelines.

## Examples

### Basic info
Analyze the status of different organizational entities within your system. This is useful to understand the overall health and activity status of various organizations for effective management.

```sql+postgres
select
  id,
  org_id,
  org_handle,
  status
from
  pipes_organization;
```

```sql+sqlite
select
  id,
  org_id,
  org_handle,
  status
from
  pipes_organization;
```