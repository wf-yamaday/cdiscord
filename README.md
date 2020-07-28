# cdiscord

This tool inspired by [cslack](https://github.com/tkyshm/cslack).

```
usage: cdiscord [<flags>] <command> [<args> ...]

notify to discord cli tool

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  version
    display version

  message --webhook-url=WEBHOOK-URL [<flags>]
    send message to discord
```

## sample

```sh
echo "Hello, World" | cdiscord message -w ${your_discord_webhookURL}
```
