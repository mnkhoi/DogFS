# ROADMAP - DogFS

This roadmap defines the phases for building DogFS:
a distributed file system inspired by HDFS, using Raft for metadata consistency

---

## Phase 1 - Local Prototype

### Goal

- Build a single-node prototype for learning
- Implement basic file PUT/GET with erasure coding and metadata tracking

### Features

- Metadata stored in-memory or BoltDB/Badger DB
- `PutObject(objectName, file)` splits file into `k+m` shards
- `GetObject(objectName)` reconstructs file from shards
- gRPC API for client-server communication

### Acceptance Criteria

- Upload and download work end-to-end
- Metadata correctly maps objectName - shard list
- Deleting a shard still allows reconstruction if -k shards exist

---

## Phase 2 - Multi-node Chunk Servers

### Goal

- Distribute block shards across multiple chunk servers
- Demonstrate fault tolerance with erasure coding

### Features

- Multiple chunk servers (Docker containers or processes)
- Metadata service assigns block IDs and placement
- Client uploads blocks to chunk servers
- Chunk servers send heartbeat to metadata service

### Acceptance Criteria

- Upload places shards across -3 chunk servers
- Download works as long as -k chunk servers respond
- Killing a chunk server does not break file reads

---

## Phase 3 - Raft Metadata Cluster

### Goal

- Replace single metadata service with a Raft cluster
- Provide fault-tolerant, strongly consistent metadata

### Features

- Metadata replicated across 3 Raft nodes
- Leader election and log replication
- Clients discover leader to perform metadata operations
- Metadata stores file namespace and block - chunk server mapping

### Acceptance Criteria

- Metadata survives failure of one node
- Writes acknowledged after Raft majority commit
- Cluster halts if Raft majority unavailable
- Clients can PUT/GET files while a metadata node is down

---

## Phase 4 - Scaling & Enhancements

### Goal

- Improve scalability and reliability

### Features

- Namespace sharding
- Automatic repair of lost shards
- Automatic rebalancing of blocks across chunk servers
- Security (TLS + authentication)

### Acceptance Criteria

- Metadata scaling via sharding
- Repair pipeline works under failures
- System supports adding/removing chunk servers without downtime
