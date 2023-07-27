# Table: pipes_user

Users can manage connections, organizations, and workspaces.

The `pipes_user` table returns information about the user whose token is used for authentication.

## Examples

### Basic info

```sql
select
  id,
  display_name,
  status,
  handle
from
  pipes_user;
```
