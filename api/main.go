package main

import (
	"amsvault/api/config"
	"amsvault/api/server"
	"fmt"
)

func main() {
	cfg := config.Get()

	err := server.Instance().Run(cfg)
	endAsError(err, "error starting server: ")

}

func endAsError(err error, messaage string) {
	if err != nil {
		fmt.Println(messaage, err)
	}
}
