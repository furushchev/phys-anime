package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	w, err := NewWindow(Title, WindowWidth, WindowHeight, FrameRate)
	if err != nil {
		panic(err)
	}
	NewTask1(w)
	w.Exec()
}