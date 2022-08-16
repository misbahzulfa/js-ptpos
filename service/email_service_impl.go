package service

import (
	//"fmt"
	"fmt"
	"js-ptpos/exception"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
	"js-ptpos/repository"
	"js-ptpos/util"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/leekchan/accounting"
)

func NewEmailService(emailRepository *repository.EmailRepository) EmailService {
	return &emailServiceImpl{
		EmailRepository: *emailRepository,
	}
}

type emailServiceImpl struct {
	EmailRepository repository.EmailRepository
}

func (service *emailServiceImpl) SendEmail(params *request.SendEmailRequest) (*response.SendEmailResponse, error) {
	res, err := service.EmailRepository.SendEmail(params)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	if res.StatusCode == "200" {
		result := &response.SendEmailResponse{
			StatusCode: "200",
			Message:    "Sukses Kirim Email",
		}
		return result, nil
	} else {
		panic(exception.EmailError{
			Message: res.Message,
		})
	}
}

func (service *emailServiceImpl) SendEmailAfterClaim(ctx *fiber.Ctx, requestId string, params *request.GetDataAfterClaimRequest) (*response.SendEmailResponse, error) {
	resExistKanal4043, err := service.EmailRepository.CheckExsistKanal40_43(params)
	if err != nil {
		message := err.Error()
		logStop := util.LogResponse(ctx, message, requestId)
		fmt.Println(logStop)
		panic(exception.GeneralError{
			Message: message,
		})
	}

	//fmt.Println("resExistKanal4043.StatusCode = " + resExistKanal4043.StatusCode)

	//Cek apakah klaim dari kanal 40||43, (exist in bpjstku.asik_konfirmasi)
	if resExistKanal4043.StatusCode == "0" {
		result := response.SendEmailResponse{
			StatusCode: "204",
			Message:    "Data tidak ditemukan di ASIK Klaim dan Konfirmasi, Tidak Kirim Email",
		}
		return &result, nil
	} else {
		if resExistKanal4043.KanalLayanan == "40" {
			//call api jmo untuk send email and push notif
			res, err := service.EmailRepository.GetDataAfterClaimPTPOS(params)
			if err != nil {
				message := err.Error()
				logStop := util.LogResponse(ctx, message, requestId)
				fmt.Println(logStop)
				panic(exception.GeneralError{
					Message: message,
				})
			}

			paramReqNotif := &request.DataNotifJMORequest{
				ClaimCode:      res.KodeKlaim,
				Email:          res.DataEmail[0].Email,
				JenisPengajuan: "JP",
				NoProses:       params.NoProses,
				UserSmile:      params.UserSmile,
			}

			sendEmail, err := service.EmailRepository.JMOSendNotifPaymentJP(paramReqNotif)
			if err != nil {
				message := err.Error()
				logStop := util.LogResponse(ctx, message, requestId)
				fmt.Println(logStop)
				panic(exception.GeneralError{
					Message: message,
				})
			}

			// if sendEmail.Message == "Success" {
			// 	result := response.SendEmailResponse{
			// 		StatusCode: "200",
			// 		Message:    "Kirim Email dan Push Notif Bayar JMO",
			// 	}
			// 	return &result, nil
			// } else {
			// 	return nil, err
			// }

			result := response.SendEmailResponse{
				StatusCode: "200",
				Message:    "Kirim Email dan Push Notif Bayar JMO |.." + sendEmail.Message + "..| ",
			}
			return &result, nil

		} else {
			res, err := service.EmailRepository.GetDataAfterClaimPTPOS(params)
			if err != nil {
				message := err.Error()
				logStop := util.LogResponse(ctx, message, requestId)
				fmt.Println(logStop)
				panic(exception.GeneralError{
					Message: message,
				})
			}

			resIdSurvey, err := service.EmailRepository.GetIDSurvey(params)
			if err != nil {
				message := err.Error()
				logStop := util.LogResponse(ctx, message, requestId)
				fmt.Println(logStop)
				panic(exception.GeneralError{
					Message: message,
				})
			}

			formatRupiah := accounting.Accounting{Symbol: "Rp ", Precision: 2}
			totalTransfer, _ := strconv.ParseFloat(res.DataPenerima[0].JumlahPembayaran, 64)
			//fmt.Println("###### ID Survey ->" + resIdSurvey.Encode)
			//idSurvey := `alhPWTdmMmtYVktFZGlTMER2UkgvditydVI0QkdLd1hMb3gzdUJRaVcvVT0`
			linkSurvey := `https://esurvey-testing.bpjsketenagakerjaan.go.id/?id=` + resIdSurvey.Encode

			var strSubject string
			var strContent string

			if res.TipePengajuan == "JHT" {
				strSubject = strings.Replace(res.DataEmail[0].EmailSubject, ":0:", res.KPJ, 1)

				str0 := strings.Replace(res.DataEmail[0].EmailContent, ":0:", res.KPJ, 2)
				str1 := strings.Replace(str0, ":1:", res.NamaTK, 2)
				str2 := strings.Replace(str1, ":2:", res.KodePengajuan, 1)
				str3 := strings.Replace(str2, ":3:", res.DataPenerima[0].NamaPenerima, 1)
				str4 := strings.Replace(str3, ":4:", res.WaktuPembayaran, 1)
				str5 := strings.Replace(str4, ":5:", res.DataPenerima[0].NamaBank, 1)
				str6 := strings.Replace(str5, ":6:", res.DataPenerima[0].NoRekening, 1)
				str7 := strings.Replace(str6, ":7:", res.DataPenerima[0].NamaRekening, 1)
				str8 := strings.Replace(str7, ":8:", formatRupiah.FormatMoney(totalTransfer), 1)
				strContent = strings.Replace(str8, ":link:", linkSurvey, 1)

			} else if res.TipePengajuan == "JP" {
				strSubject = strings.Replace(res.DataEmail[0].EmailSubject, ":0:", res.KPJ, 1)

				str0 := strings.Replace(res.DataEmail[0].EmailContent, ":0:", res.KPJ, 2)
				str1 := strings.Replace(str0, ":1:", res.NamaTK, 2)
				str2 := strings.Replace(str1, ":2:", res.KodePengajuan, 1)
				str3 := strings.Replace(str2, ":3:", res.BLTHPengajuan, 1)
				str4 := strings.Replace(str3, ":4:", res.NamaTK, 1)
				str5 := strings.Replace(str4, ":5:", res.DataPenerima[0].NamaBank, 1)
				str6 := strings.Replace(str5, ":6:", res.DataPenerima[0].NoRekening, 1)
				str7 := strings.Replace(str6, ":7:", res.DataPenerima[0].NamaRekening, 1)
				str8 := strings.Replace(str7, ":8:", formatRupiah.FormatMoney(totalTransfer), 1)
				strContent = strings.Replace(str8, ":link:", linkSurvey, 1)

			} else if res.TipePengajuan == "BEASISWA" {
				strSubject = strings.Replace(res.DataEmail[0].EmailSubject, ":0:", res.KPJ, 1)

				str0 := strings.Replace(res.DataEmail[0].EmailContent, ":0:", res.KPJ, 2)
				str1 := strings.Replace(str0, ":1:", res.NamaTK, 2)
				str2 := strings.Replace(str1, ":2:", res.KodePengajuan, 1)
				str3 := strings.Replace(str2, ":3:", res.BLTHPengajuan, 1)
				str4 := strings.Replace(str3, ":4:", res.WaktuPembayaran, 1)
				str5 := strings.Replace(str4, ":5:", res.DataPenerima[0].NamaBank, 1)
				str6 := strings.Replace(str5, ":6:", res.DataPenerima[0].NoRekening, 1)
				str7 := strings.Replace(str6, ":7:", res.DataPenerima[0].NamaRekening, 1)
				strContent = strings.Replace(str7, ":8:", formatRupiah.FormatMoney(totalTransfer), 1)
			}

			sendEmailRequest := &request.SendEmailRequest{
				Subject: strSubject,
				Body:    strSubject,
				Message: strContent,
				Email:   res.DataEmail[0].Email,
			}

			sendEmail, err := service.EmailRepository.SendEmail(sendEmailRequest)
			if err != nil {
				message := err.Error()
				logStop := util.LogResponse(ctx, message, requestId)
				fmt.Println(logStop)
				panic(exception.GeneralError{
					Message: message,
				})
			}

			if sendEmail.StatusCode == "200" && (params.JenisPengajuan == "JP" || params.JenisPengajuan == "BEASISWA") {
				updateTanggaalStatusBayar, err := service.EmailRepository.UpdateTanggalStatusBayar(params.ClaimCode, params.JenisPengajuan, params.UserSmile, params.NoProses)
				if err != nil {
					message := err.Error()
					logStop := util.LogResponse(ctx, message, requestId)
					fmt.Println(logStop)
					panic(exception.GeneralError{
						Message: message,
					})
				}
				result := response.SendEmailResponse{
					StatusCode: "200",
					Message:    "Sukses Kirim Email dan " + updateTanggaalStatusBayar.Message,
				}
				return &result, nil
			}

			result := response.SendEmailResponse{
				StatusCode: "200",
				Message:    "Sukses Kirim Email",
			}
			return &result, nil
		}
	}
}

