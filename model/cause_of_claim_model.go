package model

type CauseOfClaimModel struct {
	SegmenCode       string `json:"segmenCode"`
	ClaimTypeCode    string `json:"claimTypeCode"`
	CauseOfClaimCode string `json:"causeOfClaimCode"`
	CauseOfClaimName string `json:"causeOfClaimName"`
	ReceiverTypeCode string `json:"receiverTypeCode"`
	Number           string `json:"number"`
}
