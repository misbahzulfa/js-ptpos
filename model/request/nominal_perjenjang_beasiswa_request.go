package request

type NominalPerjenjangBeasiswaRequest struct {
	NikPelapor        string `json:"nikPelapor"`
	NamaPelapor       string `json:"namaPelapor"`
	NikPeserta        string `json:"nikPeserta"`
	JenisBeasiswa     string `json:"jenisBeasiswa"`
	JenjangPendidikan string `json:"jenjangPendidikan"`
	ChannelID         string `json:"channelID"`
	ChId              string `json:"chId"`
	ReqId             string `json:"reqId"`
	KodeKantorPTPOS   string `json:"kodeKantorPTPOS"`
	PetugasRekamPTPOS string `json:"petugasRekamPTPOS"`
}
