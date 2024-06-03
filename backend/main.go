package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"scoreboard/internal/ingest_simulator"
	"scoreboard/internal/models"
	"scoreboard/internal/serial"

	bug_serial "go.bug.st/serial"
)

func main() {
	mode := &bug_serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		Parity:   bug_serial.NoParity,
		StopBits: bug_serial.OneStopBit,
	}

	port, err := bug_serial.Open("COM4", mode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(port);
	defer port.Close()


	go ingest_simulator.Start("trimmed.txt", port)
	output := make(chan *models.ScoreboardData)
	go serial.Start(context.TODO(), output, port)

	// buffer := make([]*models.ScoreboardData, 1)
	for {
		println("main for")
		scoreboardData := <-output
		// buffer = append(buffer, scoreboardData)
		fmt.Println(scoreboardData)

		f, err := os.OpenFile("parsed.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()

		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte(fmt.Sprintln(scoreboardData))); err != nil {
			log.Fatal(err)
	}
		// if len(buffer) > 0 {
		// 	fmt.Println(buffer)
		// }
	}

}
