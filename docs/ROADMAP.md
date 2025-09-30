# ROADMAP - DogFS

This roadmap defines the phases, goals, features, and acceptance criteria for building a learning-focused distributed object store in Go.
The project is designed for one developer and emphasizes distributed systems learning over production readiness.

---

## Phase 1 - Single-node Prototype

### Goal

- Build a local object store that demonstrates erasure coding and metadata tracking
- Lay the foundation for multi-node expansion later

### Features

- `PutObject(objectName, file)`: splits file into `k+m` shards, stores locally
- `GetObject(objectName)`: reconstructs file from shards
- Metadata stored in a local BoltDB/Badger DB
- Simple gRPC API (client <-> server)

### Acceptance Criteria

- Uploading and downloading an object works end-to-end
- Shards are written as separate files on disk
- Deleting one shard still allows reconstruction if òk shards remain
- Metadata correctly maps `objectName  shard list`

---

## Phase 2 - Multi-node Storage with Erasure Coding

### Goal

- Distribute shards across multiple storage nodes and demonstrate fault tolerance

### Features

- Multiple chunk servers (each with local disk storage)
- Metadata server tracks which node stores which shard
- Client uploads shards to multiple servers in parallel
- Client downloads shards directly from chunk servers using metadata
- Simulate node failures (kill a chunk server) and still retrieve objects

### Acceptance Criteria

- File upload places shards across ò3 chunk servers
- Download works as long as òk chunk servers respond
- Killing a chunk server does not break read availability (up to `m` failures)
- Manual repair tool can regenerate missing shards onto a replacement node

---

## Phase 3 - Distributed Metadata with Raft

### Goal

- Make metadata fault-tolerant and consistent using Raft consensus

### Features

- Metadata replicated across 3 metadata nodes using Raft
- Raft ensures linearizable updates (strong consistency)
- Clients connect to metadata leader (discovered via Raft)
- Metadata is sharded: `hash(objectName) -> metadata group` (for scaling)
- Failure handling:
  - Kill one metadata node -> cluster still serves reads/writes
  - Kill 2 of 3 -> cluster halts (demonstrating CAP tradeoff)

### Acceptance Criteria

- Metadata persists even if 1 node fails
- Writes acknowledged only after Raft majority commit
- Clients can PUT/GET while one metadata node is down
- Cluster halts if Raft majority unavailable

---

## Stretch Goals (Optional)

- Background shard repair
- Object listing across shards
- Security (TLS + auth)
- Automatic rebalancing when adding/removing nodes
