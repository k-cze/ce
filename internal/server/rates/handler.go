package rates

import (
	"ce/internal/server/rates/request"
	"ce/internal/server/rates/response"
	ratesvc "ce/internal/services/rates"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service ratesvc.Service
}

func NewHandler(svc ratesvc.Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) GetRatesHandler(c *gin.Context) {
	req, err := request.New(request.RatesRequestWithCurrencies(c.Query("currencies")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	matrix, err := h.service.GetRates(req.Currencies)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	resp := response.New(response.RatesResponseWithMatrix(matrix))
	c.JSON(http.StatusOK, resp)
}
