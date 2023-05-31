package main

import (
	"go_serial/internal/pkg"
	"log"
)

func main() {
	// filename := "../../scripts/basic_src/print_loop.txt"
	filename := "../../scripts/basic_src/printloop_with_count.txt"
	p, err := pkg.OpenPort()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Port.Close()
	program := pkg.ReadProgram(filename)
	p.ProgramExecute(program)
	p.PrintLoop()
}
