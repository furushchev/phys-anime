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
	onUpdateCallbacks []func()
	objects []Drawable
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
	}

	win.SetFramebufferSizeCallback(window.onResize)
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
	gl.ClearColor(1,1,1,1)
}

func (this *Window)AddObject(o *Drawable) {
	this.objects = append(this.objects, *o)
}

func (this *Window)drawObjects() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Enable(gl.BLEND)
	gl.Enable(gl.POINT_SMOOTH)
	gl.Enable(gl.LINE_SMOOTH)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.LoadIdentity()

	for _, obj := range this.objects {
		obj.Draw()
	}
}

func (this *Window)update(dt float32) {
	for _, obj := range this.objects {
		obj.Update(dt)
	}
}

func (this *Window)Exec() {
	runtime.LockOSThread()
	glfw.SwapInterval(1)

	ticker := time.NewTicker(time.Duration(1.0 * float64(time.Second) / float64(this.FrameRate)))
	defer ticker.Stop()

	for !this.window.ShouldClose() {
		for _, f := range this.onUpdateCallbacks {
			f()
		}
		this.drawObjects()
		this.update(1.0 / float32(this.FrameRate))
		this.window.SwapBuffers()
		glfw.PollEvents()
		<- ticker.C
	}
}