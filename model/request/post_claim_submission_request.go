package request

type PostClaimSubmissionRequest struct {
	SubmissionCode string `json:"submissionCode"`
	ClaimCode      string `json:"claimCode"`
	Email          string `json:"email"`
}
