package request

type UpdateKonfirmasiJPPNtoECRequest struct {
	StatusKirim              string `json:"statusKirim"`
	BLTHManfaatBulanBerjalan string `json:"blthManfaatBulanBerjalan"`
	JumlahBulanKeterlambatan string `json:"jumlahBulanKeterlambatan"`
	JumlahBulanRapel         string `json:"jumlahBulanRapel"`
	NilaiRapel               string `json:"nilaiRapel"`
	NilaiKompensasi          string `json:"nilaiKompensasi"`
	NilaiManfaat             string `json:"nilaiManfaat"`
	NilaiTOTAL               string `json:"nilaiTotal"`
	StatusPembayaran         string `json:"statusPembayaran"`
	JenisNotifikasi          string `json:"jenisNotifikasi"`
	KodeKlaim                string `json:"kodeKlaim"`
	KodeKlaimInduk           string `json:"kodeKlaimInduk"`
	StatusPengajuan          string `json:"statusPengajuan"`
}
