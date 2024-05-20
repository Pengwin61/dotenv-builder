package core

import (
	"fmt"
	"strings"
)

const ENV_FILE = ".env"

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
	case strings.Contains(path, "docker"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Docker\n"))
	case strings.Contains(path, "project"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "App\n"))
	case strings.Contains(path, "app"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "App\n"))
	case strings.Contains(path, "s3"):
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "S3\n"))
	default:
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s", "\n"))
		WriteFileEnv(ENV_FILE, fmt.Sprintf("%s %v\n", "##", "Other\n"))
	}
}
