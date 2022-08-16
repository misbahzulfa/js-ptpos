package request

type DaftarSegmenRequest struct {
	ChannelId         string `json:"ChannelId"`
	KodeKantorPPTOS   string `json:"KodeKantorPTPOS"`
	PetugasRekamPTPOS string `json:"PetugasPTPOS"`
	ChId              string `json:"chId"`
	ReqId             string `json:"reqId"`
}
