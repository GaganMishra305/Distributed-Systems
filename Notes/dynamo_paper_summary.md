# Amazon Dynamo Paper Summary (2007)
- **High Availability**: Always writeable shopping cart.
- **Key Techniques**:
  - Partitioning: Consistent Hashing with virtual nodes.
  - High availability for writes: Vector clocks with client-side reconciliation.
  - Temporary failure handling: Sloppy Quorum and Hinted Handoff.
  - Permanent failure synchronization: Anti-entropy using Merkle trees.
  - Membership: Gossip-based membership and failure detection.
