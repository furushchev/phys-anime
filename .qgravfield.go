package main

import m "github.com/go-gl/mathgl/mgl64"

type QTree struct {
	width float64
	height float64
	root *QNode
}

type QNode struct {
	leftBottom *m.Vec2
	rightTop *m.Vec2
	element *Particle
	children [4]*QNode
}

func NewQNode(leftBottom, rightTop *m.Vec2) *QNode {
	return &QNode{
		leftBottom: leftBottom,
		rightTop: rightTop,
	}
}

func (this *QNode)Check(p *Particle) bool {
	if this.leftBottom.X() <= p.Pos.X() && p.Pos.X() < this.rightTop.X() &&
	this.leftBottom.Y() <= p.Pos.Y() && p.Pos.Y() < this.rightTop.Y() {
		return true
	} else {
		return false
	}
}

func (this *QNode)Add(p *Particle) {
	if this.element {
		diag := this.rightTop.Sub(this.leftBottom).Mul(0.5)
		x1 := m.Vec2{ diag.X(), 0 }
		y1 := m.Vec2{ 0, diag.Y() }
		this.children[0] = NewQNode(this.leftBottom, this.leftBottom.Add(diag))
		this.children[1] = NewQNode(this.leftBottom.Add(x1), this.leftBottom.Add(x1).Add(diag))
		this.children[2] = NewQNode(this.leftBottom.Add(y1), this.leftBottom.Add(y1).Add(diag))
		this.children[3] = NewQNode(this.leftBottom.Add(diag), this.rightTop)
		for _, c := range this.children {
			if c.Check(this.element) {
				c.Add(this.element)
				this.element = nil
				break
			}
		}
		for _, c := range this.children {
			if c.Check(p) {
				c.Add(p)
				return
			}
		}
	} else {
		this.element = p
	}
}

func (this *QTree)AddParticle(p *Particle) {
	this.root.Add(p)
}