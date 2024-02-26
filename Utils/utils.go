package utils

import (
	"fmt"
	"os"
)

func WriteFile(filepath string) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	fmt.Println(string(data))
}
