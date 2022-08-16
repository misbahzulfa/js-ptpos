package model

type PraClaim struct {
	StatusUpdateData         string `json:"statusUpdateData"`
	StatusMembership         string `json:"statusMembership"`
	StatusNpwp               string `json:"statusNpwp"`
	StatusMaksimumBalanceJht string `json:"statusMaksimumBalanceJht"`
	StatusPensiun            string `json:"statusPensiun"`
	MaximumBalanceJht        string `json:"maximumBalanceJht"`
}
