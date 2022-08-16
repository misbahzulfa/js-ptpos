package request

type SimilarityRequest struct {
	ChId  string `json:"chId"`
	ReqId string `json:"reqId"`
	Nama1 string `json:"nama1"`
	Nama2 string `json:"nama2"`
	Kode  string `json:"kode"`
}
