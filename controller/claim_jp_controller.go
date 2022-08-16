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

type ClaimJPController struct {
	ClaimJPService service.ClaimJPService
	StorageService service.StorageService
	EmailService   service.EmailService
}

func NewClaimJPController(ClaimJPService *service.ClaimJPService) ClaimJPController {
	return ClaimJPController{ClaimJPService: *ClaimJPService}
}

func (controller *ClaimJPController) Route(app *fiber.App) {
	app.Get("/JSPTPOS/health-checks", controller.HealtcheckJP)
	app.Post("/JSPTPOS/CariDataPelaporKonfirmasiJPBerkala", controller.CheckEligibleJP)
	app.Post("/JSPTPOS/InsertKonfirmasiJPBerkala", controller.InsertConfirmationJP)
	// app.Post("/JSPTPOS/CekJumlahKlaimJpBerkalaByNikPelapor", controller.CheckJumlahClaimJPByNIKPelapor)

}

func (controller *ClaimJPController) HealtcheckJP(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("Api Running")
}

func (controller *ClaimJPController) CheckEligibleJP(ctx *fiber.Ctx) error {
	var request request.CheckEligibleJPBerkalaRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStart := util.LogRequest(ctx, string(requestData), request.ReqId)
	fmt.Println(logStart)

	response, err := controller.ClaimJPService.CheckEligibleJP(ctx, request.ReqId, &request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStop := util.LogResponse(ctx, string(responseData), request.ReqId)
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}

// func (controller *ClaimJPController) CheckJumlahClaimJPByNIKPelapor(ctx *fiber.Ctx) error {
// 	var request request.CheckJumlahKlaimJPBerkalaNikPelaporRequest
// 	err := ctx.BodyParser(&request)
// 	if err != nil {
// 		exception.PanicIfNeeded(err)
// 	}

// 	requestData, err := json.Marshal(request)
// 	if err != nil {
// 		exception.PanicIfNeeded(err)
// 	}

// 	// requestId := guuid.New()
// 	logStart := util.LogRequest(ctx, string(requestData), requestId.String())
// 	fmt.Println(logStart)

// 	response, err := controller.ClaimJPService.CheckClaimByNIKPelapor(ctx, requestId.String(), &request)
// 	if err != nil {
// 		exception.PanicIfNeeded(err)
// 	}

// 	responseData, err := json.Marshal(response)
// 	if err != nil {
// 		exception.PanicIfNeeded(err)
// 	}

// 	logStop := util.LogResponse(ctx, string(responseData), requestId.String())
// 	fmt.Println(logStop)

// 	return ctx.Status(200).JSON(response)
// }

func (controller *ClaimJPController) InsertConfirmationJP(ctx *fiber.Ctx) error {
	var request request.InsertJPConfirmationRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	// requestId := guuid.New()
	logStart := util.LogRequest(ctx, string(requestData), request.ReqId)
	fmt.Println(logStart)

	response, err := controller.ClaimJPService.InsertConfirmationJPBerkala(ctx, request.ReqId, &request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStop := util.LogResponse(ctx, string(responseData), request.ReqId)
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}

func (controller *ClaimJPController) UploadDokumen(ctx *fiber.Ctx) error {
	var request request.UploadDocRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	response, err := controller.StorageService.UploadDokumen(ctx, &request)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	logStart := util.LogRequest(ctx, "string(requestData)", request.ReqId)
	fmt.Println(logStart)

	responseData, err := json.Marshal(response)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStop := util.LogResponse(ctx, string(responseData), request.ReqId)
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}
