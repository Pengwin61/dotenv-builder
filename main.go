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

		res := fmt.Sprintf("%s%s", *vaultSecretPath, v)
		writeLatestVersion(kv2, ctx, res)
		// printLatestVersion(kv2, ctx, res)
		// printTagretVersion(kv2, ctx, res, getCountOldVersion(kv2, ctx, res))

		writeTagretVersion(kv2, ctx, res, getCountOldVersion(kv2, ctx, res))
	}

}

func getCountOldVersion(kv2 *api.KVv2, ctx context.Context, path string) int {
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
		// log.Println("CURRENT VERSION:", array)
		// log.Println("OLD VERSION:", oldVersion)
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
