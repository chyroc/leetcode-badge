package main

import (
	"flag"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

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
		err = fmt.Errorf("name is empty")
		c.JSON(400, map[string]interface{}{"message": err})
		return nil, err
	}

	data, err = internal.FetchLeetcodeData(name)
	if err != nil {
		c.JSON(400, map[string]interface{}{"message": err})
		return nil, err
	} else if data == nil || len(data) == 0 {
		err = fmt.Errorf("fetch leetcode data of account: %s, result: nil", name)
		c.JSON(400, map[string]interface{}{"message": err})
		return nil, err
	}

	return data, nil
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
		if err != nil {
			c.JSON(400, map[string]interface{}{"message": err})
			return
		}

		c.Writer.WriteHeader(200)
		c.Writer.Header().Add("Content-Type", "image/svg+xml; charset=utf-8")
		c.Writer.WriteString(shieldsData)

		return
	})

	app.Run(fmt.Sprintf(":%d", port))
}
