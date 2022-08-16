package request

type CheckEligibleJPBerkalaRequest struct {
	ReqId               string `json:"reqId"`
	ChId                string `json:"chId"`
	ChannelID           string `json:"channelID"`
	KodeKantorPtPos     string `json:"kodeKantorPtPos"`
	PetugasRekamPtPos   string `json:"petugasRekamPtPos"`
	NikPelapor          string `json:"nikPelapor"`
	NamaPelapor         string `json:"namaPelapor"`
	NikPenerimaManfaat  string `json:"nikPenerimaManfaat"`
	NamaPenerimaManfaat string `json:"namaPenerimaManfaat"`
	NikPeserta          string `json:"nikPeserta"`
}
