package internal

import (
	"log"
)

var defaultLoger *log.Logger

func logVisitor(msg string) {
	if defaultLoger != nil {
		defaultLoger.SetPrefix("visitor ")
		defaultLoger.Printf(msg)
	}
}
