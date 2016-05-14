package main

type CollisionEffector struct {
	Damping float64
}

func NewCollisionEffector(damping float64) *CollisionEffector {
	return &CollisionEffector{
		Damping: damping,
	}
}

func (this *CollisionEffector)Update(particles []*Particle, dt float64) {
	for i, p := range particles {
		for _, other := range append(particles[:i], particles[i + 1:]...) {
			thisV := p.Pos.Sub(p.PrevPos) // prev -> pos
			otherV := other.Pos.Sub(other.PrevPos)
			toOther := other.Pos.Sub(p.Pos) // p -> other
			if toOther.Len() < p.Radius + other.Radius {
				// collide!
				p.Pos = p.Pos.Sub(toOther.Mul(0.5))
				other.Pos = other.Pos.Add(toOther.Mul(0.5))
				sqDiff := toOther.Len() * toOther.Len()
				pF := p.Pos.Dot(thisV) * this.Damping / sqDiff
				otherF := other.Pos.Dot(otherV) * this.Damping / sqDiff
				thisV = thisV.Add(toOther.Mul(otherF - pF))
				otherV = thisV.Add(toOther.Mul(pF - otherF))
				p.PrevPos = p.Pos.Sub(thisV)
				other.PrevPos = other.PrevPos.Sub(otherV)
			}
		}
	}
}
