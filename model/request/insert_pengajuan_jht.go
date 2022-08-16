package request

type InsertPengajuanJHTRequest struct {
	Nama                string             `json:"Nama"`
	NIK                 string             `json:"NIK"`
	KPJ                 string             `json:"KPJ"`
	Email               string             `json:"Email"`
	NoHP                string             `json:"NoHP"`
	TempatLahir         string             `json:"TempatLahir"`
	TanggalLahir        string             `json:"TanggalLahir"`
	NamaIbuKandung      string             `json:"NamaIbuKandung"`
	Alamat              string             `json:"Alamat"`
	KodePos             string             `json:"KodePos"`
	KodeKelurahan       int                `json:"KodeKelurahan"`
	KodeKecamatan       int                `json:"KodeKecamatan"`
	KodeKabupaten       int                `json:"KodeKabupaten"`
	KodeProvinsi        int                `json:"KodeProvinsi"`
	KodeBank            string             `json:"KodeBank"`
	NomorRekening       string             `json:"NomorRekening"`
	NamaRekening        string             `json:"NamaRekening"`
	NPWP                string             `json:"NPWP"`
	Dokumen             []DokumenDetil     `json:"Dokumen"`
	TanggalPengajuan    string             `json:"TanggalPengajuan"`
	KodeBillingPTPos    string             `json:"KodeBillingPTPos"`
	ScoreFaceMatch      string             `json:"ScoreFaceMatch"`
	ScoreSimilarityNama string             `json:"ScoreSimilarityNama"`
	JawabanProbing      string             `json:"JawabanProbing"`
	DataProbing         []DataProbingDetil `json:"DataProbing"`
	KeteranganPengajuan string             `json:"KeteranganPengajuan"`
	ChannelId           string             `json:"ChannelId"`
	KodeKantorPPTOS     string             `json:"KodeKantorPTPOS"`
	PetugasRekamPTPOS   string             `json:"PetugasPTPOS"`
	ChId                string             `json:"chId"`
	ReqId               string             `json:"reqId"`
}

type DokumenDetil struct {
	KodeDokumen    string `json:"KodeDokumen"`
	PathUrlDokumen string `json:"PathUrlDokumen"`
}

type DataProbingDetil struct {
	KodeProbing     string `json:"KodeProbing"`
	NomorUrut       string `json:"NomorUrut"`
	ResponseProbing string `json:"ResponseProbing"`
	JawabanProbing  string `json:"JawabanProbing"`
}
