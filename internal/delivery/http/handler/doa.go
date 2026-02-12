package handler

import (
	"net/http"
	"strconv"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/mapper"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/response"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/usecase"
	"github.com/gin-gonic/gin"
)

type DoaHandler struct {
	useCase usecase.DoaUseCase
}

func NewDoaUsecase(useCase usecase.DoaUseCase) *DoaHandler {
	return &DoaHandler{useCase: useCase}
}

func (h *DoaHandler) GetAll(c *gin.Context) {
	var res []response.Doa
	doas, err := h.useCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, v := range doas {
		res = append(res, *mapper.ToDoaResponse(&v))
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *DoaHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	param, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.useCase.GetById(uint(param))
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *DoaHandler) GetRandom(c *gin.Context) {
	doa, err := h.useCase.GetRandom()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	res := mapper.ToDoaResponse(doa)
	c.JSON(200, gin.H{
		"data": res,
	})
}
