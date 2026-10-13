# Split-Brain Prevention Strategies
1. **Majority Quorums**: Require $N/2 + 1$ node agreement. At most one partition can form a majority.
2. **Fencing Tokens**: Sequential numbers issued by lock coordinator. Server rejects writes with older fencing tokens.
3. **STONITH (Shoot The Other Node In The Head)**: Node physically cuts power/network of unresponsive peer via IPMI/ILO to ensure it is dead before taking over.
