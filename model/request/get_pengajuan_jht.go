package request

type GetPengajuanJHTRequest struct {
	KodePengajuan string `json:"KodePengajuan"`
	NIK           string `json:"NIK"`
	NamaLengkap   string `json:"NamaLengkap"`
	KPJ           string `json:"KPJ"`
}
