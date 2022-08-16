package request

type CheckOperasionalRequest struct {
	BankCode    string `json:"bankCode"`
	RequestDate string `json:"requestDate"`
}
