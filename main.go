package main

import (
	"dotenv-builder/internal/core"
	"fmt"
)

func main() {
	envFile := core.ListEnv("credenticals")

	core.CheckDotEnv(core.ENV_FILE)

	for _, v := range envFile {

		core.WriteHeaders(v)

		secret, err := core.ReadFile(v)
		if err != nil {
			fmt.Println("Error reading file:", err)
		}
		secret.ReadJson(v)
		defer secret.Close()
	}

}
