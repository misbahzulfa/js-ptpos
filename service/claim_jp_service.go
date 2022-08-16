package service

import (
	"js-ptpos/model/request"
	"js-ptpos/model/response"

	"github.com/gofiber/fiber/v2"
)

type ClaimJPService interface {
	CheckEligibleJP(ctx *fiber.Ctx, requestId string, params *request.CheckEligibleJPBerkalaRequest) (*response.CheckEligibleJPBerkalaResponse, error)
	CheckClaimByNIKPelapor(ctx *fiber.Ctx, requestId string, params *request.CheckJumlahKlaimJPBerkalaNikPelaporRequest) (*response.CheckJumlahKlaimJPBerkalaNikPelaporResponse, error)
	InsertConfirmationJPBerkala(ctx *fiber.Ctx, requestId string, params *request.InsertJPConfirmationRequest) (*response.InsertKonfirmasiJPResponse, error)
	// UpdateKonfirmasiJPECtoPN(ctx *fiber.Ctx, requestId string, params *request.UpdateKonfirmasiJPECtoPNRequest) (*response.UpdateKonfirmasiJPECtoPNResponse, error)
}
