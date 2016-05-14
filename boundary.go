package main

import (
	m "github.com/go-gl/mathgl/mgl64"
)

type BoundaryEffector struct {
	Width float64
	Height float64
	Ceiling bool
}

func NewBoundaryEffector(width, height float64, ceiling bool) *BoundaryEffector {
	return &BoundaryEffector{
		Width: width,
		Height: height,
		Ceiling: ceiling,
	}
}

func (this *BoundaryEffector)updateParticle(p *Particle, dt float64) {
	// floor
	if p.Pos.Y() - p.Radius < 0 {
		p.Pos = m.Vec2{p.Pos.X(), 2.0 * p.Radius - p.Pos.Y() }
		p.PrevPos = m.Vec2{p.PrevPos.X(), 2.0 * p.Radius - p.PrevPos.Y() }
	}

	// left wall
	if p.Pos.X() - p.Radius < 0 {
		p.Pos = m.Vec2{2.0 * p.Radius - p.Pos.X(), p.Pos.Y() }
		p.PrevPos = m.Vec2{2.0 * p.Radius - p.PrevPos.X(), p.PrevPos.Y() }
	}

	// right wall
	if p.Pos.X() + p.Radius > this.Width {
		p.Pos = m.Vec2{2.0 * (this.Width - p.Radius) - p.Pos.X(), p.Pos.Y() }
		p.PrevPos = m.Vec2{2.0 * (this.Width - p.Radius) - p.PrevPos.X(), p.PrevPos.Y() }
	}

	// ceiling
	if this.Ceiling {
		if p.Pos.Y() + p.Radius > this.Height {
			p.Pos = m.Vec2{p.Pos.X(), 2.0 * (this.Height - p.Radius) - p.Pos.Y()}
			p.PrevPos = m.Vec2{p.PrevPos.X(), 2.0 * (this.Height - p.Radius) - p.PrevPos.Y()}
		}
	}
}

func (this *BoundaryEffector)Update(particles []*Particle, dt float64) {
	ApplyEffectToParticle(particles, dt, this.updateParticle)
}