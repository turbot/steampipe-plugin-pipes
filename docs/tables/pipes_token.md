# Table: pipes_token

API tokens can be used to access the Turbot Pipes API or to connect to Turbot Pipes workspaces from the Steampipe CLI.

## Examples

### Basic info

```sql
select
  id,
  user_id,
  status,
  last4
from
  pipes_token;
```

### List inactive tokens

```sql
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

```sql
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
