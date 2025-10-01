package filesystem

import (
	"context"
	"log/slog"
	"net"
)

type CommandType int

const (
	GET CommandType = iota
	PUT
)

type Command struct {
	Type      CommandType
	Path      string // Path that will be uploaded to filesystem
	LocalPath string // Local filepath
}

type Client struct {
	config  ClientConfig
	conn    net.Conn
	context context.Context
}

func NewClient(ctx context.Context, config ClientConfig) *Client {
	return &Client{
		config:  config,
		context: ctx,
	}
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) Connect() {
	conn, err := net.Dial("tcp", c.config.IpAddress+":"+c.config.Socket)
	if err != nil {
		slog.Error("Connection Error: ", "error", err)
		return
	}
	c.conn = conn
}

/*
Commands to support:

PUT <path> <local file path>
GET <drive path> <path to download to>
*/
func (c *Client) Handle() {
	ctx, cancel := context.WithCancel(c.context)
	go c.handleConnection(ctx)

}

func (c *Client) handleConnection(ctx context.Context) {
	defer c.Close()
	for {
		select {
		case <-ctx.Done():
			slog.Info("Connection Thread: Context canceled, exiting.")
		default:
		}
	}

}

