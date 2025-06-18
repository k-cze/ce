package server

import (
	"ce/internal/config"
	"ce/internal/server/router"
)

func Run(cfg config.Config) {
	r := router.Setup(cfg)
	r.Run(":" + cfg.GetPort())
}
