package serial

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"scoreboard/internal/models"

	"go.bug.st/serial"
)

func Start(ctx context.Context, output chan *models.ScoreboardData, port serial.Port) error {
	dataChannel := make(chan []byte)

	go readSerial(ctx, port, dataChannel)
	go processPackets(ctx, dataChannel, output)

	return nil
}

func readSerial(ctx context.Context, port serial.Port, data chan []byte) {
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
	buffer := make([]byte, 0)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Closing packet processor")
			return
		default:
			data := <-dataChannel
			buffer = append(buffer, data...)
			if bytes.HasPrefix(buffer, prefix) && len(buffer) >= 45 {
				// log.Printf("serial.go -- parsing packet '%s'", string(buffer[:45]))
				scoreboardData := parsePacket(buffer[:45])
				if scoreboardData != nil {
					scoreboardDataChannel <- scoreboardData
				}
				buffer = buffer[45:]
			} else {
				//Trim from the front until the prefix is found or the buffer is empty.
				for len(buffer) > 0 && !bytes.HasPrefix(buffer, prefix) {
					buffer = buffer[1:]
				}
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

	//TODO: Unsure if these are the correct bytes for shots on goal.
	// This should be an easy change if it is not correct.
	data.Home.ShotsOnGoal = strings.ReplaceAll(string(packet[18:20]), ":", "")
	data.Away.ShotsOnGoal = strings.ReplaceAll(string(packet[20:22]), ":", "")

	data.Home.Penalties = parsePenalties(packet[22:32])
	data.Away.Penalties = parsePenalties(packet[32:42])

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
	//Penalty is 10 bytes long. First 5 are one penalty, second 5 are another
	penalties := make([]models.Penalty, 0)
	penalty := parsePenalty(data[0:5])
	if penalty != nil {
		penalties = append(penalties, *penalty)
	}
	penalty = parsePenalty(data[5:10])
	if penalty != nil {
		penalties = append(penalties, *penalty)
	}
	return penalties
}

func parsePenalty(data []byte) *models.Penalty {
	//Penalty is 5 bytes long. First two describe player number, second two describe time.
	playerNumber := string(data[0:2])
	playerNumber = strings.Trim(playerNumber, ":")
	if playerNumber == "" {
		return nil
	}
	penalty := models.Penalty{
		PlayerNumber: string(data[0:2]),
		Time: parsePenaltyTime(data[2:5]),
	}
	return &penalty
}

func parsePenaltyTime(data []byte) string {
	sep := ":"
	if data[0]&0x80 == 0 {
		sep = "."
	}

	data[0] = data[0] & 0x7F // Clear upper colon bit
	data[1] = data[1] & 0x7F // Clear lower colon bit

	min := strings.ReplaceAll(string(data[:1]), ":", "")
	sec := strings.ReplaceAll(string(data[1:]), ":", "")
	time := fmt.Sprintf("%s%s%s", min, sep, sec)
	return time
}

func printPacket(packet []byte, places ...int) {
	packetStr := hex.EncodeToString(packet)
	for i, place := range places {
		pos := (place * 2) + i
		packetStr = fmt.Sprintf("%s %s ", packetStr[:pos], packetStr[pos:])
	}
	fmt.Println(packetStr)
}
