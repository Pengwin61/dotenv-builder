# Dotenv-Builder


# About
В первой итерации билдер парсит секреты из файла `.json` который были сохранены из Hashicorp Vault CLI 
example:
`vault kv get -format=json  secret/projects/test-project/django > credenticals/
celery.json` 
собирает из разных файлов `credenticals/*.json` в единый файл `.env`

# Roadmap

* Добавить поддержку Vault нативно не из файлов.
