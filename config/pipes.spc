connection "pipes" {
  plugin = "pipes"

  # Turbot Pipes API token. If `token` is not specified, it will be loaded
  # from the `STEAMPIPE_CLOUD_TOKEN` environment variable and if not found
  # there will fallback to the `PIPES_TOKEN` environment variable. If both
  # are set simultaneously, `STEAMPIPE_CLOUD_TOKEN` will take preference.
  # token = "tpt_thisisnotarealtoken_123"

  # Turbot Pipes host URL. This defaults to "https://pipes.turbot.com".
  # You only need to set this if connecting to a remote Turbot Pipes database
  # not hosted in "https://pipes.turbot.com".
  # If `host` is not specified, it will be loaded from the `STEAMPIPE_CLOUD_HOST`
  # environment variable and if not found there will fallback to the
  # `PIPES_HOST` environment variable. If both are set simultaneously,
  # `STEAMPIPE_CLOUD_HOST` will take preference.
  # host = "https://pipes.turbot.com"
}