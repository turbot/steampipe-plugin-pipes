---
title: "Steampipe Table: pipes_user_email - Query Pipes User Emails using SQL"
description: "Allows users to query User Emails in Pipes, specifically the email address and associated user details, providing insights into user email data and potential anomalies."
folder: "User"
---

# Table: pipes_user_email - Query Pipes User Emails using SQL

Pipes allows users to connect and move data between various cloud and on-premises applications. It provides a simplified way to set up and manage integrations for various Steampipe plugins. Pipes helps you stay informed about the health and performance of your data integrations and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `pipes_user_email` table provides insights into user emails within Pipes service. As a data analyst or IT administrator, explore user-specific email details through this table, including associated user details and email addresses. Utilize it to uncover information about users, such as their email address, the associated user details, and the verification of user email data.

## Examples

### Basic info
Explore which users have registered their emails and when, to gain insights into user activity and status. This can help in understanding the growth and engagement levels of your user base.

```sql+postgres
select
  id,
  email,
  status,
  created_at
from
  pipes_user_email;
```

```sql+sqlite
select
  id,
  email,
  status,
  created_at
from
  pipes_user_email;
```