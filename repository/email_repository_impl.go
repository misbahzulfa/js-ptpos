package repository

import (
	"errors"
	"fmt"
	"js-ptpos/config"
	db "js-ptpos/config"
	"js-ptpos/entity"
	"js-ptpos/exception"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
	"net/http"
	"strings"
	"time"

	b64 "encoding/base64"
	"encoding/json"
	// "encoding/json"
)

func NewEmailRepository() EmailRepository {
	return &emailRepositoryImpl{}
}

type emailRepositoryImpl struct {
}

func (repository *emailRepositoryImpl) SendEmail(params *request.SendEmailRequest) (*response.SendEmailResponse, error) {
	client := &http.Client{}
	configuration := config.New()
	encodeBase64 := b64.StdEncoding.EncodeToString([]byte(params.Message))
	payload := strings.NewReader(`<x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/" xmlns:bpj="http://bpjs.com">
									<x:Header/>
									<x:Body>
										<bpj:sendEmail>
											<bpj:cfg>` + configuration.Get("COMMON_EMAIL_SERVICE_CFG") + `</bpj:cfg>
											<bpj:from>BPJS Ketenagakerjaan &lt;noreply@bpjsketenagakerjaan.go.id&gt;</bpj:from>
											<bpj:to>` + params.Email + `</bpj:to>
											<bpj:cc></bpj:cc>
											<bpj:bcc></bpj:bcc>
											<bpj:subject>` + params.Subject + `</bpj:subject>
											<bpj:body>` + params.Body + `</bpj:body>
											<bpj:isHTML>Y</bpj:isHTML>
											<bpj:bodyHTML>` + encodeBase64 + `</bpj:bodyHTML>
											<bpj:isAttach></bpj:isAttach>
											<bpj:attach></bpj:attach>
											<bpj:attachName></bpj:attachName>
											<bpj:avl></bpj:avl>
										</bpj:sendEmail>
									</x:Body>
								</x:Envelope>`)
	req, err := http.NewRequest("POST", configuration.Get("COMMON_EMAIL_SERVICE_URL")+"/WSCom/services/Main?wsdl", payload)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPAction", "")

	sendEmailResponse := &response.SendEmailResponseXml{}
	res, err := client.Do(req)
	if err != nil {
		exception.PanicIfNeeded(err)
	}
	defer res.Body.Close()

	if sendEmailResponse.Body.SendEmailResponse.Return.Ret == "0" {
		result := &response.SendEmailResponse{
			StatusCode: "200",
			Message:    "Sukses Kirim Email",
		}
		return result, nil
	} else {
		result := &response.SendEmailResponse{
			StatusCode: "200",
			Message:    sendEmailResponse.Body.SendEmailResponse.Return.Msg,
		}
		return result, nil
	}
	//return sendEmailResponse, nil
}

