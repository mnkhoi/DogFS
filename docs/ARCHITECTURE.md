# DogFS - Architecture

## Overview

This project is a learning-focused distributed object store

- Data (objects) is split into erasure-coded shards and distributed across chunk servers
- Metadata (object -> shard mapping) is managed by a Raft cluster for strong consistency

The system prioritizes **Availability (A)** and **Partition tolerance (P)** at the data layer, but uses **Consistency (C)** for metadata to simplify correctness

---

## Components

### Metadata Service

- Runs as a Raft cluster (3 nodes)
- Stores mappings: `objectName -> [shardID, nodeID]`
- Strongly consistent (linearizable)
- Sharded for scalability (via hash partitioning)

### Chunk Servers

- Store shards as flat files on local disk
- Expose gRPC interface: `PutShard`, `GetShard`
- Stateless beyond their local storage
- Can be replaced/repaired without cluster-wide impact

### Client

- Talks to metadata leader to resolve shard locations
- Upload path:
  1 Contact metadata leader
  2 Encode file -> shards
  3 Upload shards to chunk servers
  4 Update metadata via Raft
- Download path:
  1 Contact metadata leader
  2 Fetch shards directly from chunk servers
  3 Decode file

---

## Tradeoffs

### Why Raft for Metadata?

- **Pros:** Simpler semantics, no conflicts, correctness easy to reason about
- **Cons:** Sacrifices availability of metadata if quorum lost

### Why Erasure Coding over Replication?

- **Pros:** Lower storage overhead (eg, 4+2 scheme = 15x vs 3x replication)
- **Cons:** More CPU/network overhead on encode/decode
- **Cons:** Repair traffic is heavier than simple re-replication

### Why Separate Metadata and Chunk Servers?

- **Pros:** Metadata stays small, fast, strongly consistent
- **Pros:** Chunk servers can scale independently
- **Cons:** Extra RPC hops (client must contact metadata then chunks)

### Why gRPC?

- **Pros:** Type-safe, good tooling in Go, streaming support
- **Cons:** Slightly more complex setup vs REST

### CAP Theorem Position

- **Metadata:** CP (Consistency + Partition tolerance)
- **Data (chunks):** AP (Availability + Partition tolerance) with erasure coding
- This hybrid model mirrors real-world designs (eg, HDFS, Ceph)

---

## Future Stretch Goals

- Background shard repair
- Object listing across shards
- Security (TLS + auth)
- Automatic rebalancing when adding/removing nodes
