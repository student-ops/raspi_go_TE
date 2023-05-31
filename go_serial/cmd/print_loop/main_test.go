package main

import (
	"fmt"
	"go_serial/internal/pkg"
	"log"
	"regexp"
	"strconv"
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

func TestSerialIo(t *testing.T) {
	filename := "../../scripts/basic_src/printloop_with_count.txt"
	var failcount = 0
	p, err := pkg.OpenPort()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Port.Close()
	program := pkg.ReadProgram(filename)
	p.ProgramExecute(program)
	time.Sleep(500 * time.Microsecond)

	logChannel := make(chan pkg.LogEntry, 100)
	p.PrintLoopParallel(logChannel)

	go func() {
		for {
			buf := make([]byte, 128)
			n, err := p.Port.Read(buf)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println("Received log entry:", string(buf[:n]))
			logChannel <- pkg.LogEntry{Text: string(buf[:n])}
		}
	}()

	var previousNumber int

	for logEntry := range logChannel {
		// 変数logEntryを操作する
		log.Println("Received log entry:", logEntry.Text) // Add this line
		pattern := `@(\d+)`
		r := regexp.MustCompile(pattern)
		matches := r.FindAllStringSubmatch(logEntry.Text, -1)
		if len(matches) == 0 {
			fmt.Println("No matches found in log entry") // Add this line
		}

		for _, match := range matches {
			number, err := strconv.Atoi(match[1])
			if err != nil {
				log.Println(err)
				continue
			}
			// 数字の操作を行う
			log.Println("Parsed number:", number) // Add this line

			failcount++
			if previousNumber != number-1 {
				log.Printf("packed lost count :%d", failcount)
			}

			previousNumber = number
		}
	}
}
