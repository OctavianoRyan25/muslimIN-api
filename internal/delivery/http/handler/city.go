package handler

import (
	"net/http"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/mapper"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/response"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/usecase"
	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	useCase usecase.CityUseCase
}

func NewCityHandler(useCase usecase.CityUseCase) *CityHandler {
	return &CityHandler{useCase: useCase}
}

func (h *CityHandler) GetAllCity(c *gin.Context) {
	var res []response.City
	ctx := c.Request.Context()
	cities, err := h.useCase.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, v := range cities {
		res = append(res, *mapper.ToCityResponse(&v))
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}
