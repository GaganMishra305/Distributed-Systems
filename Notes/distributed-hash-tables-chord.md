# Distributed Hash Tables: Chord Protocol

Chord is a decentralized P2P lookup protocol that maps keys to nodes on an identifier ring.

## 1. Consistent Hash Ring
- Identifiers are ordered in an $m$-bit circle modulo $2^m$.
- Key $k$ is assigned to the first node whose identifier is equal to or follows $k$ in the ring space: $\text{successor}(k)$.

## 2. Finger Tables For $O(\log N)$ Routing
- Storing only successor allows basic lookup in $O(N)$ hops.
- To accelerate search, each node maintains finger table of $m$ entries:
  $$\text{finger}[i] = \text{successor}((n + 2^{i-1}) \bmod 2^m)$$
- Lookups skip exponentially large intervals across the ring, routing queries in $O(\log N)$ hops.

## 3. Node Joins & Stabilization
- Periodic background stabilization checks verify successor's predecessor pointers.
- Rebalance keys only between the newly joined node and its immediate successor.
