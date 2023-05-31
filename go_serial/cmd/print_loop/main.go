package main

import (
	"go_serial/internal/pkg"
	"log"
)

func main() {
	filename := "../../scripts/basic_src/print_loop.txt"
	p,err := pkg.OpenPort()
	defer p.Port.Close()
	if err != nil {
		log.Fatal(err)
	}
	program := pkg.ReadProgram(filename)
	go p.ProgramExecute(program)
	p.PrintLoop()
}
