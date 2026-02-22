package mapper

import (
	"encoding/json"
	"log"
	"time"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/request"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/response"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
)

func ToJadwalSholatResponse(domain domain.JadwalSholat) *response.JadwalSholatResponse {
	var details response.AdzanRow

	err := json.Unmarshal(domain.Schedule, &details)

	if err != nil {
		log.Printf("Error unmarshaling schedule for ID %d: %v", domain.ID, err)
		return &response.JadwalSholatResponse{}
	}

	return &response.JadwalSholatResponse{
		ID:       domain.ID,
		City:     domain.City,
		Date:     domain.Date.Format("2006-01-02"),
		Schedule: details,
	}
}

func ToJadwalSholatDomain(request request.JadwalSholatRequest) *domain.JadwalSholat {
	date, _ := time.Parse("2006-01-02", request.Date)
	return &domain.JadwalSholat{
		City: request.City,
		Date: date,
	}
}
