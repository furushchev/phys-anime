package main

import (
	m "github.com/go-gl/mathgl/mgl64"
	"github.com/go-gl/gl/v2.1/gl"
	"math"
)

type ForceFunc func(float64, float64) m.Vec2

type Ball struct {
	radius float64
	slices int
	mass float64
	pos m.Vec2
	prevPos m.Vec2
	color m.Vec4
	time float64
	forceFunc ForceFunc
}

func NewBall(x,y float64, r float64, c *Color) *Ball {
	return &Ball{
		radius: r,
		slices: 60,
		pos: m.Vec2{x,y},
		prevPos: m.Vec2{x,y},
		mass: 0.05,
		color: m.Vec4{c.R(), c.G(), c.B(), c.A()},
		forceFunc: func(_, _ float64) m.Vec2 {
			return m.Vec2{0.0,0.0}
		},
	}
}

func (this *Ball)RegisterForceFunc(f func(time float64, mass float64) m.Vec2) {
	this.forceFunc = f
}

func (this *Ball)Update(dt, screenWidth, screenHeight float64) {
	this.time +=  dt

	pos_n := m.Vec2{ this.pos.X(), this.pos.Y() }
	pos_n_1 := m.Vec2{ this.prevPos.X(), this.prevPos.Y() }
	this.prevPos = pos_n
	fvec := this.forceFunc(this.time, this.mass).Mul(dt * dt)
	this.pos = pos_n.Mul(2.0).Sub(pos_n_1).Add(fvec)
	this.checkBoundary(screenWidth, screenHeight)
}

func (this *Ball)checkBoundary(w, h float64) {
	// bound
	if this.pos.Y() - this.radius < 0 {
		this.pos = m.Vec2{ this.pos.X(), 2.0 * this.radius - this.pos.Y() }
		this.prevPos = m.Vec2{ this.prevPos.X(), 2.0 * this.radius -this.prevPos.Y() }
	}

	// left wall
	if this.pos.X() - this.radius < 0 {
		this.pos = m.Vec2{ 2.0 * this.radius - this.pos.X(), this.pos.Y() }
		this.prevPos = m.Vec2{ 2.0 * this.radius - this.prevPos.X(), this.prevPos.Y() }
	}

	// right wall
	if this.pos.X() + this.radius > w {
		this.pos = m.Vec2{ 2.0 * (w - this.radius) - this.pos.X(), this.pos.Y() }
		this.prevPos = m.Vec2{ 2.0 * (w - this.radius) - this.prevPos.X(), this.prevPos.Y() }
	}
}

func (this *Ball)CheckCollision(others interface{}) {
	for _, other := range []*Ball(others) {

	}

}

func (this *Ball)Draw() {
	gl.Color4d(this.color.X(), this.color.Y(), this.color.Z(), this.color.W())
	gl.PushMatrix()
	gl.Translated(this.pos.X(), this.pos.Y(), 0.0)
	gl.Begin(gl.POLYGON)
	for a := 0.0; a < 2 * math.Pi; a += (2 * math.Pi / float64(this.slices)) {
		gl.Vertex2d(math.Sin(a) * this.radius, math.Cos(a) * this.radius)
	}
	gl.End()
	gl.PopMatrix()
}
