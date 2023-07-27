# Table: pipes_user_email

The user email table allows users to manage their email.

The `pipes_user_email` table returns a list of emails added by users to their profile.

## Examples

### Basic info

```sql
select
  id,
  email,
  status,
  created_at
from
  pipes_user_email;
```
