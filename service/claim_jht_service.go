package service

import (
	model "js-ptpos/model"
	"js-ptpos/model/request"
	"js-ptpos/model/response"

	"github.com/gofiber/fiber/v2"
)

type ClaimJHTService interface {
	CheckEligibleJHT(ctx *fiber.Ctx, requestId string, params *request.CheckEligibleJHTRequest) (*response.CheckJHTEligibleResponse, error)
	InsertPengajuanJHT(ctx *fiber.Ctx, requestId string, params *request.InsertPengajuanJHTRequest) (*model.InsertPengajuanJHT, error)
	UpdatePengajuanJHT(ctx *fiber.Ctx, requestId string, params *request.UpdatePengajuanJHTRequest) (*model.UpdatePengajuanJHT, error)
	GetPengajuanJHT(ctx *fiber.Ctx, requestId string, params *request.GetPengajuanJHTRequest) (*model.PengajuanJHTResponse, error)
	DaftarSegmen(ctx *fiber.Ctx, requestId string, params *request.DaftarSegmenRequest) (*model.DaftarSegmenResponse, error)
	DaftarSebabKlaim(ctx *fiber.Ctx, requestId string, params *request.DaftarSebabKlaimRequest) (*model.DaftarSebabKlaimResponse, error)
	DaftarDokumenSebabKlaim(ctx *fiber.Ctx, requestId string, params *request.DaftarDokumenSebabKlaimRequest) (*model.DaftarDokumenSebabKlaimResponse, error)
}
