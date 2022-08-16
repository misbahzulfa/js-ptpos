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

type ClaimController struct {
	ClaimService service.ClaimService
}

func NewClaimController(ClaimService *service.ClaimService) ClaimController {
	return ClaimController{ClaimService: *ClaimService}
}

func (controller *ClaimController) Route(app *fiber.App) {
	app.Get("/health-checks", controller.Healtcheck)
	// app.Post("/membership/pra-claim", controller.PraClaim)
	// app.Post("/membership/check-rsjht", controller.CheckRsjht)
	// app.Post("/membership/contribution", controller.Contribution)
	// app.Post("/membership/check-eligible", controller.CheckEligible)
	// app.Post("/membership/cause-of-claim", controller.CauseOfClaim)
	// app.Post("/membership/detail-employee", controller.DetailEmployee)
	// app.Post("/membership/claim-benefit-detail", controller.ClaimBenefitDetail)
	app.Post("/JSPTPOS/InsertPengajuanJHT", controller.InsertClaimJht)
	app.Post("/JSPTPOS/CheckBankAccount", controller.AccountBankCheck)
	// app.Post("/membership/send-email-after-claim", controller.SendEmailAfterClaim)
	// app.Post("/membership/tracking-claim", controller.TrackingClaim)
	// app.Post("/membership/send-email-after-claim-failed-transfer", controller.SendEmailAfterClaimFailedTransfer)
}

func (controller *ClaimController) Healtcheck(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("Api Running")
}
func (controller *ClaimController) AccountBankCheck(ctx *fiber.Ctx) error {
	var request request.CheckAccountBankRemoteRequest
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

	response, err := controller.ClaimService.CheckAccountBank(ctx, requestId.String(), &request)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	logStop := util.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)

	return ctx.Status(200).JSON(response)
}
func (controller *ClaimController) InsertClaimJht(ctx *fiber.Ctx) error {
	var request request.InsertClaimJhtRequest
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

	response, err := controller.ClaimService.InsertClaimJht(ctx, requestId.String(), &request)
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
