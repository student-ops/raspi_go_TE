package pkg

import (
	"log"
	"time"

	"go.bug.st/serial"
)

type myPort struct{
	Port serial.Port
}

func OpenPort() (*myPort,error) {
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal(err)
	}
	return  &myPort{Port: port}, err
}

func (p *myPort)PortWrite(s string){
	p.Port.Write([]byte(s + "\r"))
	time.Sleep(100 * time.Millisecond)
}

func (p *myPort)PortWriteCommand(s []string){
	for _,v := range s{
		p.PortWrite(v)
	}
}


func (p *myPort)ProgramExecute(program string){
	// delete program
	commnads := []string{"edit 1","New","psave","edit 0","run"}
	p.PortWriteCommand(commnads)

	p.Port.Write([]byte("edit 1\r"))
	n, err := p.Port.Write([]byte(program))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sent %v bytes \n", n)
	p.Port.Write([]byte("own = 1\r"))
	p.Port.Write([]byte("dst = 0\r"))
	p.Port.Write([]byte("Auto=\"pload:run\"\r"))
	p.Port.Write([]byte("ssave\r"))
	p.Port.Write([]byte("psave\r"))
	p.Port.Write([]byte("edit 0\r"))
}
