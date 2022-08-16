package request

type CheckAccountBankRequest struct {
	ChId          string `json:"chId"`
	Email         string `json:"email"`
	KodeBank      string `json:"kodeBank"`
	NamaBank      string `json:"namaBank"`
	NomorRekening string `json:"nomorRekening"`
	NamaRekening  string `json:"namaRekening"`
	NamaPeserta   string `json:"namaPeserta"`
	Signature     string `json:"signature"`
}

type CheckAccountBankRemoteRequest struct {
	ChID        string `json:"chId"`
	ReqID       string `json:"reqId"`
	Bank        string `json:"bank"`
	KODEBANKATB string `json:"KODE_BANK_ATB"`
	NOREKTUJUAN string `json:"NOREK_TUJUAN"`
	NAMAREK     string `json:"NAMA_REK"`
}
