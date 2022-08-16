package entity

type TrackingHeaderClaimEntity struct {
	ClaimCode      string `json:"claimCode"`
	ClaimTypeCode  string `json:"claimTypeCode"`
	IdentityNumber string `json:"identityNumber"`
	Kpj            string `json:"kpj"`
	FullName       string `json:"fullName"`
}
