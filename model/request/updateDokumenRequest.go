package request

type UpdateDokumenRequest struct {
	ReqId             string `json:"reqId"`
	ChId              string `json:"chId"`
	ChannelID         string `json:"channelID"`
	KodeKantorPtPos   string `json:"kodeKantorPtPos"`
	PetugasRekamPtPos string `json:"petugasRekamPtPos"`
	KodePengajuan     string `json:"kodePengajuan"`
	JenisLayanan      string `json:"jenisPengajuan"`
	KodeDokumen       string `json:"kodeDokumen"`
	Mimetype          string `json:"mimeType"`
	PathURL           string `json:"pathUrl"`
}
