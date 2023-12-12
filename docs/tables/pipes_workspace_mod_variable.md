---
title: "Steampipe Table: pipes_workspace_mod_variable - Query Pipes Workspace Mod Variables using SQL"
description: "Allows users to query Pipes Workspace Mod Variables, providing detailed information about each variable in the workspace."
---

# Table: pipes_workspace_mod_variable - Query Pipes Workspace Mod Variables using SQL

Pipes Workspace Mod Variable is a feature within the Pipes service that allows you to manage and monitor variables within your workspace. It provides a centralized way to set up and manage variables for various resources within your workspace. Pipes Workspace Mod Variable helps you stay informed about the status and details of your variables and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `pipes_workspace_mod_variable` table provides insights into the variables within a Pipes workspace. As a DevOps engineer, explore variable-specific details through this table, including names, descriptions, and associated metadata. Utilize it to uncover information about variables, such as their current values, default values, and whether they are required or optional.

## Examples

### List basic information for all variables for a mod in a workspace
Analyze the settings to understand the default and set values of all variables for a specific mod in a workspace. This is useful to validate the configuration and ensure it aligns with the intended setup.

```sql+postgres
select
  id,
  name,
  description,
  value_default,
  value_setting,
  value,
  type
from
  pipes_workspace_mod_variable
where
  workspace_id = 'w_cafeina2ip835d2eoacg' 
  and mod_alias = 'aws_thrifty';
```

```sql+sqlite
select
  id,
  name,
  description,
  value_default,
  value_setting,
  value,
  type
from
  pipes_workspace_mod_variable
where
  workspace_id = 'w_cafeina2ip835d2eoacg' 
  and mod_alias = 'aws_thrifty';
```

### List all variables which have an explicit setting in a workspace mod
Discover the segments that have explicit settings in a workspace mod. This is useful in understanding which variables have been specifically configured, aiding in better management and control of your workspace mods.

```sql+postgres
select
  id,
  name,
  description,
  value_default,
  value_setting,
  value,
  type
from
  pipes_workspace_mod_variable
where
  workspace_id = 'w_cafeina2ip835d2eoacg'
  and mod_alias = 'aws_thrifty' 
  and value_setting is not null;
```

```sql+sqlite
select
  id,
  name,
  description,
  value_default,
  value_setting,
  value,
  type
from
  pipes_workspace_mod_variable
where
  workspace_id = 'w_cafeina2ip835d2eoacg'
  and mod_alias = 'aws_thrifty' 
  and value_setting is not null;
```

### List details about a particular variable in a workspace mod
Explore the specifics of a certain variable in a workspace mod to understand its default and current values, along with its type. This is useful in assessing the configuration of the variable and its impact on the workspace mod.

```sql+postgres
select
  id,
  name,
  description,
  value_default,
  value_setting,
  value,
  type
from
  pipes_workspace_mod_variable
where
  workspace_id = 'w_cafeina2ip835d2eoacg'
  and mod_alias = 'aws_tags' 
  and name = 'mandatory_tags';
```

```sql+sqlite
select
  id,
  name,
  description,
  value_default,
  value_setting,
  value,
  type
from
  pipes_workspace_mod_variable
where
  workspace_id = 'w_cafeina2ip835d2eoacg'
  and mod_alias = 'aws_tags' 
  and name = 'mandatory_tags';
```