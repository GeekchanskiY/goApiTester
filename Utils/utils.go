package utils

import (
	"log"
	"os"
)

func WriteFile(filepath string, data string) (bool, error) {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}

func ReadFile(filepath string) string {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Println("File reading error", err)
	}
	return string(data)
}
