package core

import (
	"fmt"
	"os"
	"strings"
)

func ListEnv(folderPath string) []string {
	var listEnv []string

	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	for _, file := range files {
		listEnv = append(listEnv, fmt.Sprintf("%s/%s", folderPath, file.Name()))
	}
	return listEnv
}

func WriteHeaders(path string) {

	switch {
	case strings.Contains(path, "server"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Server\n"))
	case strings.Contains(path, "database"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Database\n"))
	case strings.Contains(path, "django"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Django\n"))
	case strings.Contains(path, "redis"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Redis\n"))
	case strings.Contains(path, "celery"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Celery\n"))
	case strings.Contains(path, "build"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Build\n"))
	default:
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Other\n"))
	}
}
