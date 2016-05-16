package main

import (
	"math/rand"
	m "github.com/go-gl/mathgl/mgl64"
)

func SetupTask2(w *Window) {
	// initial randomize
	for i := 0; i < NumParticles; i++ {
		c := RandomColor(i)
		p := NewParticle(rand.Float64() * float64(w.Width), rand.Float64() * float64(w.Height), 10, c)
		w.AddParticle(p)
	}

	w.AddEffector(NewGravityEffector(func(time, mass float64) m.Vec2 {
		f := m.Vec2{0.0, -9.8}
		return f.Mul(1.0 / mass)
	}))

	w.AddEffector(NewBoundaryEffector(float64(w.Width), float64(w.Height), false))
}
