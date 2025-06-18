package rates

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"ce/internal/server/rates/response"
	ratesvc "ce/internal/services/rates"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
	response ratesvc.ExchangeMatrix
	err      error
}

func (m *mockService) GetRates(currencies []string) (ratesvc.ExchangeMatrix, error) {
	return m.response, m.err
}

func TestGetRatesHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mock := &mockService{
		response: ratesvc.ExchangeMatrix{
			"USD": {"EUR": 0.85},
			"EUR": {"USD": 1.18},
		},
	}

	handler := NewHandler(mock)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodGet, "/api/rates?currencies=USD,EUR", nil)
	c.Request = req

	handler.GetRatesHandler(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var actual []*response.RateEntry
	err := json.Unmarshal(w.Body.Bytes(), &actual)
	assert.NoError(t, err)

	expected := []*response.RateEntry{
		{From: "USD", To: "EUR", Rate: 0.85},
		{From: "EUR", To: "USD", Rate: 1.18},
	}

	assert.ElementsMatch(t, expected, actual)
}

func TestGetRatesHandler_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mock := &mockService{
		err: errors.New("simulated error"),
	}

	handler := NewHandler(mock)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodGet, "/api/rates?currencies=USD,EUR", nil)
	c.Request = req

	handler.GetRatesHandler(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{}", w.Body.String())
}

func TestGetRatesHandler_MissingParam(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mock := &mockService{}

	handler := NewHandler(mock)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodGet, "/api/rates", nil)
	c.Request = req

	handler.GetRatesHandler(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "{}", w.Body.String())
}

func TestGetRatesHandler_InvalidParamCount(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mock := &mockService{}

	handler := NewHandler(mock)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodGet, "/api/rates?currencies=USD", nil)
	c.Request = req

	handler.GetRatesHandler(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "{}", w.Body.String())
}

func TestGetRatesHandler_EmptyResult(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mock := &mockService{
		response: ratesvc.ExchangeMatrix{},
	}

	handler := NewHandler(mock)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodGet, "/api/rates?currencies=USD,EUR", nil)
	c.Request = req

	handler.GetRatesHandler(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var actual []*response.RateEntry
	err := json.Unmarshal(w.Body.Bytes(), &actual)
	assert.NoError(t, err)
	assert.Empty(t, actual)
}
