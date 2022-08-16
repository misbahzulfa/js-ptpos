package service

import (
	model "js-ptpos/model"
	"js-ptpos/model/request"
	"js-ptpos/model/response"

	"github.com/gofiber/fiber/v2"
)

type BeasiswaService interface {
	CheckEligibleBeasiswa(ctx *fiber.Ctx, requestId string, params *request.CheckEligibleBeasiswaRequest) (*response.CheckEligibleBeasiswaResponse, error)
	InsertKonfirmasiBeasiswa(ctx *fiber.Ctx, requestId string, params *request.InsertKonfirmasiBeasiswaRequest) (*response.InsertKonfirmasiBeasiswaResponse, error)
	DaftarJenisBeasiswa(ctx *fiber.Ctx, requestId string, params *request.DaftarJenisBeasiswaRequest) (*response.DaftarJenisBeasiswaResponse, error)
	DaftarJenjangPendidikan(ctx *fiber.Ctx, requestId string, params *request.DaftarJenjangPendidikanRequest) ([]model.DaftarJenjangPendidikan, error)
	DaftarPenerimaBeasiswa(ctx *fiber.Ctx, requestId string, params *request.DaftarPenerimaBeasiswaRequest) (*response.DaftarPenerimaBeasiswaResponse, error)
	NominalPerjenjangBeasiswa(ctx *fiber.Ctx, requestId string, params *request.NominalPerjenjangBeasiswaRequest) (*response.NominalPerjenjangBeasiswaResponse, error)
}
