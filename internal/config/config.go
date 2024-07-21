package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	VaultServerAddr       string
	VaultEnginesName      string
	VaultSecretPath       string
	VaultPathWithMetaData string
}

func InitConfig(vaultSecretPath *string) Config {

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

	// Проверяем путь на префикс
	checkSecretPath(vaultSecretPath)

	// Формируем путь для метаданных
	vaultPathWithMetadata := getMetaPath(vaultSecretPath, &vaultEnginesName)

	return Config{vaultServerAddr, vaultEnginesName, *vaultSecretPath, vaultPathWithMetadata}
}

func getMetaPath(vaultSecretPath *string, vaultEnginesName *string) (vaultPathWithMetadata string) {

	return filepath.Join(*vaultEnginesName, "metadata", *vaultSecretPath)
}

func checkSecretPath(vaultSecretPath *string) {
	var res string
	if !strings.HasSuffix(*vaultSecretPath, "/") {
		res = fmt.Sprint(*vaultSecretPath + "/")
		*vaultSecretPath = res
	}
}
