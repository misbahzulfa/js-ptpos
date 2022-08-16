package service

import (
	"fmt"
	"js-ptpos/exception"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
	"js-ptpos/repository"
	"js-ptpos/util"
	"js-ptpos/validation"

	"github.com/gofiber/fiber/v2"
)

func NewClaimJPService(claimJPRepository *repository.ClaimJPRepository) ClaimJPService {
	return &claimJPServiceImpl{
		ClaimJPRepository: *claimJPRepository,
	}
}

type claimJPServiceImpl struct {
	ClaimJPRepository repository.ClaimJPRepository
}

func (service *claimJPServiceImpl) CheckEligibleJP(ctx *fiber.Ctx, requestId string, params *request.CheckEligibleJPBerkalaRequest) (*response.CheckEligibleJPBerkalaResponse, error) {
	validation.CheckEligibleJPValidate(ctx, *params, requestId)
	res, err := service.ClaimJPRepository.JPCheckEligible(params)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	result := &response.CheckEligibleJPBerkalaResponse{
		StatusCode:       res.StatusCode,
		StatusDesc:       res.StatusDesc,
		DataPenJPBerkala: res.DataPenJPBerkala,
	}

	return result, nil
}

func (service *claimJPServiceImpl) CheckClaimByNIKPelapor(ctx *fiber.Ctx, requestId string, params *request.CheckJumlahKlaimJPBerkalaNikPelaporRequest) (*response.CheckJumlahKlaimJPBerkalaNikPelaporResponse, error) {

	res, err := service.ClaimJPRepository.JPCheckJumlahKlaimBerkalaByNIKPelapor(params)

	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	result := &response.CheckJumlahKlaimJPBerkalaNikPelaporResponse{
		StatusLebihDariSatu: res.StatusLebihDariSatu,
	}

	return result, nil
}

func (service *claimJPServiceImpl) InsertConfirmationJPBerkala(ctx *fiber.Ctx, requestId string, params *request.InsertJPConfirmationRequest) (*response.InsertKonfirmasiJPResponse, error) {
	validation.InsertJPConfirmationValidate(ctx, *params, requestId)

	insertConfirmationRequest := &request.InsertJPConfirmationRequest{
		ChannelID:                     params.ChannelID,
		KodeKantorPtPos:               params.KodeKantorPtPos,
		PetugasRekamPtPos:             params.PetugasRekamPtPos,
		NikPelapor:                    params.NikPelapor,
		NamaPelapor:                   params.NamaPelapor,
		NikPenerimaManfaat:            params.NikPenerimaManfaat,
		NamaPenerimaManfaat:           params.NamaPenerimaManfaat,
		TanggalLahirPenerimaManfaat:   params.TanggalLahirPenerimaManfaat,
		NikPeserta:                    params.NikPeserta,
		TanggalLahirPelapor:           params.TanggalLahirPelapor,
		EmailPelapor:                  params.EmailPelapor,
		NoHPPelapor:                   params.NoHPPelapor,
		TanggalPengajuan:              params.TanggalPengajuan,
		KodeBillingPTPos:              params.KodeBillingPTPos,
		ScoreFaceMatch:                params.ScoreFaceMatch,
		SimilarityNamaPTPOSkeAdminduk: params.SimilarityNamaPTPOSkeAdminduk,
		KeteranganKonfirmasi:          params.KeteranganKonfirmasi,
		DataDokumen:                   params.DataDokumen,
		DataProbing:                   params.DataProbing,
	}

	insertConfirmationJP, err := service.ClaimJPRepository.InsertConfirmationJPBerkala(insertConfirmationRequest)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	return insertConfirmationJP, nil
}

// func (service *claimJPServiceImpl) UpdateKonfirmasiJPECtoPN(ctx *fiber.Ctx, requestId string, params *request.UpdateKonfirmasiJPECtoPNRequest) (*response.UpdateKonfirmasiJPECtoPNResponse, error) {

// 	var dokuments []request.DataDokumenECtoPN
// 	var dokument request.DataDokumenECtoPN
// 	for _, list := range params.DataDokumen {
// 		dokument.KodeDokumen = list.KodeDokumen
// 		dokument.PathURL = list.PathURL
// 		dokuments = append(dokuments, dokument)
// 	}

// 	updateKonfirmasiECtoPN := &request.UpdateKonfirmasiJPECtoPNRequest{
// 		KPJ:                   params.KPJ,
// 		NikPeserta:            params.NikPeserta,
// 		NamaPeserta:           params.NamaPeserta,
// 		KodeTK:                params.KodeTK,
// 		KodeKlaim:             params.KodeKlaim,
// 		HubunganDenganPelapor: params.HubunganDenganPelapor,
// 		NamaHubunganLainnya:   params.NamaHubunganLainnya,
// 		NamaPelapor:           params.NamaPelapor,
// 		NikPelapor:            params.NikPelapor,
// 		TempatLahirPelapor:    params.TempatLahirPelapor,
// 		TanggalLahirPelapor:   params.TanggalLahirPelapor,
// 		JenisKelaminPelapor:   params.JenisKelaminPelapor,
// 		GolonganDarahPelapor:  params.GolonganDarahPelapor,
// 		AlamatDomisiliPelapor: params.AlamatDomisiliPelapor,
// 		KodePosPelapor:        params.KodePosPelapor,
// 		KelurahanPelapor:      params.KelurahanPelapor,
// 		KecamatanPelapor:      params.KecamatanPelapor,
// 		KabupatenPelapor:      params.KabupatenPelapor,
// 		Email:                 params.Email,
// 		NoHP:                  params.NoHP,
// 		NamaBank:              params.NamaBank,
// 		NomorRekening:         params.NomorRekening,
// 		NamaRekening:          params.NamaRekening,
// 		FaceMatchScore:        params.FaceMatchScore,
// 		DataDokumen:           dokuments,
// 	}

// 	insertConfirmationJP, err := service.ClaimJPRepository.UpdateKonfirmasiJPECtoPN(updateKonfirmasiECtoPN)
// 	if err != nil {
// 		message := err.Error()
// 		logStop := util.LogResponse(ctx, message, requestId)
// 		fmt.Println(logStop)

// 		panic(exception.GeneralError{
// 			Message: message,
// 		})
// 	}

// 	return insertConfirmationJP, nil
// }
