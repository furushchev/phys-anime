package main

import (
	m "github.com/go-gl/mathgl/mgl64"
	"github.com/go-gl/gl/v2.1/gl"
	"math"
)

type Particle struct {
	Radius    float64
	slices    int
	Mass      float64
	Pos       m.Vec2
	PrevPos   m.Vec2
	color     m.Vec4
}

func NewParticle(x,y float64, r float64, c *Color) *Particle {
	return &Particle{
		Radius: r,
		slices: 60,
		Pos: m.Vec2{x,y},
		PrevPos: m.Vec2{x,y},
		Mass: 0.05,
		color: m.Vec4{c.R(), c.G(), c.B(), c.A()},
	}
}

func (this *Particle)Draw() {
	gl.Color4d(this.color.X(), this.color.Y(), this.color.Z(), this.color.W())
	gl.PushMatrix()
	gl.Translated(this.Pos.X(), this.Pos.Y(), 0.0)
	gl.Begin(gl.POLYGON)
	for a := 0.0; a < 2 * math.Pi; a += (2 * math.Pi / float64(this.slices)) {
		gl.Vertex2d(math.Sin(a) * this.Radius, math.Cos(a) * this.Radius)
	}
	gl.End()
	gl.PopMatrix()
}
