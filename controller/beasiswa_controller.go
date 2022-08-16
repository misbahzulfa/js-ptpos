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

type BeasiswaController struct {
	BeasiswaService service.BeasiswaService
}

func NewBeasiswaController(BeasiswaService *service.BeasiswaService) BeasiswaController {
	return BeasiswaController{BeasiswaService: *BeasiswaService}
}

func (controller *BeasiswaController) Route(app *fiber.App) {

	app.Post("/JSPTPOS/CariDataPelaporKonfirmasiBeasiswa", controller.CheckEligibleBeasiswa)
	app.Post("/JSPTPOS/InsertKonfirmasiBeasiswa", controller.InsertKonfirmasiBeasiswa)
	app.Post("/JSPTPOS/DaftarJenisBeasiswa", controller.DaftarJenisBeasiswa)
	app.Post("/JSPTPOS/DaftarJenjangPendidikan", controller.DaftarJenjangPendidikan)
	app.Post("/JSPTPOS/DaftarPenerimaBeasiswa", controller.DaftarPenerimaBeasiswa)
	app.Post("/JSPTPOS/CariNominalPerJenjang", controller.NominalPerjenjangBeasiswa)

}

func (controller *BeasiswaController) CheckEligibleBeasiswa(ctx *fiber.Ctx) error {
	var request request.CheckEligibleBeasiswaRequest
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

	response, err := controller.BeasiswaService.CheckEligibleBeasiswa(ctx, requestId.String(), &request)
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

func (controller *BeasiswaController) InsertKonfirmasiBeasiswa(ctx *fiber.Ctx) error {
	var request request.InsertKonfirmasiBeasiswaRequest
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

	response, err := controller.BeasiswaService.InsertKonfirmasiBeasiswa(ctx, requestId.String(), &request)
	if err != nil {
		fmt.Println(err)
		exception.PanicIfNeeded(err)
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		exception.PanicIfNeeded(err)
	}

	logStop := util.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}

func (controller *BeasiswaController) DaftarJenisBeasiswa(ctx *fiber.Ctx) error {
	var request request.DaftarJenisBeasiswaRequest
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

	response, err := controller.BeasiswaService.DaftarJenisBeasiswa(ctx, requestId.String(), &request)
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

func (controller *BeasiswaController) DaftarJenjangPendidikan(ctx *fiber.Ctx) error {
	var request request.DaftarJenjangPendidikanRequest
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

	response, err := controller.BeasiswaService.DaftarJenjangPendidikan(ctx, requestId.String(), &request)
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

func (controller *BeasiswaController) DaftarPenerimaBeasiswa(ctx *fiber.Ctx) error {
	var request request.DaftarPenerimaBeasiswaRequest
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

	response, err := controller.BeasiswaService.DaftarPenerimaBeasiswa(ctx, requestId.String(), &request)
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

func (controller *BeasiswaController) NominalPerjenjangBeasiswa(ctx *fiber.Ctx) error {
	var request request.NominalPerjenjangBeasiswaRequest
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

	response, err := controller.BeasiswaService.NominalPerjenjangBeasiswa(ctx, requestId.String(), &request)
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
