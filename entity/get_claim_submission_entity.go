package entity

type GetClaimSubmissionEntity struct {
	KodePengajuan          string `json:"kodePengajuan"`
	Email                  string `json:"email"`
	NoHp                   string `json:"noHp"`
	KodeTk                 string `json:"kodeTk"`
	Kpj                    string `json:"kpj"`
	NomorIdentitas         string `json:"nomorIdentitas"`
	JenisIdentitas         string `json:"jenisIdentitas"`
	StatusValidIdentitas   string `json:"statusValidIdentitas"`
	NamaTk                 string `json:"namaTk"`
	TglLahir               string `json:"tglLahir"`
	TempatLahir            string `json:"tempatLahir"`
	JenisKelamin           string `json:"jenisKelamin"`
	NamaIbuKandung         string `json:"namaIbuKandung"`
	KodeKantor             string `json:"kodeKantor"`
	KodeSegmen             string `json:"kodeSegmen"`
	KodePerusahaan         string `json:"kodePerusahaan"`
	KodeDivisi             string `json:"kodeDivisi"`
	KodeKepesertaan        string `json:"kodeKepesertaan"`
	TglKepesertaan         string `json:"tglKepesertaann"`
	TglAktif               string `json:"tglAktif"`
	TglNonaktif            string `json:"tglNonaktif"`
	KodeNa                 string `json:"kodeNa"`
	SumberBlthNonAktif     string `json:"sumberBlthNonAktif"`
	FlagSumberNonAktif     string `json:"flagSumberNonAktif"`
	TglFlagSumberNonAktif  string `json:"tglFlagSumberNonAktif"`
	KodeTipeKlaim          string `json:"kodeTipeKlaim"`
	KodeSebabKlaim         string `json:"kodeSebabKlaim"`
	KodePelaporan          string `json:"kodePelaporan"`
	TglKlaim               string `json:"tglKlaim"`
	Npwp                   string `json:"npwp"`
	KodeBank               string `json:"kodeBank"`
	NamaBank               string `json:"namaBank"`
	NoRekening             string `json:"noRekening"`
	NamaRekening           string `json:"namaRekening"`
	Alamat                 string `json:"alamat"`
	KodeKelurahan          string `json:"kodeKelurahan"`
	KodeKecamatan          string `json:"kodeKecamatan"`
	KodeKabupaten          string `json:"kodeKabupaten"`
	KodePropinsi           string `json:"kodePropinsi"`
	KodePos                string `json:"kodePos"`
	StatusValidNama        string `json:"statusValidNama"`
	StatusValidTglLahir    string `json:"statusValidTglLahir"`
	StatusValidTempatLahir string `json:"statusValidTempatLahir"`
	StatusValidFoto        string `json:"statusValidFoto"`
	ScoreFace              string `json:"scoreFace"`
	ScoreFaceLiveness      string `json:"scoreFaceLiveness"`
	IsAlive                string `json:"isAlive"`
	KodeManfaat            string `json:"kodeManfaat"`
	KodeTipePenerima       string `json:"kodeTipePenerima"`
	KdPrg                  string `json:"kdPrg"`
}