func (repository *emailRepositoryImpl) GetDataAfterClaimPTPOS(params *request.GetDataAfterClaimRequest) (*entity.GetDataAfterClaimPTPOSEntity, error) {
	var data entity.GetDataAfterClaimPTPOSEntity
	var qryEmail string
	var qry string

	if params.JenisPengajuan == "JHT" {
		qryEmail = `select * from kit.vas_email_content where vas_email_content_id='56'`
		qry = `select
					nama_tk,
					kode_pengajuan,
					kode_klaim,
					kpj, 
					kode_bank, 
					nama_bank,
					no_rekening,
					nama_rekening,
					email,
					to_char(tgl_bayar,'DD-MM-YYYY HH24:MM:SS') tgl_bayar,
					nom_manfaat_netto
				from bpjstku.asik_klaim 
				where kode_pengajuan = ? or kode_klaim = ?`
		//

	} else if params.JenisPengajuan == "JP" {
		qryEmail = `select * from kit.vas_email_content where vas_email_content_id='57'`
		qry = `select 
					nama_pelapor nama_tk,
					kode_penerima_berkala,
					kode_pengajuan,
					kode_klaim,
					(select kpj from pn.pn_klaim@to_kn where kode_klaim = a.kode_klaim) kpj, 
					(select kode_bank_penerima 
						from pn.pn_klaim_penerima_berkala@to_kn 
						where kode_penerima_berkala = a.kode_penerima_berkala 
						and status_layak = 'Y' and kode_klaim = a.kode_klaim) kode_bank,
					(select bank_penerima 
						from pn.pn_klaim_penerima_berkala@to_kn 
						where kode_penerima_berkala = a.kode_penerima_berkala 
						and status_layak = 'Y' and kode_klaim = a.kode_klaim) nama_bank,
					(select no_rekening_penerima 
						from pn.pn_klaim_penerima_berkala@to_kn 
						where kode_penerima_berkala = a.kode_penerima_berkala 
						and status_layak = 'Y' and kode_klaim = a.kode_klaim) no_rekening,
					(select nama_rekening_penerima 
						from pn.pn_klaim_penerima_berkala@to_kn 
						where kode_penerima_berkala = a.kode_penerima_berkala 
						and status_layak = 'Y' and kode_klaim = a.kode_klaim) nama_rekening,
					email_pelapor email,
					(select sum(nom_manfaat_netto) from pn.pn_klaim_berkala_detil@to_kn
						where kode_klaim = a.kode_klaim and no_konfirmasi = a.no_konfirmasi and no_proses = to_number(` + params.NoProses + `) ) nom_manfaat_netto,
					(select to_char(blth_awal,'DD-YYYY') from pn.pn_klaim_berkala@to_kn
							where kode_klaim = a.kode_klaim and no_konfirmasi = a.no_konfirmasi) blth_pengajuan
				from bpjstku.asik_konfirmasi a 
				where kode_tipe_manfaat = 'F001' and ( kode_pengajuan = ? or kode_klaim = ?)`
		//

	} else if params.JenisPengajuan == "BEASISWA" {
		qryEmail = `select * from kit.vas_email_content where vas_email_content_id='58'`
		//query jp belum fix
		qry = `select 
					nama_pelapor nama_tk,
					kode_penerima_berkala,
					kode_pengajuan,
					kode_klaim,
					(select kpj from pn.pn_klaim@to_kn where kode_klaim = a.kode_klaim) kpj, 
					(select kode_bank_penerima 
						from pn.pn_klaim_penerima_berkala@to_kn 
						where kode_penerima_berkala = a.kode_penerima_berkala 
						and status_layak = 'Y' and kode_klaim = a.kode_klaim) kode_bank,
					(select bank_penerima 
						from pn.pn_klaim_penerima_berkala@to_kn 
						where kode_penerima_berkala = a.kode_penerima_berkala 
						and status_layak = 'Y' and kode_klaim = a.kode_klaim) nama_bank,
					(select no_rekening_penerima 
						from pn.pn_klaim_penerima_berkala@to_kn 
						where kode_penerima_berkala = a.kode_penerima_berkala 
						and status_layak = 'Y' and kode_klaim = a.kode_klaim) no_rekening,
					(select nama_rekening_penerima 
						from pn.pn_klaim_penerima_berkala@to_kn 
						where kode_penerima_berkala = a.kode_penerima_berkala 
						and status_layak = 'Y' and kode_klaim = a.kode_klaim) nama_rekening,
					email_pelapor email,
					(select sum(nom_manfaat_netto) from pn.pn_klaim_berkala_detil@to_kn
						where kode_klaim = a.kode_klaim and no_konfirmasi = a.no_konfirmasi ) nom_manfaat_netto,
					(select to_char(blth_awal,'DD-YYYY') from pn.pn_klaim_berkala@to_kn
							where kode_klaim = a.kode_klaim and no_konfirmasi = a.no_konfirmasi) blth_pengajuan
				from bpjstku.asik_konfirmasi a 
				where kode_tipe_manfaat = 'F001' and ( kode_pengajuan = ? or kode_klaim = ?)`
	}

	resultEmail, err := db.EngineOltp.Query(qryEmail)
	if err != nil {
		return nil, err
	}

	result, err := db.EngineEcha.Query(
		qry,
		params.ClaimCode,
		params.ClaimCode,
	)
	if err != nil {
		return nil, err
	}

	data = entity.GetDataAfterClaimPTPOSEntity{
		KodePengajuan:   string(result[0]["KODE_PENGAJUAN"]),
		TipePengajuan:   params.JenisPengajuan,
		KodeKlaim:       string(result[0]["KODE_KLAIM"]),
		KPJ:             string(result[0]["KPJ"]),
		NamaTK:          string(result[0]["NAMA_TK"]),
		WaktuPembayaran: string(result[0]["TGL_BAYAR"]),
		BLTHPengajuan:   string(result[0]["BLTH_PENGAJUAN"]),
		DataPenerima: []entity.GetDataPenerima{{
			NamaPenerima:     string(result[0]["NAMA_TK"]),
			NamaBank:         string(result[0]["NAMA_BANK"]),
			NoRekening:       string(result[0]["NO_REKENING"]),
			NamaRekening:     string(result[0]["NAMA_REKENING"]),
			JumlahPembayaran: string(result[0]["NOM_MANFAAT_NETTO"]),
		},
			{
				NamaPenerima:     string(result[0]["NAMA_TK2"]),
				NamaBank:         string(result[0]["NAMA_BANK"]),
				NoRekening:       string(result[0]["NO_REKENING"]),
				NamaRekening:     string(result[0]["NAMA_REKENING"]),
				JumlahPembayaran: string(result[0]["NOM_MANFAAT_NETTO"]),
			},
		},
		DataEmail: []entity.GetEmailData{{
			EmailSubject: string(resultEmail[0]["SUBJECT"]),
			EmailContent: string(resultEmail[0]["HTML_CONTENT"]),
			Email:        string(result[0]["EMAIL"]),
		}},
	}

	return &data, nil
}

