# Gossip Protocols And Decentralized Failure Detection

## 1. Peer-to-Peer State Dissemination
- Instead of centralized coordinator, every node periodically selects $k$ random peers and exchanges state.
- **Anti-entropy**: Compares entire dataset (using Merkle Trees) to identify discrepancies.
- **Rumor mongering (Dissemination)**: High-speed propagation of updates; epidemic spread in $O(\log N)$ rounds.

## 2. SWIM Failure Detector Protocol
- Structured Weakly-consistent Infection-style Process Group Membership Protocol.
- Separates failure detection from membership dissemination.
1. **Direct Probe**: Node $A$ pings Node $B$. If $B$ acknowledges (`ack`), healthy.
2. **Indirect Probe**: If $B$ times out, $A$ picks $k$ random members ($C, D$) and asks them to ping $B$ directly. Defends against localized network partition between $A$ and $B$.
3. **Suspicion Mechanism**: If indirect probes fail, mark $B$ as `Suspect` with a grace period timer. If no alive message is received before expiry, declare $B$ `Dead`.
