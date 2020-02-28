package main

import (
	"fmt"
	"runtime"

	"github.com/AjdinHalac/GoPortScanner/port"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Port Scanner in Go")

	results := port.InitialScan("localhost")
	fmt.Println(results)
}
