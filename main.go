package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	w, err := NewWindow(Title, WindowWidth, WindowHeight, FrameRate)
	if err != nil {
		panic(err)
	}
	//SetupTask1(w)
	//SetupTask2(w)
	SetupTask4(w)
	w.Exec()
}