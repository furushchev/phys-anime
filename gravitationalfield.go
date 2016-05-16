package main

import (
	"math"
	m "github.com/go-gl/mathgl/mgl64"
	"sync"
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

func (this *GravitationalFieldEffector)Update(particles []*Particle, dt float64) {
	this.time += dt
	wg := sync.WaitGroup{}
	for i, p := range particles {
		wg.Add(1)
		go func(p *Particle, idx int, ops []*Particle) {
			defer wg.Done()
			F := m.Vec2{0, 0 }
			for j, op := range ops {
				if i == j {
					continue
				}
				k := this.G * p.Mass * op.Mass / math.Pow(this.epsilon + op.Pos.Sub(p.Pos).Len(), 3)
				f := op.Pos.Sub(p.Pos).Mul(k)
				F = F.Add(f)
			}
			prev2Pos, prevPos := p.PrevPos, p.Pos
			p.PrevPos = p.Pos
			F = F.Add(m.Vec2{0.0, -9.8 }) // gravity
			vf := F.Mul(dt * dt / p.Mass)
			p.Pos = prevPos.Mul(2.0).Sub(prev2Pos).Add(vf)
		}(p, i, particles)
	}
	wg.Wait()
}