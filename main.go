package main

import (
	"fmt"

	"github.com/AjdinHalac/GoPortScanner/port"
)

func main() {
	fmt.Println("Port Scanner in Go")

	results := port.InitialScan("localhost")
	fmt.Println(results)
}