// func (service *emailServiceImpl) SendEmailAfterClaimJPBeasiswa(ctx *fiber.Ctx, requestId string, params *request.GetDataAfterClaimRequest) (*response.SendEmailResponse, error) {
// 	res, err := service.EmailRepository.GetDataAfterClaimPTPOS(params)
// 	if err != nil {
// 		message := err.Error()
// 		logStop := util.LogResponse(ctx, message, requestId)
// 		fmt.Println(logStop)

// 		panic(exception.GeneralError{
// 			Message: message,
// 		})
// 	}

// 	formatRupiah := accounting.Accounting{Symbol: "Rp ", Precision: 2}

// 	totalTransfer, _ := strconv.ParseFloat(res.TotalTransfer, 64)

// 	var body = `Yth. ` + res.Fullname + `,
// 		Klaim Jaminan Hari Tua (JHT) Anda telah berhasil
// 		Surel ini dikirimkan secara otomatis dan tidak untuk dibalas. Terima kasih.
// 		Salam hormat kami,
// 		BPJS Ketenagakerjaan`

// 	var template = `<!doctype html>
// 					<html>
// 					<head>
// 						<meta name="viewport" content="width=device-width" />
// 						<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
// 					</head>
// 					<body>
// 						<div bgcolor="#FFFFFF" style="margin:0;padding:0;font-family:&quot;Helvetica Neue&quot;,&quot;Helvetica&quot;,Helvetica,Arial,sans-serif;height:100%;font-size:14px;color:#404040;width:100%">
// 							<table class="m_6059976100216939223accent-wrap" style="margin:0;padding:0;font-family:&quot;Helvetica Neue&quot;,&quot;Helvetica&quot;,Helvetica,Arial,sans-serif;max-width:100%;background-color:transparent;border-collapse:collapse;border-spacing:0;width:100%">
// 								<tbody>
// 								<tr style="margin:0;padding:0">
// 									<td style="margin:0;padding:0"></td>
// 									<td class="m_6059976100216939223container" style="margin:0 auto;padding:0;display:block;max-width:600px;clear:both">
// 										<div class="m_6059976100216939223content" style="margin:0 auto;padding:0;max-width:600px;display:block;border-collapse:collapse;border:0">
// 											<table bgcolor="#fff" style="margin:0;padding:0;font-family:&quot;Helvetica Neue&quot;,&quot;Helvetica&quot;,Helvetica,Arial,sans-serif;max-width:100%;background-color:transparent;border-collapse:collapse;border-spacing:0;width:100%">
// 											<tbody>
// 												<tr style="margin:0;padding:0">
// 													<td height="4" bgcolor="#59BA52" style="background-color:#59BA52!important;line-height:4px;font-size:4px;margin:0;padding:0">&nbsp;</td>
// 													<td height="4" bgcolor="#CCDC3C" style="background-color:#CCDC3C!important;line-height:4px;font-size:4px;margin:0;padding:0">&nbsp;</td>
// 													<td height="4" bgcolor="#2693D6" style="background-color:#2693D6!important;line-height:4px;font-size:4px;margin:0;padding:0">&nbsp;</td>
// 												</tr>
// 											</tbody>
// 											</table>
// 										</div>
// 									</td>
// 									<td style="margin:0;padding:0"></td>
// 								</tr>
// 								</tbody>
// 							</table>
// 							<table class="m_6059976100216939223body-wrap" style="margin:0;padding:0;font-family:&quot;Helvetica Neue&quot;,&quot;Helvetica&quot;,Helvetica,Arial,sans-serif;max-width:100%;background-color:transparent;border-collapse:collapse;border-spacing:0;width:100%">
// 								<tbody>
// 								<tr style="margin:0;padding:0">
// 									<td style="margin:0;padding:0"></td>
// 									<td class="m_6059976100216939223container" bgcolor="#FFFFFF" style="margin:0 auto;padding:0;display:block;max-width:600px;clear:both">
// 										<div class="m_6059976100216939223content" style="margin:0 auto;padding:30px 15px;max-width:600px;display:block;border-collapse:collapse;border:1px solid #e7e7e7">
// 											<table style="margin:0;padding:0;font-family:&quot;Helvetica Neue&quot;,&quot;Helvetica&quot;,Helvetica,Arial,sans-serif;max-width:100%;background-color:transparent;border-collapse:collapse;border-spacing:0;width:100%">
// 											<tbody>
// 												<tr style="margin:0;padding:0">
// 													<td style="margin:0;padding:0">
// 														<div style="margin:0;padding:0">
// 														<table width="100%" border="0" cellspacing="0" cellpadding="0" style="background:#fff;margin:0;padding:0;font-family:&quot;Helvetica Neue&quot;,&quot;Helvetica&quot;,Helvetica,Arial,sans-serif;max-width:100%;background-color:transparent;border-collapse:collapse;border-spacing:0;width:100%">
// 															<tbody style="margin:0;padding:0">
// 																<tr style="margin:0;padding:0">
// 																	<td align="center" style="vertical-align:middle;margin:0;padding:0">
// 																	<h5 style="margin-bottom:20px;margin:0;padding:0;font-family:&quot;HelveticaNeue-Light&quot;,&quot;Helvetica Neue Light&quot;,&quot;Helvetica Neue&quot;,Helvetica,Arial,&quot;Lucida Grande&quot;,sans-serif;line-height:1.1;color:#000;font-weight:900;font-size:20px">Pembayaran KLAIM JHT</h5>
// 																	</td>
// 																</tr>
// 															</tbody>
// 														</table>
// 														</div>
// 														<hr style="margin:20px 0;padding:0;border:0;border-top:3px solid #d0d0d0;border-bottom:1px solid #ffffff">
// 														<div style="margin:10;padding:0;font-size:14px;">
// 														<p>Yth. </p>
// 														<p>Bapak/Ibu ` + res.Fullname + `</p>

