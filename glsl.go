// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx2d

import "azul3d.org/gfx.v2-dev"

var GLSLVert = []byte(`
#version 120

attribute vec3 Vertex;
attribute vec2 TexCoord0;

uniform mat4 MVP;

varying vec2 tc0;

void main()
{
	tc0 = TexCoord0;
	gl_Position = MVP * vec4(Vertex, 1.0);
}
`)

var GLSLFrag = []byte(`
#version 120

varying vec2 tc0;

uniform sampler2D Texture0;
uniform bool BinaryAlpha;

void main()
{
	gl_FragColor = texture2D(Texture0, tc0);
	if(BinaryAlpha && gl_FragColor.a < 0.5) {
		discard;
	}
}
`)

// GLSLShader is a simple textured card GLSL shader.
var GLSLShader = &gfx.Shader{
	Name: "draw2d.GLSLShader",
	GLSL: &gfx.GLSLShader{
		Vertex:   GLSLVert,
		Fragment: GLSLFrag,
	},
}
