# distributed-systems-challenges
a project to practice distributed system challenges

Some goals 
1. CAP, PACELC, and Real-World Tradeoffs
Not theory — how each constraint shows up in production.

2. Leader Election Algorithms
Raft, Paxos, Zab — and when consensus is not the answer.

3. Failure Detection & Heartbeating
Perfect vs eventually-perfect detectors, phi-accrual.

4. Idempotency & Exactly-Once Semantics
The difference between delivery vs processing guarantees.

5. Distributed Log Internals
Kafka ISR, Leader/Follower, commit index, truncation mechanics.

6. Backpressure Propagation
Preventing your system from melting under load.

7. Contention & Hotspot Mitigation
Sharding keys, consistent hashing, rendezvous hashing.

8. Distributed Transactions
2PC, Sagas, compensating steps, outbox pattern.

9. Concurrency Control
Pessimistic, optimistic, MVCC, OCC read/write conflicts.

10. Vector Clocks & Version Vectors
Solving real-world conflict resolution.

11. Service Coordination Patterns
Leases, fencing tokens, distributed locks (and why most are wrong).

12. Eventual Consistency Done Right
Anti-entropy, CRDTs, read-repair, hinted handoff.

13. Time in Distributed Systems
NTP drift, monotonic clocks, Lamport clocks, hybrid logical clocks.

14. Retry, Backoff, and Jitter Strategies
How to avoid thundering herds and cascading failures.

15. Circuit Breakers & Bulkheads
Save one bad upstream from killing your entire cluster.

16. Data Compaction & Storage Engine Internals
LSMs, SSTables, WAL, compaction strategies.

17. Network Realities
Packets drop, reorder, duplicate — assume the network lies to you.

18. Queue Semantics
At-most-once, at-least-once, exactly-once, dead-letter design.

19. Distributed Caching
Locality, cache invalidation, tiered caches, TTL discipline.

20. Hot Path Optimizations
Tail latency, p99 debugging, kernel bypass, gRPC/Protobuf tuning.

21. Operational Predictability
Capacity planning, SLOs, error budgets, on-call design.

22. Observability as a First-Class Skill
Tracing, structured logging, span correlations, flame graphs.

23. Disaster Testing
Chaos engineering, partition testing, fault injection.

24. Prod-First Mentality
Knowing how your system behaves — not just how it should work.