# Dotenv-Builder


# About
Получение секретов из Hashicorp Vault.

APP_IMAGE - Имя doсker image:TAG
OLD_APP_IMAGE - Пердыдущий TAG

# Run 
dotenv-builder create env --path="path/to/secrets"
dotenv-builder create structure --path="path/to/secrets"



`go run main.go create env --path="projects/$CI_PROJECT_ROOT_NAMESPACE/$CI_PROJECT_NAME/$CI_ENVIRONMENT_NAME"`

# For testing
```
export VAULT_ADDR="https://vault.domain.com"
export VAULT_TOKEN="TOKEN"
export VAULT_SECRETS_ENGINES_NAME="secret"

export CI_ENVIRONMENT_NAME=development
export CI_PROJECT_ROOT_NAMESPACE=infrastructure
export CI_PROJECT_NAME=test-pipline
```