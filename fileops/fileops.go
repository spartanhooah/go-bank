package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("failed to find or read file")
	}

	value, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 0, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	os.WriteFile(fileName, []byte(fmt.Sprintf("%.2f", value)), 0644)
}
