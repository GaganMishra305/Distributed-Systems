# Real-World Database Classifications Under PACELC
- **DynamoDB**: Default PA/EL (prioritizes availability under partition, latency normally). With strong consistency reads: PC/EC.
- **Apache Cassandra**: PA/EL. Configurable quorum reads/writes (`QUORUM`).
- **Google Spanner**: PC/EC. Synchronous Paxos replication across zones.
- **MongoDB**: PC/EC under `majority` write concern; PA/EL under default unacknowledged writes.
