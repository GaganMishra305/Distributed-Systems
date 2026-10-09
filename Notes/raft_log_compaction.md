# Raft Log Compaction And Snapshotting
- Without log compaction, the append-only log grows indefinitely and exhausts disk space.
- **Snapshotting**: Each server independently takes a snapshot of its state machine up to index $X$.
- Discards all log entries up to index $X$.
- **InstallSnapshot RPC**: Sent by leader when a follower lags so far behind that the leader has already discarded the needed log entries.
