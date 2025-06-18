package router

import (
	"ce/internal/config"
	"ce/internal/external/openexchangerates"
	exchangehandler "ce/internal/server/exchange"
	rateshandler "ce/internal/server/rates"
	exchangesvc "ce/internal/services/exchange"
	ratessvc "ce/internal/services/rates"

	"github.com/gin-gonic/gin"
)

func Setup(cfg config.Config) *gin.Engine {
	r := gin.Default()

	client := openexchangerates.NewClient(cfg.GetOpenExchangeAPIKey())

	ratesService := ratessvc.NewService(client)
	exchangeService := exchangesvc.NewService(cfg)

	ratesHandler := rateshandler.NewHandler(ratesService)
	exchangeHandler := exchangehandler.NewHandler(exchangeService)

	api := r.Group("/api")
	rateshandler.RegisterRoutes(api.Group("/rates"), ratesHandler)
	exchangehandler.RegisterRoutes(api.Group("/exchange"), exchangeHandler)

	return r
}
