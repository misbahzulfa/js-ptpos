package validation

import (
	"fmt"
	"js-ptpos/exception"
	"js-ptpos/model/request"
	"js-ptpos/util"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
)

func GetPengajuanJHTValidate(ctx *fiber.Ctx, params request.GetPengajuanJHTRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.KodePengajuan, validation.Required),
		validation.Field(&params.NIK, validation.Required),
		validation.Field(&params.NamaLengkap, validation.Required),
		validation.Field(&params.KPJ, validation.Required),
	)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
}
func GetDokumenBySebabKlaimJHTValidate(ctx *fiber.Ctx, params request.DaftarDokumenSebabKlaimRequest, requestId string) {

	err := validation.ValidateStruct(&params,
		validation.Field(&params.KodeSebabKlaim, validation.Required),
		validation.Field(&params.ChannelId, validation.Required),
		validation.Field(&params.KodeKantorPPTOS, validation.Required),
		validation.Field(&params.PetugasRekamPTPOS, validation.Required),
		validation.Field(&params.ChId, validation.Required),
		validation.Field(&params.ReqId, validation.Required),
	)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
}
func DaftarSegmenJHTValidate(ctx *fiber.Ctx, params request.DaftarSegmenRequest, requestId string) {

	err := validation.ValidateStruct(&params,
		validation.Field(&params.ChannelId, validation.Required),
		validation.Field(&params.KodeKantorPPTOS, validation.Required),
		validation.Field(&params.PetugasRekamPTPOS, validation.Required),
		validation.Field(&params.ChId, validation.Required),
		validation.Field(&params.ReqId, validation.Required),
	)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
}

func CheckEligibleJHT(ctx *fiber.Ctx, params request.CheckEligibleJHTRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.Nik, validation.Required),
		validation.Field(&params.Kpj, validation.Required),
		validation.Field(&params.Fullname, validation.Required),
		validation.Field(&params.ChannelId, validation.Required),
		validation.Field(&params.KodeKantorPPTOS, validation.Required),
		validation.Field(&params.PetugasRekamPTPOS, validation.Required),
		validation.Field(&params.ChId, validation.Required),
		validation.Field(&params.ReqId, validation.Required),
	)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
}

func GetDaftarSebabKlaimValidate(ctx *fiber.Ctx, params request.DaftarSebabKlaimRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.KodeSegmen, validation.Required),
		validation.Field(&params.ChannelId, validation.Required),
		validation.Field(&params.KodeKantorPPTOS, validation.Required),
		validation.Field(&params.PetugasRekamPTPOS, validation.Required),
		validation.Field(&params.ChId, validation.Required),
		validation.Field(&params.ReqId, validation.Required),
	)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
}

func CauseOfClaimValidate(ctx *fiber.Ctx, params request.CauseOfClaimRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.BirthDate, validation.Required),
		validation.Field(&params.SegmenCode, validation.Required),
		validation.Field(&params.WorkerCode, validation.Required),
		validation.Field(&params.Program, validation.Required),
	)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
}

func GetDataEmployeeValidate(ctx *fiber.Ctx, params request.GetDataEmployeeRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.BirthDate, validation.Required),
		validation.Field(&params.SegmenCode, validation.Required),
		validation.Field(&params.WorkerCode, validation.Required),
		validation.Field(&params.FullName, validation.Required),
		validation.Field(&params.Kpj, validation.Required),
		validation.Field(&params.IdentityNumber, validation.Required),
	)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
}

func GetClaimBenefitDetailValidate(ctx *fiber.Ctx, params request.ClaimBenefitDetailRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.BirthDate, validation.Required),
		validation.Field(&params.SegmenCode, validation.Required),
		validation.Field(&params.WorkerCode, validation.Required),
		validation.Field(&params.FullName, validation.Required),
		validation.Field(&params.Kpj, validation.Required),
		validation.Field(&params.IdentityNumber, validation.Required),
	)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
}

func InsertClaimJhtValidate(ctx *fiber.Ctx, params request.InsertClaimJhtRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.Address, validation.Required),
		validation.Field(&params.ChannelId, validation.Required),
		validation.Field(&params.DataDetailProbing, validation.Required),
		validation.Field(&params.PathUrl, validation.Required),
		validation.Field(&params.Email, validation.Required),
		validation.Field(&params.JawabanProbing, validation.Required),
		validation.Field(&params.Kpj, validation.Required),
		validation.Field(&params.KeteranganPengajuan, validation.Required),
		validation.Field(&params.BankCode, validation.Required),
		validation.Field(&params.KodeBillingPTPOS, validation.Required),
		validation.Field(&params.KabupatenCode, validation.Required),
		validation.Field(&params.KodeKantorPPTOS, validation.Required),
		validation.Field(&params.KecamatanCode, validation.Required),
		validation.Field(&params.KelurahanCode, validation.Required),
		validation.Field(&params.SegmenCode, validation.Required),
		validation.Field(&params.PropinsiCode, validation.Required),
		validation.Field(&params.IdentityNumber, validation.Required),
		validation.Field(&params.Npwp, validation.Required),
		validation.Field(&params.FullName, validation.Required),
		validation.Field(&params.MotherName, validation.Required),
		validation.Field(&params.AccountBankName, validation.Required),
		validation.Field(&params.PhoneNumber, validation.Required),
		validation.Field(&params.AccountBankNumber, validation.Required),
		validation.Field(&params.PetugasRekamPTPOS, validation.Required),
		validation.Field(&params.ScoreFace, validation.Required),
		validation.Field(&params.ScoreSimilarityNama, validation.Required),
		validation.Field(&params.BirthDate, validation.Required),
		validation.Field(&params.TanggalPengajuan, validation.Required),
		validation.Field(&params.BirthPlace, validation.Required),
		validation.Field(&params.ChId, validation.Required),
		validation.Field(&params.ReqId, validation.Required),
	)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
}

func PraClaimValidate(ctx *fiber.Ctx, params request.PraClaimRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.BirthDate, validation.Required),
		validation.Field(&params.SegmenCode, validation.Required),
		validation.Field(&params.WorkerCode, validation.Required),
		validation.Field(&params.FullName, validation.Required),
		validation.Field(&params.Kpj, validation.Required),
		validation.Field(&params.IdentityNumber, validation.Required),
		validation.Field(&params.CompanyCode, validation.Required),
		validation.Field(&params.DivisionCode, validation.Required),
	)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
}

func ContributionValidate(ctx *fiber.Ctx, params request.PraClaimRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.BirthDate, validation.Required),
		validation.Field(&params.SegmenCode, validation.Required),
		validation.Field(&params.WorkerCode, validation.Required),
		validation.Field(&params.FullName, validation.Required),
		validation.Field(&params.Kpj, validation.Required),
		validation.Field(&params.IdentityNumber, validation.Required),
		validation.Field(&params.CompanyCode, validation.Required),
		validation.Field(&params.DivisionCode, validation.Required),
	)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
}

func CheckRsjhtValidate(ctx *fiber.Ctx, params request.CheckRsjhtRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.Year, validation.Required),
		validation.Field(&params.WorkerCode, validation.Required),
		validation.Field(&params.Kpj, validation.Required),
	)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
}

func TrackingClaimValidate(ctx *fiber.Ctx, params request.TrackingClaimRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.Kpj, validation.Required),
	)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
}
