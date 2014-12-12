// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Demonstrates basic draw2d usage.
package main

import (
	"image"
	"image/color"

	"code.google.com/p/draw2d/draw2d"
	"github.com/slimsag/gfx2d"
)

func app(win *gfx2d.Window) {
	draw2d.SetFontFolder("font/")

	// Render loop.
	for {
		// Clear the window to black.
		win.SetFillColor(image.Black)
		win.Clear()

		// Set font data and size.
		win.SetFontData(draw2d.FontData{"luxi", draw2d.FontFamilyMono, draw2d.FontStyleBold | draw2d.FontStyleItalic})
		win.SetFontSize(18)

		// Draw the green text.
		win.SetFillColor(color.RGBA{0, 255, 0, 255})
		win.FillStringAt("Hello World!", 50, 50)

		// Render the frame.
		win.Render()
	}
}

func main() {
	gfx2d.Props.SetTitle("2D Graphics Demo")
	gfx2d.Run(app)
}
