package main

import (
	"context"
	"dotenv-builder/internal/vault"
	"fmt"

	"log"

	"github.com/hashicorp/vault/api"
)

func main() {
	// ctx := context.Background()
	vaultEnginesName := "secret"
	metaPath := "metadata"
	secretPath := "test/prj/prj1/development/"

	fullPath := vaultEnginesName + "/" + metaPath + "/" + secretPath

	// vault_addr := os.Getenv("VAULT_ADDR")
	// vaultToken := os.Getenv("VAULT_TOKEN")
	// vaultEnginesName := os.Getenv("VAULT_SECRETS_ENGINES_NAME")

	vaultClient, err := vault.NewClient()
	if err != nil {
		panic(err)
	}
	//
	// sec := vaultClient.KVv2(vaultEnginesName)

	//
	liiis := getListSecrets(vaultClient, fullPath)

	for _, v := range liiis {
		fmt.Println(v)
		// res := fmt.Sprintf("%s%s", "secret/test/prj/prj1/development", v)
		// fmt.Println(res)
		// targetVerson := getCurrentVersion(sec, ctx, res)

		// fmt.Println(targetVerson)
	}

	//

	// vaultClient.SetToken(os.Getenv("VAULT_TOKEN"))

	// targetVerson := getOldVersion(sec, ctx, path)
	// getCurrentVersion(sec, ctx, path)
	// getTagretVersion(sec, ctx, path, targetVerson)

}

func getOldVersion(sec *api.KVv2, ctx context.Context, path string) int {
	list, err := sec.GetVersionsAsList(ctx, path)
	if err != nil {
		fmt.Println(err)
	}
	oldVersion := len(list) - 1
	fmt.Println("OLD VERSION:", oldVersion)

	return oldVersion
}

func getCurrentVersion(sec *api.KVv2, ctx context.Context, path string) {
	key, err := sec.Get(ctx, path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("---")
	fmt.Println("CURRENT DATA")
	for k, v := range key.Data {
		fmt.Println(k, v)
	}
}

func getTagretVersion(sec *api.KVv2, ctx context.Context, path string, version int) {
	old, err := sec.GetVersion(ctx, path, version)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---")
	fmt.Println("VERSION:", version)
	for k, v := range old.Data {
		fmt.Println(k, v)
	}
	fmt.Println("---")
}

func getListSecrets(vaultClient *api.Client, path string) []string {

	var array []string
	list, err := vaultClient.Logical().List(path)
	if err != nil {
		log.Println(err)
	}

	for _, v := range list.Data {
		a := fmt.Sprintln(v)
		array = append(array, a)
	}

	return array
}
