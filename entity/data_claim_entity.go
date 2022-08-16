package entity

type DataClaimEntity struct {
	Number         string `json:"number"`
	SubmissionCode string `json:"submissionCode"`
	Email          string `json:"email"`
	BankCode       string `json:"bankCode"`
}
