package main

import (
	m "github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v2.1/gl"
	"math"
)

type Ball struct {
	radius float64
	slices int
	pos m.Vec2
	color m.Vec4
}

func NewBall(x,y float32, r float64) *Ball {
	return &Ball{
		radius: r,
		slices: 60,
		pos: m.Vec2{x,y},
		color: m.Vec4{1.0, 0.0, 0.0, 1.0},

	}
}

func (this *Ball)Update(dt float32) {

}

func (this *Ball)Draw() {
	gl.Color4f(this.color.X(), this.color.Y(), this.color.Z(), this.color.W())
	gl.PushMatrix()
	gl.Translatef(this.pos.X(), this.pos.Y(), 0.0)
	gl.Begin(gl.LINE_LOOP)
	for a := 0.0; a < 2 * math.Pi; a += (2 * math.Pi / float64(this.slices)) {
		gl.Vertex2d(math.Sin(a) * this.radius, math.Cos(a) * this.radius)
	}
	gl.Vertex3f(0, 0, 0)
	gl.End()
	gl.PopMatrix()
}
