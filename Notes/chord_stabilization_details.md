# Chord Periodic Stabilization Mechanics
1. Node $n$ asks its immediate successor $s$ for $s$'s predecessor $p$.
2. If $p$ falls between $n$ and $s$ on the ring:
   $$p \in (n, s)$$
   then $n$ updates its successor to $p$.
3. Node $n$ notifies $s$ of its presence so $s$ can update its predecessor if appropriate.
