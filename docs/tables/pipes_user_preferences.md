# Table: pipes_user_preferences

User Preferences represents various preferences settings for a user e.g. email settings.

The `pipes_user_preferences` table returns preferences of a user whose token is used for authentication.

## Examples

### Basic info

```sql
select
  id,
  communication_community_updates,
  communication_product_updates,
  communication_tips_and_tricks,
  created_at
from
  pipes_user_preferences;
```
