# TUI

Open the diagnostic terminal interface:

```bash
seata-ctl
login --ip 127.0.0.1 --port 7091 --username seata --password seata
tui
```

Pages:

- `1` or `Tab`: diagnostics
- `2`: global transactions
- `3`: global locks

Keys:

- `r`: refresh now
- `a`: toggle auto refresh
- `q` or `Ctrl+C`: quit

Flags:

- `--refresh 5s`
- `--page-size 20`
- `--check-db`
- `--db-address host:port`
