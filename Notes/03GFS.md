# Google File System

_PAPER: https://static.googleusercontent.com/media/research.google.com/en//archive/gfs-sosp2003.pdf_

## 1. Why is it hard to store big data in distributed fashion?
1. Performance -> sharding
2. Faults      -> tolerance
3. Tolernace   -> replication
4. Replication -> inconsistency
5. Inconsisten -> low-performance
_So the above loop makes it harder to develop permant, fault-tolerant distributed storage systems_


## 2. GFS Goals:
1. Big, fast
2. Global
3. Sharding
4. Automatic recovery


## 3. GFS use limits: 
1. Single data centers
2. Internal use
3. Sequential access


## 4. General Structure
```
        MASTER
c1  ->    C1
c2  ->    C1
c3  ->    C1  
.
.
.
cn  ->    Cm

```

```
MASTER DATA:
    - filename -> array of chunk handles
    - handle   -> list of chunks, version, primary, lease expiration
    - LOG, CHECKPOINT on disk
```



