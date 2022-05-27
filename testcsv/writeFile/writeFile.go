package writeFile

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var columnToValidate = []string{"email", "Email", "EMAIL"}

const (
	Email = "email"
)

func CreateFileWithValidInvalidData(data [][]string) {
	validFile, errValidFile := CreateNewFile("valid")
	checkError("can not create valid file", errValidFile)

	inValidFile, errInvalidFile := CreateNewFile("invalid")
	checkError("can not create in valid file", errInvalidFile)

	validWriter := csv.NewWriter(validFile)
	defer validWriter.Flush()

	inValidWriter := csv.NewWriter(inValidFile)
	defer inValidWriter.Flush()

	var ColumnNameToValidate int
	for index, record := range data {
		if index == 0 {
			ColumnNameToValidate = getValidateColumnIndex(record)
			log.Println("get column name for validate the value of it")
			errWriteValidFile := validWriter.Write(record)
			checkError("unable to insert record in file", errWriteValidFile)
			errWriteInValidFile := inValidWriter.Write(record)
			checkError("unable to insert record in invalid file", errWriteInValidFile)

		} else {
			isValidEmail := Validate(record[ColumnNameToValidate])
			if isValidEmail {
				log.Printf("insert valid data %d \n", index)
				errWriteValidFile := validWriter.Write(record)
				checkError("unable to insert record in valid file", errWriteValidFile)

			} else {
				log.Printf("insert in valid data %d \n", index)
				errWriteInValidFile := inValidWriter.Write(record)
				checkError("unable to insert record in in valid file", errWriteInValidFile)

			}

		}

	}
}

func CreateNewFile(fileName string) (io.Writer, error) {
	// file, err := os.Create(fileName + ".csv")
	// if err != nil {
	// 	return nil, err
	// }
	file, err := os.OpenFile(fileName+".csv", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func getValidateColumnIndex(columns []string) int {
	for colIndex, colVal := range columns {
		if colVal = strings.ToLower(colVal); colVal == Email {
			return colIndex
		}
	}
	return 0
}

func Validate(value string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(value)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
		os.Exit(0)
	}
}
