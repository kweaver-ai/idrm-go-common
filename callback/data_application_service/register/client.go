package register

import grpc "google.golang.org/grpc"

type Client struct {
	conn grpc.ClientConnInterface
}

func New(conn grpc.ClientConnInterface) *Client { return &Client{conn: conn} }

// UserService implements Interface.
func (c *Client) UserService() UserServiceClient {
	return NewUserServiceClient(c.conn)
}

var _ Interface = &Client{}
