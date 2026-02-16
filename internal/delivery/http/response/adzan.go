package response

type AdzanRow struct {
	Tanggal string `json:"tanggal"`
	Imsyak  string `json:"imsyak"`
	Shubuh  string `json:"shubuh"`
	Terbit  string `json:"terbit"`
	Dhuha   string `json:"dhuha"`
	Dzuhur  string `json:"dzuhur"`
	Ashr    string `json:"ashr"`
	Magrib  string `json:"magrib"`
	Isya    string `json:"isya"`
}
