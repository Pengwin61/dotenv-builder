# Dotenv-Builder


# About
Получение секретов из Hashicorp Vault

# Run 
dotenv-builder --path="path/to/secrets"

`go run main.go --path="projects/$CI_PROJECT_ROOT_NAMESPACE/$CI_PROJECT_NAME/$CI_ENVIRONMENT_NAME"`

# For testing
```
export VAULT_ADDR="https://vault.domain.com"
export VAULT_TOKEN="TOKEN"
export VAULT_SECRETS_ENGINES_NAME="secret"

export CI_ENVIRONMENT_NAME=development
export CI_PROJECT_ROOT_NAMESPACE=prj
export CI_PROJECT_NAME=prj1
```