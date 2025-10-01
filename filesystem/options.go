package filesystem

type Connection struct {
	IpAddress string
	Socket    string
}

type ClientConfig struct {
	Connection
}

type MetadataNodeConfig struct {
	Connection
}

type ChunkingNodeConfig struct {
	Connection
}
