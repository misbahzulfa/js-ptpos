package request

type ClaimBenefitDetailRequest struct {
	SegmenCode       string `json:"segmenCode"`
	Kpj              string `json:"kpj"`
	WorkerCode       string `json:"workerCode"`
	IdentityNumber   string `json:"identityNumber"`
	FullName         string `json:"fullName"`
	BirthDate        string `json:"birthDate"`
	CauseOfClaimCode string `json:"causeOfClaimCode"`
}
