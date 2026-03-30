# RPC and Threads

## 1. Threads
- Each thread has its own separate stack-space and program-counter. (withing the program space)
- **WHY?**
    - I/O concurrency
    - Parallelism
    - Convenience (like background tasks)
- "Event-driven-programming" uses a single thread and waits for event-triggers, threads are better coz: _more compute utilization and simpler implementaion_

## 2. Thread Challenges
Challnges when writing code with mutithreading.
- RACE condition (resolved using LOCKs)
- Coordination   (resolved using channels[no shared memory], condition-variables[shared memory], waitGroup)
- Deadlock       (program efforts)

## 3. Misc-threads
- waitGroup() waits for a number of child threads.
- channels() avoid the usage of mutexes but adds to code complexity sometimes
- code: Practice/crawlers.go

## 4. RPC (remote procedure call)
- it uses stub-block infront of the server and the client
- code: Practice/kv.go

## 5. RPC semantics under failures
- at least once (we try to)
- at most once (checkin for duplicates)
- exactly once (ideally, not real)

## IMP:
- _GOAL of concurrency:_ is to improve performance by adding concurrent procdure to existing design.
- https://www.youtube.com/watch?v=oV9rvDllKEg
