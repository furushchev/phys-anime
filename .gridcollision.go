package main

type GridCollisionEffector struct {
	damping float64
	width float64
	height float64
}

func NewGridCollisionEffector(damping float64, width, height, particleNum int) *GravityEffector {
	return &GridCollisionEffector{
		damping: damping,
		width: width,
		height: height,
	}
}

func (this *GridCollisionEffector)Update(particles []*Particle) {
	tree := NewQuadTree(2, this.width, this.height)
	tree.SetParticles(particles)
	tree.
}

