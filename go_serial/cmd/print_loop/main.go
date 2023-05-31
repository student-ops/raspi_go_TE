package main

import (
	"go_serial/internal/pkg"
	"log"
)

func main() {
	filename := "../../scripts/basic_src/print_loop.txt"
	p, err := pkg.OpenPort()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Port.Close()
	program := pkg.ReadProgram(filename)
	go p.ProgramExecute(program)
	p.PrintLoopPararel()
}
