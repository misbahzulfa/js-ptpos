package request

type CheckEligibleJHTRequest struct {
	Nik               string `json:"Nik"`
	Kpj               string `json:"Kpj"`
	Fullname          string `json:"NamaLengkap"`
	ChannelId         string `json:"ChannelId"`
	KodeKantorPPTOS   string `json:"KodeKantorPTPOS"`
	PetugasRekamPTPOS string `json:"PetugasPTPOS"`
	ChId              string `json:"chId"`
	ReqId             string `json:"reqId"`
	TglLahir          string `json:"tglLahir"`
}
