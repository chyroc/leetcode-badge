package main

import (
	"fmt"

	"github.com/Chyroc/leetcode-badge/internal"
)

func main() {
	data, err := internal.FetchLeetcodeData("todo")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", data)
}
