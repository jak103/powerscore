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

//Reads hex data from "trimmed.txt", parses scoreboard data, outputs to "parsed.txt"
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
	defer port.Close()

	//Read hex from the file "trimmed.txt"
	go ingest_simulator.Start("trimmed.txt", port)
	output := make(chan *models.ScoreboardData)
	go serial.Start(context.TODO(), output, port)

	os.Create("parsed.txt")
	f, err := os.OpenFile("parsed.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for {
		scoreboardData := <-output
		fmt.Println(scoreboardData)

		if _, err := f.Write([]byte(fmt.Sprintln(scoreboardData))); err != nil {
			log.Fatal(err)
		}
	}
}