// 														<p>Pembayaran klaim JHT Anda melalui aplikasi JMO telah berhasil ditransfer dari rekening BPJS Ketenagakerjaan ke nomor rekening yang telah diajukan dengan rincian sebagai berikut:
// 														</p>
// 														<br/>
// 														<table align="center" style="background-color:#eee;border-collapse:collapse;">
// 															<tr>
// 																<td align="right" style="background-color:#fff; font-weight:bold; font-size:10pt; padding:5px;border:1px solid #000;">Nomor Pengajuan</td>
// 																<td align="left" style="background-color:#fff; font-size:10pt; padding:5px;border:1px solid #000;">` + res.SubmissionCode + `</td>
// 															</tr>
// 															<tr>
// 																<td align="right" style="background-color:#fff; font-weight:bold; font-size:10pt; padding:5px;border:1px solid #000;">Nama Peserta</td>
// 																<td align="left" style="background-color:#fff; font-size:10pt; padding:5px;border:1px solid #000;">` + res.Fullname + `</td>
// 															</tr>
// 															<tr>
// 																<td align="right" style="background-color:#fff; font-weight:bold; font-size:10pt; padding:5px;border:1px solid #000;">Nomor Identitas Kependudukan</td>
// 																<td align="left" style="background-color:#fff; font-size:10pt; padding:5px;border:1px solid #000;">` + res.IdentityNumber + `</td>
// 															</tr>
// 															<tr>
// 																<td align="right" style="background-color:#fff; font-weight:bold; font-size:10pt; padding:5px;border:1px solid #000;">Nomor Kartu BPJS Ketenagakerjaan</td>
// 																<td align="left" style="background-color:#fff; font-size:10pt; padding:5px;border:1px solid #000;">` + res.Kpj + `</td>
// 															</tr>
// 															<tr>
// 																<td align="right" style="background-color:#fff; font-weight:bold; font-size:10pt; padding:5px;border:1px solid #000;">Nama Bank</td>
// 																<td align="left" style="background-color:#fff; font-size:10pt; padding:5px;border:1px solid #000;">` + res.BankName + `</td>
// 															</tr>
// 															<tr>
// 																<td align="right" style="background-color:#fff; font-weight:bold; font-size:10pt; padding:5px;border:1px solid #000;">Nomor Rekening</td>
// 																<td align="left" style="background-color:#fff; font-size:10pt; padding:5px;border:1px solid #000;">` + res.AccountBankNumber + `</td>
// 															</tr>
// 															<tr>
// 																<td align="right" style="background-color:#fff; font-weight:bold; font-size:10pt; padding:5px;border:1px solid #000;">Nama Rekening</td>
// 																<td align="left" style="background-color:#fff; font-size:10pt; padding:5px;border:1px solid #000;">` + res.AccountBankName + `</td>
// 															</tr>
// 															<tr>
// 																<td align="right" style="background-color:#fff; font-weight:bold; font-size:10pt; padding:5px;border:1px solid #000;">Tanggal Transfer</td>
// 																<td align="left" style="background-color:#fff; font-size:10pt; padding:5px;border:1px solid #000;">` + res.PaymentDate + ` WIB</td>
// 															</tr>
// 															<tr>
// 																<td align="right" style="background-color:#eee; font-weight:bold; font-size:10pt;  padding:5px;border:1px solid #000;">Jumlah Transfer</td>
// 																<td align="right" style="background-color:#eee; font-size:10pt; font-weight:bold;  padding:5px;border:1px solid #000;;">` + formatRupiah.FormatMoney(totalTransfer) + `</td>
// 															</tr>
// 														</table>
// 														<br/>
// 														<p>Salam hangat,
// 															<br>BPJS Ketenagakerjaan
// 														</p>
// 														</div>
// 														<p class="m_6059976100216939223footnote" style="margin:40px 0 0 0;padding:10px 0 0 0;margin-bottom:20px;font-weight:normal;font-size:14px;line-height:1.6;border-top:3px solid #d0d0d0">
// 														<small class="m_6059976100216939223muted" style="margin:0;padding:0;color:#999">
// 														Surel ini dikirimkan secara otomatis dan tidak untuk dibalas. Terima kasih.
// 														</small>
// 														</p>
// 													</td>
// 												</tr>
// 											</tbody>
// 											</table>
// 										</div>
// 									</td>
// 									<td style="margin:0;padding:0"></td>
// 								</tr>
// 								</tbody>
// 							</table>
// 							<table style="max-width:100%;border-collapse:collapse;border-spacing:0;width:100%;clear:both!important;background-color:transparent;margin:0 0 60px;padding:0" bgcolor="transparent">
// 								<tbody>
// 								<tr style="margin:0;padding:0">
// 									<td style="margin:0;padding:0"></td>
// 									<td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0">
// 										<div style="max-width:600px;display:block;border-collapse:collapse;margin:0 auto;padding:20px 15px;border-color:#e7e7e7;border-style:solid;border-width:0 1px 1px">
// 											<table width="100%" style="max-width:100%;border-collapse:collapse;border-spacing:0;width:100%;background-color:transparent;margin:0;padding:0" bgcolor="transparent">
// 											<tbody style="margin:0;padding:0">
// 												<tr style="margin:0;padding:0">
// 													<td valign="top" style="margin:0;padding:0;width:50%">
// 														<span style="font-size:12px;margin-bottom:6px;display:inline-block;color: #8c8c8c;">Download Aplikasi <span class="il">JMO (Jamsostek Mobile)</span></span>
// 														<div style="text-align:left">
// 														<a style="color:#008000" href="https://play.google.com/store/apps/details?id=com.ptpos&hl=en" target="_blank">
// 														<img src="https://www.bpjsketenagakerjaan.go.id/images/socialmedia/playstore.png" style="border:0;min-height:auto;max-width:120px;outline:0" width="200" height="65" alt="Download Android App" class="CToWUd"></a>
// 														</div>
// 													</td>
// 													<td valign="top" align="right" style="margin:0;padding:0">
// 														<span style="font-size:12px;margin-bottom:6px;display:inline-block;color: #8c8c8c;">Stay Connected</span>
// 														</br>
// 														</br>
// 														<div valign="center" align="right">
// 														<a href="https://www.facebook.com/BPJSTKInfo" style="color:#008000;display:inline-block" target="_blank"><img src="https://www.bpjsketenagakerjaan.go.id/images/socialmedia/facebook.png" alt="Facebook" width="28" height="28" class="CToWUd" style="border:0;min-height:auto;max-width:100%;outline:0"></a>
// 														<a href="https://www.twitter.com/bpjstkinfo" style="color:#008000;display:inline-block" target="_blank"><img src="https://www.bpjsketenagakerjaan.go.id/images/socialmedia/Twitter.png" alt="Twitter" width="30" class="CToWUd" style="border:0;min-height:auto;max-width:100%;outline:0"></a>
// 														<a href="https://www.youtube.com/bpjstkinfo" style="color:#008000;display:inline-block" target="_blank"><img src="https://www.bpjsketenagakerjaan.go.id/images/socialmedia/youtube2.png" alt="Youtube" height="29" width="40" class="CToWUd" style="border:0;min-height:auto;max-width:100%;outline:0"></a>
// 														<a href="https://www.instagram.com/bpjs.ketenagakerjaan/" style="color:#008000;display:inline-block" target="_blank"><img src="https://www.bpjsketenagakerjaan.go.id/images/socialmedia/instagram.png" alt="Youtube" height="30" width="30" class="CToWUd" style="border:0;min-height:auto;max-width:100%;outline:0"></a>
// 														</div>
// 													</td>
// 												</tr>
// 											</tbody>
// 											</table>
// 										</div>
// 									</td>
// 									<td style="margin:0;padding:0"></td>
// 								</tr>
// 								<tr style="margin:0;padding:0">
// 									<td style="margin:0;padding:0"></td>
// 									<td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0">
// 										<div style="max-width:600px;display:block;border-collapse:collapse;background-color:#f7f7f7;margin:0 auto;padding:20px 15px;border-color:#e7e7e7;border-style:solid;border-width:0 1px 1px">
// 											<table width="100%" style="max-width:100%;border-collapse:collapse;border-spacing:0;width:100%;background-color:transparent;margin:0;padding:0" bgcolor="transparent">
// 											<tbody style="margin:0;padding:0">
// 												<tr style="margin:0;padding:0">
// 													<td valign="middle" style="margin:0;padding:0;width:15%">
// 														<img src="https://bpjsketenagakerjaan.go.id/images/logo-bpjs.png" alt="BPJS Ketenagakerjaan" style="width:100px;">
// 													</td>
// 													<td valign="middle" style="margin:0;padding:0;width:53%">
// 														<p style="color:#91908e;font-size:10px;line-height:150%;font-weight:normal;margin:0px;padding:0px">
// 														Jika butuh bantuan, silahkan hubungi 175 Care Center Kami.
// 														<br style="margin:0;padding:0"> Copyright &copy; 2021 BPJS Ketenagakerjaan
// 														</p>
// 													</td>
// 												</tr>
// 											</tbody>
// 											</table>
// 										</div>
// 									</td>
// 									<td style="margin:0;padding:0"></td>
// 								</tr>
// 								</tbody>
// 							</table>
// 						</div>
// 					</body>
// 					</html>`

