package handler

import (
	"fmt"
	"net/http"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/mapper"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/request"
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
	city := c.Param("city")
	if city == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Param needed"})
		return
	}
	jadwalSholat, err := h.useCase.GetToday(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := mapper.ToJadwalSholatResponse(*jadwalSholat)

	c.JSON(http.StatusOK, gin.H{"Data": res})
}

func (h *JadwalSholatHandler) GetByDate(c *gin.Context) {
	var jadwalReq request.JadwalSholatRequest
	err := c.ShouldBindJSON(&jadwalReq)
	fmt.Println(jadwalReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	jadwalSholat := mapper.ToJadwalSholatDomain(jadwalReq)
	jadwal, err := h.useCase.GetByDate(jadwalSholat.City, jadwalSholat.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := mapper.ToJadwalSholatResponse(*jadwal)

	c.JSON(http.StatusOK, gin.H{"Data": res})
}
