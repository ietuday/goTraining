package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)



func WriteFloatToFile(value float64, filename string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644) // Owner Read/write and Others Read only

}

func GetFloatFromFile(filename string) (float64, error) {
	data, error := os.ReadFile(filename)

	if error != nil {
		return 1000, errors.New("Failed to find file")
	}

	valueText := string(data)
	value, error := strconv.ParseFloat(valueText, 64)

	if error != nil {
		return 1000, errors.New("Failed to parse stored value")
	}

	return value, nil
}
