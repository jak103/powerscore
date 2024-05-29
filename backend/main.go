package main

import (
	"context"
	"fmt"
	"log"
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


	go ingest_simulator.Start("output.txt", port)
	output := make(chan *models.ScoreboardData)
	go serial.Start(context.TODO(), output, port)

	buffer := make([]*models.ScoreboardData, 1)
	for {
		println("main for")
		scoreboardData := <-output
		buffer = append(buffer, scoreboardData)

		if len(buffer) > 0 {
			fmt.Println(buffer)
		}
	}

}
