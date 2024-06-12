package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	VaultServerAddr       string
	VaultEnginesName      string
	VaultSecretPath       string
	VaultPathWithMetaData string
}

func InitConfig() Config {

	vaultServerAddr := os.Getenv("VAULT_ADDR")
	if vaultServerAddr == "" {
		log.Fatal("VAULT_ADDR is empty")
	}

	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		log.Fatal("VAULT_TOKEN is empty")
	}

	// Получаем название engines из env
	vaultEnginesName := os.Getenv("VAULT_SECRETS_ENGINES_NAME")
	if vaultEnginesName == "" {
		log.Fatal("VAULT_SECRETS_ENGINES_NAME is empty")
	}

	// Получаем путь к секретам из флага
	vaultSecretPath := flag.String("path", "projects/", "path to secrets")
	if *vaultSecretPath == "" {
		log.Fatal("path is empty, i can`t get secrets without path")
	}
	flag.Parse()

	// Проверяем путь на префикс
	checkSecretPath(vaultSecretPath)

	// Формируем путь для метаданных
	vaultPathWithMetadata := vaultEnginesName + "/" + "metadata" + "/" + *vaultSecretPath

	return Config{vaultServerAddr, vaultEnginesName, *vaultSecretPath, vaultPathWithMetadata}
}

func checkSecretPath(vaultSecretPath *string) {
	var res string
	if !strings.HasSuffix(*vaultSecretPath, "/") {
		res = fmt.Sprint(*vaultSecretPath + "/")
		*vaultSecretPath = res
	}
}
