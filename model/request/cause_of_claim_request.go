package request

type CauseOfClaimRequest struct {
	SegmenCode string `json:"segmenCode"`
	WorkerCode string `json:"workerCode"`
	Program    string `json:"program"`
	BirthDate  string `json:"birthDate"`
}
