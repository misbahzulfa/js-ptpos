package service

import (
	"fmt"
	"js-ptpos/exception"
	model "js-ptpos/model"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
	"js-ptpos/repository"
	"js-ptpos/util"
	"js-ptpos/validation"

	"github.com/gofiber/fiber/v2"
)

func NewBeasiswaService(BeasiswaRepository *repository.BeasiswaRepository) BeasiswaService {
	return &beasiswaServiceImpl{
		BeasiswaRepository: *BeasiswaRepository,
	}
}

type beasiswaServiceImpl struct {
	BeasiswaRepository repository.BeasiswaRepository
}

func (service *beasiswaServiceImpl) CheckEligibleBeasiswa(ctx *fiber.Ctx, requestId string, params *request.CheckEligibleBeasiswaRequest) (*response.CheckEligibleBeasiswaResponse, error) {
	validation.CheckEligibleBeasiswaValidate(ctx, *params, requestId)

	res, err := service.BeasiswaRepository.CheckEligibleBeasiswa(params)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	result := &response.CheckEligibleBeasiswaResponse{
		StatusCode:            res.StatusCode,
		StatusDesc:            res.StatusDesc,
		CheckEligibleBeasiswa: res.CheckEligibleBeasiswa,
	}

	return result, nil
}

func (service *beasiswaServiceImpl) InsertKonfirmasiBeasiswa(ctx *fiber.Ctx, requestId string, params *request.InsertKonfirmasiBeasiswaRequest) (*response.InsertKonfirmasiBeasiswaResponse, error) {
	validation.InsertKonfirmasiBeasiswa(ctx, *params, requestId)

	// res, err := service.BeasiswaRepository.InsertKonfirmasiBeasiswa(params)
	// if err != nil {
	// 	message := err.Error()
	// 	logStop := util.LogResponse(ctx, message, requestId)
	// 	fmt.Println(logStop)

	// 	panic(exception.GeneralError{
	// 		Message: message,
	// 	})
	// }

	//var statusKirim string
	//var keteranganSukses string

	//statusKirim = ""
	//keteranganSukses = ""

	InsertKonfirmasiBeasiswa, err := service.BeasiswaRepository.InsertKonfirmasiBeasiswa(params)
	if err != nil {

		message := err.Error()
		//statusKirim = "GAGAL"
		//keteranganSukses = message

		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})

	}
	// else {
	// 	//statusKirim = "BERHASIL"
	// 	keteranganSukses = "PROSES INSERT BERHASIL"
	// }

	// result := &response.InsertKonfirmasiBeasiswaResponse{
	// 	//StatusCode:               statusKirim,
	// 	StatusDesc:               keteranganSukses,
	// 	InsertKonfirmasiBeasiswa: InsertKonfirmasiBeasiswa.InsertKonfirmasiBeasiswa,
	// }

	return InsertKonfirmasiBeasiswa, nil

}

func (service *beasiswaServiceImpl) DaftarJenisBeasiswa(ctx *fiber.Ctx, requestId string, params *request.DaftarJenisBeasiswaRequest) (*response.DaftarJenisBeasiswaResponse, error) {
	validation.DaftarJenisBeasiswa(ctx, *params, requestId)

	res, err := service.BeasiswaRepository.DaftarJenisBeasiswa(params)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	result := &response.DaftarJenisBeasiswaResponse{
		StatusCode:        res.StatusCode,
		StatusDesc:        res.StatusDesc,
		DataJenisBeasiswa: res.DataJenisBeasiswa,
		// StatusPencarian:     res.StatusPencarian,
		// KeteranganPencarian: res.KeteranganPencarian,
		// DaftarDokumen:       res.DaftarDokumen,
		// DaftarProbing:       res.DaftarProbing,
	}

	// var DaftarJenisBeasiswaResult model.DaftarJenisBeasiswa
	// var DaftarJenisBeasiswaResults []model.DaftarJenisBeasiswa

	// // result := &model.DaftarJenisBeasiswa{
	// // 	KodeJenisBeasiswa: res.KodeJenisBeasiswa,
	// // 	NamaJenisBeasiswa: res.NamaJenisBeasiswa,
	// // }

	// for _, item := range res {
	// 	DaftarJenisBeasiswaResult.KodeJenisBeasiswa = item.KodeJenisBeasiswa
	// 	DaftarJenisBeasiswaResult.NamaJenisBeasiswa = item.NamaJenisBeasiswa
	// 	DaftarJenisBeasiswaResults = append(DaftarJenisBeasiswaResults, DaftarJenisBeasiswaResult)
	// }

	return result, nil
}

func (service *beasiswaServiceImpl) DaftarJenjangPendidikan(ctx *fiber.Ctx, requestId string, params *request.DaftarJenjangPendidikanRequest) ([]model.DaftarJenjangPendidikan, error) {
	validation.DaftarJenjangPendidikan(ctx, *params, requestId)

	res, err := service.BeasiswaRepository.DaftarJenjangPendidikan(params)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	var DaftarJenjangPendidikanResult model.DaftarJenjangPendidikan
	var DaftarJenjangPendidikanResults []model.DaftarJenjangPendidikan

	// result := &model.DaftarJenisBeasiswa{
	// 	KodeJenisBeasiswa: res.KodeJenisBeasiswa,
	// 	NamaJenisBeasiswa: res.NamaJenisBeasiswa,
	// }

	for _, item := range res {
		DaftarJenjangPendidikanResult.KodeJenisBeasiswa = item.KodeJenisBeasiswa
		DaftarJenjangPendidikanResult.KodeJenjangPendidikan = item.KodeJenjangPendidikan
		DaftarJenjangPendidikanResult.NamaJenjangPendidikan = item.NamaJenjangPendidikan
		DaftarJenjangPendidikanResults = append(DaftarJenjangPendidikanResults, DaftarJenjangPendidikanResult)
	}

	return DaftarJenjangPendidikanResults, nil
}

func (service *beasiswaServiceImpl) DaftarPenerimaBeasiswa(ctx *fiber.Ctx, requestId string, params *request.DaftarPenerimaBeasiswaRequest) (*response.DaftarPenerimaBeasiswaResponse, error) {
	validation.DaftarPenerimaBeasiswa(ctx, *params, requestId)

	res, err := service.BeasiswaRepository.DaftarPenerimaBeasiswa(params)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	return res, nil

}
func (service *beasiswaServiceImpl) NominalPerjenjangBeasiswa(ctx *fiber.Ctx, requestId string, params *request.NominalPerjenjangBeasiswaRequest) (*response.NominalPerjenjangBeasiswaResponse, error) {
	validation.CariNominalPerJenjang(ctx, *params, requestId)

	res, err := service.BeasiswaRepository.NominalPerjenjangBeasiswa(params)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	return res, nil

}
