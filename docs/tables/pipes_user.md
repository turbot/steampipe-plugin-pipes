---
title: "Steampipe Table: pipes_user - Query Pipes User using SQL"
description: "Allows users to query Pipes Users, specifically providing insights into user details, including user ID, name, and email."
---

# Table: pipes_user - Query Pipes User using SQL

Pipes User is a resource within the Pipes platform that allows for user identification and management. It provides a centralized way to manage user details within the Pipes platform, including user ID, name, and email. Pipes User helps users to stay informed about the details and status of various users within the platform.

## Table Usage Guide

The `pipes_user` table provides insights into user details within the Pipes platform. As a system administrator, explore user-specific details through this table, including user ID, name, and email. Utilize it to uncover information about users, such as their identification details and email addresses, facilitating better user management and communication.

## Examples

### Basic info
Explore the status and identifiers of users in your system. This can be useful for understanding user activity and maintaining account security.

```sql+postgres
select
  id,
  display_name,
  status,
  handle
from
  pipes_user;
```

```sql+sqlite
select
  id,
  display_name,
  status,
  handle
from
  pipes_user;
```