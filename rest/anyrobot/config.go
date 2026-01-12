package anyrobot

// Config contains information about how to communicate with a AnyRobot cluster.
type Config struct {
	// Server is the address of the AnyRobot cluster (https://hostname:port).
	Server string `json:"server,omitempty"`
}
