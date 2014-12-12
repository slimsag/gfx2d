// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx2d

import (
	"image"
	"time"

	"azul3d.org/gfx.v2-dev"
	"azul3d.org/gfx.v2-dev/window"

	"code.google.com/p/draw2d/draw2d"
)

// Window represents a window and a draw2d graphic context.
type Window struct {
	window.Window
	draw2d.GraphicContext

	// Redraw is sent a single event whenever the window is resized or damaged.
	//
	// Your application should handle this event by redrawing the window and
	// calling Render.
	Redraw chan window.Event

	app    func(w *Window)
	render chan struct{}
}

// Render renders the image drawn to the graphic context to this window.
func (w *Window) Render() {
	// Signal to the graphics loop to render a frame.
	w.render <- struct{}{}

	// Wait for the render to complete.
	<-w.render
}

func (w *Window) gfxLoop(gfxWin window.Window, d gfx.Device) {
	w.Window = gfxWin

	// Have the window notify our redraw channel whenever a framebuffer resize
	// event or a window damage event occurs.
	gfxWin.Notify(w.Redraw, window.FramebufferResizedEvents|window.DamagedEvents)

	// Create the draw2d image graphics context.
	img := image.NewRGBA(d.Bounds())
	w.GraphicContext = draw2d.NewGraphicContext(img)
	card := NewCard(img)

	// Spawn the app.
	go w.app(w)

	for {
		// Wait for the app to ask us to render.
		<-w.render

		// Re-upload the card's texture to the graphics device.
		card.Textures[0].Loaded = false

		// Clear the color buffer.
		d.Clear(d.Bounds(), gfx.Color{1, 1, 1, 1})

		// Draw the textured card.
		d.Draw(d.Bounds(), card, nil)

		// Ask the device to render the frame.
		d.Render()

		// If the device's bounds has changed (i.e. the window has been
		// resized) then we must recreate the image with the new bounds and
		// recreate the graphic context.
		if img.Bounds() != d.Bounds() {
			img = image.NewRGBA(d.Bounds())
			w.GraphicContext = draw2d.NewGraphicContext(img)
			card = NewCard(img)

			// Need one more redraw.
			ev := &window.FramebufferResized{
				T:      time.Now(),
				Width:  d.Bounds().Dx(),
				Height: d.Bounds().Dy(),
			}
			select {
			case w.Redraw <- ev:
			default:
			}
		}

		// Signal that the render is complete.
		w.render <- struct{}{}
	}
}

// Props is the set of properties that the window will be opened with.
var Props = window.NewProps()

func init() {
	Props.SetTitle("")
}

// Run opens a window and runs the given application.
func Run(app func(w *Window)) {
	w := &Window{
		Redraw: make(chan window.Event, 128),
		app:    app,
		render: make(chan struct{}),
	}
	window.Run(w.gfxLoop, Props)
}
