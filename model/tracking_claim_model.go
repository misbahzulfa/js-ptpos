package model

type TrackingClaim struct {
	ClaimCode         string              `json:"claimCode"`
	ClaimTypeCode     string              `json:"claimTypeCode"`
	IdentityNumber    string              `json:"identityNumber"`
	Kpj               string              `json:"kpj"`
	TrackingDataClaim []TrackingDataClaim `json:"data"`
}

type TrackingDataClaim struct {
	ClaimCode       string `json:"claimCode"`
	ClaimTypeCode   string `json:"claimTypeCode"`
	IdentityNumber  string `json:"identityNumber"`
	Kpj             string `json:"kpj"`
	Fullname        string `json:"fullname"`
	Step            string `json:"step"`
	StepInformation string `json:"stepInformation"`
	CreatedAt       string `json:"createdAt"`
	Notes           string `json:"notes"`
	StepName        string `json:"stepName"`
}
