package service

import (
	"js-ptpos/model/request"
	"js-ptpos/model/response"

	"github.com/gofiber/fiber/v2"
)

type ClaimService interface {
	InsertClaimJht(ctx *fiber.Ctx, requestId string, params *request.InsertClaimJhtRequest) (*response.InsertClaimJhtResponse, error)
	CheckEligibleInsert(ctx *fiber.Ctx, requestId string, params *request.CheckEligibleJHTRequest) (*response.CheckJHTEligibleResponse, error)
	//checkbank
	CheckAccountBank(ctx *fiber.Ctx, requestId string, params *request.CheckAccountBankRemoteRequest) (*response.CheckAccountBankRemoteResponse, error)
	//SimilairtyAccountBankName(params *request.SimilarityRequest) (*response.CommonSimilarityResponse, error)
}
