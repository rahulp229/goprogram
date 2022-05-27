package readFile

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"testcsv/writeFile"
)

func OpenFile(fileName string) (io.Reader, error) {
	return os.Open(fileName)
}

func ReadFile(f io.Reader) error {
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}
	log.Println("start writing file data...")
	writeFile.CreateFileWithValidInvalidData(data)
	return nil
}
