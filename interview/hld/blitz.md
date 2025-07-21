# System design blitz
Speed designs & deep dive points of each systems.

## Chat System
- Hold chats in Cassandra and have Partition Key as composite (ChatID, Time_bucket):
Key (room_id, time_bucket) determines partitions, message_id determine order
```
CREATE TABLE messages (
    message_id timeuuid,
    room_id timeuuid,
    time_bucket int,
    content text,
    ...
    PRIMARY KEY ((room_id, time_bucket), message_id) 
) WITH CLUSTERING ORDER BY (message_id DESC);
```

- For 1-1, thers also: 
CREATE TABLE user_messages (
    conversation_key text, -- format: "smaller_uid:bigger_uid"
    message_timestamp bigint,
    sender_uid bigint,
    target_uid bigint,
    message_content text,
    message_type text,
    PRIMARY KEY (conversation_key, message_timestamp)
) WITH CLUSTERING ORDER BY (message_timestamp DESC);

## Collab Doc

### General consideration
- Consider tradeoff between delta sync + SSE / WS.
- If storing operational logs, make sure to use c*
- Metadata storing can be in pq.
- Sticky session/document hashing/coord for directing Doc-to-Server
- Spark doesn't make sense since operations are relatively simple.

### Operational Transformation (OT) vs Conflict-Free Replicated Data Types (CRDT)
- CRDTs provide automatic conflict resolution without coordination
{
  id: "unique_operation_id",
  type: "insert",
  position: 42,
  character: "x",
  author: "user_123",
  timestamp: "2025-01-20T15:30:00Z",
  vectorClock: {user_123: 15, user_456: 23}
}

Pros: 
- Peer-to-peer capable
- Lower coordination needs for network
- Horizontal scaling friendly

Cons: 
- Harder to debug conflicts.
- Complex data structures


- OT works by transforming operations against each other to maintain consistency. OT algorithms transform these operations so they can be applied in any order while producing the same result. OT contains: insert, delete and retains. OT performs

Pros: 
- Easier to reason about & debug.
- Strong eventual consistency. 

Cons: 
- Server-side coordinations required
- High algorithmic complexity
- Higher network coordination overhead
- Vertical scaling challenges

## Newsfeed 

### DB modelling
Hotspot: 
- Trade off: shard by user -> hot node/ shard by posts -> joins and slower queries
Use reference table (user)
- Try to make it so that complex joins are co-located (within a shard/ node)
- Try to CDC.

### Content generations
- Consider starting as Pull-based model (creates on requests) and move on to Push-based model.
- Can consider hybrid approach for pull-based and push-based. (Pull for celebs)

Pull-based:
User requests feed
Query social graph to get followees
Fetch recent posts from each followee
Merge, rank, and return results

Push-based: 
User creates a post
System identifies all followers
Pre-generate and cache feeds for all followers
User feed requests served from cache

## Payment system
- Requires law complicances when handling data
- Prefer Idempotency when dealings. Instead of "reduce 20" do "set 80", adding sessions id to avoid double works. 
- Consider rollback (SAGA) or distributed transactions

Adaptive Authentication
- Low-risk transactions: Single-factor authentication with device fingerprinting
- Medium-risk transactions: SMS/email verification with velocity checks
- High-risk transactions: Multi-factor authentication with manual review

Progressive Friction Model:
- Session trust building: Reduce security friction as user establishes pattern
- Behavioral scoring: Machine learning models determine authentication requirements
- Context-aware security: Location, device, and time-based risk assessment

Security Performance Impact
- Encryption overhead: 10-50ms per operation
- Token validation: 5-15ms database lookup
- Risk scoring: 20-100ms ML inference
- MFA challenges: 30-60 second user workflow

## Search completion & suggestions
- Consider multi level caching & warming cache with predictions. 
- Approximation algorithms over exact ranking for performance. 
- Stale data tolerance to avoid real-time computation bottlenecks.
- Circuit breaker patterns with fallback to cached results.
- Consider multi stages pipelines: trie/bloom filter -> Elastic Search -> Ranking Algorithm.

### Types of searches
Bloom Filters: Query pre-filtering mechanism to eliminate expensive lookups for non-existent query prefixes.
In-mem Trie: Real-time auto complete suggestions. Service-level in-memory storage for suggestions matching. 
Inverted Index: Secondary lookup mechanism for Multi-word, Fuzzy Matching, Categories filters, etc...

## Metric monitor & alert
- Consider Push/Pull based metric collections
- Also consider push/pull for alerts valuations.
- Consider adding kafka to handle loads/sink processing.
- Generally partitions data by timestamps/timebucket.
- For log, use somesorts of inverted index. 