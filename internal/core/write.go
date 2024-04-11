package core

import (
	"log"
	"os"
)

func WriteFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func CheckDotEnv(path string) {
	_, err := os.Stat(path)

	if !os.IsNotExist(err) {
		os.Remove(path)
	}
}

func WriteFileEnv(path string, data string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
}
