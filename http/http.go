package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	math.Sqrt(2)
	OS := runtime.GOOS
	fmt.Println("os:", OS)
}
