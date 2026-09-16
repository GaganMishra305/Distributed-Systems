# Lamport Logical Clocks Summary
- Partial ordering of events in distributed systems.
- Happens-before relation: $a \to b$.
- If $a$ and $b$ are in same process and $a$ occurs before $b$, $a \to b$.
- If $a$ is sending of message and $b$ is receipt, $a \to b$.
- If $a \to b$ and $b \to c$, then $a \to c$ (transitivity).
