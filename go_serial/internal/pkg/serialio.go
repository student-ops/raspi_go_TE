package pkg

import (
	"log"
	"strconv"
	"sync"
	"time"

	"go.bug.st/serial"
)

type myPort struct {
	Port serial.Port
}

func OpenPort() (*myPort, error) {
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal(err)
	}
	return &myPort{Port: port}, err
}

func (p *myPort) PortWrite(s string) error {
	_, err := p.Port.Write([]byte(s + "\r"))
	if err != nil {
		return err
	}
	time.Sleep(100 * time.Millisecond)
	return nil
}

func (p *myPort) PortWriteCommand(s []string) error {
	for _, v := range s {
		err := p.PortWrite(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *myPort) ProgramExecute(program string) {
	p.PortWrite("edit 1")
	p.PortWrite(program)
	p.PortWrite("edit 0")
	p.PortWrite("run")
}

func (p *myPort) VuoyExecute(file string) {
	// delete program
	commnads := []string{"edit 1", "New", "psave", "edit 0", "run"}
	p.PortWriteCommand(commnads)

	p.Port.Write([]byte("edit 1\r"))
	program := ReadProgram(file)
	err := p.PortWrite(program)
	if err != nil {
		log.Fatal(err)
	}
	commnads = []string{"own =1", "dst = 0", "Auto=\"pload:run\"", "ssave", "psave", "edit 0"}
	p.PortWriteCommand(commnads)
}

func (p *myPort) PrintLoop() {
	for {
		buf := make([]byte, 128)
		n, err := p.Port.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(string(buf[:n]))
	}
}

func (p *myPort) PrintLoopPararel() {
	count := 0
	buffer := make(chan []byte, 100) // Create a channel to store data

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			buf := make([]byte, 128)
			n, err := p.Port.Read(buf)
			if err != nil {
				log.Println(err)
				close(buffer)
				return
			}
			count += 1
			buffer <- buf[:n] // Send the data to the channel
		}
	}()

	go func() {
		defer wg.Done()
		for buf := range buffer {
			log.Print(strconv.Itoa(count) + ":" + string(buf))
		}
	}()

	wg.Wait()
}
