package request

type DaftarSebabKlaimRequest struct {
	KodeSegmen        string `json:"kodeSegmen"`
	ChannelId         string `json:"channelId"`
	KodeKantorPPTOS   string `json:"kodeKantorPTPOS"`
	PetugasRekamPTPOS string `json:"petugasPTPOS"`
	ChId              string `json:"chId"`
	ReqId             string `json:"reqId"`
}
