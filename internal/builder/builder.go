package builder

import (
	"context"
	"dotenv-builder/internal/config"
	"dotenv-builder/internal/core"
	"dotenv-builder/internal/vault"
	vaultwrapper "dotenv-builder/internal/vaultWrapper"
	"log"
	"path/filepath"
	"strings"
)

func RunBuilder(cfg config.Config) {
	ctx := context.Background()

	// Проверяем существует ли файл .env
	core.CheckDotEnv(core.ENV_FILE)

	// Init Vault client
	vaultClient, err := vault.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Init KV v2 secrets engine
	kv2 := vaultClient.KVv2(cfg.VaultEnginesName)

	// Get list of secrets
	secretList := vaultwrapper.GetListSecrets(vaultClient, cfg.VaultPathWithMetaData)
	for _, v := range secretList {

		core.WriteHeaders(v)

		fullPath := filepath.Join(cfg.VaultSecretPath, v)

		vaultwrapper.WriteLatestVersion(kv2, ctx, fullPath)
		log.Println("read secret to path:", fullPath)

		if strings.Contains(v, "build") {
			vaultwrapper.GetOtherVersion(kv2, ctx, fullPath)
		}
	}
}
