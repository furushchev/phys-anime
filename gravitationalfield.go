package main

import (
	"math"
	m "github.com/go-gl/mathgl/mgl64"
	"sync"
)

type GravitationalFieldEffector struct {
	G float64
	epsilon float64
}

func NewGravitationalFieldEffector(G, epsilon float64) *GravitationalFieldEffector {
	return &GravitationalFieldEffector{
		G: G,
		epsilon: epsilon,
	}
}

func (this *GravitationalFieldEffector)updateParticle(p *Particle, others []*Particle, dt float64) {
	F := m.Vec2{ 0, 0 }
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	for _, otherParticle := range others {
		wg.Add(1)
		go func(op *Particle) {
			defer wg.Done()
			k := this.G * p.Mass * op.Mass / math.Pow(this.epsilon + op.Pos.Sub(p.Pos).Len(), 3)
			f := op.Pos.Sub(p.Pos).Mul(k)
			mu.Lock()
			F = F.Add(f)
			mu.Unlock()
		}(otherParticle)
	}
	wg.Wait()
	println(F.X(), F.Y())
	prev2Pos, prevPos := p.PrevPos, p.Pos
	p.PrevPos = prevPos
	p.Pos = prevPos.Mul(2.0).Sub(prev2Pos).Add(F)
}


func (this *GravitationalFieldEffector)Update(particles []*Particle, dt float64) {
	ApplyEffectToParticleWithOthers(particles, dt, this.updateParticle)
}