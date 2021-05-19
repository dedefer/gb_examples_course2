package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

const (
	filesCount     = 1_000_000
	filesDir       = "go_touch"
	filenamePrefix = "file"
)

func panicHandler() {
	val := recover()
	if val != nil {
		trace := make([]byte, 1024)
		runtime.Stack(trace, false)
		fmt.Printf("panic raised: %v\n%s\n", val, trace)
	}
}

func main() {
	filesPath := path.Join(os.TempDir(), filesDir)
	allFiles := make([]*os.File, 0, filesCount)

	defer panicHandler()

	for i := 0; i < filesCount; i++ {
		name := fmt.Sprintf("%s%07d", filenamePrefix, i)
		file, err := os.Create(path.Join(filesPath, name))
		if err != nil {
			fmt.Println(err)
			break
		}
		allFiles = append(allFiles, file)
	}

	fmt.Printf("touched %d files\n", len(allFiles))
}
