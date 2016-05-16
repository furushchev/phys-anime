package main

import (
	"math/rand"
	m "github.com/go-gl/mathgl/mgl64"
)

func SetupTask3(w *Window) {

	lineStart := m.Vec2{ Task3LineStartPos[0], Task3LineStartPos[1] }
	lineEnd := m.Vec2{ Task3LineEndPos[0], Task3LineEndPos[1] }
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

