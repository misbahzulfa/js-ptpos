package request

type CheckEligibleRequest struct {
	SegmenCode     string `json:"segmenCode"`
	WorkerCode     string `json:"workerCode"`
	Kpj            string `json:"kpj"`
	IdentityNumber string `json:"identityNumber"`
	FullName       string `json:"fullName"`
	BirthDate      string `json:"birthDate"`
	Email          string `json:"email"`
}
