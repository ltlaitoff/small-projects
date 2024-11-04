package main

import (
	"bytes"
	"math/rand"
	"strings"
	"image"
	"image/png"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/fogleman/gg"
)

func getImage(width int, height int, matrix string) image.Image {
	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.SetRGB(1, 1, 1)
	dc.Fill()

	rows := strings.Split(matrix, ",")
	rowHeigth := height / len(rows)

	for rowIndex, element := range rows {
		cols := strings.Split(element, "")
		colsWidth := width / len(cols)

		for colIndex := range cols {
			dc.DrawRectangle(
				float64(colsWidth) * float64(colIndex),
				float64(rowHeigth) * float64(rowIndex),
				float64(colsWidth)*float64(colIndex)+float64(colsWidth),
				float64(rowHeigth)*float64(rowIndex)+float64(rowHeigth),
			)

			dc.SetRGB(rand.Float64(), rand.Float64(), rand.Float64())
			dc.Fill()
		}
	}

	return dc.Image()
}

func getImageRoute(c *gin.Context) {
	width, err := strconv.Atoi(c.DefaultQuery("width", "480"))

	if err != nil {
		log.Fatal(err)
	}

	height, err := strconv.Atoi(c.DefaultQuery("height", "480"))

	if err != nil {
		log.Fatal(err)
	}

	matrix := c.DefaultQuery("matrix", "01,10")

	image := getImage(width, height, matrix)
	if image == nil {
		log.Fatal("Image is nil")
	}

	buf := new(bytes.Buffer)

	if err := png.Encode(buf, image); err != nil {
		log.Fatal(err)
	}

	c.Data(http.StatusOK, "image/png", buf.Bytes())
}

func main() {
	router := gin.Default()

	router.GET("/image", getImageRoute)

	router.Run(":8080")
}
