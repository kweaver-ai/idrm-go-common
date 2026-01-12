package v1

import grpc "google.golang.org/grpc"

type Client struct {
	conn grpc.ClientConnInterface
}

func New(conn grpc.ClientConnInterface) *Client { return &Client{conn: conn} }

// DataPush implements Interface.
func (c *Client) DataPush() DataPushCallbackServiceClient {
	return NewDataPushCallbackServiceClient(c.conn)
}

var _ Interface = &Client{}
