package response

type CheckAccountBankRemoteResponse struct {
	Ret  string `json:"ret"`
	Data struct {
		NAMAREKTUJUAN string `json:"NAMA_REK_TUJUAN"`
		BANKTUJUAN    string `json:"BANK_TUJUAN"`
		NOREKTUJUAN   string `json:"NOREK_TUJUAN"`
	} `json:"data"`
	Msg string `json:"msg"`
}
