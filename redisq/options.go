package redisq

const (
	DefaultAddr = "https://redisq.zkillboard.com/listen.php"
)

// Options is used to configure the RedisQ Client.
type Options struct {
	Addr string
}

// DefaultOptions for the package client.
func DefaultOptions() *Options {
	return &Options{
		Addr: DefaultAddr,
	}
}
