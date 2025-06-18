package exchange

import (
	"ce/internal/server/exchange/request"
	"ce/internal/server/exchange/response"
	exchangeservice "ce/internal/services/exchange"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service exchangeservice.Service
}

func NewHandler(svc exchangeservice.Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) GetExchangeHandler(c *gin.Context) {
	req, err := request.New(
		request.ExchangeRequestWithFrom(c.Query("from")),
		request.ExchangeRequestWithTo(c.Query("to")),
		request.ExchangeRequestWithAmount(c.Query("amount")),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	amount, err := h.service.GetExchange(req.From, req.To, float64(req.Amount))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	resp := response.New(
		response.ExchangeResponseWithFrom(req.From),
		response.ExchangeResponseWithTo(req.To),
		response.ExchangeResponseWithAmount(*amount),
	)
	c.JSON(http.StatusOK, resp)
}
