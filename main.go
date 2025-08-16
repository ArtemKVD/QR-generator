package main

import (
	"image"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	gozxing "github.com/makiuchi-d/gozxing"
	decoder "github.com/makiuchi-d/gozxing/qrcode"
	"github.com/skip2/go-qrcode"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/qr", generateQR)
	r.POST("/decode", decodeQR)

	r.Run(":8080")
}

func generateQR(c *gin.Context) {
	text := c.Query("text")
	if text == "" {
		log.Printf("error get text")
		return
	}

	png, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		log.Printf("error encode")
		return
	}

	c.Data(200, "image/png", png)
}

func decodeQR(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("error get file")
		return
	}

	qr, err := file.Open()
	if err != nil {
		log.Printf("error open file")
		return
	}
	defer qr.Close()

	img, _, err := image.Decode(qr)
	if err != nil {
		log.Printf("decode error")
		return
	}

	image, _ := gozxing.NewBinaryBitmapFromImage(img)
	qrReader := decoder.NewQRCodeReader()
	result, err := qrReader.Decode(image, nil)
	if err != nil {
		log.Printf("error decode qr")
		return
	}

	c.JSON(200, gin.H{"text": result.GetText()})
}
