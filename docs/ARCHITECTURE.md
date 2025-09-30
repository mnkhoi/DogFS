# DogFS Architecture – HDFS-like with Raft Metadata

## Overview

DogFS is a learning-focused distributed file system inspired by HDFS but enhanced with **Raft consensus for metadata**
This design removes the NameNode single point of failure while retaining strong consistency for metadata

DogFS uses:

- **Raft consensus** for metadata
- **Erasure coding** for efficient, fault-tolerant block storage
- **Separated metadata and chunk servers**

For project milestones, see [ROADMAP.md](docs/ROADMAP.md)

---

## Components

### 1. Metadata Raft Cluster

- Stores filesystem namespace: directories, files, and block mapping
- Stores block → chunk server mapping
- Strong consistency via Raft
- Leader election for fault tolerance
- Handles file creation, deletion, and block placement decisions

### 2. Chunk Servers

- Store fixed-size blocks (e.g., 64 MB)
- Use erasure coding for durability
- Serve block read/write requests from clients
- Send heartbeat messages to metadata cluster

### 3. Client

- Discovers Raft leader metadata node
- Requests metadata for files
- Uploads/downloads blocks directly to/from chunk servers

---

## Data Flow

### File Write

1. Client → Metadata leader: request file creation
2. Metadata leader → Raft cluster: commit metadata entry
3. Metadata returns block IDs + chunk server assignments
4. Client uploads blocks to chunk servers in parallel
5. Chunk servers acknowledge uploads → metadata leader commits completion

### File Read

1. Client → Metadata leader: request file metadata
2. Metadata returns block locations
3. Client downloads blocks from chunk servers and reconstructs file

---

## Trade-offs

### Metadata

- **Consistency**: Strong (Raft consensus)
- **Availability**: Limited by Raft quorum requirements
- **Partition tolerance**: Raft tolerates partitions with leader election

### Data Layer

- **Consistency**: Eventual for block content (erasure coding + quorum reads)
- **Availability**: High — data available as long as k of k+m shards exist
- **Partition tolerance**: High

---

## CAP Theorem Position

- **Metadata layer**: CP (Consistency + Partition tolerance)
- **Data layer**: AP (Availability + Partition tolerance)

---

## Future Enhancements

- Namespace sharding for scaling metadata
- Background repair of lost shards
- Automatic rebalancing of blocks
- Security (TLS, authentication)
