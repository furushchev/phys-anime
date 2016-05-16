package main

import (
	"math/rand"
	m "github.com/go-gl/mathgl/mgl64"
)

func SetupTask3(w *Window) {

	lineStart := m.Vec2{ 200, 200 }
	lineEnd := m.Vec2{ 600, 500 }
	v := lineEnd.Sub(lineStart)

	// initial randomize
	for i := 0; i < NumParticles; i++ {
		c := RandomColor(i)
		pos := lineStart.Add(v.Mul(rand.Float64()))
		p := NewParticle(pos.X(), pos.Y(), 10, c)
		w.AddParticle(p)
	}

	w.AddEffector(NewBeadEffector(lineStart, lineEnd, func(time, mass float64) m.Vec2 {
		f := m.Vec2{ 0.0, -9.8 }
		return f.Mul(1.0 / mass)
	}))

}

