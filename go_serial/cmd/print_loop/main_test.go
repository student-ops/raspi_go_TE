package main

import (
	"go_serial/internal/pkg"
	"log"
	"testing"
)

func TestReadProgram(t *testing.T) {
	p,err := pkg.OpenPort()
	defer p.Port.Close()
	if err != nil {
		log.Fatal(err)
	}
	commnads := []string{"edit 1","New","psave","edit 0","run"}
	p.PortWriteCommand(commnads)
}

