# Vector Clocks And Logical Causality

## 1. The Problem: Physical Clocks Can't Be Trusted
- Quartz oscillators drift due to temperature, hardware variation, and clock skew.
- NTP synchronization can jump backwards or freeze, causing false timestamps across distributed nodes.
- Physical timestamps cannot reliably establish happens-before relationships ($A \to B$).

## 2. Lamport Timestamps
- Each process $P_i$ maintains counter $C_i$.
- Before local event: $C_i = C_i + 1$.
- When sending message $m$, attach timestamp $C_i$.
- On receiving $(m, C_{msg})$: $C_j = \max(C_j, C_{msg}) + 1$.
- **Limitation**: If $L(A) < L(B)$, we *cannot* conclude that $A \to B$ (could be concurrent).

## 3. Vector Clocks Algorithm
- In a system of $N$ processes, each process maintains vector $V[1..N]$:
  1. Before local event, $P_i$ increments $V_i[i]$.
  2. When sending message, attach $V_i$.
  3. When receiving message with vector $V_{msg}$:
     $$V_j[k] = \max(V_j[k], V_{msg}[k]) \quad \forall k$$
     $$V_j[j] = V_j[j] + 1$$

## 4. Detecting Concurrency
- $V_A < V_B$ iff $\forall k, V_A[k] \le V_B[k]$ and $\exists k, V_A[k] < V_B[k]$ (causally related).
- If neither $V_A \le V_B$ nor $V_B \le V_A$, events $A$ and $B$ are **concurrent** ($A \parallel B$), requiring conflict resolution (e.g. CRDTs or Dynamo siblings).
