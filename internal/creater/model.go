package creater

func buildMain(pathSplit []string) (map[string]interface{}, map[string]interface{}) {

	var keysBuild = map[string]interface{}{
		"APP_IMAGE":     "null",
		"APP_STATE":     pathSplit[3],
		"PROJECT_GROUP": pathSplit[1],
		"PROJECT_NAME":  pathSplit[2],
	}

	var keysDocker = map[string]interface{}{
		"EXT_NGINX_PORT": "80",
	}
	return keysBuild, keysDocker
}

func buildBackend(pathSplit []string) (map[string]interface{}, map[string]interface{}, map[string]interface{}) {

	var keysDatabase = map[string]interface{}{
		"DB_HOST":     "db",
		"DB_NAME":     "postgres",
		"DB_PASSWORD": "postgres",
		"DB_PORT":     "5432",
		"DB_USER":     "postgres",
	}

	var keysRedis = map[string]interface{}{
		"REDIS_CACHE_URL": "null",
	}

	var keysDjango = map[string]interface{}{
		"DEBUG":                  "True",
		"ALLOWED_HOSTS":          "null",
		"SETTINGS_CONFIGURATION": pathSplit[3],
	}

	return keysDatabase, keysRedis, keysDjango
}

func buildFrontend() map[string]interface{} {

	var keysFront = map[string]interface{}{
		"API_URL": "null",
	}
	return keysFront
}
