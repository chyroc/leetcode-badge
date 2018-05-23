package main

import (
	"fmt"
	"flag"

	"github.com/gin-gonic/gin"

	"github.com/Chyroc/leetcode-badge/internal"
)

var port int

func init() {
	flag.IntVar(&port, "port", 9090, "app port")
	flag.Parse()
}

func main() {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(400, map[string]interface{}{"message": "name is empty"})
			return
		}

		data, err := internal.FetchLeetcodeData(name)
		if err != nil {
			c.JSON(400, map[string]interface{}{"message": err})
			return
		} else if data == nil || len(data) == 0 {
			c.JSON(400, map[string]interface{}{"message": fmt.Sprintf("fetch leetcode data of account: %s, result: nil", name)})
			return
		}

		c.JSON(200, data)
		return
	})

	app.Run(fmt.Sprintf(":%d", port))
}
