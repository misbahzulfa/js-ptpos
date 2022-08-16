package service

import (
	"fmt"
	"js-ptpos/exception"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
	"js-ptpos/repository"
	"js-ptpos/util"
	"js-ptpos/validation"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func NewClaimService(claimRepository *repository.ClaimRepository) ClaimService {
	return &claimServiceImpl{
		ClaimRepository: *claimRepository,
	}
}

type claimServiceImpl struct {
	ClaimRepository repository.ClaimRepository
}

func (service *claimServiceImpl) CheckAccountBank(ctx *fiber.Ctx, requestId string, params *request.CheckAccountBankRemoteRequest) (*response.CheckAccountBankRemoteResponse, error) {

	//params.Email = claims.Email

	//validation.CheckAccountBankValidate(ctx, *params, requestId)

	checkAccountBankRemoteRequest := &request.CheckAccountBankRemoteRequest{
		ChID:        params.ChID,
		ReqID:       params.ReqID,
		Bank:        params.Bank,
		KODEBANKATB: params.KODEBANKATB,
		NOREKTUJUAN: params.NOREKTUJUAN,
		NAMAREK:     params.NAMAREK,
	}

	res, err := service.ClaimRepository.CheckAccountBank(checkAccountBankRemoteRequest)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}

	if res.Ret == "0" {
		resRek := &response.CheckAccountBankRemoteResponse{
			Ret:  "200",
			Msg:  "Rekening Valid",
			Data: res.Data,
		}
		return resRek, nil
	} else {
		message := "Pastikan Nomor Rekening Benar"
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
	// return resRek, nil
}
func (service *claimServiceImpl) InsertClaimJht(ctx *fiber.Ctx, requestId string, params *request.InsertClaimJhtRequest) (*response.InsertClaimJhtResponse, error) {
	validation.InsertClaimJhtValidate(ctx, *params, requestId)

	CheckEligibleJHTRequest := &request.CheckEligibleJHTRequest{
		Nik:               params.IdentityNumber,
		Kpj:               params.Kpj,
		Fullname:          params.FullName,
		ChannelId:         params.ChannelId,
		KodeKantorPPTOS:   params.KodeKantorPPTOS,
		PetugasRekamPTPOS: params.KodeKantorPPTOS,
		ChId:              params.ChId,
		ReqId:             params.ReqId,
		TglLahir:          params.BirthDate,
	}

	resEligible, err := service.ClaimRepository.CheckEligibleInsert(CheckEligibleJHTRequest)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)

		panic(exception.GeneralError{
			Message: message,
		})
	}
	if resEligible.Data[0].KodeKelayakan == "JHTA000" {

		checkAccountBankRemoteRequest := &request.CheckAccountBankRemoteRequest{
			ChID:        params.ChId,
			ReqID:       params.ReqId,
			Bank:        params.BankName,
			KODEBANKATB: params.BankCode,
			NOREKTUJUAN: params.AccountBankNumber,
			NAMAREK:     params.AccountBankName,
		}

		res, err := service.ClaimRepository.CheckAccountBank(checkAccountBankRemoteRequest)
		if err != nil {
			message := err.Error()
			logStop := util.LogResponse(ctx, message, requestId)
			fmt.Println(logStop)

			panic(exception.GeneralError{
				Message: message,
			})
		}

		fmt.Println(res, "service")

		if res.Ret == "0" {
			accountNameResult := strings.ToUpper(res.Data.NAMAREKTUJUAN)
			accountNameParam := strings.ToUpper(params.AccountBankName)

			accountNameParamRune := []rune(accountNameParam)
			accountName := string(accountNameParamRune[0:len(accountNameResult)])

			fmt.Println(accountName, accountNameResult)
			// if accountNameResult == accountName {
			if 1 == 1 {
				similaryRequest := &request.SimilarityRequest{
					ChId:  params.ChId,
					ReqId: params.ReqId,
					Nama1: res.Data.NAMAREKTUJUAN,
					Nama2: params.FullName,
					Kode:  "1",
				}
				similaryResponse, err := service.ClaimRepository.SimilairtyAccountBankName(similaryRequest)
				if err != nil {
					message := err.Error()
					logStop := util.LogResponse(ctx, message, requestId)
					fmt.Println(logStop)

					panic(exception.GeneralError{
						Message: message,
					})
				}

				// fmt.Println(similaryResponse, "score rekening")
				if similaryResponse.Score < similaryResponse.MinScore {
					message := "similarity nama rekening dan nama peserta kurang dari 75 %"
					logStop := util.LogResponse(ctx, message, requestId)
					fmt.Println(logStop)

					result := &response.InsertClaimJhtResponse{
						StatusCode: 200,
						StatusDesc: "OK",
						Data: []response.InsertClaimJht{{
							StatusKirim:        "T",
							KeteranganKirim:    message,
							KodePengajuanKirim: "",
						}},
					}
					return result, nil
				}

				//start
				getDataEmployeeRequest := &request.GetDataEmployeeRequest{
					SegmenCode:     params.SegmenCode,
					Kpj:            params.Kpj,
					WorkerCode:     params.WorkerCode,
					IdentityNumber: params.IdentityNumber,
					FullName:       params.FullName,
					BirthDate:      params.BirthDate,
				}
				// fmt.Println(getDataEmployeeRequest, "data employee")
				resultDetailAsik, err := service.ClaimRepository.GetDataEmployee(getDataEmployeeRequest)
				if err != nil {
					message := err.Error()
					logStop := util.LogResponse(ctx, message, requestId)
					fmt.Println(logStop)

					panic(exception.GeneralError{
						Message: message,
					})
				}
				// fmt.Println("lanjut")

				claimBenefitDetailRequest := &request.ClaimBenefitDetailRequest{
					SegmenCode:       params.SegmenCode,
					Kpj:              params.Kpj,
					WorkerCode:       params.WorkerCode,
					IdentityNumber:   params.IdentityNumber,
					FullName:         params.FullName,
					BirthDate:        params.BirthDate,
					CauseOfClaimCode: params.ClaimCauseCode,
				}

				resultDetailBenefit, err := service.ClaimRepository.GetClaimBenefitDetail(claimBenefitDetailRequest)
				if err != nil {
					message := err.Error()
					logStop := util.LogResponse(ctx, message, requestId)
					fmt.Println(logStop)

					panic(exception.GeneralError{
						Message: message,
					})
				}

				var sumberBlthNonAktif string
				var flagSumberBlthNonAktif string
				var tglSumberBlthNonAktif string
				currentTime := time.Now()

				if resultDetailAsik.FlagSipp == "1" {
					sumberBlthNonAktif = "1"
					flagSumberBlthNonAktif = "T"
					tglSumberBlthNonAktif = currentTime.Format("02-01-2006")
				} else {
					sumberBlthNonAktif = "0"
					flagSumberBlthNonAktif = "Y"
					tglSumberBlthNonAktif = currentTime.Format("02-01-2006")
				}

				kodeProgram, _ := strconv.Atoi(resultDetailBenefit.ProgramCode)
				ratePengembangan, err := strconv.ParseFloat(resultDetailBenefit.RatePengembangan, 64)
				nominalSaldoAwalTahun, _ := strconv.ParseFloat(resultDetailBenefit.NominalSaldoAwalTahun, 64)
				nominalSaldoPengembangan, _ := strconv.ParseFloat(resultDetailBenefit.NominalSaldoPengembangan, 64)
				nominalSaldoTotal, _ := strconv.ParseFloat(resultDetailBenefit.NominalSaldoTotal, 64)
				nominalIuranTahunBerjalan, _ := strconv.ParseFloat(resultDetailBenefit.NominalIuranTahunBerjalan, 64)
				nominalIuranPengembangan, _ := strconv.ParseFloat(resultDetailBenefit.NominalIuranPengembangan, 64)
				nominalIuranTotal, _ := strconv.ParseFloat(resultDetailBenefit.NominalIuranTotal, 64)
				nominalSaldoIuranTotal, _ := strconv.ParseFloat(resultDetailBenefit.NominalSaldoIuranTotal, 64)
				persentasePengambilan, _ := strconv.ParseFloat(resultDetailBenefit.PersentasePengambilan, 64)
				nominalManfaatMaxBisaDiAmbil, _ := strconv.ParseFloat(resultDetailBenefit.NominalManfaatMaxBisaDiAmbil, 64)
				nominalManfaatDiAmbil, _ := strconv.ParseFloat(resultDetailBenefit.NominalManfaatDiAmbil, 64)
				nominalManfaatGross, _ := strconv.ParseFloat(resultDetailBenefit.NominalManfaatGross, 64)
				nominalPPH, _ := strconv.ParseFloat(resultDetailBenefit.NominalPPH, 64)
				nominalPembulatan, _ := strconv.ParseFloat(resultDetailBenefit.NominalPembulatan, 64)
				nominalManfaatNetto, _ := strconv.ParseFloat(resultDetailBenefit.NominalManfaatNetto, 64)

				insertClaimRequest := &request.InsertClaimRequest{
					Email:                     params.Email,
					PhoneNumber:               params.PhoneNumber,
					WorkerCode:                resultDetailAsik.WorkerCode,
					Kpj:                       params.Kpj,
					IdentityNumber:            params.IdentityNumber,
					IdentityType:              resultDetailAsik.IdentityType,
					ValidIdentity:             resultDetailAsik.ValidIdentity,
					FullName:                  params.FullName,
					BirthDate:                 params.BirthDate,
					BirthPlace:                params.BirthPlace,
					Gender:                    resultDetailAsik.Gender,
					MotherName:                params.MotherName,
					OfficeCode:                resultDetailAsik.OfficeCode,
					SegmenCode:                params.SegmenCode,
					CompanyCode:               resultDetailAsik.CompanyCode,
					DivisionCode:              resultDetailAsik.DivisionCode,
					MembershipCode:            resultDetailAsik.MembershipCode,
					MembershipDate:            resultDetailAsik.MembershipDate,
					ActiveDate:                resultDetailAsik.ActiveDate,
					NonActiveDate:             resultDetailAsik.NonActiveDate,
					NonActiveCode:             resultDetailAsik.NonActiveCode,
					SumberBlthNonAktif:        sumberBlthNonAktif,
					FlagSubmerNonAktif:        flagSumberBlthNonAktif,
					TglFlagSumberNonAktif:     tglSumberBlthNonAktif,
					ClaimTypeCode:             "JHT01",
					ClaimCauseCode:            params.ClaimCauseCode,
					ReportCode:                resultDetailBenefit.ReportCode,
					Npwp:                      params.Npwp,
					BankCode:                  params.BankCode,
					BankName:                  params.BankName,
					AccountBankNumber:         params.AccountBankNumber,
					AccountBankName:           params.AccountBankName,
					Address:                   params.Address,
					KelurahanCode:             params.KelurahanCode,
					KecamatanCode:             params.KecamatanCode,
					KabupatenCode:             params.KabupatenCode,
					PropinsiCode:              params.PropinsiCode,
					PostalCode:                resultDetailAsik.PostalCode,
					ScoreFace:                 params.ScoreFace,
					ScoreFaceLiveness:         params.ScoreFaceLiveness,
					Live:                      params.Live,
					TrxId:                     params.TrxId,
					RefId:                     params.RefId,
					BenefitCode:               resultDetailBenefit.BenefitCode,
					ReceiverTypeCode:          resultDetailBenefit.ReceiverTypeCode,
					ProgramCode:               kodeProgram,
					TanggalPengembangan:       resultDetailBenefit.TanggalPengembangan,
					RatePengembangan:          ratePengembangan,
					TanggalSaldoAwalTahun:     resultDetailBenefit.TanggalSaldoAwalTahun,
					NominalSaldoAwalTahun:     nominalSaldoAwalTahun,
					NominalSaldoPengembangan:  nominalSaldoPengembangan,
					NominalSaldoTotal:         nominalSaldoTotal,
					NominalIuranTahunBerjalan: nominalIuranTahunBerjalan,
					NominalIuranPengembangan:  nominalIuranPengembangan,
					NominalIuranTotal:         nominalIuranTotal,
					NominalSaldoIuranTotal:    nominalSaldoIuranTotal,
					PersentasePengambilan:     persentasePengambilan,
					NominalManfaatBisaDiAmbil: nominalManfaatMaxBisaDiAmbil,
					NominalManfaatDiAmbil:     nominalManfaatDiAmbil,
					NominalManfaatGross:       nominalManfaatGross,
					NominalPPH:                nominalPPH,
					NominalPembulatan:         nominalPembulatan,
					NominalManfaatNetto:       nominalManfaatNetto,
					KodeKantorPengajuan:       resultDetailAsik.OfficeCode,
					PetugasRekam:              params.Email,
					PathUrl:                   params.PathUrl,
					DataDetailProbing:         params.DataDetailProbing,
					Platform:                  params.Platform,
					TanggalPengajuan:          params.TanggalPengajuan,
					KodeBillingPTPOS:          params.KodeBillingPTPOS,
					ScoreSimilarityNama:       params.ScoreSimilarityNama,
					KeteranganPengajuan:       params.KeteranganPengajuan,
					ChannelId:                 params.ChannelId,
					KodeKantorPTPOS:           params.KodeKantorPPTOS,
					PetugasPTPOS:              params.PetugasRekamPTPOS,
					ReqId:                     params.ReqId,
					ChId:                      params.ChId,
				}
				// fmt.Println("lanjut3")

				insertClaimJht, err := service.ClaimRepository.InsertClaimJht(insertClaimRequest)
				if err != nil {
					message := err.Error()
					logStop := util.LogResponse(ctx, message, requestId)
					fmt.Println(logStop)

					// panic(exception.GeneralError{
					// 	Message: message,
					// })
					// fmt.Println(submisstionCode)
					result := &response.InsertClaimJhtResponse{
						StatusCode: 400,
						StatusDesc: message,
						Data: []response.InsertClaimJht{{
							StatusKirim:        "T",
							KeteranganKirim:    "Gagal",
							KodePengajuanKirim: "",
						}},
					}
					return result, nil

				}
				// fmt.Println("lanjut4")

				submisstionCode := insertClaimJht

				sendEmailRequest := &request.SendEmailRequest{
					Email:            params.Email,
					OfficeCode:       resultDetailAsik.OfficeCode,
					FullName:         params.FullName,
					KodeKantorPPTOS:  params.KodeKantorPPTOS,
					KodePengajuan:    submisstionCode,
					TanggalPengajuan: params.TanggalPengajuan,
					Kpj:              params.Kpj,
				}
				// fmt.Println("lanjut6")
				sendEmail, err := service.ClaimRepository.SendEmail(sendEmailRequest)
				if err != nil {
					message := err.Error()
					logStop := util.LogResponse(ctx, message, requestId)
					fmt.Println(logStop)

					panic(exception.GeneralError{
						Message: message,
					})
				}
				// fmt.Println("lanjut7")

				if sendEmail.Message == "Success" {
					message := sendEmail.Message
					logStop := util.LogResponse(ctx, message, requestId)
					fmt.Println(logStop)
				}
				// fmt.Println(submisstionCode)
				result := &response.InsertClaimJhtResponse{
					StatusCode: 200,
					StatusDesc: "OK",
					Data: []response.InsertClaimJht{{
						StatusKirim:        "Y",
						KeteranganKirim:    "Data berhasil disimpan",
						KodePengajuanKirim: submisstionCode,
					}},
				}

				return result, nil
			} else {
				message := "parameter nama rekening tidak sama dengan nama rekening bank"
				logStop := util.LogResponse(ctx, message, requestId)
				fmt.Println(logStop)

				result := &response.InsertClaimJhtResponse{
					StatusCode: 200,
					StatusDesc: "OK",
					Data: []response.InsertClaimJht{{
						StatusKirim:        "T",
						KeteranganKirim:    message,
						KodePengajuanKirim: "",
					}},
				}
				return result, nil
			}
			// end
		} else {
			message := "Pastikan Nomor Rekening Benar"
			logStop := util.LogResponse(ctx, message, requestId)
			fmt.Println(logStop)

			// panic(exception.GeneralError{
			// 	Message: message,
			// })
			result := &response.InsertClaimJhtResponse{
				StatusCode: 200,
				StatusDesc: "OK",
				Data: []response.InsertClaimJht{{
					StatusKirim:        "T",
					KeteranganKirim:    message,
					KodePengajuanKirim: "",
				}},
			}
			return result, nil
		}
	} else {
		result := &response.InsertClaimJhtResponse{
			StatusCode: 200,
			StatusDesc: "OK",
			Data: []response.InsertClaimJht{{
				StatusKirim:        "T",
				KeteranganKirim:    resEligible.Data[0].KeteranganKelayakan,
				KodePengajuanKirim: "",
			}},
		}
		return result, nil
	}
}
func (service *claimServiceImpl) CheckEligibleInsert(ctx *fiber.Ctx, requestId string, params *request.CheckEligibleJHTRequest) (*response.CheckJHTEligibleResponse, error) {
	validation.CheckEligibleJHT(ctx, *params, requestId)

	CheckEligibleJHTRequest := &request.CheckEligibleJHTRequest{
		Nik:               params.Nik,
		Kpj:               params.Kpj,
		Fullname:          params.Fullname,
		ChannelId:         params.ChannelId,
		KodeKantorPPTOS:   params.KodeKantorPPTOS,
		PetugasRekamPTPOS: params.KodeKantorPPTOS,
		ChId:              params.ChId,
		ReqId:             params.ReqId,
	}

	res, err := service.ClaimRepository.CheckEligibleInsert(CheckEligibleJHTRequest)
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
