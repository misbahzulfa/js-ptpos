package response

type UpdateKonfirmasiJPECtoPNResponse struct {
	StatusKirim              string `json:"statusKirim"`
	BLTHManfaatBulanBerjalan string `json:"blthManfaatBulanBerjalan"`
	JumlahBulanKeterlambatan string `json:"jumlahBulanKeterlambatan"`
	JumlahBulanRapel         string `json:"jumlahBulanRapel"`
	NilaiRapel               string `json:"nilaiRapel"`
	NilaiKompensasi          string `json:"nilaiKompensasi"`
	NilaiManfaat             string `json:"nilaiMafaat"`
	NilaiTOTAL               string `json:"nilaiTotal"`
	StatusPembayaran         string `json:"statusPembayaran"`
	JenisNotifikasi          string `json:"jenisNotifikasi"`
}
