package main

import m "github.com/go-gl/mathgl/mgl64"

type Morton uint32
type CellId struct {
	Level uint32
	Index uint32
	Morton Morton
}

func Pow(a, b uint32) uint32 {
	p := uint32(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

type TreeCell struct {
	id uint32
	elements []*Particle
	damping float64
}

func (this *TreeCell)SearchChildren() {

}

func SolveCollision(p1, p2 *Particle, damping float64) {
	v1 := p1.Pos.Sub(p1.PrevPos)
	v2 := p2.Pos.Sub(p2.PrevPos)
	p1p2 := p2.Pos.Sub(p1.Pos)
	if p1p2.Len() < p1.Radius + p2.Radius {
		// collide
		p1.Pos = p1.Pos.Sub(p1p2.Mul(0.5))
		p2.Pos = p2.Pos.Add(p1p2.Mul(0.5))
		f1 := p1p2.Dot(v1) * damping / (p1p2.Len() * p1p2.Len())
		f2 := p1p2.Dot(v2) * damping / (p1p2.Len() * p1p2.Len())
		v1 = v1.Add(p1p2.Mul(f2 - f1))
		v2 = v2.Add(p1p2.Mul(f1 - f2))
		p1.PrevPos = p1.Pos.Sub(v1)
		p2.PrevPos = p2.Pos.Sub(v2)
	}
}

func (this *TreeCell)CheckCollision(upperStack []*Particle) {
	for i, p1 := range this.elements {
		for _, p2 := range this.elements[i+1:] {
			SolveCollision(p1, p2, this.damping)
		}
	}

	if (upperStack) {
		for _, p1 := range this.elements {
			for _, p2 := range upperStack {
				SolveCollision(p1, p2, this.damping)
			}
		}
	}

	this.
}

type QuadTree struct {
	level uint32
	rootWidth uint32
	rootHeight uint32
	minWidth uint32
	minHeight uint32
	m map[Morton]*Particle `Linear Quaternary Tree`
}

func NewQuadTree(level, width, height uint32) *QuadTree {
	cn := Pow(2, level)
	return &QuadTree{
		level: level,
		rootWidth: width,
		rootHeight: height,
		minWidth: width / cn,
		minHeight: height / cn,
	}
}

func bitslide(b uint32) uint32 {
	b = (b|b<<8)&0x00ff00ff
	b = (b|b<<4)&0x0f0f0f0f
	b = (b|b<<2)&0x33333333
	b = (b|b<<1)&0x55555555
	return b
}

func (this *QuadTree)SetParticles(particles []*Particle) {
	for _, p := range particles {
		var arr := this.m[this.ParticleToCellId(p).Morton]
		this.m[this.ParticleToCellId(p).Morton] = append(arr, p)
	}
}

func (this *QuadTree)PointToMorton(p m.Vec2) Morton {
	kx, ky := uint32(p.X()) / this.minWidth, uint32(p.Y()) / this.minHeight
	return Morton(bitslide(kx) | (bitslide(ky) << uint32(1)))
}

func (this *QuadTree)ParticleToCellId(p *Particle) *CellId {
	rad := m.Vec2{ p.Radius, p.Radius }
	leftBottom := p.Pos.Sub(rad)
	rightTop := p.Pos.Add(rad)

	ma := this.PointToMorton(leftBottom)
	mb := this.PointToMorton(rightTop)

	i := uint32(0)
	m := ma ^ mb
	for i = 0; m != 0; i++ {
		m = ma ^ mb
	}

	L := this.level - i
	I := mb >> (2*i)
	M := LevelIndexToMorton(L, I)
	return &CellId{
		Level: L,
		Index: I,
		Morton: M,
	}
}

func LevelIndexToMorton(L, I uint32) Morton {
	return Morton((Pow(4,L)-1)/(4-1)+I)
}

func (this *QuadTree)UpperSearch(c CellId) {
	var ret []Morton
	m := c.Morton
	for l := c.Level - 1; l >= 0; l-- {
		m = Morton((m-1) / 4)
		ret = append(ret, m)
	}
}

func (this *QuadTree)LowerCellId(c CellId) []CellId {

}