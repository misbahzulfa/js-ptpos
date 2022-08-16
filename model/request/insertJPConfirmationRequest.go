package request

type InsertJPConfirmationRequest struct {
	ReqId                         string                        `json:"reqId"`
	ChId                          string                        `json:"chId"`
	ChannelID                     string                        `json:"channelID"`
	KodeKantorPtPos               string                        `json:"kodeKantorPtPos"`
	PetugasRekamPtPos             string                        `json:"petugasRekamPtPos"`
	NikPelapor                    string                        `json:"nikPelapor"`
	NamaPelapor                   string                        `json:"namaPelapor"`
	TanggalLahirPelapor           string                        `json:"tanggalLahirPelapor"`
	NikPenerimaManfaat            string                        `json:"nikPenerimaManfaat"`
	NamaPenerimaManfaat           string                        `json:"namaPenerimaManfaat"`
	TanggalLahirPenerimaManfaat   string                        `json:"tanggalLahirPenerimaManfaat"`
	NikPeserta                    string                        `json:"nikPeserta"`
	EmailPelapor                  string                        `json:"emailPelapor"`
	NoHPPelapor                   string                        `json:"noHpPelapor"`
	TanggalPengajuan              string                        `json:"tanggalPengajuan"`
	KodeBillingPTPos              string                        `json:"kodeBillingPTPos"`
	ScoreFaceMatch                string                        `json:"scoreFaceMatch"`
	SimilarityNamaPTPOSkeAdminduk string                        `json:"similarityNamaPTPOSkeAdminduk"`
	KeteranganKonfirmasi          string                        `json:"keteranganKonfirmasi"`
	DataDokumen                   []DokumenInsertJPConfirmation `json:"dataDokumen"`
	DataProbing                   []DataProbingJPConfirmation   `json:"dataProbing"`
}

type DokumenInsertJPConfirmation struct {
	KodeDokumen string `json:"kodeDokumen"`
	PathURL     string `json:"pathURL"`
}

type DataProbingJPConfirmation struct {
	KodeProbing    string `json:"kodeProbing"`
	NoUrut         string `json:"noUrut"`
	ResponProbing  string `json:"responProbing"`
	JawabanProbing string `json:"jawabanProbing"`
}
