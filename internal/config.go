package internal

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

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

func InitConfig() error {
	if Conf.Cache {
		defaultCache = cache.New(Conf.CacheTTL, Conf.CacheTTL*2)
	}

	if Conf.LogPath != "" {
		f, err := os.OpenFile(Conf.LogPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open logfile err: %v\n", err)
			return nil
		}
		defaultLoger = log.New(f, "nopre", log.LstdFlags)
	}

	return nil
}
