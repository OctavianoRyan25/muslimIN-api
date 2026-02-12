package request

import "time"

type Doa struct {
	Nama          string    `json:"nama"`
	Lafal         string    `json:"lafal"`
	Transliterasi string    `json:"transliterasi"`
	Arti          string    `json:"arti"`
	Riwayat       string    `json:"riwayat"`
	Keterangan    string    `json:"keterangan"`
	KataKunci     []string  `json:"kata_kunci"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}
