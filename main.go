package main

import (
	"fmt"
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/Chyroc/leetcode-badge/internal"
)

var port int

func init() {
	flag.IntVar(&port, "port", 9090, "app port")
	flag.Parse()
}

func handlerFetchLeetcode(c *gin.Context) (data map[string]int, err error) {
	name := c.Query("name")
	if name == "" {
		c.JSON(400, map[string]interface{}{"message": "name is empty"})
		return
	}

	data, err = internal.FetchLeetcodeData(name)
	if err != nil {
		c.JSON(400, map[string]interface{}{"message": err})
		return
	} else if data == nil || len(data) == 0 {
		c.JSON(400, map[string]interface{}{"message": fmt.Sprintf("fetch leetcode data of account: %s, result: nil", name)})
		return
	}

	return
}

func main() {
	app := gin.Default()
	app.Use(cors.Default())

	app.GET("/data", func(c *gin.Context) {
		leetcodeData, err := handlerFetchLeetcode(c)
		if err != nil {
			return
		}

		c.JSON(200, leetcodeData)
		return
	})

	app.GET("", func(c *gin.Context) {
		leetcodeData, err := handlerFetchLeetcode(c)
		if err != nil {
			return
		}

		shieldsData, err := internal.FetchShieldsData(c, leetcodeData)
		fmt.Printf("err %v\n",err)
		if err != nil {
			c.JSON(400, map[string]interface{}{"message": err})
			return
		}

		c.Writer.WriteHeader(200)
		c.Writer.Header().Add("Content-Type", "image/svg+xml")
		c.Writer.Header().Add("Content-Type", "charset=utf-8")
		c.Writer.WriteString(shieldsData)

		return
	})

	app.Run(fmt.Sprintf(":%d", port))
}
