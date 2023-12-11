---
title: "Steampipe Table: pipes_user_preferences - Query Pipes User Preferences using SQL"
description: "Allows users to query User Preferences in Pipes, specifically the user-defined settings and preferences, providing insights into user behaviors and customization patterns."
---

# Table: pipes_user_preferences - Query Pipes User Preferences using SQL

Pipes User Preferences is a feature within Turbot Pipes that allows users to customize and manage their individual settings and preferences. It provides users with the ability to personalize their experience and interactions with the Pipes platform. Pipes User Preferences helps users maintain control over their individual settings, providing a tailored user experience.

## Table Usage Guide

The `pipes_user_preferences` table provides insights into user-defined settings within Turbot Pipes. As a system administrator, explore user-specific details through this table, including individual preferences, settings, and associated metadata. Utilize it to uncover information about user behaviors, such as customization patterns, preference trends, and the verification of user-defined settings.

## Examples

### Basic info
Explore user preferences to understand their communication preferences and the time of their preference creation. This can be useful in tailoring communication strategies and understanding user engagement over time.

```sql+postgres
select
  id,
  communication_community_updates,
  communication_product_updates,
  communication_tips_and_tricks,
  created_at
from
  pipes_user_preferences;
```

```sql+sqlite
select
  id,
  communication_community_updates,
  communication_product_updates,
  communication_tips_and_tricks,
  created_at
from
  pipes_user_preferences;
```