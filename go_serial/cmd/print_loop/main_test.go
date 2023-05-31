package main

import (
	"fmt"
	"go_serial/internal/pkg"
	"log"
	"testing"
	"time"
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

type LogEntry struct {
	Count int
	Text  string
}

func TestSerialIo(t *testing.T) {
	filename := "../../scripts/basic_src/printloop_with_count.txt"
	p, err := pkg.OpenPort()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Port.Close()
	program := pkg.ReadProgram(filename)
	p.ProgramExecute(program)
	time.Sleep(500 * time.Microsecond)
	p.PrintLoopParallel()

	logChannel := make(chan LogEntry, 100)

	go func() {
		count := 0
		for {
			buf := make([]byte, 128)
			n, err := p.Port.Read(buf)
			if err != nil {
				log.Println(err)
				return
			}
			count++
			logChannel <- LogEntry{Count: count, Text: string(buf[:n])}
		}
	}()

	for logEntry := range logChannel {
		// 変数logEntryを操作する
		fmt.Println(logEntry.Count, logEntry.Text)
	}
}
