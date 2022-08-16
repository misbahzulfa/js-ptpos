package request

type TrackingClaimRequest struct {
	Kpj            string `json:"kpj"`
	IdentityNumber string `json:"identityNumber"`
}
