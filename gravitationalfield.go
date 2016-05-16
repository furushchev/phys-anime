package main

import (
	"math"
	m "github.com/go-gl/mathgl/mgl64"
)

type GravitationalFieldEffector struct {
	G float64
	epsilon float64
	time float64
}

func NewGravitationalFieldEffector(G, epsilon float64) *GravitationalFieldEffector {
	return &GravitationalFieldEffector{
		G: G,
		epsilon: epsilon,
	}
}
/*
func (this *GravitationalFieldEffector)updateParticle(p *Particle, others []*m.Vec2, dt float64) {
	F := m.Vec2{ 0, 0 }
	for _, op := range others {
		k := this.G * p.Mass * op.Mass / math.Pow(this.epsilon + op.Pos.Sub(p.Pos).Len(), 3)
		f := op.Pos.Sub(p.Pos).Mul(k)
		F = F.Add(f)
	}

	if math.IsNaN(F.Len()) {
		F = m.Vec2{0.0, 0.0}
	}

	// gravity
	fx := 0.0
	if this.time < 1.0 {
		fx = (rand.Float64() - 0.5) * 2000.0
	}
	F = F.Add(m.Vec2{fx, -9.8})

	prev2Pos, prevPos := p.PrevPos, p.Pos
	p.PrevPos = m.Vec2{ p.Pos.X(), p.Pos.Y() }
	vf := F.Mul(dt * dt / p.Mass)
	p.Pos = prevPos.Mul(2.0).Sub(prev2Pos)//.Add(vf)
	println(p.PrevPos.X(), p.PrevPos.Y(), p.Pos.X(), p.Pos.Y(), vf.X(), vf.Y())

}
*/

func (this *GravitationalFieldEffector)Update(particles []*Particle, dt float64) {
	this.time += dt
	for i, p := range particles {
		F := m.Vec2{ 0, 0 }
		for j, op := range particles {
			if i == j {
				continue
			}
			k := this.G * p.Mass * op.Mass / math.Pow(this.epsilon + op.Pos.Sub(p.Pos).Len(), 3)
			f := op.Pos.Sub(p.Pos).Mul(k)
			F = F.Add(f)
		}
		prev2Pos, prevPos := p.PrevPos, p.Pos
		p.PrevPos = p.Pos
		F = F.Add(m.Vec2{ 0.0, -9.8 }) // gravity
		vf := F.Mul(dt * dt / p.Mass)
		p.Pos = prevPos.Mul(2.0).Sub(prev2Pos).Add(vf)
	}
}