package internal

import "time"

var Conf = new(config)

type config struct {
	Port     int
	Cache    bool
	CacheTTL time.Duration
	Release  bool
	LogPath  string
}

func init() {
	Conf = &config{
		Port:     9090,
		Cache:    false,
		CacheTTL: time.Hour * 2,
	}
}
