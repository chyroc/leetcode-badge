package internal

import "time"

var Conf = new(Config)

type Config struct {
	Port     int
	Cache    bool
	CacheTTL time.Duration
}

func init() {
	Conf = &Config{
		Port:     9090,
		Cache:    false,
		CacheTTL: time.Hour * 2,
	}
}