// 	sendEmailRequest := &request.SendEmailRequest{
// 		Subject: "Pembayaran Klaim JHT KPJ " + res.KPJ,
// 		Body:    body,
// 		Message: template,
// 		Email:   params.Email,
// 	}

// 	sendEmail, err := service.EmailRepository.SendEmail(sendEmailRequest)
// 	if err != nil {
// 		message := err.Error()
// 		logStop := util.LogResponse(ctx, message, requestId)
// 		fmt.Println(logStop)

// 		panic(exception.GeneralError{
// 			Message: message,
// 		})
// 	}

// 	if sendEmail.Message == "Success" {
// 		message := sendEmail.Message
// 		logStop := util.LogResponse(ctx, message, requestId)
// 		fmt.Println(logStop)
// 	}

// 	result := &response.SendEmailResponse{
// 		StatusCode: "200",
// 		Message:    "Sukses Kirim Email",
// 	}
// 	return result, nil
// }

// func (service *emailServiceImpl) SendEmailAfterClaimFailedTransfer(ctx *fiber.Ctx, requestId string, params *request.GetDataAfterClaimRequest) (*response.SendEmailResponse, error) {
// 	res, err := service.EmailRepository.GetDataAfterClaimPTPOS(params)
// 	if err != nil {
// 		message := err.Error()
// 		logStop := util.LogResponse(ctx, message, requestId)
// 		fmt.Println(logStop)

