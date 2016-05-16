package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"runtime"
	"time"
)

type Window struct {
	window *glfw.Window
	Width, Height int
	FrameRate int
	particles []*Particle
	effectors []Effector
	step bool
}

func NewWindow(title string, width, height int, frameRate int) (w *Window, err error) {
	runtime.LockOSThread()

	err = glfw.Init()
	if err != nil {
		return
	}

	win, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return
	}

	window := &Window{
		window: win,
		Width: width,
		Height: height,
		FrameRate: frameRate,
		step: false,
	}

	win.SetFramebufferSizeCallback(window.onResize)
	win.SetKeyCallback(window.onKeyPressed)
	win.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		return
	}

	window.onResize(win, width, height)
	w = window
	return
}

func (this *Window)Terminate() {
	this.window.SetShouldClose(true)
	glfw.Terminate()
}

func (this *Window)onResize(win *glfw.Window, w, h int) {
	w, h = this.window.GetSize()
	width, height := this.window.GetFramebufferSize()
	gl.Viewport(0,0, int32(width), int32(height))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(w), 0, float64(h), -1, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.ClearColor(0,0,0,1)
}

func (this *Window)onKeyPressed(win *glfw.Window, key glfw.Key, code int, action glfw.Action, mods glfw.ModifierKey) {
	switch key {
	case 's':
		println("pressed s")
		this.step = true
	}
}

func (this *Window)AddParticle(p *Particle) {
	this.particles = append(this.particles, p)
}

func (this *Window)AddEffector(e Effector) {
	this.effectors = append(this.effectors, e)
}

func (this *Window)update(dt float64) {
	for _, e := range this.effectors {
		e.Update(this.particles, dt)
	}
}

func (this *Window)drawObjects() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Enable(gl.BLEND)
	gl.Enable(gl.POINT_SMOOTH)
	gl.Enable(gl.LINE_SMOOTH)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.LoadIdentity()

	for _, p := range this.particles {
		p.Draw()
	}
	gl.Flush()
}

func (this *Window)Exec() {
	runtime.LockOSThread()
	glfw.SwapInterval(1)

	ticker := time.NewTicker(time.Duration(1.0 * float64(time.Second) / float64(this.FrameRate)))
	defer ticker.Stop()

	for !this.window.ShouldClose() {
		this.update(1.0 / float64(this.FrameRate))
		this.drawObjects()
		this.window.SwapBuffers()
		glfw.PollEvents()
		<- ticker.C
	}
}