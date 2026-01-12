package v1

import grpc "google.golang.org/grpc"

type Client struct {
	conn grpc.ClientConnInterface
}

func New(conn grpc.ClientConnInterface) *Client { return &Client{conn: conn} }

// WorkOrder implements Interface.
func (c *Client) WorkOrder() WorkOrderCallbackServiceClient {
	return NewWorkOrderCallbackServiceClient(c.conn)
}

var _ Interface = &Client{}
