# Atomic Commit: Two-Phase Commit (2PC) vs Saga Pattern

## 1. Two-Phase Commit (2PC)
Guarantees atomicity across multiple distinct databases/resources.
1. **Prepare Phase**: Coordinator asks all participants: `CAN COMMIT?`. Participants acquire locks, write changes to WAL, and reply `VOTE_COMMIT` or `VOTE_ABORT`.
2. **Commit Phase**: If all voted yes, coordinator logs `GLOBAL_COMMIT` and sends `COMMIT` to all; else sends `ABORT`.

*Weakness*: Blocking protocol. If coordinator crashes during commit phase, participants remain blocked holding locks indefinitely.

## 2. Saga Pattern (Eventual Consistency)
Series of local transactions. Each transaction updates database and publishes an event triggering the next step.
- **Compensating Transactions**: If Step 3 fails, the saga executes compensating actions in reverse order (e.g., Refund Payment, Cancel Reservation).
- **Execution Modes**:
  - *Choreography*: Services listen to domain events and decide autonomously.
  - *Orchestration*: Central orchestrator process drives step execution and compensation state machine.
