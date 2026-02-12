package mapper

import (
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/response"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
)

func ToDoaResponse(doa *domain.Doa) *response.Doa {
	return &response.Doa{
		ID:            doa.ID,
		Nama:          doa.Nama,
		Lafal:         doa.Lafal,
		Transliterasi: doa.Transliterasi,
		Arti:          doa.Arti,
		Riwayat:       doa.Riwayat,
		Keterangan:    doa.Keterangan,
		CreatedAt:     doa.CreatedAt,
		UpdatedAt:     doa.UpdatedAt,
	}
}
