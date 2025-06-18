package main

import (
	"ce/internal/config"
	"ce/internal/server"
)

func main() {
	cfg := config.New()
	server.Run(cfg)
}
