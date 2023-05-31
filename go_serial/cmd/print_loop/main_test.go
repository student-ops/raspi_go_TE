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

func TestResetProgram(t *testing.T) {
	p, err := pkg.OpenPort()
	if err != nil {
		t.Fatal(err)
	}
	defer p.Port.Close()
	buf := make([]byte, 128)
	p.Port.Write([]byte{0x03})
	commands := []string{"edit 1", "New", "psave", "edit 0"}
	if err = p.PortWriteCommands(commands); err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		if err = p.PortWrite("run"); err != nil {
			t.Error(err)
		}
		close(done)
	}()

	n, err := p.Port.Read(buf)
	if err != nil {
		t.Errorf("failed port read %v", err)
	}

	// Wait for the PortWrite goroutine to finish
	<-done

	output := string(buf[:n])
	t.Logf("read %s", output)
}

func TestSerialIo(t *testing.T) {
	filename := "../../scripts/basic_src/printloop_with_count.txt"
	var failcount = 0
	loopcount := 50
	p, err := pkg.OpenPort()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Port.Close()
	program := pkg.ReadProgram(filename)
	p.ProgramExecute(program)
	time.Sleep(500 * time.Microsecond)

	logChannel := make(chan pkg.LogEntry, 100)
	go p.PrintLoopParallel(logChannel, loopcount)

	previousNumber := 1

	for logEntry := range logChannel {
		// 変数logEntryを操作する
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
			fmt.Printf("prev %d; tmp %d\n", previousNumber, number)
			// 数字の操作を行う
			gap := number - previousNumber
			if gap > 1 {
				failcount += gap - 1
				log.Printf("packet lost count :%d \n", failcount)
			}

			previousNumber = number
		}
	}

	// fail if packed lost count is more than 10%
	log.Printf("packet lost rate : %.2f\n", float64(failcount)*100/float64(loopcount))
	if failcount > loopcount/10 {
		t.Errorf("packet lost rate : %.2f\n", float64(failcount)*100/float64(loopcount))
	}
}
