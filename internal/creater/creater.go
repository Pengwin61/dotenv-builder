package creater

import (
	"context"
	"dotenv-builder/internal/config"
	"dotenv-builder/internal/vault"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/hashicorp/vault/api"
)

var (
	arraySecrets  [2]string = [2]string{"build", "docker"}
	araayBackend  [3]string = [3]string{"django", "database", "redis"}
	arrayFrontend [1]string = [1]string{"app"}
)

func RunCreater(cfg config.Config, languages string) {
	pathSplit := strings.Split(cfg.VaultSecretPath, "/")

	keysBuild, keysDocker := buildMain(pathSplit)
	keysBackend, keysRedis, keysDjango := buildBackend(pathSplit)
	keysFront := buildFrontend()

	ctx := context.Background()

	// Init Vault client
	vaultClient, err := vault.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Init KV v2 secrets engine
	kv2 := vaultClient.KVv2(cfg.VaultEnginesName)

	writeBaseKeys(arraySecrets, cfg, ctx, kv2, keysDocker, keysBuild)
	writeBackendKeys(araayBackend, cfg, ctx, kv2, keysDjango, keysBackend, keysRedis, languages)
	writeFrontendKeys(arrayFrontend, cfg, ctx, kv2, keysFront, languages)

}

func writeBaseKeys(arraySecrets [2]string,
	cfg config.Config, ctx context.Context, kv2 *api.KVv2,
	keysDocker map[string]interface{}, keysBuild map[string]interface{}) {

	for _, v := range arraySecrets {
		fullPath := filepath.Join(cfg.VaultSecretPath, v)
		sec, err := kv2.Get(ctx, fullPath)
		if err != nil {
			log.Println(err)
		}
		if sec == nil {
			switch v {
			case "docker":
				_, err := kv2.Put(ctx, fullPath, keysDocker)
				if err != nil {
					fmt.Println("i cant create secret", err)
				}
			case "build":
				_, err := kv2.Put(ctx, fullPath, keysBuild)
				if err != nil {
					fmt.Println("i cant create secret", err)
				}
			}
		}
	}
}

func writeBackendKeys(araayBackend [3]string,
	cfg config.Config, ctx context.Context, kv2 *api.KVv2,
	keysDjango map[string]interface{}, keysBackend map[string]interface{},
	keysRedis map[string]interface{}, languages string) {

	if strings.Contains(languages, "python") {
		for _, v := range araayBackend {
			fullPath := filepath.Join(cfg.VaultSecretPath, v)
			sec, err := kv2.Get(ctx, fullPath)
			if err != nil {
				log.Println(err)
			}
			if sec == nil {
				switch v {
				case "django":
					_, err := kv2.Put(ctx, fullPath, keysDjango)
					if err != nil {
						fmt.Println("i cant create secret", err)
					}
				case "database":
					_, err := kv2.Put(ctx, fullPath, keysBackend)
					if err != nil {
						fmt.Println("i cant create secret", err)
					}
				case "redis":
					_, err := kv2.Put(ctx, fullPath, keysRedis)
					if err != nil {
						fmt.Println("i cant create secret", err)
					}
				}
			}
		}
	}
}

func writeFrontendKeys(arrayFrontend [1]string,
	cfg config.Config, ctx context.Context, kv2 *api.KVv2,
	keysFront map[string]interface{}, languages string) {

	if strings.Contains(languages, "javascript") {
		for _, v := range arrayFrontend {
			fullPath := filepath.Join(cfg.VaultSecretPath, v)
			sec, err := kv2.Get(ctx, fullPath)
			if err != nil {
				log.Println(err)
			}
			if sec == nil {
				switch v {
				case "app":
					_, err := kv2.Put(ctx, fullPath, keysFront)
					if err != nil {
						fmt.Println("i cant create secret", err)
					}
				}
			}
		}
	}
}
