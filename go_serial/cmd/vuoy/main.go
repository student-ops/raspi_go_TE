package main

import (
	"go_serial/internal/pkg"
	"log"
)

func main() {
	file := "../basic_src/print_loop.txt"
	p,err := pkg.OpenPort()
	defer p.Port.Close()
	if err != nil {
		log.Fatal(err)
	}
	p.VuoyExecute(file)
}
