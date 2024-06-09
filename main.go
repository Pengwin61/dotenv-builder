package main

import (
	"context"
	"dotenv-builder/internal/core"
	"dotenv-builder/internal/vault"
	"flag"
	"fmt"
	"os"
	"strings"

	"log"

	"github.com/hashicorp/vault/api"
)

func main() {
	ctx := context.Background()

	// Получаем название engines из env
	vaultEnginesName := os.Getenv("VAULT_SECRETS_ENGINES_NAME")

	// Получаем путь к секретам из флага
	vaultSecretPath := flag.String("path", "projects/", "path to secrets")
	if *vaultSecretPath == "" {
		log.Fatal("path is empty, i can`t get secrets without path")
	}
	flag.Parse()

	// Проверяем путь на префикс
	checkSecretPath(vaultSecretPath)

	// Формируем путь для метаданных
	fullPath := vaultEnginesName + "/" + "metadata" + "/" + *vaultSecretPath

	// Проверяем существует ли файл .env
	core.CheckDotEnv(core.ENV_FILE)

	// Init Vault client
	vaultClient, err := vault.NewClient()
	if err != nil {
		panic(err)
	}

	// Init KV v2 secrets engine
	kv2 := vaultClient.KVv2(vaultEnginesName)

	// Get list of secrets
	secretList := getListSecrets(vaultClient, fullPath)

	for _, v := range secretList {

		core.WriteHeaders(v)

		fullPath := fmt.Sprintf("%s%s", *vaultSecretPath, v)

		writeLatestVersion(kv2, ctx, fullPath)
		// printLatestVersion(kv2, ctx, fullPath)
		// printTagretVersion(kv2, ctx, fullPath, getIndexOldVersion(kv2, ctx, fullPath))

		if strings.Contains(v, "build") {
			getOtherVersion(kv2, ctx, fullPath)
		}
	}
}

func getIndexOldVersion(kv2 *api.KVv2, ctx context.Context, path string) int {
	array := make([]int, 0)
	var oldVersion int

	list, err := kv2.GetVersionsAsList(ctx, path)
	if err != nil {
		log.Println(err)
	}

	if strings.Contains(path, "build") {
		for _, v := range list {
			array = append(array, v.Version)
		}
		oldVersion = array[len(array)-2]
	}

	return oldVersion
}

func getListSecrets(vaultClient *api.Client, path string) []string {
	array := make([]string, 0)

	list, err := vaultClient.Logical().List(path)
	if err != nil {
		log.Println(err)
	}

	for _, key := range list.Data["keys"].([]interface{}) {
		str, ok := key.(string)
		if ok {
			array = append(array, str)
		} else {
			log.Println("Ошибка при конвертации ключа в строку")
		}
	}

	return array
}

func writeLatestVersion(kv2 *api.KVv2, ctx context.Context, path string) {

	key, err := kv2.Get(ctx, path)
	if err != nil {
		log.Println(err)
	}

	for k, v := range key.Data {
		core.WriteFileEnv(core.ENV_FILE, fmt.Sprintf("%s=%v\n", k, v))
	}
}

func getSecret(kv2 *api.KVv2, ctx context.Context, path string) string {
	var currentVersion string

	key, err := kv2.Get(ctx, path)
	if err != nil {
		log.Println(err)
	}
	for k, v := range key.Data {
		if k == "APP_IMAGE" {
			currentVersion = fmt.Sprintf("%v", v)
		}
	}
	return currentVersion
}

func getSecretVersion(kv2 *api.KVv2, ctx context.Context, path string, version int) string {
	var secret string
	old, err := kv2.GetVersion(ctx, path, version)
	if err != nil {
		log.Println(err)
	}
	for k, v := range old.Data {
		if k == "APP_IMAGE" {
			secret = fmt.Sprintf("%v", v)
		}
	}
	return secret
}

func getOtherVersion(kv2 *api.KVv2, ctx context.Context, fullPath string) {
	var indexOldVersion int
	current := getSecret(kv2, ctx, fullPath)
	indexOldVersion = getIndexOldVersion(kv2, ctx, fullPath)
	old := getSecretVersion(kv2, ctx, fullPath, indexOldVersion)

	if current == old {
		for indexOldVersion := indexOldVersion; indexOldVersion >= 1; indexOldVersion-- {
			old := getSecretVersion(kv2, ctx, fullPath, indexOldVersion)
			fmt.Println("CURRENT:", current, "OLD:", old, "DIFF:", current != old, "VERSION:", indexOldVersion)
			if current != old {
				writeTagretVersion(kv2, ctx, fullPath, indexOldVersion)
				break
			}
		}
	} else {
		writeTagretVersion(kv2, ctx, fullPath, indexOldVersion)
	}
}

func writeTagretVersion(kv2 *api.KVv2, ctx context.Context, path string, version int) {
	old, err := kv2.GetVersion(ctx, path, version)
	if err != nil {
		log.Println(err)
	}

	for k, v := range old.Data {
		if k == "APP_IMAGE" {
			core.WriteFileEnv(core.ENV_FILE, fmt.Sprintf("%s=%v\n", "OLD_APP_IMAGE", v))
		}
	}
}

func checkSecretPath(vaultSecretPath *string) {
	var res string
	if !strings.HasSuffix(*vaultSecretPath, "/") {
		res = fmt.Sprint(*vaultSecretPath + "/")
		*vaultSecretPath = res
	}
}
