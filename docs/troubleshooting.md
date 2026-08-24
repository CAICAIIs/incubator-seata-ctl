# Troubleshooting

`seata-ctl` provides read-only diagnostics for Seata Server.

## Diagnose

Run a quick health sweep:

```bash
seata-ctl
login --ip 127.0.0.1 --port 7091 --username seata --password seata
diagnose run
```

The command checks:

- server configuration
- TCP connectivity
- login token
- status endpoint
- global transaction query
- global lock query
- optional database connectivity

Use `--output table|json|yaml` to switch formats.

## Common failures

- `server address is not configured`: log in again with the right IP and port.
- `please login`: log in through the REPL first.
- `tcp connectivity failed`: check the Seata Server port and network path.
