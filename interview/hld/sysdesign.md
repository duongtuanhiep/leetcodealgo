# General speed
- RAM: 100 ns
- SSD: 10 us
- HDD: 10 ms


# Service Communication
## Comms technologies TCP
- HTTP/1.1 or 2: web browser native, slowest
- NATS: lightweight real-time communication with loose coupling. Transient messages, deliver at most or at least one.
- gRPC: strict, structured APIs for synchronous request/reply workflows.

## Comms pattern TCP
- Polling: Low freq update
- SSE: Server notifications, events
- WS: Interactive

## Comms technology UDP
Usecase: DNS, VoIP, IoT data collection, market data feed (high freq trade)
- RTP: media streaming 
- WebRTC: p2p

# Networking
## L4 LB
- High-performance, cost effective, w/o fancy load balancing, lower threat if compromised.
- Non-HTTP traffic needs to be load balanced (databases conn, game servers)

## L7 LB
- Microservices architecture requiring content-based routing
- Advanced features like caching, SSL termination, or content modification are needed
- Application-aware routing based on URLs, headers, or content
- Session persistence using application-level identifiers is necessary

# Database
Table scan: Slow reads or small table < 10k rows
Inverted-Index: Full text search (Elastic search)
Location: Geospatial index (PostGIS), geohash (Redis)
In-mem matches: Hash index (Redis)
B-Tree: Everything else (Postgres)

## PQ
- All SQL func, JsonB support, CTEs, window function, geospatial queries, better OLTP handling
- USE: ERP systems, inventory management, financial applications, social media.

### Challenges
Concurrency and Locking: 
- Pessimistic Locking & Optimistic Locking for high-lock & low-lock scenario. Essentially, use the Update + versioning either begin of a tx or at the end
- MVCC helps by Assigns transaction IDs + provides different snapshots of data.

Distributed Transactions: 
- Co-ordinator: default provided in Citus, is a 2PC.
- 2PC for writes accross nodes.

Cross-joins:
- Try to co-locate data.
- Reference table if small.

### Scaling
Try: 
- Batching & adding buffer queue.
- Queries Optimization.
- HA setup with 1 writes nodes.
- beef up instances
- caching response & queries

### Citus (MX) - Single leader and something in between.
Introduce coordinator node that can plan & break down queries and execute it parallelly.
For MX it replicates metadata to all nodes as reference so no need for coord. 
We have access to Data node / coord node.

Shard Keys:
- Tenant-based sharding for multi-tenant applications
- Time-based partitioning for time-series data
- Hash-based distribution for even data distribution

Queries:
- Are analized and sent to relevant workers based of metadata.
- Distributed queries & joins possible.
- Try to co-locate shards via shard keys.

Rebalancing:
- Citus: rebalance everything when you want to add shards/nodes.
- Citus MX: Use PQ's Logical replication (Pub (old shard location) -> Sub (new shard location))

Replication:
- At Citus level is statement based: send writes to multiple worker nodes to add redundancy.
- Citus MX use native streaming replication from Postgres.
- Within node, use postgres native streaming replication to provide further redundancy & failover.
- If strong strong consistency use sync repl (wait for ack on slaves before commit)
- If not, use async repl (commit right away once it's in WAL)

Failover/monitor:
- Patroni
- pg_autofailover

Cons:
- Doesn't support more complicated keys (GIS index) for sharding.

## Cassandra
Use: Time-series data, Multi-datacenter deployments, high write throughput.

### Challenges
Schema design and data modelling: 
- Design Around Queries, Not Entities, Table-Per-Query Pattern.
- (partition_key, clustering_col(s)): Determines partition locations, orderings.
- Increase cardinality for partition key (use composite) for better data distribution. 
- Use bucket to ensure small partition size = (some_id, some_date) 

Consistency: 
- Tunable consistency: One, local quorums, quorums, all.
- Use both reads and writes quorums when querying data.
- Only support resolution by cell timestamps. 

Data reconcilliations:
- Repair during quorum reads when coord detect inconsistency.
- Async repair via gossip protocol.
- Manual `nodetool` repair.
- BEWARE: Require sync timeclock in nodes & no manual overrides.
- BEWARE: Only guarantees latest version of each row-col, not ALL VERSIONS of same row-col

### Scaling
Gossip Protocol:
- When start, have some `seed` node for bootstraping.
- Runs async every sec.
- Exchange information with a few nodes.
- Streams missings data to nodes. 

## Redis
Use: Caching, session managements, analytics & leaderboards, rate-limit, distributed lock, msg queue.

Persistence:
- Data is persisted via snapshots / WAL (Append only file) Loggings. 
- Either shouldn't be rely on to keep data.

Transaction: 
- Multi (for locking pessimistic) 
- Watch (for optimistic)

### Usecase
Caching, session managements:
- Behave like online k-v store
- Evicting keys based on LRU/LFU or TTL

Leaderboard/ ranking:
- Operations with SortedSet `ZADD game_leaderboard 1500 "player1"`
- Range to get top K `ZREVRANGE game_leaderboard 0 9 WITHSCORES`
- TTL only apply to sets but can simulate it with K/V or another sortedSet with expiration date and manual cleanup

Distributed Ratelimit/lock:
- RateLimit: Create key `ratelimit:id:id:some-time-bucket` with ttl and count the value (K:V)
- Dist Lock: Use set with NotExist and Expire (Set)

MessageQueue & Pub/Sub:
- Use a List for message queue with `LPUSH/RPOP` (Load distribution)
- Use Pub/Sub for broadcasting (ChatRooms/live updates)

### Scaling 
Single node - Setinel:
- Replicate by Sync/Async Replication.
- In case of network failure, perform leader elections.

# Message Queue
## Kafka
Strength: 
- Persistence by default. Msgs on disk. 
- Event replay by moving timestamps. 
- Fault Tolerance. Distributed by message parttitions accross brokers. 
- High throughput. Topics data is append only logs.

Cons:
- No ordering accross multiple partitions. Try making idempotent/app level sort
- Complex & expensive since there's replications. 

## RabbitMQ
Strength:
- Complex Message Routing with (Direct, Fanout, Topic, Headers) routing
- Low Latency Message Delivery due to push-based.
- Protocol Flexibility: AMQP, MQTT, WS. 

Cons: 
- Limited throughput since routing & persistent affect cpu & i/o -> Use CloudAMQP as managed services or set up clusters.
- Broker bottleneck. also requires clusters set up & 
- No stream processings & replayabilities.

### Scaling
Setup:
- Each node must maintain connections to every other node.
- Needs to be odd num to prevent splitbrains. 
- Quorum queues with Raft Consensus if need data recovery.
- For metrics, don't even need to replicate.
- Clients can connect to any node and access any queues - RabbitMQ routes operations.

Scaling difficulty:
- Requires highly reliable, low-latency links between brokers. 
- Require user intervention for split-brains when network partition.
- Each node must maintain connections to every other node

# Special technique
- Circuit breaker
- Blue green deployments


## BQ
- Columnar storage

## Thanos - Prometheus