package response

// type RincianManfaatBeasiswaResponse struct {
// 	ChId              string             `json:"chId"`
// 	ReqId             string             `json:"reqId"`
// 	KodeKlaim         string             `json:"KODE_KLAIM"`
// 	KodeManfaat       string             `json:"KODE_MANFAAT"`
// 	NoUrut            string             `json:"NO_URUT"`
// 	KodeTipePenerima  string             `json:"KODE_TIPE_PENERIMA"`
// 	NikPenerima       string             `json:"NIK_PENERIMA"`
// 	NamaPenerima      string             `json:"NAMA_PENERIMA"`
// 	TglLahirPenerima  string             `json:"TGLLAHIR_PENERIMA"`
// 	FlagMasihSekolah  string             `json:"FLAG_MASIH_SEKOLAH"`
// 	Keterangan        string             `json:"KETERANGAN"`
// 	TglPengajuan      string             `json:"TGL_PENGAJUAN"`
// 	KondisiAkhir      string             `json:"KONDISI_AKHIR"`
// 	TglKondisiAkhir   string             `json:"TGL_KONDISI_AKHIR"`
// 	FlagDitunda       string             `json:"FLAG_DITUNDA"`
// 	FlagDihentikan    string             `json:"FLAG_DIHENTIKAN"`
// 	FlagDiterima      string             `json:"FLAG_DITERIMA"`
// 	NomBiayaDisetujui string             `json:"NOM_BIAYA_DISETUJUI"`
// 	DataRincianDetil  []DataRincianDetil `json:"dataArrDetil"`
// }

// type DataRincianDetil struct {
// 	Tahun      string `json:"TAHUN"`
// 	Jenis      string `json:"JENIS"`
// 	Jenjang    string `json:"JENJANG"`
// 	FlagTerima string `json:"FLAG_TERIMA"`
// 	Tingkat    string `json:"TINGKAT"`
// 	Lembaga    string `json:"LEMBAGA"`
// 	Keterangan string `json:"KETERANGAN"`
// 	NomManfaat string `json:"NOM_MANFAAT"`
// }

type HitungManfaatBeasiswaResponse struct {
	Ret           string  `json:"ret"`
	Msg           string  `json:"msg"`
	PSukses       string  `json:"P_SUKSES"`
	PMess         string  `json:"P_MESS"`
	PNomDisetujui float32 `json:"P_NOM_DISETUJUI"`
}
