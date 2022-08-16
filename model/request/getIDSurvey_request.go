package request

type GetIDSurveyReq struct {
	ChID  string    `json:"chId"`
	ReqID string    `json:"reqId"`
	Data  DataClaim `json:"data"`
}

type DataClaim struct {
	KodeKlaim string `json:"kodeKlaim"`
	User      string `json:"user"`
	KodeKanal string `json:"kodeKanal"`
	SendEmail string `json:"sendEmail"`
}

type IDSurveyRes struct {
	Ret     string `json:"ret"`
	Encode  string `json:"encode"`
	Message string `json:"msg"`
}
