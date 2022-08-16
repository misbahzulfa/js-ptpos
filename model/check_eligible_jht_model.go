package model

type CheckEligibleJHTLayak struct {
	SebabKlaim          string        `json:"SebabKlaim"`
	StatusKelayakan     string        `json:"StatusKelayakan"`
	KodeKelayakan       string        `json:"KodeKelayakan"`
	KeteranganKelayakan string        `json:"KeteranganKelayakan"`
	DataProbingDetil    []DataProbing `json:"DataProbing"`
	Success             string        `json:"success"`
	Message             string        `json:"message"`
}

type DataTK struct {
	KodeTK     string `json:"KodeTK"`
	KodeSegmen string `json:"KodeSegmen"`
	TglLahir   string `json:"TglLahir"`
	Email      string `json:"Email"`
}

type DataProbing struct {
	KodeProbing   string `json:"kodeProbing"`
	NoUrut        string `json:"noUrut"`
	NamaProbing   string `json:"namaProbing"`
	ResponProbing string `json:"responProbing"`
}

type CheckEligibleJHTLayakResponse struct {
	StatusCode int                     `json:"StatusCode"`
	Data       []CheckEligibleJHTLayak `json:"Data"`
}
