package serial

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/jak103/powerscore/internal/models"
	"go.bug.st/serial"
)

func Start(ctx context.Context, output chan *models.ScoreboardData) error {
	mode := &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open("COM4", mode)
	if err != nil {
		log.Fatal(err)
		return err
	}

	dataChannel := make(chan []byte)

	go readSerial(ctx, port, dataChannel)
	go processPackets(ctx, dataChannel, output)

	return nil
}

func readSerial(ctx context.Context, port serial.Port, data chan []byte) {
	// TODO These should eventually end up in a settings/config location
	// 	RS232 DTE
	// Serial port settings: Baud Rate = 9600, Data Bits = 8, Parity = None, Stop Bits = 1 and Flow Control = None
	// file, err := os.OpenFile("output.txt", os.O_CREATE, os.ModeAppend)
	// if err != nil {
	// 	fmt.Println("Failed to open file", err)
	// }

	buff := make([]byte, 100)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Closing serial listener")
			return
		default:
			n, err := port.Read(buff)
			if err != nil {
				fmt.Printf("Failed to read from serial port: %v\n", err)
				return
			}

			data <- buff[:n]
		}
	}
}

func processPackets(ctx context.Context, dataChannel chan []byte, scoreboardDataChannel chan *models.ScoreboardData) {
	prefix := []byte{0x02, 0x74}
	buffer := make([]byte, 255)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Closing packet processor")
			return
		default:
			data := <-dataChannel
			buffer = append(buffer, data...)

			if bytes.HasPrefix(buffer, prefix) && len(buffer) > 45 {
				scoreboardData := parsePacket(buffer[:45])
				if scoreboardData != nil {
					scoreboardDataChannel <- scoreboardData
				}

				buffer = buffer[45:]
			} else if len(buffer) > 45 {
				buffer = buffer[1:]
			}
		}
	}
}

func parsePacket(packet []byte) *models.ScoreboardData {
	printPacket(packet, 1, 2, 6, 7, 9, 11, 18, 20, 22, 24, 27, 29, 32, 34, 37, 39, 42, 43, 44)

	if !isValidPacket(packet) {
		return nil
	}

	data := &models.ScoreboardData{
		Home: models.TeamData{},
		Away: models.TeamData{},
	}

	data.GameTime, data.Paused = parseGameTime(packet[2:6])
	data.Period = string(packet[6:7])

	data.Home.Score = strings.ReplaceAll(string(packet[7:9]), ":", "")
	data.Away.Score = strings.ReplaceAll(string(packet[9:11]), ":", "")

	data.Home.Penalties = parsePenalties(packet[18:20])
	data.Away.Penalties = parsePenalties(packet[21:22])

	return data
}

func isValidPacket(packet []byte) bool {
	// TODO Validate the checksum
	return true
}

func parseGameTime(data []byte) (string, bool) {
	paused := data[0]&0x80 != 0

	sep := ":"
	if data[1]&0x80 == 0 {
		sep = "."
	}

	data[0] = data[0] & 0x7F // Clear the time stop bit
	data[1] = data[1] & 0x7F // Clear upper colon bit
	data[2] = data[2] & 0x7F // Clear lower colon bit

	min := strings.ReplaceAll(string(data[:2]), ":", "")
	sec := strings.ReplaceAll(string(data[2:]), ":", "")
	time := fmt.Sprintf("%s%s%s", min, sep, sec)

	return time, paused
}

func parsePenalties(data []byte) []models.Penalty {
	// TODO parse penalties
	return nil
}

func parsePenaltyTime(data []byte) string {
	// TODO Parse penalty time
	return ""
}

func printPacket(packet []byte, places ...int) {
	packetStr := hex.EncodeToString(packet)
	for i, place := range places {
		pos := (place * 2) + i
		packetStr = fmt.Sprintf("%s %s", packetStr[:pos], packetStr[pos:])
	}
	fmt.Println(packetStr)
}
