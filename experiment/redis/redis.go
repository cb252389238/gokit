package redis

type redis struct {
	Option
}

type Option struct {
	Host string
	Port int32
	Auth string
}

type RedisOption func(*Option)

func WithHost(host string) RedisOption {
	return func(option *Option) {
		option.Host = host
	}
}

func WithPort(port int32) RedisOption {
	return func(option *Option) {
		option.Port = port
	}
}

func WithAuth(auth string) RedisOption {
	return func(option *Option) {
		option.Auth = auth
	}
}

func New(opts ...RedisOption) *redis {
	o := &Option{}
	for _, opt := range opts {
		opt(o)
	}
	return &redis{*o}
}
