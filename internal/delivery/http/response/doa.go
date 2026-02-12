package response

import "time"

type Doa struct {
	ID            uint      `json:"id"`
	Nama          string    `json:"nama"`
	Lafal         string    `json:"lafal"`
	Transliterasi string    `json:"transliterasi"`
	Arti          string    `json:"arti"`
	Riwayat       string    `json:"riwayat"`
	Keterangan    *string   `json:"keterangan"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
