package main

import (
	"math/rand"
//	m "github.com/go-gl/mathgl/mgl64"
)

func SetupTask4(w *Window) {
	for i := 0; i < 100; i++ {
		c := RandomColor(i)
		p := NewParticle(rand.Float64() * float64(w.Width), rand.Float64() * float64(w.Height), 10, c)
		w.AddParticle(p)
	}
	
	w.AddEffector(NewGravitationalFieldEffector(6.67408e-11, 1))
	w.AddEffector(NewBoundaryEffector(float64(w.Width), float64(w.Height), true))
	w.AddEffector(NewCollisionEffector(0.99, float64(w.Width), float64(w.Height), 8, 4 ))
}

