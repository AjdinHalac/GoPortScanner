package main

import (
	"fmt"

	"github.com/AjdinHalac/GoPortScanner/port"
)

func main() {
	fmt.Println("Port Scanner in Go")

	open := port.ScanPort("tcp", "localhost", 80)
	fmt.Printf("Port Open: %t\n", open)

	results := port.InitialScan("localhost")
	fmt.Println(results)
}
