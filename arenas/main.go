package main

import (
	"arena"
	"fmt"
)

func main() {
	// Create an arena in the beginning of the function.
	mem := arena.NewArena()
	// Free the arena in the end.
	defer mem.Free()

	b := arena.MakeSlice[byte](mem, 10, 100)

	fmt.Printf("%+v\n", b)
}
