package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// Open the serial data file
	file, err := os.Open("full_game.txt")
	if err != nil {
		fmt.Println("Error opening serial data file:", err)
		return
	}
	defer file.Close()

	// Read data from the file
	buf := make([]byte, 100)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("EOF reached, end of serial data file")
				return
			}
			fmt.Println("Error reading from serial data file:", err)
			return
		}

		processPacket(buf[:n])
	}
}

func processPacket(packet []byte) {
	fmt.Printf("Raw packet: %x\n", packet)

	// Extract time field (assuming it's represented by 4 bytes)
	time := packet[2:6]
	// Convert each byte to a hexadecimal string and format as two digits
	tStr := fmt.Sprintf("%02x:%02x", time[2], time[3]) // Extracting last two bytes for minutes and seconds
	fmt.Printf("time: %s\n", tStr)

	// Extract period field
	period := string(packet[6:7])
	fmt.Printf("period: %s\n", period)

	// Extract home and away scores (assuming each is represented by 2 bytes)
	homeScore := fmt.Sprintf("%02x", packet[7]) // Convert byte to a hexadecimal string
	awayScore := fmt.Sprintf("%02x", packet[8]) // Convert byte to a hexadecimal string
	fmt.Printf("Home: %s, Away: %s\n", homeScore, awayScore)
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
