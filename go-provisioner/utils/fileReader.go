package utils

import (
	"bufio"
	"os"
)

func GetVariable(folderName string, name string) (string, error) {
	file, err := os.Open(folderName + "/" + name)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", scanner.Err()
}
