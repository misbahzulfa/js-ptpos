package controller

import (
	"encoding/json"
	"fmt"
	"js-ptpos/exception"
	"js-ptpos/model/request"
	"js-ptpos/service"
	"js-ptpos/util"

	"github.com/gofiber/fiber/v2"
)

type StorageController struct {
	StorageService service.StorageService
}

func NewStorageController(StorageService *service.StorageService) StorageController {
	return StorageController{StorageService: *StorageService}
}

func (controller *StorageController) Route(app *fiber.App) {
	app.Get("/JSPTPOS/UploadDokumen/health-checks", controller.HealtcheckJP)
	app.Post("/JSPTPOS/UploadDokumen", controller.UploadDokumen)

}

func (controller *StorageController) HealtcheckJP(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("Api Running")
}

func (controller *StorageController) UploadDokumen(ctx *fiber.Ctx) error {
	var request request.UploadDocRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStart := util.LogRequest(ctx, "LOG REQUEST", request.ReqId)
	fmt.Println(logStart)

	response, err := controller.StorageService.UploadDokumen(ctx, &request)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStop := util.LogResponse(ctx, string(responseData), request.ReqId)
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}
