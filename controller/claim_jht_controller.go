package controller

import (
	"encoding/json"
	"fmt"
	"js-ptpos/exception"
	"js-ptpos/model/request"
	"js-ptpos/service"
	"js-ptpos/util"

	guuid "github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

type ClaimJHTController struct {
	ClaimJHTService service.ClaimJHTService
}

func NewClaimJHTController(ClaimJHTService *service.ClaimJHTService) ClaimJHTController {
	return ClaimJHTController{ClaimJHTService: *ClaimJHTService}
}

func (controller *ClaimJHTController) Route(app *fiber.App) {
	app.Post("/JSPTPOS/EligibleJHT", controller.CheckEligibleJHT)
	// app.Post("/JSPTPOS/InsertPengajuanJHT", controller.InsertPengajuanJHT)
	app.Post("/JSPTPOS/UpdatePengajuanJHT", controller.UpdatePengajuanJHT)
	app.Post("/JSPTPOS/GetPengajuanJHT", controller.GetPengajuanJHT)
	app.Post("/JSPTPOS/DaftarSegmen", controller.DaftarSegmen)
	app.Post("/JSPTPOS/DaftarSebabKlaim", controller.DaftarSebabKlaim)
	app.Post("/JSPTPOS/DaftarDokumenBySebabKlaimJHT", controller.DaftarDokumenBySebabKlaimJHT)
}

func (controller *ClaimJHTController) CheckEligibleJHT(ctx *fiber.Ctx) error {
	var request request.CheckEligibleJHTRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestId := guuid.New()
	logStart := util.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := controller.ClaimJHTService.CheckEligibleJHT(ctx, requestId.String(), &request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStop := util.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}
func (controller *ClaimJHTController) InsertPengajuanJHT(ctx *fiber.Ctx) error {
	var request request.InsertPengajuanJHTRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestId := guuid.New()
	logStart := util.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := controller.ClaimJHTService.InsertPengajuanJHT(ctx, requestId.String(), &request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStop := util.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}
func (controller *ClaimJHTController) UpdatePengajuanJHT(ctx *fiber.Ctx) error {
	var request request.UpdatePengajuanJHTRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestId := guuid.New()
	logStart := util.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := controller.ClaimJHTService.UpdatePengajuanJHT(ctx, requestId.String(), &request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStop := util.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}
func (controller *ClaimJHTController) GetPengajuanJHT(ctx *fiber.Ctx) error {
	var request request.GetPengajuanJHTRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestId := guuid.New()
	logStart := util.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := controller.ClaimJHTService.GetPengajuanJHT(ctx, requestId.String(), &request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStop := util.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}

func (controller *ClaimJHTController) DaftarSegmen(ctx *fiber.Ctx) error {
	var request request.DaftarSegmenRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestId := guuid.New()
	logStart := util.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := controller.ClaimJHTService.DaftarSegmen(ctx, requestId.String(), &request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStop := util.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}

func (controller *ClaimJHTController) DaftarSebabKlaim(ctx *fiber.Ctx) error {
	var request request.DaftarSebabKlaimRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestId := guuid.New()
	logStart := util.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := controller.ClaimJHTService.DaftarSebabKlaim(ctx, requestId.String(), &request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStop := util.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}

func (controller *ClaimJHTController) DaftarDokumenBySebabKlaimJHT(ctx *fiber.Ctx) error {
	var request request.DaftarDokumenSebabKlaimRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestId := guuid.New()
	logStart := util.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := controller.ClaimJHTService.DaftarDokumenSebabKlaim(ctx, requestId.String(), &request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStop := util.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}
