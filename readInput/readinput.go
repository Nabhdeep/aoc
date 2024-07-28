package readinput

import (
	"fmt"
	"os"
)

func ReadFilePath(path string) (string, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Printf("Error in Readfile %s", err)
		return "", err
	}
	return string(data), nil
}

func ReadFile(path string) string {
	data, err := ReadFilePath(path)
	if err != nil {
		fmt.Println(err)
	}

	return data
}
