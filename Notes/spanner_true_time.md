# Google Spanner And TrueTime API
- Provides externally consistent (linearizable) transactions globally.
- TrueTime API represents time as an interval: $[t_{earliest}, t_{latest}]$ with bounded uncertainty $\epsilon \approx 7\text{ms}$.
- Wait-out-the-uncertainty: commit wait guarantees that transaction commit timestamp is strictly in the past before returning to client.
