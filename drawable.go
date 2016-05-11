package main

type Drawable interface {
	Update(dt float32)
	Draw()
}
