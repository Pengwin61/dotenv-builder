package main

import (
	"context"
	"dotenv-builder/internal/core"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/vault-client-go"
)

func main() {

	// Получаем путь к секретам из флага
	vaultSecretPath := flag.String("path", "projects/", "path to secrets")
	if *vaultSecretPath == "" {
		log.Fatal("path is empty, i can`t get secrets without path")
	}
	flag.Parse()

	// Проверяем существует ли файл .env
	core.CheckDotEnv(core.ENV_FILE)

	ctx := context.Background()
	vault_addr := os.Getenv("VAULT_ADDR")
	vaultToken := os.Getenv("VAULT_TOKEN")
	// vaultSecretPath := os.Getenv("VAULT_SECRET_PATH")

	// создаем клиента vault
	client, err := vault.New(
		vault.WithAddress(vault_addr),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	// авторизуемся через токен
	if err := client.SetToken(vaultToken); err != nil {
		log.Fatal(err)
	}

	// получае список секретов
	list, err := client.Secrets.KvV2List(ctx, *vaultSecretPath, vault.WithMountPath("secret"))
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range list.Data.Keys {

		secretPath := fmt.Sprint(*vaultSecretPath, "/", v)

		core.WriteHeaders(v)

		// читаем секреты
		s, err := client.Secrets.KvV2Read(ctx, secretPath, vault.WithMountPath("secret"))
		if err != nil {
			log.Fatal(err)
		}

		// записываем в файл .env
		for k, v := range s.Data.Data {
			core.WriteFileEnv(core.ENV_FILE, fmt.Sprintf("%s=%v\n", k, v))

			// fmt.Printf("%s=%v\n", k, v)
		}
	}
}
