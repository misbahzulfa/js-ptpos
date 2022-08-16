package request

type GetDataAfterClaimRequest struct {
	ClaimCode      string `json:"kodeKlaim"`
	Email          string `json:"email"`
	JenisPengajuan string `json:"jenisPengajuan"`
	NoProses       string `json:"noProses"`
	UserSmile      string `json:"kodeUser"`
}

type DataNotifJMORequest struct {
	ClaimCode      string `json:"kodeKlaim"`
	Email          string `json:"email"`
	JenisPengajuan string `json:"jenisPengajuan"`
	NoProses       string `json:"noProses"`
	UserSmile      string `json:"kodeUser"`
}
