![image](https://hub.steampipe.io/images/plugins/turbot/pipes-social-graphic.png)

# Turbot Pipes Plugin for Steampipe

Use SQL to query workspaces, connections and more from Turbot Pipes.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/pipes)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/pipes/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-pipes/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install pipes
```

Run a query:

```sql
select
  user_handle,
  email,
  status
from
  pipes_organization_member
where
  status = 'accepted'
```

```
> select user_handle, email, status from pipes_organization_member where status = 'accepted'
+-------------+------------------+----------+
| user_handle | email            | status   |
+-------------+------------------+----------+
| mario       | mario@turbot.com | accepted |
| yoshi       | yoshi@turbot.com | accepted |
+-------------+------------------+----------+
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-pipes.git
cd steampipe-plugin-pipes
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/pipes.spc
```

Try it!

```
steampipe query
> .inspect pipes
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-pipes/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Turbot Pipes Plugin](https://github.com/turbot/steampipe-plugin-pipes/labels/help%20wanted)
