package response

import (
	"time"
)

type JadwalSholatResponse struct {
	ID        uint     `json:"id"`
	City      string   `json:"city"`
	Schedule  AdzanRow `json:"schedule"`
	Date      string   `json:"date"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
