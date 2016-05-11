package main


func main() {
	w, err := NewWindow(Title, WindowWidth, WindowHeight, FrameRate)
	if err != nil {
		panic(err)
	}
	w.Exec()
}