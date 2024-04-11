package core

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const ENV_FILE = "env.custom"

var SecretFilePath string = "database.json"

type JSONData struct {
	RequestID     string                 `json:"request_id"`
	LeaseID       string                 `json:"lease_id"`
	LeaseDuration int                    `json:"lease_duration"`
	Renewable     bool                   `json:"renewable"`
	Data          map[string]interface{} `json:"data"`
	Warnings      []string               `json:"warnings"`
}

type Secrets struct {
	File *os.File
}

func GetEnvInFile(path string) map[string]string {
	var myEnv map[string]string

	reader, err := os.Open(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	myEnv, err = godotenv.Parse(reader)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return myEnv
}

func ReadFile(path string) (*Secrets, error) {

	var secretConfig *Secrets

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	secretConfig = &Secrets{
		File: file,
	}

	return secretConfig, nil
}

func (s *Secrets) Close() {
	s.File.Close()
}

func (s *Secrets) ReadJson(path string) error {

	var jsonData JSONData
	err := json.NewDecoder(s.File).Decode(&jsonData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return err
	}

	dataData := jsonData.Data["data"].(map[string]interface{})
	for key, value := range dataData {

		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s=%v\n", key, value))

	}
	return nil
}
