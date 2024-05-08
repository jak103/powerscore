package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"go.bug.st/serial"
)

func main() {
	fmt.Println("EIC Scoreboard")
	prefix := []byte{0x02, 0x74}
	dataChannel := make(chan []byte)

	go readSerial(context.TODO(), dataChannel)

	buffer := make([]byte, 100)
	for {
		data := <-dataChannel
		buffer = append(buffer, data...)
		// fmt.Printf("Len buffer: %v\n", len(buffer))

		if bytes.HasPrefix(buffer, prefix) && len(buffer) > 45 {
			processPacket(buffer[:45])
			buffer = buffer[45:]
			// os.Exit(1)
		} else if len(buffer) > 45 {
			buffer = buffer[1:]
		}
	}
}

func processPacket(packet []byte) {
	time := packet[2:6]
	tStr, _ := toTimeString(time)
	fmt.Printf("time: %s\n", tStr)

	period := string(packet[6:7])

	fmt.Printf("period: %s\n", period)

	homeScore := string(packet[7:9])
	awayScore := string(packet[9:11])

	homeScore = strings.ReplaceAll(homeScore, ":", " ")
	awayScore = strings.ReplaceAll(awayScore, ":", " ")
	fmt.Printf("Home: %s, Away: %s\n", homeScore, awayScore)

	printPacket(packet, 1, 2, 6, 7, 9, 11, 18, 20, 22, 24, 27, 29, 32, 34, 37, 39, 42, 43, 44)
}

func readSerial(ctx context.Context, data chan []byte) {
	// 	RS232 DTE
	// Serial port settings: Baud Rate = 9600, Data Bits = 8, Parity = None, Stop Bits = 1 and Flow Control = None
	// file, err := os.OpenFile("output.txt", os.O_CREATE, os.ModeAppend)
	// if err != nil {
	// 	fmt.Println("Failed to open file", err)
	// }
	mode := &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open("COM4", mode)
	if err != nil {
		log.Fatal(err)
	}

	buff := make([]byte, 100)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		data <- buff[:n]

		// file.WriteString(hex.EncodeToString(buff[:n]))
	}
}

func toTimeString(data []byte) (string, bool) {
	// fmt.Printf("bin: %b\n", data)
	paused := true
	if data[0]&0x80 == 0 {
		paused = false
	}
	data[0] = data[0] & 0x7F // Clear the time stop bit
	data[1] = data[1] & 0x7F // Clear upper colon bit
	data[2] = data[2] & 0x7F // Clear lower colon bit

	sep := "."
	if data[1]&0x80 == 0 {
		sep = ":"
	}

	// fmt.Println("hex:", hex.EncodeToString(data))
	min := string(bytes.ReplaceAll(data[:2], []byte{0x3a}, []byte{0x20}))
	sec := string(bytes.ReplaceAll(data[2:], []byte{0x3a}, []byte{0x20}))
	time := fmt.Sprintf("%s%s%s", min, sep, sec)
	// fmt.Printf("Time: %s\n", time)
	return time, paused
}

func printPacket(packet []byte, places ...int) {
	packetStr := hex.EncodeToString(packet)
	for i, place := range places {
		pos := (place * 2) + i
		packetStr = fmt.Sprintf("%s %s", packetStr[:pos], packetStr[pos:])
	}
	fmt.Println(packetStr)
}
