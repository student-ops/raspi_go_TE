package main

import (
	"go_serial/internal/pkg"
	"log"
	"testing"
)

func TestReadProgram(t *testing.T) {
	p, err := pkg.OpenPort()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Port.Close()
	commnads := []string{"edit 1", "New", "psave", "edit 0", "run"}
	p.PortWriteCommand(commnads)
}

func TestSerialIo(t *testing.T) {
	filename := "../../scripts/basic_src/printloop_with_count.txt"
	p, err := pkg.OpenPort()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Port.Close()
	program := pkg.ReadProgram(filename)
	go p.ProgramExecute(program)
	p.PrintLoopPararel()
}
