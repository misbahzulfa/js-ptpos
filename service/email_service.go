package service

import (
	"js-ptpos/model/request"
	"js-ptpos/model/response"

	"github.com/gofiber/fiber/v2"
)

type EmailService interface {
	SendEmail(params *request.SendEmailRequest) (*response.SendEmailResponse, error)
	SendEmailAfterClaim(ctx *fiber.Ctx, requestId string, params *request.GetDataAfterClaimRequest) (*response.SendEmailResponse, error)
	//SendEmailAfterClaimJPBeasiswa(ctx *fiber.Ctx, requestId string, params *request.GetDataAfterClaimRequest) (*response.SendEmailResponse, error)
	//SendEmailAfterClaimFailedTransfer(ctx *fiber.Ctx, requestId string, params *request.GetDataAfterClaimRequest) (*response.SendEmailResponse, error)
}