func (repository *emailRepositoryImpl) GetDataAfterClaimJPBeasiswaPTPOS(params *request.GetDataAfterClaimRequest) (*entity.GetDataAfterClaimEntity, error) {
	var data entity.GetDataAfterClaimEntity

	var qry = `SELECT *
					FROM BPJSTKU.ASIK_KONFIRMASI
				WHERE (   KODE_PENGAJUAN = ?
						OR KODE_KLAIM = ?)`
	result, err := db.EngineEcha.Query(
		qry,
		params.ClaimCode,
		params.ClaimCode,
	)
	if err != nil {
		return nil, err
	}

	data = entity.GetDataAfterClaimEntity{
		SubmissionCode:    string(result[0]["KODE_PENGAJUAN"]),
		ClaimCode:         string(result[0]["KODE_KLAIM"]),
		Email:             string(result[0]["EMAIL"]),
		PhoneNumber:       string(result[0]["NO_HP"]),
		Kpj:               string(result[0]["KPJ"]),
		IdentityNumber:    string(result[0]["NOMOR_IDENTITAS"]),
		Fullname:          string(result[0]["NAMA_TK"]),
		Birthdate:         string(result[0]["TGL_LAHIR"]),
		Birthplace:        string(result[0]["TEMPAT_LAHIR"]),
		Gender:            string(result[0]["JENIS_KELAMIN"]),
		Npwp:              string(result[0]["NPWP"]),
		BankCode:          string(result[0]["KODE_BANK"]),
		BankName:          string(result[0]["NAMA_BANK"]),
		AccountBankNumber: string(result[0]["NO_REKENING"]),
		AccountBankName:   string(result[0]["NAMA_REKENING"]),
		TotalTransfer:     string(result[0]["NOM_MANFAAT_NETTO"]),
		Address:           string(result[0]["ALAMAT"]),
		PaymentStatus:     string(result[0]["STATUS_BAYAR"]),
		PaymentDate:       string(result[0]["TGL_BAYAR"]),
		ActiveDate:        string(result[0]["TGL_KEPESERTAAN"]),
		NonActiveDate:     string(result[0]["TGL_AKTIF"]),
		MembershipDate:    string(result[0]["TGL_NONAKTIF"]),
	}

	return &data, nil
}

func (repository *emailRepositoryImpl) GetIDSurvey(params *request.GetDataAfterClaimRequest) (*request.IDSurveyRes, error) {
	configuration := config.New()
	client := &http.Client{}

	t := time.Now()
	requestId := t.Format("20060102150405")

	paramsData := &request.DataClaim{
		KodeKlaim: params.ClaimCode,
		User:      "PTPOS",
		SendEmail: "T",
		KodeKanal: "43",
	}

	paramSendGetIDSurvey := &request.GetIDSurveyReq{
		ChID:  "PTPOS",
		ReqID: requestId,
		Data:  *paramsData,
	}

	jsonData, err := json.Marshal(paramSendGetIDSurvey)

	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(jsonData))
	req, err := http.NewRequest("POST", configuration.Get("OPG_URL")+"/JSSurvey/GenerateRespondenKlaimByKanal", payload)
	if err != nil {
		return nil, err
	}

	IDSurveyResponse := &request.IDSurveyRes{}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	json.NewDecoder(res.Body).Decode(&IDSurveyResponse)
	defer res.Body.Close()

	fmt.Println(&IDSurveyResponse)

	idSurveyRespon := request.IDSurveyRes{
		Ret:     IDSurveyResponse.Ret,
		Encode:  IDSurveyResponse.Encode,
		Message: IDSurveyResponse.Message,
	}

	return &idSurveyRespon, nil
}

