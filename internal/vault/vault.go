package vault

import (
	"os"

	"github.com/hashicorp/vault/api"
)

func NewClient() (*api.Client, error) {
	vaultConfig := api.DefaultConfig()
	vaultConfig.Address = os.Getenv("VAULT_ADDR")
	vaultConfig.MaxRetries = 10
	return api.NewClient(vaultConfig)
}
