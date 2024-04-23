package core

import (
	"fmt"
	"strings"
)

const ENV_FILE = ".env"

// func ListEnv(folderPath string) []string {
// 	var listEnv []string

// 	files, err := os.ReadDir(folderPath)
// 	if err != nil {
// 		fmt.Println("Error reading directory:", err)
// 		return nil
// 	}

// 	for _, file := range files {

// 		if file.Name() == "server.json" {
// 			break
// 		}
// 		if file.Name() == "list.json" {
// 			break
// 		}

// 		listEnv = append(listEnv, fmt.Sprintf("%s/%s", folderPath, file.Name()))
// 	}
// 	return listEnv
// }

func WriteHeaders(path string) {

	switch {
	case strings.Contains(path, "database"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Database\n"))
	case strings.Contains(path, "django"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Django\n"))
	case strings.Contains(path, "redis") || strings.Contains(path, "cache"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Cache\n"))
	case strings.Contains(path, "celery"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Celery\n"))
	case strings.Contains(path, "build"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Build\n"))
	case strings.Contains(path, "4payments"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Four Payments\n"))
	case strings.Contains(path, "payler"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "PAYLER\n"))
	default:
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Other\n"))
	}
}