// 		panic(exception.GeneralError{
// 			Message: message,
// 		})
// 	}

// 	var body = `Yth. ` + res.Fullname + `,
// 		Klaim Jaminan Hari Tua (JHT) Anda telah berhasil
// 		Surel ini dikirimkan secara otomatis dan tidak untuk dibalas. Terima kasih.
// 		Salam hormat kami,
// 		BPJS Ketenagakerjaan`

// 	var template = `<!doctype html>
// 					<html>
// 					<head>
// 						<meta name="viewport" content="width=device-width" />
// 						<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
// 					</head>
// 					<body>
// 						<div bgcolor="#FFFFFF" style="margin:0;padding:0;font-family:&quot;Helvetica Neue&quot;,&quot;Helvetica&quot;,Helvetica,Arial,sans-serif;height:100%;font-size:14px;color:#404040;width:100%">
// 							<table class="m_6059976100216939223accent-wrap" style="margin:0;padding:0;font-family:&quot;Helvetica Neue&quot;,&quot;Helvetica&quot;,Helvetica,Arial,sans-serif;max-width:100%;background-color:transparent;border-collapse:collapse;border-spacing:0;width:100%">
// 								<tbody>
// 								<tr style="margin:0;padding:0">
// 									<td style="margin:0;padding:0"></td>
// 									<td class="m_6059976100216939223container" style="margin:0 auto;padding:0;display:block;max-width:600px;clear:both">
// 										<div class="m_6059976100216939223content" style="margin:0 auto;padding:0;max-width:600px;display:block;border-collapse:collapse;border:0">
// 											<table bgcolor="#fff" style="margin:0;padding:0;font-family:&quot;Helvetica Neue&quot;,&quot;Helvetica&quot;,Helvetica,Arial,sans-serif;max-width:100%;background-color:transparent;border-collapse:collapse;border-spacing:0;width:100%">
// 											<tbody>
// 												<tr style="margin:0;padding:0">
// 													<td height="4" bgcolor="#59BA52" style="background-color:#59BA52!important;line-height:4px;font-size:4px;margin:0;padding:0">&nbsp;</td>
// 													<td height="4" bgcolor="#CCDC3C" style="background-color:#CCDC3C!important;line-height:4px;font-size:4px;margin:0;padding:0">&nbsp;</td>
// 													<td height="4" bgcolor="#2693D6" style="background-color:#2693D6!important;line-height:4px;font-size:4px;margin:0;padding:0">&nbsp;</td>
// 												</tr>
// 											</tbody>
// 											</table>
// 										</div>
// 									</td>
// 									<td style="margin:0;padding:0"></td>
// 								</tr>
// 								</tbody>
// 							</table>
// 							<table class="m_6059976100216939223body-wrap" style="margin:0;padding:0;font-family:&quot;Helvetica Neue&quot;,&quot;Helvetica&quot;,Helvetica,Arial,sans-serif;max-width:100%;background-color:transparent;border-collapse:collapse;border-spacing:0;width:100%">
// 								<tbody>
// 								<tr style="margin:0;padding:0">
// 									<td style="margin:0;padding:0"></td>
// 									<td class="m_6059976100216939223container" bgcolor="#FFFFFF" style="margin:0 auto;padding:0;display:block;max-width:600px;clear:both">
// 										<div class="m_6059976100216939223content" style="margin:0 auto;padding:30px 15px;max-width:600px;display:block;border-collapse:collapse;border:1px solid #e7e7e7">
// 											<table style="margin:0;padding:0;font-family:&quot;Helvetica Neue&quot;,&quot;Helvetica&quot;,Helvetica,Arial,sans-serif;max-width:100%;background-color:transparent;border-collapse:collapse;border-spacing:0;width:100%">
// 											<tbody>
// 												<tr style="margin:0;padding:0">
// 													<td style="margin:0;padding:0">
// 														<div style="margin:0;padding:0">
// 														<table width="100%" border="0" cellspacing="0" cellpadding="0" style="background:#fff;margin:0;padding:0;font-family:&quot;Helvetica Neue&quot;,&quot;Helvetica&quot;,Helvetica,Arial,sans-serif;max-width:100%;background-color:transparent;border-collapse:collapse;border-spacing:0;width:100%">
// 															<tbody style="margin:0;padding:0">
// 																<tr style="margin:0;padding:0">
// 																	<td align="center" style="vertical-align:middle;margin:0;padding:0">
// 																	<h5 style="margin-bottom:20px;margin:0;padding:0;font-family:&quot;HelveticaNeue-Light&quot;,&quot;Helvetica Neue Light&quot;,&quot;Helvetica Neue&quot;,Helvetica,Arial,&quot;Lucida Grande&quot;,sans-serif;line-height:1.1;color:#000;font-weight:900;font-size:20px">Pembayaran KLAIM JHT</h5>
// 																	</td>
// 																</tr>
// 															</tbody>
// 														</table>
// 														</div>
// 														<hr style="margin:20px 0;padding:0;border:0;border-top:3px solid #d0d0d0;border-bottom:1px solid #ffffff">
// 														<div style="margin:10;padding:0;font-size:14px;">
// 														<p>Yth. </p>
// 														<p>Bapak/Ibu ` + res.Fullname + `</p>

