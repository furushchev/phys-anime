package main

import (
	"math/rand"
	m "github.com/go-gl/mathgl/mgl64"
)

func SetupTask2(w *Window) {
	for i := 0; i < 100; i++ {
		c := RandomColor(i)
		p := NewParticle(rand.Float64() * float64(w.Width), rand.Float64() * float64(w.Height), 10, c)
		w.AddParticle(p)
	}
	w.AddEffector(NewGravityEffector(func(time, mass float64) m.Vec2 {
		fx := 0.0
		if time < 1.0 {
			fx = (rand.Float64() - 0.5) * 2000.0
		}
		return m.Vec2{fx, -9.8 / mass }
	}))
	w.AddEffector(NewBoundaryEffector(float64(w.Width), float64(w.Height), false))
}
