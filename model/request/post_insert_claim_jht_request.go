package request

type PostInsertClaimJhtRequest struct {
	SubmissionCode string `json:"submissionCode"`
	Email          string `json:"email"`
}
