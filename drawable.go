package main

type Drawable interface {
	Update(dt, screenWidth, screenHeight float64)
	CheckCollision(interface{})
	Draw()
}
