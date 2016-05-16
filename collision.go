package main

type CollisionEffector struct {
	damping float64
}

func NewCollisionEffector(damping, windowWidth, windowHeight float64, gridXNum, gridYNum int) *CollisionEffector {
	return &CollisionEffector{
		damping: damping,
	}
}

func (this *CollisionEffector)Update(particles []*Particle, dt float64) {
	for i, p1 := range particles {
		for _, p2 := range particles[i+1:] {
			v1 := p1.Pos.Sub(p1.PrevPos)
			v2 := p2.Pos.Sub(p2.PrevPos)
			p1p2 := p2.Pos.Sub(p1.Pos)
			if p1p2.Len() < p1.Radius + p2.Radius {
				// collide
				p1.Pos = p1.Pos.Sub(p1p2.Mul(0.5))
				p2.Pos = p2.Pos.Add(p1p2.Mul(0.5))
				f1 := p1p2.Dot(v1) * this.damping / (p1p2.Len() * p1p2.Len())
				f2 := p1p2.Dot(v2) * this.damping / (p1p2.Len() * p1p2.Len())
				v1 = v1.Add(p1p2.Mul(f2 - f1))
				v2 = v2.Add(p1p2.Mul(f1 - f2))
				p1.PrevPos = p1.Pos.Sub(v1)
				p2.PrevPos = p2.Pos.Sub(v2)
			}
		}
	}
}
