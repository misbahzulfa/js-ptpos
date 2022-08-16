package service

import (
	"js-ptpos/model/request"
	"js-ptpos/model/response"

	"github.com/gofiber/fiber/v2"
)

type StorageService interface {
	UploadDokumen(ctx *fiber.Ctx, params *request.UploadDocRequest) (*response.UploadDokumenResponse, error)
}
