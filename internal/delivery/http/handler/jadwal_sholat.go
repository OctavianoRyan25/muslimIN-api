package handler

import (
	"net/http"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/mapper"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/usecase"
	"github.com/gin-gonic/gin"
)

type JadwalSholatHandler struct {
	useCase usecase.JadwalSholatUseCase
}

func NewJadwalSholatHandler(useCase usecase.JadwalSholatUseCase) *JadwalSholatHandler {
	return &JadwalSholatHandler{useCase: useCase}
}

func (h *JadwalSholatHandler) GetToday(c *gin.Context) {
	jadwalSholat, err := h.useCase.GetToday()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := mapper.ToJadwalSholatResponse(*jadwalSholat)

	c.JSON(http.StatusOK, gin.H{"Data": res})
}
