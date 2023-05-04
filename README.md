# go-cron 

Go-Cron is an easy to use cron service inside of a container.

# How to use?

```bash
    docker run -e CRON_TIME=1m -e COMMAND="ls -l" avhost/go-cron
```

# Variables

| Name | Default Value | Description |
| --- | --- | --- |
| COMMAND | | The command that should be executed by go-cron |
| CRON_TIME | 15m | In which interval the COMMAND has to be run (m, s, h) |
| LOGLEVEL | info | The loglevel (info, warn, error, debug) |
| ENABLE_SYSLOG | true | Enable/disable syslog |