// 														<p>Pengajuan klaim JHT KPJ ` + res.Kpj + ` tidak berhasil. Silahkan melakukan pengajuan Klaim JHT kembali dan pastikan rekening Anda masih aktif.</p>

// 														<br/>
// 														<p>Salam hangat,
// 															<br>BPJS Ketenagakerjaan
// 														</p>
// 														</div>
// 														<p class="m_6059976100216939223footnote" style="margin:40px 0 0 0;padding:10px 0 0 0;margin-bottom:20px;font-weight:normal;font-size:14px;line-height:1.6;border-top:3px solid #d0d0d0">
// 														<small class="m_6059976100216939223muted" style="margin:0;padding:0;color:#999">
// 														Surel ini dikirimkan secara otomatis dan tidak untuk dibalas. Terima kasih.
// 														</small>
// 														</p>
// 													</td>
// 												</tr>
// 											</tbody>
// 											</table>
// 										</div>
// 									</td>
// 									<td style="margin:0;padding:0"></td>
// 								</tr>
// 								</tbody>
// 							</table>
// 							<table style="max-width:100%;border-collapse:collapse;border-spacing:0;width:100%;clear:both!important;background-color:transparent;margin:0 0 60px;padding:0" bgcolor="transparent">
// 								<tbody>
// 								<tr style="margin:0;padding:0">
// 									<td style="margin:0;padding:0"></td>
// 									<td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0">
// 										<div style="max-width:600px;display:block;border-collapse:collapse;margin:0 auto;padding:20px 15px;border-color:#e7e7e7;border-style:solid;border-width:0 1px 1px">
// 											<table width="100%" style="max-width:100%;border-collapse:collapse;border-spacing:0;width:100%;background-color:transparent;margin:0;padding:0" bgcolor="transparent">
// 											<tbody style="margin:0;padding:0">
// 												<tr style="margin:0;padding:0">
// 													<td valign="top" style="margin:0;padding:0;width:50%">
// 														<span style="font-size:12px;margin-bottom:6px;display:inline-block;color: #8c8c8c;">Download Aplikasi <span class="il">JMO (Jamsostek Mobile)</span></span>
// 														<div style="text-align:left">
// 														<a style="color:#008000" href="https://play.google.com/store/apps/details?id=com.ptpos&hl=en" target="_blank">
// 														<img src="https://www.bpjsketenagakerjaan.go.id/images/socialmedia/playstore.png" style="border:0;min-height:auto;max-width:120px;outline:0" width="200" height="65" alt="Download Android App" class="CToWUd"></a>
// 														</div>
// 													</td>
// 													<td valign="top" align="right" style="margin:0;padding:0">
// 														<span style="font-size:12px;margin-bottom:6px;display:inline-block;color: #8c8c8c;">Stay Connected</span>
// 														</br>
// 														</br>
// 														<div valign="center" align="right">
// 														<a href="https://www.facebook.com/BPJSTKInfo" style="color:#008000;display:inline-block" target="_blank"><img src="https://www.bpjsketenagakerjaan.go.id/images/socialmedia/facebook.png" alt="Facebook" width="28" height="28" class="CToWUd" style="border:0;min-height:auto;max-width:100%;outline:0"></a>
// 														<a href="https://www.twitter.com/bpjstkinfo" style="color:#008000;display:inline-block" target="_blank"><img src="https://www.bpjsketenagakerjaan.go.id/images/socialmedia/Twitter.png" alt="Twitter" width="30" class="CToWUd" style="border:0;min-height:auto;max-width:100%;outline:0"></a>
// 														<a href="https://www.youtube.com/bpjstkinfo" style="color:#008000;display:inline-block" target="_blank"><img src="https://www.bpjsketenagakerjaan.go.id/images/socialmedia/youtube2.png" alt="Youtube" height="29" width="40" class="CToWUd" style="border:0;min-height:auto;max-width:100%;outline:0"></a>
// 														<a href="https://www.instagram.com/bpjs.ketenagakerjaan/" style="color:#008000;display:inline-block" target="_blank"><img src="https://www.bpjsketenagakerjaan.go.id/images/socialmedia/instagram.png" alt="Youtube" height="30" width="30" class="CToWUd" style="border:0;min-height:auto;max-width:100%;outline:0"></a>
// 														</div>
// 													</td>
// 												</tr>
// 											</tbody>
// 											</table>
// 										</div>
// 									</td>
// 									<td style="margin:0;padding:0"></td>
// 								</tr>
// 								<tr style="margin:0;padding:0">
// 									<td style="margin:0;padding:0"></td>
// 									<td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0">
// 										<div style="max-width:600px;display:block;border-collapse:collapse;background-color:#f7f7f7;margin:0 auto;padding:20px 15px;border-color:#e7e7e7;border-style:solid;border-width:0 1px 1px">
// 											<table width="100%" style="max-width:100%;border-collapse:collapse;border-spacing:0;width:100%;background-color:transparent;margin:0;padding:0" bgcolor="transparent">
// 											<tbody style="margin:0;padding:0">
// 												<tr style="margin:0;padding:0">
// 													<td valign="middle" style="margin:0;padding:0;width:15%">
// 														<img src="https://bpjsketenagakerjaan.go.id/images/logo-bpjs.png" alt="BPJS Ketenagakerjaan" style="width:100px;">
// 													</td>
// 													<td valign="middle" style="margin:0;padding:0;width:53%">
// 														<p style="color:#91908e;font-size:10px;line-height:150%;font-weight:normal;margin:0px;padding:0px">
// 														Jika butuh bantuan, silahkan hubungi 175 Care Center Kami.
// 														<br style="margin:0;padding:0"> Copyright &copy; 2021 BPJS Ketenagakerjaan
// 														</p>
// 													</td>
// 												</tr>
// 											</tbody>
// 											</table>
// 										</div>
// 									</td>
// 									<td style="margin:0;padding:0"></td>
// 								</tr>
// 								</tbody>
// 							</table>
// 						</div>
// 					</body>
// 					</html>`

// 	sendEmailRequest := &request.SendEmailRequest{
// 		Subject: "Pengajuan Klaim JHT KPJ " + res.KPJ + "Tidak berhasil",
// 		Body:    body,
// 		Message: template,
// 		Email:   params.Email,
// 	}

// 	sendEmail, err := service.EmailRepository.SendEmail(sendEmailRequest)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if sendEmail.Message == "Success" {
// 		message := sendEmail.Message
// 		logStop := util.LogResponse(ctx, message, requestId)
// 		fmt.Println(logStop)
// 	}

// 	result := &response.SendEmailResponse{
// 		StatusCode: "200",
// 		Message:    "Sukses Kirim Email",
// 	}
// 	return result, nil
// }
