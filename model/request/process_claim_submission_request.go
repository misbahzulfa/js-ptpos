package request

type ProcessClaimSubmissionRequest struct {
	SubmissionCode string `json:"submissionCode"`
	Email          string `json:"email"`
}
