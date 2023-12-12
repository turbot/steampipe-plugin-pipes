---
title: "Steampipe Table: pipes_token - Query Pipes Tokens using SQL"
description: "Allows users to query Pipes Tokens, specifically the token ID, creation time, and expiration time, providing insights into token usage and lifecycle management."
---

# Table: pipes_token - Query Pipes Tokens using SQL

Pipes Tokens are resources in the Pipes service that represent unique identifiers for authenticated sessions. They contain information about the token's creation, expiration, and the associated user. Pipes Tokens are essential for managing access and authentication within the Pipes service.

## Table Usage Guide

The `pipes_token` table provides insights into the authentication tokens within the Pipes service. As a Security Engineer, explore token-specific details through this table, including token ID, creation time, and expiration time. Utilize it to manage token lifecycle, monitor token usage, and ensure proper access controls are in place.

## Examples

### Basic info
Explore which user ID is linked with a specific ID, and gain insights into their status and last four digits of their token. This can be useful in understanding user activity and token usage patterns.

```sql+postgres
select
  id,
  user_id,
  status,
  last4
from
  pipes_token;
```

```sql+sqlite
select
  id,
  user_id,
  status,
  last4
from
  pipes_token;
```

### List inactive tokens
Discover the segments that contain inactive tokens to help maintain system security by ensuring outdated or unused tokens are identified and managed appropriately. This can be particularly useful in preventing unauthorized access or detecting potential system vulnerabilities.

```sql+postgres
select
  id,
  user_id,
  status,
  last4
from
  pipes_token
where
  status = 'inactive';
```

```sql+sqlite
select
  id,
  user_id,
  status,
  last4
from
  pipes_token
where
  status = 'inactive';
```

### List tokens older than 90 days
Identify instances where tokens have been active for more than 90 days. This could be useful for auditing purposes, ensuring tokens are not being used beyond a certain lifespan for security reasons.

```sql+postgres
select
  id,
  user_id,
  status,
  created_at,
  last4
from
  pipes_token
where
  created_at <= (current_date - interval '90' day);
```

```sql+sqlite
select
  id,
  user_id,
  status,
  created_at,
  last4
from
  pipes_token
where
  created_at <= date('now','-90 day');
```