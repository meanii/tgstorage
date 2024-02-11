package main

import (
	"github.com/meanii/tgstorage/configs"
	"github.com/meanii/tgstorage/server"
)

func main() {
	configs.InitEnvConfigs()
	server.Server()
}
