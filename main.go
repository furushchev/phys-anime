package main

import (
	"runtime"
	"os"
	"fmt"
	"strconv"
)

func printUsage() {
	fmt.Fprintln(os.Stderr, "Please specify Task ID: [1-4]")
	fmt.Println("phys-anime [TaskID]")
	os.Exit(1)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	w, err := NewWindow(Title, WindowWidth, WindowHeight, FrameRate)
	if err != nil {
		panic(err)
	}
	args := os.Args
	if len(args) != 2 {
		printUsage()
	}
	taskID, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("Please input valid taskID"))
	}
	switch taskID {
	case 1:
		SetupTask1(w)
	case 2:
		SetupTask2(w)
	case 3:
		SetupTask3(w)
	case 4:
		SetupTask4(w)
	case 5:
		SetupTask5(w)
	default:
		printUsage()
	}
	w.Exec()
}