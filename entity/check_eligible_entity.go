package entity

type CheckEligibleEntity struct {
	BalanceMaksimum    string `json:"balanceMaksimum"`
	PartialClaim       string `json:"partialClaim"`
	MessageCode        string `json:"messageCode"`
	MessageNotEligible string `json:"messageNotEligible"`
	OfficeName         string `json:"officeName"`
	Success            string `json:"success"`
	Message            string `json:"message"`
}
