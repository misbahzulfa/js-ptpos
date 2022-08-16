package request

type GetOfficeCodeRequest struct {
	SegmenCode     string `json:"segmenCode"`
	Kpj            string `json:"kpj"`
	WorkerCode     string `json:"workerCode"`
	IdentityNumber string `json:"identityNumber"`
	Fullname       string `json:"fullname"`
	Birthdate      string `json:"birthdate"`
}
