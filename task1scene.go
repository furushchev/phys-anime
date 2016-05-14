package main

import (
	"math/rand"
	"github.com/go-gl/mathgl/mgl64"
)

type Task1 struct {
	balls []*Ball
}

func NewTask1(w *Window) *Task1 {
	var balls []*Ball
	for i := 0; i < 100; i++ {
		c := RandomColor(i)
		b := NewBall(rand.Float64() * float64(w.Width), rand.Float64() * float64(w.Height), 10, c)

		b.RegisterForceFunc(func(time, mass float64) mgl64.Vec2 {
			fx := 0.0
			if time < 1.0 {
				fx = (rand.Float64() - 0.5) * 2000.0
			}
			return mgl64.Vec2{fx, -9.8 / mass }
		})
		w.AddObject(b)
		balls = append(balls, b)
	}
	return &Task1 {
		balls: balls,
	}
}
