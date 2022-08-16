package request

type CheckEligibleBeasiswaRequest struct {
	NikPelapor        string `json:"nikPelapor"`
	NamaPelapor       string `json:"namaPelapor"`
	NikPeserta        string `json:"nikPeserta"`
	ChannelID         string `json:"channelID"`
	ChId              string `json:"chId"`
	ReqId             string `json:"reqId"`
	KodeKantorPTPOS   string `json:"kodeKantorPTPOS"`
	PetugasRekamPTPOS string `json:"petugasRekamPTPOS"`
}