func (repository *emailRepositoryImpl) UpdateTanggalStatusBayar(kodePengajuan string, jenisPengajuan string, userSmile string, noKonfirmasi string) (*response.SendEmailResponse, error) {
	session := db.EngineEcha.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return nil, err
	}

	sqlUpdate := `UPDATE bpjstku.asik_konfirmasi
								SET tgl_bayar = sysdate, 
									status_bayar = 'Y'  
							WHERE kode_klaim = ?`
	if _, err = session.Exec(
		sqlUpdate,
		kodePengajuan,
	); err != nil {
		session.Rollback()
		return nil, err
	}

	if jenisPengajuan == "JP" {
		sqlUpdateDetilJP := `UPDATE bpjstku.asik_konfirmasi_berkala_detil
								SET tgl_lunas = sysdate, 
									status_lunas = 'Y',
									petugas_lunas = '` + userSmile + `' 
							WHERE kode_klaim = ? and no_proses = to_number('` + noKonfirmasi + `')`
		if _, err = session.Exec(
			sqlUpdateDetilJP,
			kodePengajuan,
		); err != nil {
			session.Rollback()
			return nil, err
		}
	} //  else if jenisPengajuan == "BEASISWA" {
	// 	sqlUpdateBea := `UPDATE bpjstku.asik_konfirmasi
	// 							SET tgl_lunas = sysdate,
	// 								status_status = 'Y',
	// 								petugas_lunas = '` + userSmile + `'
	// 						WHERE kode_klaim = ?`
	// 	if _, err = session.Exec(
	// 		sqlUpdateBea,
	// 		kodePengajuan,
	// 	); err != nil {
	// 		session.Rollback()
	// 		return nil, err
	// 	}
	// }

	err = session.Commit()
	if err != nil {
		fmt.Println(err)
	}

	statusUpdate := &response.SendEmailResponse{
		StatusCode: "200",
		Message:    "Tanggal dan Status Bayar Berhasil Di Update",
	}

	return statusUpdate, nil
}

func (repository *emailRepositoryImpl) CheckExsistKanal40_43(params *request.GetDataAfterClaimRequest) (*response.BeforeSendEmailResponse, error) {
	var qryCnt string
	if params.JenisPengajuan == "JHT" {
		qryCnt = `select nvl ((select kanal_pelayanan
						from bpjstku.asik_klaim x
					where kode_klaim = ?),
					'KANAL')
					kanal,
				(select count(*) from bpjstku.asik_klaim x 
				where kode_klaim = ?) cnt
				from dual`
	} else {
		qryCnt = `select nvl ((select kanal_pelayanan
				from bpjstku.asik_konfirmasi x
			where kode_klaim = ?),
			'KANAL')
			kanal,
		(select count(*) from bpjstku.asik_konfirmasi x 
		where kode_klaim = ?) cnt
		from dual`
	}

	result, err := db.EngineEcha.Query(
		qryCnt,
		params.ClaimCode,
		params.ClaimCode,
	)
	if err != nil {
		return nil, err
	}

	res := &response.BeforeSendEmailResponse{
		StatusCode:   string(result[0]["CNT"]),
		Message:      "Total Data",
		KanalLayanan: string(result[0]["KANAL"]),
	}

	return res, nil
}

func (repository *emailRepositoryImpl) JMOSendNotifPaymentJP(params *request.DataNotifJMORequest) (*response.InsertNotificationResponse, error) {
	configuration := config.New()
	client := &http.Client{}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(jsonData))
	req, err := http.NewRequest("POST", configuration.Get("JMO_SEND_JP_EMAIL")+"/claim-jp/send-email-notif-payment-jp", payload)

	if err != nil {
		return nil, err
	}

	insertNotificationResponse := &response.InsertNotificationResponse{}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	json.NewDecoder(res.Body).Decode(&insertNotificationResponse)
	defer res.Body.Close()

	if res.StatusCode == 200 {
		return insertNotificationResponse, nil
	} else {
		return nil, errors.New(insertNotificationResponse.Message)
	}
}
