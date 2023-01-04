//go:build goexperiment.arenas

package main

import "arena"

func main() {
	// Create an arena in the beginning of the function.
	mem := arena.NewArena()
	// Free the arena in the end.
	defer mem.Free()

	b := arena.MakeSlice[byte](mem, 10, 100)

	println("%+v", b)
}
