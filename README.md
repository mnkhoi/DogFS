# DogFS

DogFS is a **HDFS-inspired distributed file system** written in Go, enhanced with **Raft consensus for metadata**
It is designed as a personal learning project to explore core distributed systems concepts including **metadata replication, erasure coding, fault tolerance, and CAP trade-offs**.

DogFS is not intended as a production-ready file system, but as a **learning playground for distributed systems**.

---

## Project Goals

DogFS aims to:
- Replace the single NameNode bottleneck in HDFS with a Raft-based metadata cluster.
- Use erasure coding for efficient, fault-tolerant data storage.
- Learn and experiment with **consensus, metadata sharding, fault recovery, and CAP theorem trade-offs**.

See [ARCHITECTURE.md](docs/ARCHITECTURE.md) for the detailed system design
See [ROADMAP.md](docs/ROADMAP.md) for the project phases.

---

## How DogFS Works

### Components

1. **Metadata Raft Cluster** — stores namespace and block mapping with strong consistency.
2. **Chunk Servers** — store fixed-size blocks, handle read/write requests.
3. **Clients** — discover metadata leader, upload/download blocks directly.

### Data Flow

- **Write:** Client -> metadata leader -> chunk servers -> commit in Raft.
- **Read:** Client -> metadata leader -> chunk servers -> reconstruct file.

---

## Getting Started

### Prerequisites

- Go >= 1.20
- Docker
- Protobuf + gRPC

### Running the Prototype

1. Clone:

    ```bash
    git clone https://github.com/yourusername/DogFS.git
    cd DogFS
    ```

2. Build:

    ```bash
    go build ./cmd/metadata
    go build ./cmd/chunk
    ```

3. Run:

    ```bash
    ./metadata
    ./chunk
    ```

4. Use the client API to PUT/GET objects.

---

## Learning Outcomes

By building DogFS, you will gain hands-on experience in:

- Distributed consensus (Raft).
- Erasure coding for storage efficiency.
- Metadata separation and sharding.
- Client/server design with gRPC.
- Fault injection and recovery strategies.

---

## References

- [ARCHITECTURE.md](docs/ARCHITECTURE.md) — System design and trade-offs.
- [ROADMAP.md](docs/ROADMAP.md) — Project phases and milestones.

---

## License

GNUv3 License.
