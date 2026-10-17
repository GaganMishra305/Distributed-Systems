# Flavors of Eventual Consistency
- **Read-Your-Writes Consistency**: User always sees updates submitted by themselves.
- **Monotonic Reads**: If a user sees a value at time $t$, they will never see an older value at $t' > t$.
- **Monotonic Writes**: Writes by the same user are serialized in the order they are issued.
- **Consistent Prefix Reads**: If a sequence of writes occurs in an order, anyone reading those writes will see them appear in the same order.
