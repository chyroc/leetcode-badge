package internal

import (
	"log"
	"os"
	"sync"
)

var defaultLoger *log.Logger
var defaultLogerOnce = new(sync.Once)

func logVisitor(msg string) {
	defaultLogerOnce.Do(func() {
		if Conf.LogPath != "" {
			f, err := os.OpenFile(Conf.LogPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
			if err != nil {
				return
			}
			defaultLoger = log.New(f, "nopre", log.LstdFlags)
		}
	})

	if defaultLoger != nil {
		defaultLoger.SetPrefix("visitor ")
		defaultLoger.Printf(msg)
	}
}
