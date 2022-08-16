package controller

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"js-ptpos/exception"
	"js-ptpos/model/request"
	"js-ptpos/service"
	"js-ptpos/util"

	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
)

type EmailController struct {
	EmailService service.EmailService
}

func NewEmailController(emailService *service.EmailService) EmailController {
	return EmailController{EmailService: *emailService}
}

func (controller *EmailController) Route(app *fiber.App) {
	app.Post("/JSPTPOS/SendEmail", controller.SendEmail)
	app.Post("/JSPTPOS/SendEmailSuccessTransfer", controller.SendEmailAfterClaimPTPOS)
	// app.Post("/JSPTPOS/SendEmailFailedTransfer", controller.SendEmailAfterClaimFailedTransferPTPOS)
	app.Post("/JSPTPOS/SendEmailHealthcek", controller.HealtcheckEmail)
}

func (controller *EmailController) HealtcheckEmail(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("Api Running")
}

func (controller *EmailController) SendEmail(ctx *fiber.Ctx) error {
	var request request.SendEmailRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	response, err := controller.EmailService.SendEmail(&request)

	if err != nil {
		exception.PanicIfNeeded(err)
	}

	requestId := guuid.New()
	logStart := util.LogRequest(ctx /*string(requestData)*/, "Request", requestId.String())
	fmt.Println(logStart)
	logStop := util.LogResponse(ctx /*string(responseData)*/, "Response", requestId.String())
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}

func (controller *EmailController) SendEmailAfterClaimPTPOS(ctx *fiber.Ctx) error {
	var request request.GetDataAfterClaimRequest
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

	response, err := controller.EmailService.SendEmailAfterClaim(ctx, requestId.String(), &request)
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

// func (controller *EmailController) SendEmailAfterClaimFailedTransferPTPOS(ctx *fiber.Ctx) error {
// 	var request request.GetDataAfterClaimRequest
// 	err := ctx.BodyParser(&request)
// 	if err != nil {
// 		exception.PanicIfNeeded(err)
// 	}

// 	requestData, err := json.Marshal(request)
// 	if err != nil {
// 		exception.PanicIfNeeded(err)
// 	}

// 	requestId := guuid.New()
// 	logStart := util.LogRequest(ctx, string(requestData), requestId.String())
// 	fmt.Println(logStart)

// 	response, err := controller.EmailService.SendEmailAfterClaimFailedTransfer(ctx, requestId.String(), &request)
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
