package entity

// type CheckJHTEligibleEntity struct {
// 	SebabKlaim          string `json:"SebabKlaim"`
// 	StatusKelayakan     string `json:"StatusKelayakan"`
// 	KodeKelayakan       string `json:"KodeKelayakan"`
// 	KeteranganKelayakan string `json:"KeteranganKelayakan"`
// 	DataProbing         string `json:"DataProbing"`
// }

type CheckJHTEligibleEntity struct {
	SebabKlaim          []DataSebabKlaim `json:"SebabKlaim"`
	StatusKelayakan     string           `json:"StatusKelayakan"`
	KodeKelayakan       string           `json:"KodeKelayakan"`
	KeteranganKelayakan string           `json:"KeteranganKelayakan"`
	DataProbingDetil    []DataProbing    `json:"DataProbingDetil"`
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

type DataSebabKlaim struct {
	KodeSebabKlaim string `json:"KodeSebabKlaim"`
	NamaSebabKlaim string `json:"NamaSebabKlaim"`
}

type CheckJHTEligibleEntityResponse struct {
	StatusCode string                   `json:"StatusCode"`
	StatusDesc string                   `json:"StatusDesc"`
	Data       []CheckJHTEligibleEntity `json:"Data"`
}
