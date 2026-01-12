package v1

import grpc "google.golang.org/grpc"

type Client struct {
	conn grpc.ClientConnInterface
}

// Notification implements Interface.
func (c *Client) Notification() NotificationServiceClient {
	return NewNotificationServiceClient(c.conn)
}

var _ Interface = &Client{}

func New(conn grpc.ClientConnInterface) *Client { return &Client{conn: conn} }
