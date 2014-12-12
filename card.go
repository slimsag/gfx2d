// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx2d

import (
	"image"

	"azul3d.org/gfx.v2-dev"
)

// NewCard creates and returns a new card object for the given image.
func NewCard(img image.Image) *gfx.Object {
	// Create the card's texture.
	tex := gfx.NewTexture()
	tex.Source = img
	tex.MinFilter = gfx.LinearMipmapLinear
	tex.MagFilter = gfx.Linear
	tex.Format = gfx.DXT1RGBA
	tex.KeepDataOnLoad = true

	// Create the card's mesh.
	cardMesh := gfx.NewMesh()
	cardMesh.Vertices = []gfx.Vec3{
		// Bottom-left triangle.
		{-1, 0, -1},
		{1, 0, -1},
		{-1, 0, 1},

		// Top-right triangle.
		{-1, 0, 1},
		{1, 0, -1},
		{1, 0, 1},
	}
	cardMesh.TexCoords = []gfx.TexCoordSet{
		{
			Slice: []gfx.TexCoord{
				{0, 1},
				{1, 1},
				{0, 0},

				{0, 0},
				{1, 1},
				{1, 0},
			},
		},
	}

	// Create the card object.
	card := gfx.NewObject()
	card.State.DepthTest = false
	card.State.DepthWrite = false
	card.AlphaMode = gfx.AlphaToCoverage
	card.Shader = GLSLShader
	card.Textures = []*gfx.Texture{tex}
	card.Meshes = []*gfx.Mesh{cardMesh}
	return card
}
