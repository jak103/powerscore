package ingest_simulator

import (
	"encoding/hex"
	"io"
	"log"
	"os"

	"go.bug.st/serial"
)

// Reads hexidecimal data from a file at `data_filepath`. Writes the byte data to the serial port `port`.
// Reports an error if the file `data_filepath` cannot be read as hexidecimal or data cannot be written to the serial port.
func Start(data_filepath string, port serial.Port) error {
	const BUFFER_SIZE = 90
	file, err := os.Open(data_filepath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	fileStats, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileBuffer := make([]byte, fileStats.Size())
	bytesread, err := file.Read(fileBuffer)
	log.Printf("write.go -- read %v bytes from file\n", bytesread)

	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
			return err
		}
	}
	serialBuffer, err := hex.DecodeString(string(fileBuffer))

	// After experiencing an issue where the first byte is dropped and the second is repeated,
	// this line of code solved that issue because serial.go truncates the first byte received
	// until the prefix is found.
	serialBuffer = append([]byte{0x0,0x0}, serialBuffer...)

	if err != nil {
		println("failed to decode hex")
	}

	n, err := port.Write(serialBuffer)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("write.go -- Sent %v bytes to serial port\n", n)

	return nil
}

