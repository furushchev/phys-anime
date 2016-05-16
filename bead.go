package main

import m "github.com/go-gl/mathgl/mgl64"

type BeadEffector struct {
	LineStart m.Vec2
	LineEnd m.Vec2
	Force ForceFunc
	ElapsedTime float64
}

func NewBeadEffector(lineStart, lineEnd m.Vec2, force ForceFunc) *BeadEffector {
	return &BeadEffector{
		LineStart: lineStart,
		LineEnd: lineEnd,
		Force: force,
		ElapsedTime: 0.0,
	}
}

func (this *BeadEffector)updateParticle(p *Particle, dt float64) {
	f := this.Force(this.ElapsedTime, p.Mass)
	lvec := this.LineEnd.Sub(this.LineStart)
	fline := lvec.Normalize().Mul(lvec.Dot(f) / lvec.Len())
	prev2Pos, prevPos := p.PrevPos, p.Pos
	p.PrevPos = prevPos
	p.Pos = prevPos.Mul(2.0).Sub(prev2Pos).Add(fline.Mul(dt * dt))

	if p.Pos.Sub(this.LineStart).Dot(lvec) < 0 {
		d := p.Pos.Sub(this.LineStart).Mul(-1.0)
		p.Pos = this.LineStart.Add(d)
		dp := p.PrevPos.Sub(this.LineStart).Mul(-1.0)
		p.PrevPos = this.LineStart.Add(dp)
	} else if p.Pos.Sub(this.LineEnd).Dot(lvec) > 0 {
		d := p.Pos.Sub(this.LineEnd).Mul(-1.0)
		p.Pos = this.LineEnd.Add(d)
		dp := p.PrevPos.Sub(this.LineEnd).Mul(-1.0)
		p.PrevPos = this.LineEnd.Add(dp)
	}
}

func (this *BeadEffector)Update(particles []*Particle, dt float64) {
	this.ElapsedTime += dt
	ApplyEffectToParticle(particles, dt, this.updateParticle)
}