package request

type CheckRsjhtRequest struct {
	Year       string `json:"year"`
	Kpj        string `json:"kpj"`
	WorkerCode string `json:"workerCode"`
}
