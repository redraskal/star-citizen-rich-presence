package main

import (
	"image"

	"github.com/JamesHovious/w32"
	"github.com/kbinani/screenshot"
)

func CaptureWindow(hwnd w32.HWND) (*image.RGBA, error) {
	dim := w32.GetWindowRect(hwnd)
	println("DIM:", dim.Left, dim.Top, dim.Right, dim.Bottom)
	return screenshot.CaptureRect(image.Rectangle{
		image.Point{
			int(dim.Left),
			int(dim.Top),
		},
		image.Point{
			int(dim.Right),
			int(dim.Bottom),
		},
	})
}
