package main

import "github.com.br/GregoryLacerda/AMSVault/configs"

func main() {

	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	println(cfg.DBName)
}
