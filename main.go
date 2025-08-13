package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/qr", func(c *gin.Context) {
		text := c.Query("text")

		code, err := qrcode.Encode(text, qrcode.Medium, 256)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.Data(200, "image/png", code)
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.Run(":8080")
}
