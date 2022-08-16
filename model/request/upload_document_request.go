package request

type UploadDocRequest struct {
	ReqId             string `json:"reqId"`
	ChId              string `json:"chId"`
	ChannelID         string `json:"channelID"`
	KodeKantorPtPos   string `json:"kodeKantorPtPos"`
	PetugasRekamPtPos string `json:"petugasRekamPtPos"`
	KodePengajuan     string `json:"kodePengajuan"`
	JenisLayanan      string `json:"jenisPengajuan"`
	KodeDokumen       string `json:"kodeDokumen"`
	File              []File `json:"file"`
}

type File struct {
	Mime string `json:"mime"`
	Data string `json:"data"`
}
