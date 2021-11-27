package bmp_test

import (
	"image/color"
	"image/draw"
	"os"

	"github.com/sergeymakinen/go-bmp"
)

func Example() {
	f, _ := os.Open("file.bmp")
	img, _ := bmp.Decode(f)
	for i := 10; i < 20; i++ {
		img.(draw.Image).Set(i, i, color.NRGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})
	}
	f.Truncate(0)
	bmp.Encode(f, img)
	f.Close()
}
