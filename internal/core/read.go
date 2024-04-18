package core

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type JSONData struct {
// 	RequestID     string                 `json:"request_id"`
// 	LeaseID       string                 `json:"lease_id"`
// 	LeaseDuration int                    `json:"lease_duration"`
// 	Renewable     bool                   `json:"renewable"`
// 	Data          map[string]interface{} `json:"data"`
// 	Warnings      []string               `json:"warnings"`
// }
// type JSONDataList struct {
// 	Data map[string]interface{} `json:"data"`
// }

// type Secrets struct {
// 	File *os.File
// }

// func ReadFile(path string) (*Secrets, error) {

// 	var secretConfig *Secrets

// 	file, err := os.Open(path)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return nil, err
// 	}

// 	secretConfig = &Secrets{
// 		File: file,
// 	}

// 	return secretConfig, nil
// }

// func (s *Secrets) Close() {
// 	s.File.Close()
// }

// func (s *Secrets) ReadJson(path string) error {

// 	var jsonData JSONData
// 	err := json.NewDecoder(s.File).Decode(&jsonData)
// 	if err != nil {
// 		fmt.Println("Error decoding JSON:", err)
// 		return err
// 	}

// 	dataData := jsonData.Data["data"].(map[string]interface{})
// 	for key, value := range dataData {

// 		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s=%v\n", key, value))

// 	}
// 	return nil
// }

// func ReadList(path string) {
// 	file, err := os.ReadFile(path)
// 	if err != nil {
// 		fmt.Println("Error reading JSON file:", err)
// 		return
// 	}

// 	var data []string
// 	err = json.Unmarshal(file, &data)
// 	if err != nil {
// 		fmt.Println("Error parsing JSON:", err)
// 		return
// 	}
// 	fmt.Println(data)
// }
