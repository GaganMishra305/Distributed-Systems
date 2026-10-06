# CAP And PACELC Theorems

## 1. CAP Theorem Breakdown
In any asynchronous distributed data store across an unreliable network, you can guarantee at most two:
- **Consistency (Linearizability)**: Every read receives the most recent write or an error.
- **Availability**: Every non-failing node returns a non-error response for every request.
- **Partition Tolerance**: System functions despite network packet loss or split-brain.

Because network partitions ($P$) are physical inevitabilities, the real choice is:
- Under partition: choose **Consistency (CP)** (drop requests to preserve correctness) OR **Availability (AP)** (accept writes, risk divergence).

## 2. The PACELC Theorem Extension
Addresses the normal operation when **no partition** is occurring:
$$If \; P \implies choose \; A \; or \; C, \quad else \implies choose \; L \; (Latency) \; or \; C \; (Consistency)$$

- **PC/EC**: (e.g. Bigtable, Spanner) - CP during partition; prioritizes consistency over latency normally.
- **PA/EL**: (e.g. Cassandra, DynamoDB default) - AP during partition; prioritizes low latency over strong consistency normally.
