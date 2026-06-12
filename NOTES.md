# Redis Notes

## Video 1 — What makes Redis fast

- **In-memory** — all data lives in RAM; persistence (RDB snapshots / AOF log) is opt-in
- **Single-threaded** — no locks, no context switches; atomicity of individual commands is free
- **I/O multiplexing** — uses `epoll`/`kqueue`; single thread never blocks on I/O, handles thousands of connections (same model as nginx)
- **Atomic operations** — single commands are atomic; multi-step logic needs `MULTI`/`EXEC` + `WATCH` (optimistic locking) or a Lua script
  - `MULTI`/`EXEC` has no rollback — failed commands don't stop the rest from executing
- **Core insight** — Redis bets that avoiding coordination overhead (locks, context switches) beats raw parallelism for I/O-bound workloads
