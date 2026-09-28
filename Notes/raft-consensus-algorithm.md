# Raft Consensus Algorithm

Raft is designed for understandability while providing equivalent safety guarantees to Multi-Paxos.

## 1. Node Roles and State Transitions
- **Leader**: Handles all client requests, replicates log entries to followers.
- **Candidate**: Transitory state when initiating leader election.
- **Follower**: Passive; responds to incoming RPCs from leaders and candidates.

```text
       [ Follower ] --(times out without heartbeat)--> [ Candidate ]
            ^                                                |
            |                                       (wins election)
            |                                                v
            +-----------(discovers higher term)------- [ Leader ]
```

## 2. Terms and Election Safety
- Time is divided into arbitrary numbered **terms**.
- Each node votes for at most one candidate per term on a first-come-first-served basis.
- **Election Restriction**: Candidate's log must be at least as up-to-date as the voter's log (comparing last term, then last index). Prevents nodes missing committed entries from winning.

## 3. Log Replication & Commit Index
- Leader sends `AppendEntries` RPC with:
  - `term`, `leaderId`
  - `prevLogIndex`, `prevLogTerm` (consistency check)
  - `entries[]`, `leaderCommit`
- Entry is committed once safely replicated on a majority (quorum) of nodes.
