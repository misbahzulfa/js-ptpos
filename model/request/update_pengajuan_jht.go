package request

type UpdatePengajuanJHTRequest struct {
	KodeSegmen            string `json:"KodeSegmen"`
	NPP                   string `json:"NPP"`
	KodeTK                string `json:"KodeTK"`
	KodeDivisi            string `json:"KodeDivisi"`
	KodePerusahaan        string `json:"KodePerusahaan"`
	KodeKepesertaan       string `json:"KodeKepesertaan"`
	NamaPerusahaan        string `json:"NamaPerusahaan"`
	TanggalAktif          string `json:"TanggalAktif"`
	TanggalNonAktif       string `json:"TanggalNonAktif"`
	KodeNA                string `json:"KodeNA"`
	KodeSebabNA           string `json:"KodeSebabNA"`
	KodeKantorKepesertaan string `json:"KodeKantorKepesertaan"`
	KodePembina           string `json:"KodePembina"`
	NominalSaldo          string `json:"NominalSaldo"`
}
