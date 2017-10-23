package redisq

// Options is used to configure the RedisQ Client.
type Options struct {
	// Addr of the target RedisQ service. When left empty, defaults to DefaultAddr.
	Addr string
	// QueueID to identify the client. Set this to a unique value if you have
	// multiple services hitting the RedisQ API at the same time.
	//
	// NOTE: Save this value somewhere if you don't want to miss killmails!
	QueueID string
}

// DefaultOptions for the package client.
func DefaultOptions() *Options {
	return &Options{
		Addr: DefaultAddr,
	}
}
