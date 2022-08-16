package repository

import (
	"js-ptpos/entity"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
)

type ClaimRepository interface {
	CauseOfClaim(params *request.CauseOfClaimRequest) ([]entity.CauseOfClaimEntity, error)
	GetDataEmployee(params *request.GetDataEmployeeRequest) (*entity.DataEmployeeEntity, error)
	GetClaimBenefitDetail(params *request.ClaimBenefitDetailRequest) (*entity.ClaimBenefitDetailEntity, error)
	InsertClaimJht(params *request.InsertClaimRequest) (string, error)
	SendEmail(params *request.SendEmailRequest) (*response.SendEmailResponse, error)
	CheckEligibleInsert(params *request.CheckEligibleJHTRequest) (*response.CheckJHTEligibleResponse, error)
	//check bank
	CheckAccountBank(params *request.CheckAccountBankRemoteRequest) (*response.CheckAccountBankRemoteResponse, error)
	SimilairtyAccountBankName(params *request.SimilarityRequest) (*response.CommonSimilarityResponse, error)
}
