package validation

import (
	"fmt"
	"js-ptpos/exception"
	"js-ptpos/model/request"
	"js-ptpos/util"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
)

func CheckEligibleBeasiswaValidate(ctx *fiber.Ctx, params request.CheckEligibleBeasiswaRequest, requestId string) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.NikPelapor, validation.Required),
		validation.Field(&params.NamaPelapor, validation.Required),
		validation.Field(&params.NikPeserta, validation.Required),
		validation.Field(&params.ChannelID, validation.Required),
		validation.Field(&params.KodeKantorPTPOS, validation.Required),
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

func InsertKonfirmasiBeasiswa(ctx *fiber.Ctx, params request.InsertKonfirmasiBeasiswaRequest, requestId string) {

	// for _, penerimaBeasiswa := range params.PenerimaBeasiswa {
	// 	validation.Field(penerimaBeasiswa.NamaPenerimaBeasiswa, validation.Required)
	// }

	err := validation.ValidateStruct(&params,
		validation.Field(&params.NikPelapor, validation.Required),
		validation.Field(&params.NamaPelapor, validation.Required),
		validation.Field(&params.NikPeserta, validation.Required),
		validation.Field(&params.ChannelID, validation.Required),
		validation.Field(&params.KodeKantorPTPOS, validation.Required),
		validation.Field(&params.PetugasRekamPTPOS, validation.Required),
		validation.Field(&params.ChId, validation.Required),
		validation.Field(&params.ReqId, validation.Required),
		validation.Field(&params.NikPelapor, validation.Required),
		validation.Field(&params.NamaPelapor, validation.Required),
		validation.Field(&params.NikPeserta, validation.Required),
		validation.Field(&params.NikPelapor, validation.Required),
		validation.Field(&params.TglLahirPelapor, validation.Required),
		validation.Field(&params.EmailPelapor, validation.Required),
		validation.Field(&params.TglLahirPelapor, validation.Required),
		validation.Field(&params.HandphonePelapor, validation.Required),
		validation.Field(&params.TglPengajuan, validation.Required),
		validation.Field(&params.TglLahirPelapor, validation.Required),
		validation.Field(&params.KodeBilling, validation.Required),
		validation.Field(&params.SkorFace, validation.Required),
		validation.Field(&params.KodeBilling, validation.Required),
		validation.Field(&params.KemiripanNamaPelapor, validation.Required),
		validation.Field(&params.KeteranganApproval, validation.Required),
		validation.Field(&params.PenerimaBeasiswa, validation.Required),
		validation.Field(&params.DataDokumenPelapor, validation.Required),
		validation.Field(&params.DataProbingInsertKonfBeasiswa, validation.Required),
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

func DaftarJenisBeasiswa(ctx *fiber.Ctx, params request.DaftarJenisBeasiswaRequest, requestId string) {

	// for _, penerimaBeasiswa := range params.PenerimaBeasiswa {
	// 	validation.Field(penerimaBeasiswa.NamaPenerimaBeasiswa, validation.Required)
	// }

	err := validation.ValidateStruct(&params,
		validation.Field(&params.NikPelapor, validation.Required),
		validation.Field(&params.NamaPelapor, validation.Required),
		validation.Field(&params.NikPeserta, validation.Required),
		validation.Field(&params.ChannelID, validation.Required),
		validation.Field(&params.KodeKantorPTPOS, validation.Required),
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

func DaftarJenjangPendidikan(ctx *fiber.Ctx, params request.DaftarJenjangPendidikanRequest, requestId string) {

	// for _, penerimaBeasiswa := range params.PenerimaBeasiswa {
	// 	validation.Field(penerimaBeasiswa.NamaPenerimaBeasiswa, validation.Required)
	// }

	err := validation.ValidateStruct(&params,
		validation.Field(&params.NikPelapor, validation.Required),
		validation.Field(&params.NamaPelapor, validation.Required),
		validation.Field(&params.NikPeserta, validation.Required),
		validation.Field(&params.ChannelID, validation.Required),
		validation.Field(&params.KodeKantorPTPOS, validation.Required),
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

func DaftarPenerimaBeasiswa(ctx *fiber.Ctx, params request.DaftarPenerimaBeasiswaRequest, requestId string) {

	// for _, penerimaBeasiswa := range params.PenerimaBeasiswa {
	// 	validation.Field(penerimaBeasiswa.NamaPenerimaBeasiswa, validation.Required)
	// }

	err := validation.ValidateStruct(&params,
		validation.Field(&params.NikPelapor, validation.Required),
		validation.Field(&params.NamaPelapor, validation.Required),
		validation.Field(&params.NikPeserta, validation.Required),
		validation.Field(&params.ChannelID, validation.Required),
		validation.Field(&params.KodeKantorPTPOS, validation.Required),
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

func CariNominalPerJenjang(ctx *fiber.Ctx, params request.NominalPerjenjangBeasiswaRequest, requestId string) {

	// for _, penerimaBeasiswa := range params.PenerimaBeasiswa {
	// 	validation.Field(penerimaBeasiswa.NamaPenerimaBeasiswa, validation.Required)
	// }

	err := validation.ValidateStruct(&params,
		validation.Field(&params.NikPelapor, validation.Required),
		validation.Field(&params.NamaPelapor, validation.Required),
		validation.Field(&params.NikPeserta, validation.Required),
		validation.Field(&params.ChannelID, validation.Required),
		validation.Field(&params.KodeKantorPTPOS, validation.Required),
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
