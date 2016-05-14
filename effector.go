package main

import "sync"

type Effector interface {
	Update([]*Particle, float64)
}

func ApplyEffectToParticle(particles []*Particle, dt float64, f func(*Particle, float64)) {
	wg := sync.WaitGroup{}
	for _, particle := range particles {
		wg.Add(1)
		go func (p *Particle) {
			defer wg.Done()
			f(p, dt)
		}(particle)
	}
	wg.Wait()

}