package validation

import (
	"fmt"
	"js-ptpos/exception"
	"js-ptpos/model/request"
	"js-ptpos/util"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
)

func CheckEligibleJPValidate(ctx *fiber.Ctx, params request.CheckEligibleJPBerkalaRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.ChannelID, validation.Required),
		validation.Field(&params.KodeKantorPtPos, validation.Required),
		validation.Field(&params.PetugasRekamPtPos, validation.Required),
		validation.Field(&params.NikPelapor, validation.Required),
		validation.Field(&params.NamaPelapor, validation.Required),
		validation.Field(&params.NikPenerimaManfaat, validation.Required),
		validation.Field(&params.NamaPenerimaManfaat, validation.Required),
		validation.Field(&params.NikPeserta, validation.Required),
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

func InsertJPConfirmationValidate(ctx *fiber.Ctx, params request.InsertJPConfirmationRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.ChannelID, validation.Required),
		validation.Field(&params.KodeKantorPtPos, validation.Required),
		validation.Field(&params.PetugasRekamPtPos, validation.Required),
		validation.Field(&params.NikPelapor, validation.Required),
		validation.Field(&params.NamaPelapor, validation.Required),
		validation.Field(&params.TanggalLahirPelapor, validation.Required),
		validation.Field(&params.NikPenerimaManfaat, validation.Required),
		validation.Field(&params.NamaPenerimaManfaat, validation.Required),
		validation.Field(&params.TanggalLahirPenerimaManfaat, validation.Required),
		validation.Field(&params.NikPeserta, validation.Required),
		validation.Field(&params.EmailPelapor, validation.Required),
		validation.Field(&params.NoHPPelapor, validation.Required),
		validation.Field(&params.TanggalPengajuan, validation.Required),
		validation.Field(&params.KodeBillingPTPos, validation.Required),
		validation.Field(&params.ScoreFaceMatch, validation.Required),
		validation.Field(&params.SimilarityNamaPTPOSkeAdminduk, validation.Required),
		validation.Field(&params.KeteranganKonfirmasi, validation.Required),
		validation.Field(&params.DataDokumen, validation.Required),
		validation.Field(&params.DataProbing, validation.Required),
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
