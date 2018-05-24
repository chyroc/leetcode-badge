package internal

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func APP() *gin.Engine {
	app := gin.Default()
	app.Use(cors.Default())

	app.GET("", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(400, "name is empty")
			return
		}

		leetcodeData, err := FetchLeetcodeData(name)
		if err != nil {
			c.JSON(400, err)
			return
		} else if len(leetcodeData) == 0 {
			c.JSON(400, fmt.Errorf("fetch leetcode data of account: %s, result: nil", name))
			return
		}

		shieldsData, err := FetchShieldsData(c, leetcodeData)
		if err != nil {
			c.JSON(400, err)
			return
		}

		c.Writer.WriteHeader(200)
		c.Writer.Header().Add("Content-Type", "image/svg+xml; charset=utf-8")
		c.Writer.WriteString(shieldsData)

		return
	})

	return app
}
