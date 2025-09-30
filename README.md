# DogFS

DogFS is a **learning-focused distributed object store** written in Go
The project is designed as a hands-on exploration of distributed systems concepts, including **erasure coding**, **Raft consensus**, and **partitioned metadata**

DogFS is not intended as a production-ready file system - it is a personal learning project to gain practical experience in building distributed storage systems.

---

## Project Goals

The main goals of DogFS are to:

- Explore **CAP theorem trade-offs** for metadata and data layers.
- Learn how **erasure coding** improves storage efficiency and fault tolerance.
- Build a small-scale **distributed metadata service** using Raft consensus.
- Understand how to design a system with **separated metadata and chunk storage**.
- Implement failure scenarios and repair mechanisms.

For a detailed project roadmap, refer to [ROADMAP.md](docs/ROADMAP.md)
For a complete architectural overview and trade-offs, refer to [ARCHITECTURE.md](docs/ARCHITECTURE.md).

---

## Key Concepts

DogFS is designed around several core concepts:

1. **Metadata vs Chunk Storage Separation**
   Metadata servers track object locations and system state, while chunk servers store the actual object shards.

2. **Erasure Coding**
   Files are split into `k+m` shards using Reed-Solomon encoding, enabling data recovery from up to `m` lost shards.

3. **Raft Consensus**
   Metadata servers use Raft to maintain strong consistency and fault tolerance for object metadata.

4. **CAP Theorem Trade-off**
   - Metadata layer: CP (Consistency + Partition tolerance) using Raft
   - Data layer: AP (Availability + Partition tolerance) with erasure coding.

---

## Project Structure

The project will follow a phased development roadmap:

- **Phase 1:** Single-node prototype with erasure coding and metadata tracking.
- **Phase 2:** Multi-node storage with erasure coding and basic fault tolerance.
- **Phase 3:** Distributed metadata using Raft for strong consistency.

Refer to [ROADMAP.md](docs/ROADMAP.md) for full details.

---

## Getting Started

### Prerequisites
- Go >= 1.20
- Docker (for multi-node testing)
- Protobuf + gRPC plugins (for API generation)

### Running Phase 1 Prototype
1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/DogFS.git
    cd DogFS
    ```

2. Build:
    ```bash
    go build ./cmd/metadata
    go build ./cmd/chunk
    ```

3. Run a metadata server and chunk server locally:
    ```bash
    ./metadata
    ./chunk
    ```

4. Use the gRPC client to PUT and GET objects.

---

## Learning Outcomes

By building DogFS, you will gain hands-on experience with:
- Distributed consensus and Raft implementation.
- Erasure coding algorithms and trade-offs.
- Distributed metadata architecture.
- Client/server gRPC communication.
- Fault injection and repair strategies.

---

## References

- [ARCHITECTURE.md](docs/ARCHITECTURE.md) - Architectural design and trade-offs.
- [ROADMAP.md](docs/ROADMAP.md) - Work plan and phase breakdown.

---

## License

This project is for personal learning and is released under the MIT License.
