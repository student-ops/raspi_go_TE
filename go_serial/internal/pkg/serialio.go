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

func (p *myPort)PortWrite(s string)error{
	_, err := p.Port.Write([]byte(s + "\r"))
	if err != nil{
		return err
	}
	time.Sleep(100 * time.Millisecond)
	return	nil
}

func (p *myPort)PortWriteCommand(s []string)error {
	for _,v := range s{
		err := p.PortWrite(v)
		if err != nil{
			return err
		}
	}
	return nil
}

func (p *myPort)ProgramExecute(program string){
	p.PortWrite("edit 1")
	p.PortWrite(program)
	p.PortWrite("edit 0")
	p.PortWrite("run")
}

func (p *myPort)VuoyExecute(file string){
	// delete program
	commnads := []string{"edit 1","New","psave","edit 0","run"}
	p.PortWriteCommand(commnads)

	p.Port.Write([]byte("edit 1\r"))
	program := ReadProgram(file)
	err := p.PortWrite(program)
	if err != nil {
		log.Fatal(err)
	}
	commnads = []string{"own =1","dst = 0","Auto=\"pload:run\"","ssave","psave","edit 0"}
	p.PortWriteCommand(commnads)
}


func (p *myPort)PrintLoop(){
	for {
		buf := make([]byte, 128)
		_, err := p.Port.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		go (func(){
			sbuf := make([]byte, 128)
			_, serr := p.Port.Read(buf)
			if serr != nil {
				log.Fatal(err)
			}
			log.Print(string(buf[:n]))
		})
		log.Print(string(buf[:n]))
	}
}