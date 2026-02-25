package mapper

import (
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/response"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
)

func ToCityResponse(city *domain.City) *response.City {
	return &response.City{
		ID:   city.Id,
		Name: city.Name,
	}
}
