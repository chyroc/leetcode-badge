package main

import (
	"flag"
	"fmt"

	"github.com/Chyroc/leetcode-badge/internal"
)

func init() {
	flag.IntVar(&internal.Conf.Port, "port", 9090, "app port")
	flag.BoolVar(&internal.Conf.Cache, "cache", false, "use cache")
	flag.BoolVar(&internal.Conf.Release, "release", false, "release or debug")
	flag.Parse()
}

func main() {
	app := internal.APP()
	app.Run(fmt.Sprintf(":%d", internal.Conf.Port))
}
