package main

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
)

func PrepareImageForOCR(img image.Image) image.Image {
	img = imaging.Invert(img)
	img = imaging.Grayscale(img)
	img = imaging.Sharpen(img, 10)
	img = imaging.AdjustFunc(img, func(c color.NRGBA) color.NRGBA {
		if c.R < 60 && c.G < 60 && c.B < 60 {
			return c
		} else {
			return color.NRGBA{255, 255, 255, 255}
		}
	})
	return img
}
