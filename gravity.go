package main

import (
	m "github.com/go-gl/mathgl/mgl64"
)

type ForceFunc func(float64, float64) m.Vec2

type GravityEffector struct {
	Force ForceFunc
	ElapsedTime float64
}

func NewGravityEffector(force ForceFunc) *GravityEffector {
	return &GravityEffector{
		Force: force,
		ElapsedTime: 0.0,
	}
}

func (this *GravityEffector)updateParticle(p *Particle, dt float64) {
	prev2Pos, prevPos := p.PrevPos, p.Pos

	fVec := this.Force(this.ElapsedTime, p.Mass).Mul(dt * dt)
	p.PrevPos = prevPos
	p.Pos = prevPos.Mul(2.0).Sub(prev2Pos).Add(fVec)
}

func (this *GravityEffector)Update(particles []*Particle, dt float64) {
	this.ElapsedTime += dt
	ApplyEffectToParticle(particles, dt , this.updateParticle)
}