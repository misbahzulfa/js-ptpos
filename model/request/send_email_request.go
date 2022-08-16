package request

type SendEmailRequest struct {
	Subject          string `json:"subject"`
	Body             string `json:"body"`
	Message          string `json:"message"`
	Email            string `json:"email"`
	OfficeCode       string `json:"OfficeCode"`
	FullName         string `json:"FullName"`
	KodeKantorPPTOS  string `json:"KodeKantorPPTOS"`
	KodePengajuan    string `json:"KodePengajuan"`
	TanggalPengajuan string `json:"TanggalPengajuan"`
	Kpj              string `json:"Kpj"`
}
