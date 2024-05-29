package ingest_simulator

import (
	"encoding/hex"
	"io"
	"log"
	"os"

	"go.bug.st/serial"
)

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

	// On Matt's machine, the serial port will drop the first byte and duplicate the second byte.
	//		Setting the first two bytes to zero works around this issue and ensures no data is lost.
	serialBuffer = append([]byte{0x0,0x0}, serialBuffer...)

	if err != nil {
		println("fails to decode hex")
		// break
	}

	n, err := port.Write(serialBuffer)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("write.go -- Sent %v bytes to serial port\n", n)

	return nil
}

