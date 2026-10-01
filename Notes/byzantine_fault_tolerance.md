# Byzantine Fault Tolerance (BFT)
- Assumes nodes can fail arbitrarily or act maliciously (sending conflicting messages).
- To tolerate $f$ Byzantine failures, a consensus network requires at least $3f + 1$ total nodes:
  $$N \ge 3f + 1$$
- Contrasts with crash-fault-tolerant (CFT) protocols (Raft, Paxos) which require $2f + 1$ nodes.
