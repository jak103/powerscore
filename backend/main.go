package main

import (
	"context"
	"fmt"

	"github.com/jak103/powerscore/internal/models"
	"github.com/jak103/powerscore/internal/serial"
)

func main() {
	fmt.Println("Powerscore v0.0.0") // TODO get real version here

	serialContext := context.Background()
	scoreboardDataChannel := make(chan *models.ScoreboardData)
	serial.Start(serialContext, scoreboardDataChannel)

	for {
		data := <-scoreboardDataChannel
		fmt.Println(data)
	}
}
