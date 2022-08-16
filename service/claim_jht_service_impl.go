package service

import (
	"fmt"
	// "js-ptpos/entity"

	"js-ptpos/exception"
	model "js-ptpos/model"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
	"js-ptpos/repository"
	"js-ptpos/util"

	"js-ptpos/validation"

	"github.com/gofiber/fiber/v2"
)

func NewJHTClaimService(ClaimJHTRepository *repository.ClaimJHTRepository) ClaimJHTService {
	return &claimJHTServiceImpl{
		ClaimJHTRepository: *ClaimJHTRepository,
	}
}

type claimJHTServiceImpl struct {
	ClaimJHTRepository repository.ClaimJHTRepository
}

func (service *claimJHTServiceImpl) CheckEligibleJHT(ctx *fiber.Ctx, requestId string, params *request.CheckEligibleJHTRequest) (*response.CheckJHTEligibleResponse, error) {
	validation.CheckEligibleJHT(ctx, *params, requestId)

	res, err := service.ClaimJHTRepository.CheckJHTEligible(params)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	result := &response.CheckJHTEligibleResponse{
		StatusCode: res.StatusCode,
		StatusDesc: res.StatusDesc,
		Data:       res.Data,
	}

	return result, nil
}

func (service *claimJHTServiceImpl) GetPengajuanJHT(ctx *fiber.Ctx, requestId string, params *request.GetPengajuanJHTRequest) (*model.PengajuanJHTResponse, error) {
	validation.GetPengajuanJHTValidate(ctx, *params, requestId)

	res, err := service.ClaimJHTRepository.GetPengajuanJHT(params)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	var data []model.PengajuanJHT
	var dataDetil model.PengajuanJHT

	if len(res) < 1 {
		message := "Data tidak ditemukan"
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.DataNotFoundError{
			Message: message,
		})

	}

	for _, item := range res {
		dataDetil.KodeSegmen = item.KodeSegmen
		dataDetil.NPP = item.NPP
		dataDetil.KodeTK = item.KodeTK
		data = append(data, dataDetil)
	}
	result := &model.PengajuanJHTResponse{
		StatusCode: 200,
		Data:       data,
	}

	return result, nil

}

func (service *claimJHTServiceImpl) InsertPengajuanJHT(ctx *fiber.Ctx, requestId string, params *request.InsertPengajuanJHTRequest) (*model.InsertPengajuanJHT, error) {
	// validation.InsertClaimJhtValidate(ctx, *params, requestId)

	res, err := service.ClaimJHTRepository.InsertPengajuanJHT(params)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	result := &model.InsertPengajuanJHT{
		StatusCode:         res.StatusCode,
		StatusKirim:        res.StatusKirim,
		KeteranganKirim:    res.KeteranganKirim,
		KodePengajuanKlaim: res.KodePengajuanKlaim,
	}

	return result, nil
}
func (service *claimJHTServiceImpl) DaftarSegmen(ctx *fiber.Ctx, requestId string, params *request.DaftarSegmenRequest) (*model.DaftarSegmenResponse, error) {
	validation.DaftarSegmenJHTValidate(ctx, *params, requestId)

	res, err := service.ClaimJHTRepository.DaftarSegmen(params)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	var data []model.DaftarSegmen
	var dataDetil model.DaftarSegmen

	if len(res) < 1 {
		message := "Data tidak ditemukan"
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.DataNotFoundError{
			Message: message,
		})

	}

	for _, item := range res {
		dataDetil.Kode = item.Kode
		dataDetil.Keterangan = item.Keterangan
		data = append(data, dataDetil)
	}
	result := &model.DaftarSegmenResponse{
		StatusCode: 200,
		StatusDesc: "Sukses",
		Data:       data,
	}

	return result, nil
}
func (service *claimJHTServiceImpl) DaftarSebabKlaim(ctx *fiber.Ctx, requestId string, params *request.DaftarSebabKlaimRequest) (*model.DaftarSebabKlaimResponse, error) {
	validation.GetDaftarSebabKlaimValidate(ctx, *params, requestId)

	res, err := service.ClaimJHTRepository.DaftarSebabKlaim(params)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	var data []model.DaftarSebabKlaim
	var dataDetil model.DaftarSebabKlaim

	if len(res) < 1 {
		message := "Data tidak ditemukan"
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.DataNotFoundError{
			Message: message,
		})

	}

	for _, item := range res {
		dataDetil.KodeSebabKlaim = item.KodeSebabKlaim
		dataDetil.NamaSebabKlaim = item.NamaSebabKlaim
		data = append(data, dataDetil)
	}
	result := &model.DaftarSebabKlaimResponse{
		StatusCode: 200,
		StatusDesc: "Sukses",
		Data:       data,
	}

	return result, nil
}
func (service *claimJHTServiceImpl) DaftarDokumenSebabKlaim(ctx *fiber.Ctx, requestId string, params *request.DaftarDokumenSebabKlaimRequest) (*model.DaftarDokumenSebabKlaimResponse, error) {
	validation.GetDokumenBySebabKlaimJHTValidate(ctx, *params, requestId)

	res, err := service.ClaimJHTRepository.DaftarDokumenSebabKlaim(params)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	var data []model.DaftarDokumenSebabKlaim
	var dataDetil model.DaftarDokumenSebabKlaim

	if len(res) < 1 {
		message := "Data tidak ditemukan"
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.DataNotFoundError{
			Message: message,
		})

	}

	for _, item := range res {
		dataDetil.KodeDokumen = item.KodeDokumen
		dataDetil.NamaDokumen = item.NamaDokumen
		dataDetil.FlagMandatory = item.FlagMandatory
		data = append(data, dataDetil)
	}
	result := &model.DaftarDokumenSebabKlaimResponse{
		StatusCode: 200,
		StatusDesc: "Sukses",
		Data:       data,
	}

	return result, nil
}
func (service *claimJHTServiceImpl) UpdatePengajuanJHT(ctx *fiber.Ctx, requestId string, params *request.UpdatePengajuanJHTRequest) (*model.UpdatePengajuanJHT, error) {
	// validation.GetPengajuanJHTValidate(ctx, *params, requestId)

	res, err := service.ClaimJHTRepository.UpdatePengajuanJHT(params)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	result := &model.UpdatePengajuanJHT{
		StatusKirim:     res.StatusKirim,
		KeteranganKirim: res.KeteranganKirim,
	}

	return result, nil
}
