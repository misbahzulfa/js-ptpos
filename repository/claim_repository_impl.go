package repository

import (
	"database/sql"
	"encoding/json"
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
)

func NewClaimRepository(configuration *config.Config) ClaimRepository {
	return &claimRepositoryImpl{
		Configuration: *configuration,
	}
}

type claimRepositoryImpl struct {
	Configuration config.Config
}

func (repository *claimRepositoryImpl) CheckAccountBank(params *request.CheckAccountBankRemoteRequest) (*response.CheckAccountBankRemoteResponse, error) {
	client := &http.Client{}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, exception.JSONParseExceptionMessage
	}

	payload := strings.NewReader(string(jsonData))
	fmt.Println(payload)
	req, err := http.NewRequest("POST", repository.Configuration.Get("OPG_URL")+"/JSOPG/GetAccountInfo", payload)

	if err != nil {
		return nil, exception.CallApiExceptionMessage
	}

	checkAccountBankRemoteResponse := &response.CheckAccountBankRemoteResponse{}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, exception.ClientDoExceptionMessage
	}

	json.NewDecoder(res.Body).Decode(&checkAccountBankRemoteResponse)
	defer res.Body.Close()
	fmt.Println(checkAccountBankRemoteResponse, "response ws")
	return checkAccountBankRemoteResponse, nil
}
func (repository *claimRepositoryImpl) SimilairtyAccountBankName(params *request.SimilarityRequest) (*response.CommonSimilarityResponse, error) {
	client := &http.Client{}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, exception.JSONParseExceptionMessage
	}

	payload := strings.NewReader(string(jsonData))
	req, err := http.NewRequest("POST", repository.Configuration.Get("OPG_URL")+"/JSSISLA/SimilarityLapas", payload)

	if err != nil {
		return nil, exception.CallApiExceptionMessage
	}

	similarityResponse := &response.CommonSimilarityResponse{}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, exception.ClientDoExceptionMessage
	}
	json.NewDecoder(res.Body).Decode(&similarityResponse)
	defer res.Body.Close()

	return similarityResponse, nil
}

func (repository *claimRepositoryImpl) CheckEligible(params *request.CheckEligibleRequest) (*entity.CheckEligibleEntity, error) {
	var p_param_maksimal_saldo_jht string
	var p_klaim_sebagian string
	var p_kode_pesan string
	var p_pesan_tidak_layak string
	var p_nama_kanal_kantor_cabang string
	var p_sukses string
	var p_mess string

	query := `begin
				ptpos.p_ptpos_asik_klaim.x_cek_kelayakan_jht(?,
																 ?,
																 ?,
																 ?,
																 ?,
																 TO_DATE(? ,'dd-mm-yyyy'),
																 ?,
																 :p_param_maksimal_saldo_jht,
																 :p_klaim_sebagian,
																 :p_kode_pesan,
																 :p_nama_kanal_kantor_cabang,
																 :p_pesan_tidak_layak,
																 :p_sukses,
																 :p_mess);
			 end;`
	if _, err := db.EngineEcha.Exec(query,
		params.SegmenCode,
		params.Kpj,
		params.WorkerCode,
		params.IdentityNumber,
		params.FullName,
		params.BirthDate,
		params.Email,
		sql.Named("p_param_maksimal_saldo_jht", sql.Out{Dest: &p_param_maksimal_saldo_jht}),
		sql.Named("p_klaim_sebagian", sql.Out{Dest: &p_klaim_sebagian}),
		sql.Named("p_kode_pesan", sql.Out{Dest: &p_kode_pesan}),
		sql.Named("p_nama_kanal_kantor_cabang", sql.Out{Dest: &p_nama_kanal_kantor_cabang}),
		sql.Named("p_pesan_tidak_layak", sql.Out{Dest: &p_pesan_tidak_layak}),
		sql.Named("p_sukses", sql.Out{Dest: &p_sukses}),
		sql.Named("p_mess", sql.Out{Dest: &p_mess})); err != nil {
		return nil, err
	}

	result := &entity.CheckEligibleEntity{
		BalanceMaksimum:    p_param_maksimal_saldo_jht,
		PartialClaim:       p_klaim_sebagian,
		MessageCode:        p_kode_pesan,
		MessageNotEligible: p_pesan_tidak_layak,
		OfficeName:         p_nama_kanal_kantor_cabang,
		Success:            p_sukses,
		Message:            p_mess,
	}

	return result, nil
}

func (repository *claimRepositoryImpl) CauseOfClaim(params *request.CauseOfClaimRequest) ([]entity.CauseOfClaimEntity, error) {
	resultQry, err := db.EngineEcha.Query(
		`select * from table (ptpos.p_ptpos_asik_klaim.f_get_list_sebab_klaim(?, ?, ?,  TO_DATE(?, 'dd-mm-yyyy')))`,
		params.SegmenCode,
		params.WorkerCode,
		params.Program,
		params.BirthDate,
	)

	if err != nil {
		return nil, err
	}

	objects := make([]entity.CauseOfClaimEntity, 0)
	for _, item := range resultQry {
		object := entity.CauseOfClaimEntity{
			SegmenCode:       string(item["KODE_SEGMEN"]),
			ClaimTypeCode:    string(item["KODE_TIPE_KLAIM"]),
			CauseOfClaimCode: string(item["KODE_SEBAB_KLAIM"]),
			CauseOfClaimName: string(item["NAMA_SEBAB_KLAIM"]),
			ReceiverTypeCode: string(item["KODE_TIPE_PENERIMA"]),
			Number:           string(item["NO_URUT"]),
		}
		objects = append(objects, object)
	}

	return objects, nil
}

func (repository *claimRepositoryImpl) GetDataEmployee(params *request.GetDataEmployeeRequest) (*entity.DataEmployeeEntity, error) {
	var employee entity.DataEmployeeEntity
	var kodeTK string
	var tglLahir string
	// fmt.Println("repo")
	var qry = `
	select kode_tk,kode_segmen,tgl_lahir,email
    from (
        select a.kode_tk,a.kode_segmen,to_char(a.tgl_lahir,'dd-mm-yyyy')tgl_lahir,a.email
        from kn.vw_kn_tk a where kpj =?
        )xx 
	group by xx.kode_tk,kode_segmen,tgl_lahir,email
	`
	// fmt.Println(qry, "1")
	resultKdTK, err := db.EngineOltp.Query(qry, params.Kpj)
	if err != nil {
		return nil, err
	}

	if len(resultKdTK) > 0 {
		kodeTK = string(resultKdTK[0]["KODE_TK"])
		tglLahir = string(resultKdTK[0]["TGL_LAHIR"])
	}
	var sql = `select 	kode_asik,
							nik,
							jenis_identitas,
							status_valid_identitas,
							nama_lengkap,
							jenis_kelamin,
							tempat_lahir,
							to_char(tgl_lahir,'dd-mm-yyyy') tgl_lahir,
							nama_ibu_kandung,
							status_kawin,
							alamat,
							kode_kelurahan,
							kode_kecamatan,
							kode_kabupaten,
							kode_propinsi,
							kode_pos,
							no_hp,
							email,
							no_npwp,
							no_passport,
							to_char(tgl_berlaku_passport, 'dd-mm-yyyy') tgl_berlaku_passport,
							kode_bank,
							nama_bank,
							no_rekening,
							nama_rekening,
							kode_pendidikan_terakhir,
							kode_agama,
							golongan_darah,
							nama_kontak_darurat,
							no_hp_kontak_darurat,
							alamat_kontak_darurat,
							kode_hubungan_kontak_darurat,
							kode_kelurahan_kontak_darurat,
							kode_kecamatan_kontak_darurat,
							kode_kabupaten_kontak_darurat,
							kode_propinsi_kontak_darurat,
							kode_pos_kontak_darurat,
							kode_tk,
							kpj,
							kode_segmen,
							kode_divisi,
							kode_perusahaan,
							kode_kepesertaan,
							kode_kantor,
							npp,
							to_char(tgl_kepesertaan, 'dd-mm-yyyy') tgl_kepesertaan,
							to_char(tgl_aktif, 'dd-mm-yyyy') tgl_aktif,
							to_char(tgl_na, 'dd-mm-yyyy') tgl_na,
							kode_na,
							sumber_blth_nonaktif,
							aktif_tk,
							to_char(saldo_jht) saldo_jht
			from table(bpjstku.p_ptpos_asik_klaim.f_get_pengkinian_data(?, ?, ?, ?, ?, to_date(?, 'dd-mm-yyyy')))`
	result, err := db.EngineEcha.Query(
		sql,
		params.SegmenCode,
		params.Kpj,
		kodeTK,
		params.IdentityNumber,
		params.FullName,
		tglLahir,
	)
	// fmt.Println(params.SegmenCode, params.Kpj, kodeTK, params.IdentityNumber, params.FullName, params.BirthDate)
	if err != nil {
		return nil, err
	}
	if len(result) > 0 {
		employee = entity.DataEmployeeEntity{
			AsikCode:                         string(result[0]["KODE_ASIK"]),
			IdentityNumber:                   string(result[0]["NIK"]),
			IdentityType:                     string(result[0]["JENIS_IDENTITAS"]),
			ValidIdentity:                    string(result[0]["STATUS_VALID_IDENTITAS"]),
			FullName:                         string(result[0]["NAMA_LENGKAP"]),
			Gender:                           string(result[0]["JENIS_KELAMIN"]),
			BirthPlace:                       string(result[0]["TEMPAT_LAHIR"]),
			BirthDate:                        string(result[0]["TGL_LAHIR"]),
			MotherName:                       string(result[0]["NAMA_IBU_KANDUNG"]),
			MaritalStatus:                    string(result[0]["STATUS_KAWIN"]),
			Address:                          string(result[0]["ALAMAT"]),
			KelurahanCode:                    string(result[0]["KODE_KELURAHAN"]),
			KecamatanCode:                    string(result[0]["KODE_KECAMATAN"]),
			KabupatenCode:                    string(result[0]["KODE_KABUPATEN"]),
			PropinsiCode:                     string(result[0]["KODE_PROPINSI"]),
			PostalCode:                       string(result[0]["KODE_POS"]),
			PhoneNumber:                      string(result[0]["NO_HP"]),
			Email:                            string(result[0]["EMAIL"]),
			Npwp:                             string(result[0]["NO_NPWP"]),
			PassportNumber:                   string(result[0]["NO_PASSPORT"]),
			PassportExpired:                  string(result[0]["TGL_BERLAKU_PASSPORT"]),
			BankCode:                         string(result[0]["KODE_BANK"]),
			BankName:                         string(result[0]["NAMA_BANK"]),
			AccountBankNumber:                string(result[0]["NOMOR_REKENING"]),
			AccountBankName:                  string(result[0]["NAMA_REKENING"]),
			LastEducationCode:                string(result[0]["KODE_PENDIDIKAN_TERAKHIR"]),
			ReligionCode:                     string(result[0]["KODE_AGAMA"]),
			BloodGroup:                       string(result[0]["GOLONGAN_DARAH"]),
			EmergencyContactName:             string(result[0]["NAMA_KONTAK_DARURAT"]),
			EmergencyContactPhoneNumber:      string(result[0]["NO_HP_KONTAK_DARURAT"]),
			EmergencyContactAddress:          string(result[0]["ALAMAT_KONTAK_DARURAT"]),
			EmergencyContactRelationshipCode: string(result[0]["KODE_HUBUNGAN_KONTAK_DARURAT"]),
			EmergencyContactKelurahanCode:    string(result[0]["KODE_KELURAHAN_KONTAK_DARURAT"]),
			EmergencyContactKecamatanCode:    string(result[0]["KODE_KECAMATAN_KONTAK_DARURAT"]),
			EmergencyContactKabupatenCode:    string(result[0]["KODE_KABUPATEN_KONTAK_DARURAT"]),
			EmergencyContactPropinsiCode:     string(result[0]["KODE_PROPINSI_KONTAK_DARURAT"]),
			EmergencyContactPostalCode:       string(result[0]["KODE_POS_KONTAK_DARURAT"]),
			WorkerCode:                       string(result[0]["KODE_TK"]),
			Kpj:                              string(result[0]["KPJ"]),
			SegmenCode:                       string(result[0]["KODE_SEGMEN"]),
			DivisionCode:                     string(result[0]["KODE_DIVISI"]),
			CompanyCode:                      string(result[0]["KODE_PERUSAHAAN"]),
			MembershipCode:                   string(result[0]["KODE_KEPESERTAAN"]),
			OfficeCode:                       string(result[0]["KODE_KANTOR"]),
			Npp:                              string(result[0]["NPP"]),
			MembershipDate:                   string(result[0]["TGL_KEPESERTAAN"]),
			ActiveDate:                       string(result[0]["TGL_AKTIF"]),
			NonActiveDate:                    string(result[0]["TGL_NA"]),
			NonActiveCode:                    string(result[0]["KODE_NA"]),
			FlagSipp:                         string(result[0]["SUMBER_BLTH_NONAKTIF"]),
			Active:                           string(result[0]["AKTIF_TK"]),
			JhtBalance:                       string(result[0]["SALDO_JHT"]),
		}
	} else {
		employee = entity.DataEmployeeEntity{}
	}

	// fmt.Println(employee, "xxxxx")
	return &employee, nil
}

func (repository *claimRepositoryImpl) GetClaimBenefitDetail(params *request.ClaimBenefitDetailRequest) (*entity.ClaimBenefitDetailEntity, error) {
	var benefitDetail entity.ClaimBenefitDetailEntity
	var kodeTK string
	// fmt.Println("repo")
	var qry = `
	select kode_tk,kode_segmen,tgl_lahir,email
    from (
        select a.kode_tk,a.kode_segmen,to_char(a.tgl_lahir,'dd-mm-yyyy')tgl_lahir,a.email
        from kn.vw_kn_tk a where kpj =?
        )xx 
	group by xx.kode_tk,kode_segmen,tgl_lahir,email
	`
	// fmt.Println(qry, "1")
	resultKdTK, err := db.EngineOltp.Query(qry, params.Kpj)
	if err != nil {
		return nil, err
	}

	if len(resultKdTK) > 0 {
		kodeTK = string(resultKdTK[0]["KODE_TK"])
	}
	var sql = `
			select 	kode_manfaat,
							kode_tipe_penerima,
							nvl(kode_pelaporan, '-') kode_pelaporan,
							kd_prg,
							to_char(tgl_pengembangan, 'dd-mm-yyyy') tgl_pengembangan,
							rate_pengembangan,
							to_char(tgl_saldo_awal_tahun, 'dd-mm-yyyy') tgl_saldo_awal_tahun,
							to_char(nvl(nom_saldo_awal_tahun, 0)) nom_saldo_awal_tahun,
							to_char(nvl(nom_saldo_pengembangan, 0)) nom_saldo_pengembangan,
							to_char(nvl(nom_saldo_total, 0)) nom_saldo_total,
							to_char(nvl(nom_iuran_tahun_berjalan, 0)) nom_iuran_tahun_berjalan,
							to_char(nvl(nom_iuran_pengembangan, 0)) nom_iuran_pengembangan,
							to_char(nvl(nom_iuran_total, 0)) nom_iuran_total,
							to_char(nvl(nom_saldo_iuran_total, 0)) nom_saldo_iuran_total,
							to_char(nvl(persentase_pengambilan, 0)) persentase_pengambilan,
							to_char(nvl(nom_manfaat_maxbisadiambil, 0)) nom_manfaat_maxbisadiambil,
							to_char(nvl(nom_manfaat_diambil, 0)) nom_manfaat_diambil,
							to_char(nvl(nom_manfaat_gross, 0)) nom_manfaat_gross,
							to_char(nvl(nom_pph, 0)) nom_pph,
							to_char(nvl(nom_pembulatan, 0)) nom_pembulatan,
							to_char(nvl(nom_manfaat_netto, 0)) nom_manfaat_netto
			from 		table(bpjstku.p_ptpos_asik_klaim.f_get_klaim_manfaat_detil(?, ?, ?, ?, ?, to_date(?,'dd-mm-yyyy'), ?))
	`
	result, err := db.EngineEcha.Query(
		sql,
		params.SegmenCode,
		params.Kpj,
		kodeTK,
		params.IdentityNumber,
		params.FullName,
		params.BirthDate,
		params.CauseOfClaimCode,
	)

	if err != nil {
		return nil, err
	}
	benefitDetail = entity.ClaimBenefitDetailEntity{
		BenefitCode:                  string(result[0]["KODE_MANFAAT"]),
		ReceiverTypeCode:             string(result[0]["KODE_TIPE_PENERIMA"]),
		ProgramCode:                  string(result[0]["KD_PRG"]),
		ReportCode:                   string(result[0]["KODE_PELAPORAN"]),
		TanggalPengembangan:          string(result[0]["TGL_PENGEMBANGAN"]),
		RatePengembangan:             string(result[0]["RATE_PENGEMBANGAN"]),
		TanggalSaldoAwalTahun:        string(result[0]["TGL_SALDO_AWAL_TAHUN"]),
		NominalSaldoAwalTahun:        string(result[0]["NOM_SALDO_AWAL_TAHUN"]),
		NominalSaldoPengembangan:     string(result[0]["NOM_SALDO_PENGEMBANGAN"]),
		NominalSaldoTotal:            string(result[0]["NOM_SALDO_TOTAL"]),
		NominalIuranTahunBerjalan:    string(result[0]["NOM_IURAN_TAHUN_BERJALAN"]),
		NominalIuranPengembangan:     string(result[0]["NOM_IURAN_PENGEMBANGAN"]),
		NominalIuranTotal:            string(result[0]["NOM_IURAN_TOTAL"]),
		NominalSaldoIuranTotal:       string(result[0]["NOM_SALDO_IURAN_TOTAL"]),
		PersentasePengambilan:        string(result[0]["PRESENTASE_PENGAMBILAN"]),
		NominalManfaatMaxBisaDiAmbil: string(result[0]["NOM_MANFAAT_MAXBISADIAMBIL"]),
		NominalManfaatDiAmbil:        string(result[0]["NOM_MANFAAT_DIAMBIL"]),
		NominalManfaatGross:          string(result[0]["NOM_MANFAAT_GROSS"]),
		NominalPPH:                   string(result[0]["NOM_PPH"]),
		NominalPembulatan:            string(result[0]["NOM_PEMBULATAN"]),
		NominalManfaatNetto:          string(result[0]["NOM_MANFAAT_NETTO"]),
	}

	return &benefitDetail, nil
}

func (repository *claimRepositoryImpl) InsertClaimJht(params *request.InsertClaimRequest) (string, error) {
	var kodeTK string
	var tglLahir string
	// fmt.Println("repo")
	var qry = `
	select kode_tk,kode_segmen,tgl_lahir,email
    from (
        select a.kode_tk,a.kode_segmen,to_char(a.tgl_lahir,'dd-mm-yyyy')tgl_lahir,a.email
        from kn.vw_kn_tk a where kpj =?
        )xx 
	group by xx.kode_tk,kode_segmen,tgl_lahir,email
	`
	// fmt.Println(qry, "1")
	resultKdTK, err := db.EngineOltp.Query(qry, params.Kpj)
	if err != nil {
		return "", err
	}

	if len(resultKdTK) > 0 {
		kodeTK = string(resultKdTK[0]["KODE_TK"])
		tglLahir = string(resultKdTK[0]["TGL_LAHIR"])
	}
	checkDataExsist := repository.CheckStatusTkExsist(kodeTK)

	if checkDataExsist >= 1 {
		return "", errors.New("Data sudah pernah diinputkan sebelumnya")
	}

	getSubmissionCode := repository.GetSubmissionCode()

	getOfficeCodeRequest := &request.GetOfficeCodeRequest{
		SegmenCode:     params.SegmenCode,
		Kpj:            params.Kpj,
		WorkerCode:     kodeTK,
		IdentityNumber: params.IdentityNumber,
		Fullname:       params.FullName,
		Birthdate:      tglLahir,
	}

	officeCode := repository.GetOfficeCode(getOfficeCodeRequest)

	session := db.EngineEcha.NewSession()
	defer session.Close()

	err2 := session.Begin()
	if err2 != nil {
		return "", err2
	}

	query := `insert /*+append*/ into bpjstku.asik_klaim(
										kode_pengajuan,
										email,
										no_hp,
										kode_tk,
										kpj,
										nomor_identitas,
										jenis_identitas,
										status_valid_identitas,
										nama_tk,
										tgl_lahir,
										tempat_lahir,
										jenis_kelamin,
										nama_ibu_kandung,
										kode_kantor,
										kode_segmen,
										kode_perusahaan,
										kode_divisi,
										kode_kepesertaan,
										tgl_kepesertaan,
										tgl_aktif,
										tgl_nonaktif,
										kode_na,
										sumber_blth_nonaktif,
										flag_sumber_nonaktif,
										tgl_flag_sumber_nonaktif,
										kode_tipe_klaim,
										kode_sebab_klaim,
										kode_pelaporan,
										tgl_klaim,
										npwp,
										kode_bank,
										nama_bank,
										no_rekening,
										nama_rekening,
										alamat,
										kode_kelurahan,
										kode_kecamatan,
										kode_kabupaten,
										kode_propinsi,
										kode_pos,
										status_valid_nama,
										status_valid_tgl_lahir,
										status_valid_tempat_lahir,
										status_valid_foto,
										score_face,
										score_face_liveness,
										is_alive,
										trx_id,
										ref_id,
										kode_manfaat,
										kode_tipe_penerima,
										kd_prg,
										tgl_pengembangan,
										rate_pengembangan,
										tgl_saldo_awal_tahun,
										nom_saldo_awal_tahun,
										nom_saldo_pengembangan,
										nom_saldo_total,
										nom_iuran_tahun_berjalan,
										nom_iuran_pengembangan,
										nom_iuran_total,
										nom_saldo_iuran_total,
										persentase_pengambilan,
										nom_manfaat_maxbisadiambil,
										nom_manfaat_diambil,
										nom_manfaat_gross,
										nom_pph,
										nom_pembulatan,
										nom_manfaat_netto,
										kanal_pelayanan,
										kode_kantor_pengajuan,
										status_pengajuan,
										tgl_pengajuan,
										status_submit_pengajuan,
										tgl_submit_pengajuan,
										tgl_rekam,
										petugas_rekam,
										status_submit_dokumen,
										tgl_submit_dokumen,
										platform,
										kode_billing,
										kemiripan_nama_pelapor,
										kode_pointer_asal
								) values (
										?, --kode_pengajuan
										?, --email
										?,--no_hp
										?,--kode_tk
										?,--kpj
										?,--nomor_identitas
										?,--jenis_identitas
										?,--status_valid_identitas
										?,--nama_tk
										to_date(?,'dd-mm-yyyy'),--tgl_lahir
										?,--tempat_lahir
										?,--jenis_kelamin
										?,--nama_ibu_kandung
										?,--kode_kantor
										?,--kode_segmen
										?,--kode_perusahaan
										?,--kode_divisi
										?,--kode_kepesertaan
										to_date(?,'dd-mm-yyyy'),--tgl_kepesertaan
										to_date(?,'dd-mm-yyyy'),--tgl_aktif
										to_date(?,'dd-mm-yyyy'),--tgl_nonaktif
										?,--kode_na
										?,--sumber_blth_nonaktif
										?,--flag_sumber_nonaktif
										to_date(?,'dd-mm-yyyy'),--tgl_flag_sumber_nonaktif
										?,--kode_tipe_klaim
										?,--kode_sebab_klaim
										?,--kode_pelaporan
										sysdate,--tgl_klaim
										?,--npwp
										?,--kode_bank
										?,--nama_bank
										?,--no_rekening
										?,--nama_rekening
										?,--alamat
										?,--kode_kelurahan
										?,--kode_kecamatan
										?,--kode_kabupaten
										?,--kode_propinsi
										?,--kode_pos
										'Y',--status_valid_nama
										'Y',--status_valid_tgl_lahir
										'Y',--status_valid_tempat_lahir
										'Y',--status_valid_foto
										?,--score_face
										?,--score_face_liveness
										'Y',--is_alive
										?,--trx_id
										?,--ref_id
										?,--kode_manfaat
										?,--kode_tipe_penerima
										?,--kd_prg
										to_date(?,'dd-mm-yyyy'),--tgl_pengembangan
										?,--rate_pengembangan
										to_date(?,'dd-mm-yyyy'),--tgl_saldo_awal_tahun
										?,--nom_saldo_awal_tahun
										?,--nom_saldo_pengembangan
										?,--nom_saldo_total
										?,--nom_iuran_tahun_berjalan
										?,--nom_iuran_pengembangan
										?,--nom_iuran_total
										?,--nom_saldo_iuran_total
										?,--persentase_pengambilan
										?,--nom_manfaat_maxbisadiambil
										?,--nom_manfaat_diambil
										?,--nom_manfaat_gross
										?,--nom_pph
										?,--nom_pembulatan
										?,--nom_manfaat_netto,
										'43',--kanal_pelayanan,
										?,--kode_kantor_pengajuan
										'KLA1',--status_pengajuan
										sysdate,--tgl_pengajuan
										'Y',--status_submit_pengajuan,
										sysdate,--tgl_submit_pengajuan
										sysdate,--tgl_rekam
										?,--petugas_rekam
										'Y',--status_submit_dokumen
										sysdate, --tgl_submit_dokumen,
										?, --platform
										?, --kodebilling
										?, --kemiripannama pelapor
										? --kode_pointerasal
									)
				`

	if _, err = session.Exec(
		query,
		getSubmissionCode,
		params.Email,
		params.PhoneNumber,
		kodeTK,
		params.Kpj,
		params.IdentityNumber,
		params.IdentityType,
		params.ValidIdentity,
		params.FullName,
		params.BirthDate,
		params.BirthPlace,
		params.Gender,
		params.MotherName,
		params.OfficeCode,
		params.SegmenCode,
		params.CompanyCode,
		params.DivisionCode,
		params.MembershipCode,
		params.MembershipDate,
		params.ActiveDate,
		params.NonActiveDate,
		params.NonActiveCode,
		params.SumberBlthNonAktif,
		params.FlagSubmerNonAktif,
		params.TglFlagSumberNonAktif,
		params.ClaimTypeCode,
		params.ClaimCauseCode,
		params.ReportCode,
		params.Npwp,
		params.BankCode,
		params.BankName,
		params.AccountBankNumber,
		params.AccountBankName,
		params.Address,
		params.KelurahanCode,
		params.KecamatanCode,
		params.KabupatenCode,
		params.PropinsiCode,
		params.PostalCode,
		params.ScoreFace,
		params.ScoreFaceLiveness,
		params.TrxId,
		params.RefId,
		params.BenefitCode,
		params.ReceiverTypeCode,
		params.ProgramCode,
		params.TanggalPengembangan,
		params.RatePengembangan,
		params.TanggalSaldoAwalTahun,
		params.NominalSaldoAwalTahun,
		params.NominalSaldoPengembangan,
		params.NominalSaldoTotal,
		params.NominalIuranTahunBerjalan,
		params.NominalIuranPengembangan,
		params.NominalIuranTotal,
		params.NominalSaldoIuranTotal,
		params.PersentasePengambilan,
		params.NominalManfaatBisaDiAmbil,
		params.NominalManfaatDiAmbil,
		params.NominalManfaatGross,
		params.NominalPPH,
		params.NominalPembulatan,
		params.NominalManfaatNetto,
		officeCode,
		params.PetugasPTPOS,
		params.Platform,
		params.KodeBillingPTPOS,
		params.ScoreSimilarityNama,
		params.KodeKantorPTPOS,
	); err != nil {
		session.Rollback()
		return "", err
	}

	for _, itemDokumen := range params.PathUrl {
		query := `insert /*+append*/ into bpjstku.asik_klaim_dokumen(kode_pengajuan, kode_dokumen, path_url_mitra)
						 values(?, ?, ?)`
		if _, err = session.Exec(query, getSubmissionCode, itemDokumen.KodeDokumen, itemDokumen.PathUrlDokumen); err != nil {
			session.Rollback()
			return "", err
		}
	}

	fmt.Println(params.DataDetailProbing, "xxxxxx")
	for _, itemProbing := range params.DataDetailProbing {
		query := `insert /*+append*/ into bpjstku.asik_probing(kode_pengajuan, kode_probing, no_urut,respon_probing,jawaban_probing,tgl_rekam,petugas_rekam)
						 values(?, ?, ?,?,?,sysdate,'PT POS')`
		if _, err = session.Exec(
			query,
			getSubmissionCode,
			itemProbing.KodeProbing,
			itemProbing.NomorUrut,
			itemProbing.ResponseProbing,
			itemProbing.JawabanProbing,
		); err != nil {
			session.Rollback()
			return "", err
		}
	}

	err = session.Commit()
	if err != nil {
		fmt.Println(err)
	}

	return getSubmissionCode, nil
}

func (repository *claimRepositoryImpl) GetClaimSubmission(params *request.GetClaimSubmissionRequest) (*entity.GetClaimSubmissionEntity, error) {
	var claimSubmission entity.GetClaimSubmissionEntity

	var sql = `
			select 	kode_pengajuan,
							email,
							no_hp,
							kode_tk,
							kpj,
							nomor_identitas,
							jenis_identitas,
							status_valid_identitas,
							nama_tk,
							to_char(tgl_lahir,'dd-mm-yyyy') tgl_lahir,
							tempat_lahir,
							jenis_kelamin,
							nama_ibu_kandung,
							kode_kantor_pengajuan,
							kode_segmen,
							kode_perusahaan,
							kode_divisi,
							kode_kepesertaan,
							to_char(tgl_kepesertaan,'dd-mm-yyyy') tgl_kepesertaan,
							to_char(tgl_aktif,'dd-mm-yyyy') tgl_aktif,
							to_char(tgl_nonaktif,'dd-mm-yyyy') tgl_nonaktif,
							kode_na,
							sumber_blth_nonaktif,
							flag_sumber_nonaktif,
							to_char(tgl_flag_sumber_nonaktif,'dd-mm-yyyy') tgl_flag_sumber_nonaktif,
							kode_tipe_klaim,
							kode_sebab_klaim,
							kode_pelaporan,
							to_char(tgl_klaim,'dd-mm-yyyy') tgl_klaim,
							npwp,
							kode_bank,
							nama_bank,
							no_rekening,
							nama_rekening,
							alamat,
							kode_kelurahan,
							kode_kecamatan,
							kode_kabupaten,
							kode_propinsi,
							kode_pos,
							status_valid_nama,
							status_valid_tgl_lahir,
							status_valid_tempat_lahir,
							status_valid_foto,
							score_face,
							score_face_liveness,
							is_alive,
							kode_manfaat,
							kode_tipe_penerima,
							kd_prg
			from 		bpjstku.asik_klaim 
			where 	kode_pengajuan=? 
							and status_batal='T'`
	result, err := db.EngineEcha.Query(
		sql,
		params.SubmissionCode,
	)
	if err != nil {
		return nil, err
	}
	claimSubmission = entity.GetClaimSubmissionEntity{
		KodePengajuan:          string(result[0]["KODE_PENGAJUAN"]),
		Email:                  string(result[0]["EMAIL"]),
		NoHp:                   string(result[0]["NO_HP"]),
		KodeTk:                 string(result[0]["KODE_TK"]),
		Kpj:                    string(result[0]["KPJ"]),
		NomorIdentitas:         string(result[0]["NOMOR_IDENTITAS"]),
		JenisIdentitas:         string(result[0]["JENIS_IDENTITAS"]),
		StatusValidIdentitas:   string(result[0]["STATUS_VALID_IDENTITAS"]),
		NamaTk:                 string(result[0]["NAMA_TK"]),
		TglLahir:               string(result[0]["TGL_LAHIR"]),
		TempatLahir:            string(result[0]["TEMPAT_LAHIR"]),
		JenisKelamin:           string(result[0]["JENIS_KELAMIN"]),
		NamaIbuKandung:         string(result[0]["NAMA_IBU_KANDUNG"]),
		KodeKantor:             string(result[0]["KODE_KANTOR_PENGAJUAN"]),
		KodeSegmen:             string(result[0]["KODE_SEGMEN"]),
		KodePerusahaan:         string(result[0]["KODE_PERUSAHAAN"]),
		KodeDivisi:             string(result[0]["KODE_DIVISI"]),
		KodeKepesertaan:        string(result[0]["KODE_KEPESERTAAN"]),
		TglKepesertaan:         string(result[0]["TGL_KEPESERTAAN"]),
		TglAktif:               string(result[0]["TGL_AKTIF"]),
		TglNonaktif:            string(result[0]["TGL_NONAKTIF"]),
		KodeNa:                 string(result[0]["KODE_NA"]),
		SumberBlthNonAktif:     string(result[0]["SUMBER_BLTH_NONAKTIF"]),
		FlagSumberNonAktif:     string(result[0]["FLAG_SUMBER_NONAKTIF"]),
		TglFlagSumberNonAktif:  string(result[0]["TGL_FLAG_SUMBER_NONAKTIF"]),
		KodeTipeKlaim:          string(result[0]["KODE_TIPE_KLAIM"]),
		KodeSebabKlaim:         string(result[0]["KODE_SEBAB_KLAIM"]),
		KodePelaporan:          string(result[0]["KODE_PELAPORAN"]),
		TglKlaim:               string(result[0]["TGL_KLAIM"]),
		Npwp:                   string(result[0]["NPWP"]),
		KodeBank:               string(result[0]["KODE_BANK"]),
		NamaBank:               string(result[0]["NAMA_BANK"]),
		NoRekening:             string(result[0]["NO_REKENING"]),
		NamaRekening:           string(result[0]["NAMA_REKENING"]),
		Alamat:                 string(result[0]["ALAMAT"]),
		KodeKelurahan:          string(result[0]["KODE_KELURAHAN"]),
		KodeKecamatan:          string(result[0]["KODE_KECAMATAN"]),
		KodeKabupaten:          string(result[0]["KODE_KABUPATEN"]),
		KodePropinsi:           string(result[0]["KODE_PROPINSI"]),
		KodePos:                string(result[0]["KODE_POS"]),
		StatusValidNama:        string(result[0]["STATUS_VALID_NAMA"]),
		StatusValidTglLahir:    string(result[0]["STATUS_VALID_TGL_LAHIR"]),
		StatusValidTempatLahir: string(result[0]["STATUS_VALID_TEMPAT_LAHIR"]),
		StatusValidFoto:        string(result[0]["STATUS_VALID_FOTO"]),
		ScoreFace:              string(result[0]["SCORE_FACE"]),
		ScoreFaceLiveness:      string(result[0]["SCORE_FACE_LIVENESS"]),
		IsAlive:                string(result[0]["IS_ALIVE"]),
		KodeManfaat:            string(result[0]["KODE_MANFAAT"]),
		KodeTipePenerima:       string(result[0]["KODE_TIPE_PENERIMA"]),
		KdPrg:                  string(result[0]["KD_PRG"]),
	}

	return &claimSubmission, nil
}

func (repository *claimRepositoryImpl) DetailContribution(params *request.PraClaimRequest, companyName string) ([]entity.DetailContributionEntity, error) {
	var sql = `
			select 	to_char(blth, 'dd-mm-yyyy') blth, 
					to_char(nom_iuran) nom_iuran
			from 		table(bpjstku.p_ptpos_asik_klaim.f_get_rincian_iuran_berjalan(?, ?, ?, ? )) where nama_perusahaan = ?
	`
	resultQry, err := db.EngineEcha.Query(
		sql,
		params.SegmenCode,
		params.CompanyCode,
		params.DivisionCode,
		params.WorkerCode,
		companyName,
	)
	if err != nil {
		return nil, err
	}

	entities := make([]entity.DetailContributionEntity, 0)
	for _, item := range resultQry {
		headerContribution := entity.DetailContributionEntity{
			Blth:                string(item["BLTH"]),
			NominalContribution: string(item["NOM_IURAN"]),
		}
		entities = append(entities, headerContribution)
	}

	return entities, nil
}

func (repository *claimRepositoryImpl) HeaderContribution(params *request.PraClaimRequest) ([]entity.HeaderContributionEntity, error) {
	var sql = `
			select 	distinct nama_perusahaan 
			from 		table(bpjstku.p_ptpos_asik_klaim.f_get_rincian_iuran_berjalan(?, ?, ?, ?))
	`
	resultQry, err := db.EngineEcha.Query(
		sql,
		params.SegmenCode,
		params.CompanyCode,
		params.DivisionCode,
		params.WorkerCode,
	)
	if err != nil {
		return nil, err
	}

	entities := make([]entity.HeaderContributionEntity, 0)
	for _, item := range resultQry {
		headerContribution := entity.HeaderContributionEntity{
			CompanyName: string(item["NAMA_PERUSAHAAN"]),
		}
		entities = append(entities, headerContribution)
	}

	return entities, nil
}

func (repository *claimRepositoryImpl) CheckDataUpdate(params *request.PraClaimRequest) (*entity.DataUpdateEntity, error) {
	var check entity.DataUpdateEntity

	var qry = `select 	bpjstku.p_ptpos_asik_klaim.f_cek_pengkinian_data(?, ?, ?, ?, ?, to_date(?, 'dd-mm-yyyy')) status
			from 		dual`
	if _, err := db.EngineEcha.SQL(
		qry,
		params.SegmenCode,
		params.Kpj,
		params.WorkerCode,
		params.IdentityNumber,
		params.FullName,
		params.BirthDate,
	).Get(&check.Status); err != nil {
		return nil, err
	}

	return &check, nil
}

func (repository *claimRepositoryImpl) CheckMembershipStatus(params *request.PraClaimRequest) (*entity.MembershipStatusEntity, error) {
	var check entity.MembershipStatusEntity

	var qry = `select bpjstku.p_ptpos_asik_klaim.f_cek_status_keps_nonaktif(
					?, 
					?, 
					?, 
					?,
					?,
					to_date(?, 'dd-mm-yyyy')) status
				from dual`
	if _, err := db.EngineEcha.SQL(qry, params.SegmenCode, params.Kpj, params.WorkerCode, params.IdentityNumber, params.FullName, params.BirthDate).Get(&check.Status); err != nil {
		return nil, err
	}

	return &check, nil
}

func (repository *claimRepositoryImpl) CheckNpwpInformation(params *request.PraClaimRequest) (*entity.NpwpInformationEntity, error) {
	var check entity.NpwpInformationEntity

	var qry = `select bpjstku.p_ptpos_asik_klaim.f_cek_informasi_npwp(
						?, 
						?, 
						?, 
						?,
						?,
						to_date(?, 'dd-mm-yyyy')) status
					from dual`
	if _, err := db.EngineEcha.SQL(qry, params.SegmenCode, params.Kpj, params.WorkerCode, params.IdentityNumber, params.FullName, params.BirthDate).Get(&check.Status); err != nil {
		return nil, err
	}

	return &check, nil
}

func (repository *claimRepositoryImpl) CheckRsjht(params *request.CheckRsjhtRequest) (*entity.CheckRsjhtEntity, error) {
	var check entity.CheckRsjhtEntity

	var qry = `select bpjstku.p_ptpos_asik_klaim.f_cek_ada_rsjhtjp(?, ?, ?) rsjht from dual`
	if _, err := db.EngineEcha.SQL(qry, params.Year, params.Kpj, params.WorkerCode).Get(&check.Status); err != nil {
		return nil, err
	}

	return &check, nil
}

func (repository *claimRepositoryImpl) GetDataAfterClaim(params *request.GetDataAfterClaimRequest) (*entity.GetDataAfterClaimEntity, error) {
	var data entity.GetDataAfterClaimEntity

	var qry = `
			select 	kode_pengajuan,
							kode_klaim,
							email,
							no_hp,
							kpj,
							nomor_identitas,
							nama_tk,
							to_char(tgl_lahir, 'dd-mm-yyyy') tgl_lahir,
							tempat_lahir,
							jenis_kelamin,
							npwp,
							kode_bank,
							nama_bank,
							no_rekening,
							nama_rekening,
							nom_manfaat_netto,
							alamat,
							status_bayar,
							to_char(tgl_bayar, 'dd-mm-yyyy HH24:MI') tgl_bayar,
							to_char(tgl_kepesertaan, 'dd-mm-yyyy') tgl_kepesertaan,
							to_char(tgl_aktif, 'dd-mm-yyyy') tgl_aktif,
							to_char(tgl_nonaktif, 'dd-mm-yyyy') tgl_nonaktif
			from 		table(bpjstku.p_ptpos_asik_klaim.f_get_informasi_pasca_klaim(?))`
	result, err := db.EngineEcha.Query(
		qry,
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

func (repository *claimRepositoryImpl) TrackingClaim(params *request.TrackingClaimRequest) ([]entity.TrackingClaimEntity, error) {
	var qry = `
			select 	kode_klaim,
							kode_tipe_klaim,
							nik,
							kpj,
							nama_tk,
							tahap,
							ket_tahap,
							to_char(tgl_rekam, 'dd-mm-yyyy HH24:MI:SS') tgl_rekam,
							notes,
							nama_tahap
			from 		table (pn.f_tracking_klaim(?, ?))
	`
	resultQry, err := db.EngineOltp.Query(
		qry,
		params.Kpj,
		params.IdentityNumber,
	)
	if err != nil {
		return nil, err
	}

	objects := make([]entity.TrackingClaimEntity, 0)
	for _, item := range resultQry {
		object := entity.TrackingClaimEntity{
			ClaimCode:       string(item["KODE_KLAIM"]),
			ClaimTypeCode:   string(item["KODE_TIPE_KLAIM"]),
			IdentityNumber:  string(item["NIK"]),
			Kpj:             string(item["KPJ"]),
			Fullname:        string(item["NAMA_TK"]),
			Step:            string(item["TAHAP"]),
			StepInformation: string(item["KET_TAHAP"]),
			CreatedAt:       string(item["TGL_REKAM"]),
			Notes:           string(item["NOTES"]),
			StepName:        string(item["NAMA_TAHAP"]),
		}
		objects = append(objects, object)
	}
	return objects, nil
}

func (repository *claimRepositoryImpl) TrackingHeaderClaim(params *request.TrackingClaimRequest) (*entity.TrackingHeaderClaimEntity, error) {
	var data entity.TrackingHeaderClaimEntity

	var qry = `
			select 	kode_klaim,
							kode_tipe_klaim,
							nik,
							kpj,
							nama_tk
			from 		table (pn.f_tracking_klaim(?, ?)) group by kode_klaim, kode_tipe_klaim, nik, kpj, nama_tk
	`
	resultQry, err := db.EngineOltp.Query(
		qry,
		params.Kpj,
		params.IdentityNumber,
	)
	if err != nil {
		return nil, err
	}

	for _, item := range resultQry {
		data = entity.TrackingHeaderClaimEntity{
			ClaimCode:      string(item["KODE_KLAIM"]),
			ClaimTypeCode:  string(item["KODE_TIPE_KLAIM"]),
			IdentityNumber: string(item["NIK"]),
			Kpj:            string(item["KPJ"]),
			FullName:       string(item["NAMA_TK"]),
		}

	}

	return &data, nil
}

func (repository *claimRepositoryImpl) CheckMaximumBalanceJht(params *request.PraClaimRequest) (*entity.CheckSubmissionMaximumBalanceEntity, error) {
	var check entity.CheckSubmissionMaximumBalanceEntity

	var qry = `select bpjstku.p_ptpos_asik_klaim.f_cek_pengajuan_maksimal_jht(
						?, 
						?, 
						?, 
						?,
						?,
						to_date(?, 'dd-mm-yyyy')) status
					from dual`
	if _, err := db.EngineEcha.SQL(qry, params.SegmenCode, params.Kpj, params.WorkerCode, params.IdentityNumber, params.FullName, params.BirthDate).Get(&check.Status); err != nil {
		return nil, err
	}

	return &check, nil
}

func (repository *claimRepositoryImpl) GetParameterSystem(params *request.GetParamterSystemRequest) (*entity.ParameterSystemEntity, error) {
	var system entity.ParameterSystemEntity

	var qry = `
			select 	kode_parameter, 
							nama_parameter, 
							nilai 
			from 		bpjstku.asik_parameter_sistem 
			where 	kode_parameter=?
							and nvl(AKTIF,'X') = 'Y'
	`
	result, err := db.EngineEcha.Query(qry, params.ParameterCode)
	if err != nil {
		return nil, err
	}

	system = entity.ParameterSystemEntity{
		ParameterCode: string(result[0]["KODE_PARAMETER"]),
		ParameterName: string(result[0]["NAMA_PARAMETER"]),
		Value:         string(result[0]["NILAI"]),
	}

	return &system, nil
}

func (repository *claimRepositoryImpl) CheckStatusPensiun(params *request.CheckStatusPensiunRequest) (*entity.CheckStatusPensiunEntity, error) {
	var check entity.CheckStatusPensiunEntity

	var qry = `select bpjstku.p_ptpos_asik_klaim.f_cek_usia_pensiun_jht(to_date(?, 'dd-mm-yyyy')) from dual`
	if _, err := db.EngineEcha.SQL(qry, params.Birthdate).Get(&check.Status); err != nil {
		return nil, err
	}

	return &check, nil
}

func (repository *claimRepositoryImpl) CheckStatusTkExsist(workerCode string) int {
	var total int

	var qry = `
			select 	count(1) 
			from 		bpjstku.asik_klaim x 
			where 	kode_tk = ? 
							and nvl(status_batal, 'X') = 'T' 
							and status_pengajuan not in ('KLA5','KLA6')
	`
	if _, err := db.EngineEcha.SQL(qry, workerCode).Get(&total); err != nil {
		fmt.Println(err)
	}

	return total
}

func (repository *claimRepositoryImpl) GetSubmissionCode() string {
	var claimCode string

	var qry = `select bpjstku.p_ptpos_asik_klaim.f_gen_kode_pengajuan from dual`
	if _, err := db.EngineEcha.SQL(qry).Get(&claimCode); err != nil {
		fmt.Println(err)
	}

	return claimCode
}

func (repository *claimRepositoryImpl) GetOfficeCode(params *request.GetOfficeCodeRequest) string {
	var officeCode string

	var qry = `select bpjstku.p_ptpos_asik_klaim.f_get_kode_kantor_pengajuan
															(
																?,
																?,
																?,
																?,
																?,
																?
															)
														from dual`
	if _, err := db.EngineEcha.SQL(
		qry,
		params.SegmenCode,
		params.Kpj,
		params.WorkerCode,
		params.IdentityNumber,
		params.Fullname,
		params.Birthdate,
	).Get(&officeCode); err != nil {
		fmt.Println(err)
	}

	return officeCode
}

func (repository *claimRepositoryImpl) InsertNotifications(params *request.InsertNotificationRequest) (*response.InsertNotificationResponse, error) {
	configuration := config.New()
	client := &http.Client{}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(jsonData))
	req, err := http.NewRequest("POST", configuration.Get("JMO_NOTIFICATION_URL")+"/notification", payload)

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

func (repository *claimRepositoryImpl) SendEmail(params *request.SendEmailRequest) (*response.SendEmailResponse, error) {
	configuration := config.New()
	client := &http.Client{}

	//send Email
	qryEmail := `select * from kit.vas_email_content where vas_email_content_id=?`
	resultEmail, err := db.EngineOltp.Query(qryEmail, "52")
	if err != nil {
		return nil, err
	}

	var tgl_rekam string
	var qry = `select to_char(tgl_rekam,'DD-MM-YYYY HH24:mm:ss')||' WIB' tgl_rekam from bpjstku.asik_klaim where kode_pengajuan='` + params.KodePengajuan + `'`
	if _, err := db.EngineEcha.SQL(qry).Get(&tgl_rekam); err != nil {
		fmt.Println(err)
	}

	ContentEmail := string(resultEmail[0]["HTML_CONTENT"])
	string1 := strings.Replace(ContentEmail, ":0:", params.Kpj, 1)
	string2 := strings.Replace(string1, ":1:", params.FullName, 1)
	kanalLayanan := strings.Split(params.KodeKantorPPTOS, "|")
	string3 := strings.Replace(string2, ":2:", strings.ToUpper(kanalLayanan[1]), 1)
	string4 := strings.Replace(string3, ":3:", params.KodePengajuan, 1)
	string5 := strings.Replace(string4, ":4:", tgl_rekam, 1)
	string6 := strings.Replace(string5, ":5:", "PENGAJUAN", 1)

	SubjectContent := string(resultEmail[0]["SUBJECT"])
	SubjectContent1 := strings.Replace(SubjectContent, ":0:", params.Kpj, 1)

	paramSendEmail := &request.SendEmailRequest{
		Subject: SubjectContent1,
		Body:    SubjectContent1,
		Message: string6,
		Email:   params.Email,
	}

	jsonDataEmail, err := json.Marshal(paramSendEmail)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(jsonDataEmail))
	req, err := http.NewRequest("POST", configuration.Get("JS_PTPOS_URL")+"/JSPTPOS/SendEmail", payload)
	if err != nil {
		return nil, err
	}

	sendEmailResponse := &response.SendEmailResponse{}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	json.NewDecoder(res.Body).Decode(&sendEmailResponse)
	defer res.Body.Close()
	//end of send email

	if res.StatusCode == 200 {
		return sendEmailResponse, nil
	} else {
		return nil, errors.New(sendEmailResponse.Message)
	}
}

func (repository *claimRepositoryImpl) CheckEligibleInsert(params *request.CheckEligibleJHTRequest) (*response.CheckJHTEligibleResponse, error) {
	var p_param_maksimal_saldo_jht string
	var p_klaim_sebagian string
	var p_kode_pesan string
	var p_pesan_tidak_layak string
	var p_nama_kanal_kantor_cabang string
	var p_sukses string
	var p_mess string
	var kodeTK string
	var kodeSegmen string
	var tglLahir string
	var email string
	// var cntKodeTK int

	// cntKodeTK = repository.CheckTotalKodeTK(params.Kpj)
	// // fmt.Println(cntKodeTK, "0")
	// if cntKodeTK > 1 {
	// 	result := &entity.CheckJHTEligibleEntity{
	// 		StatusKelayakan: "T",
	// 	}
	// 	return
	// }

	var qry = `
	select kode_tk,kode_segmen,tgl_lahir,email
    from (
        select a.kode_tk,a.kode_segmen,to_char(a.tgl_lahir,'dd-mm-yyyy')tgl_lahir,a.email
        from kn.vw_kn_tk a where kpj =?
        )xx 
	group by xx.kode_tk,kode_segmen,tgl_lahir,email
	`
	// fmt.Println(qry, "1")
	result, err := db.EngineOltp.Query(qry, params.Kpj)
	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		kodeTK = string(result[0]["KODE_TK"])
		kodeSegmen = string(result[0]["KODE_SEGMEN"])
		email = string(result[0]["EMAIL"])
	}

	var query = `begin
					bpjstku.p_ptpos_asik_klaim.x_cek_kelayakan_jht(?,
																 ?,
																 ?,
																 ?,
																 ?,
																 TO_DATE(? ,'dd-mm-yyyy'),
																 ?,
																 :p_param_maksimal_saldo_jht,
																 :p_klaim_sebagian,
																 :p_kode_pesan,
																 :p_nama_kanal_kantor_cabang,
																 :p_pesan_tidak_layak,
																 :p_sukses,
																 :p_mess);
			 end;`
	// fmt.Println(qry, "4")

	fmt.Println(kodeSegmen, params.Kpj, kodeTK, params.Nik, params.Fullname, params.TglLahir, email)
	if _, err := db.EngineEcha.Exec(query,
		kodeSegmen,
		params.Kpj,
		kodeTK,
		params.Nik,
		params.Fullname,
		tglLahir,
		email,
		sql.Named("p_param_maksimal_saldo_jht", sql.Out{Dest: &p_param_maksimal_saldo_jht}),
		sql.Named("p_klaim_sebagian", sql.Out{Dest: &p_klaim_sebagian}),
		sql.Named("p_kode_pesan", sql.Out{Dest: &p_kode_pesan}),
		sql.Named("p_nama_kanal_kantor_cabang", sql.Out{Dest: &p_nama_kanal_kantor_cabang}),
		sql.Named("p_pesan_tidak_layak", sql.Out{Dest: &p_pesan_tidak_layak}),
		sql.Named("p_sukses", sql.Out{Dest: &p_sukses}),
		sql.Named("p_mess", sql.Out{Dest: &p_mess})); err != nil {
		return nil, err
	}
	switch p_kode_pesan {
	case "JHTA000":
		p_pesan_tidak_layak = "Peserta LAYAK mengajukan klaim via PT POS, silahkan melajutkan proses verifikasi biometrik dan proses upload data peserta"
	case "JHTA002":
		p_pesan_tidak_layak = "Hasil :Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : peserta memiliki lebih dari 1 kartu peserta aktif dan membutuhkan verifikasi lebih lanjut di kantor cabang BPJS Ketenagakerjaan |Tindak Lanjut : Silahkan melakukan proses klaim JHT melalui Kantor Cabang BPJS Ketenagakerjaan terdekat  atau melalui website lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA003":
		p_pesan_tidak_layak = "Hasil :Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : peserta memiliki lebih dari 1 kartu peserta aktif dan membutuhkan verifikasi lebih lanjut di kantor cabang BPJS Ketenagakerjaan |Tindak Lanjut : Silahkan melakukan proses klaim JHT melalui Kantor Cabang BPJS Ketenagakerjaan terdekat  atau melalui website lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA004":
		p_pesan_tidak_layak = "Hasil:Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab :Peserta sudah mengajukan klaim di kanal layanan lain|Tindak Lanjut : Silahkan menghubungi Kantor Cabang BPJS Ketenagakerjaan terdekat"
	case "JHTA005":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Saldo peserta diatas batas kewenangan pengajuan klaim via PT POS|Tindak Lanjut : Silahkan melakukan proses klaim JHT melalui Kantor Cabang BPJS Ketenagakerjaan terdekat  atau melalui website lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA006":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS| Sebab : Data nomor kepesertaan yang di cari tidak valid dan tidak ditemukan pada data kepesertaan kami|Tindak Lanjut : Silahkan melakukan proses klaim JHT melalui Kantor Cabang BPJS Ketenagakerjaan terdekat  atau melalui website lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA007":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Peserta memiliki beberapa saldo JHT dalam satu kartu yang belum tergabung, dan membutuhkan verifikasi lebih lanjut di kantor cabang BPJS Ketenagakerjaan|Tindak Lanjut : Silahkan melakukan proses klaim JHT melalui Kantor Cabang BPJS Ketenagakerjaan terdekat  atau melalui website lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA008":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Peserta memiliki beberapa saldo JHT dalam satu kartu yang belum tergabung, dan membutuhkan verifikasi lebih lanjut di kantor cabang BPJS Ketenagakerjaan|Tindak Lanjut : Silahkan melakukan proses klaim JHT melalui Kantor Cabang BPJS Ketenagakerjaan terdekat  atau melalui website lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA009":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Data peserta sedang dalam proses perbaikan data kepesertaan, dan saat ini tidak dapat melanjutkan proses klain JHT sampai proses perbaikan selesai dilakukan|Tindak Lanjut : Untuk informasi lebih lanjut  silakan  menghubungi perusahaan Anda atau Kantor Cabang BPJS Ketenagakerjaan terdekat, untuk proses pengajuan klaim anda dapat memanfaatkan aplikasi JMO dan aplikasi Lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA010":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Data peserta sedang dalam proses perbaikan data kepesertaan, dan saat ini tidak dapat melanjutkan proses klain JHT sampai proses perbaikan selesai dilakukan|Tindak Lanjut : Untuk informasi lebih lanjut  silakan  menghubungi perusahaan Anda atau Kantor Cabang BPJS Ketenagakerjaan terdekat, untuk proses pengajuan klaim anda dapat memanfaatkan aplikasi JMO dan aplikasi Lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA011":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Data peserta sedang dalam proses perbaikan data kepesertaan, dan saat ini tidak dapat melanjutkan proses klain JHT sampai proses perbaikan selesai dilakukan|Tindak Lanjut : Untuk informasi lebih lanjut  silakan  menghubungi perusahaan Anda atau Kantor Cabang BPJS Ketenagakerjaan terdekat, untuk proses pengajuan klaim anda dapat memanfaatkan aplikasi JMO dan aplikasi Lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA012":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Data peserta sedang dalam proses perbaikan data kepesertaan, dan saat ini tidak dapat melanjutkan proses klain JHT sampai proses perbaikan selesai dilakukan|Tindak Lanjut : Untuk informasi lebih lanjut  silakan  menghubungi perusahaan Anda atau Kantor Cabang BPJS Ketenagakerjaan terdekat, untuk proses pengajuan klaim anda dapat memanfaatkan aplikasi JMO dan aplikasi Lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA013":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Data peserta sedang dalam proses perbaikan data kepesertaan, dan saat ini tidak dapat melanjutkan proses klain JHT sampai proses perbaikan selesai dilakukan|Tindak Lanjut : Untuk informasi lebih lanjut  silakan  menghubungi perusahaan Anda atau Kantor Cabang BPJS Ketenagakerjaan terdekat, untuk proses pengajuan klaim anda dapat memanfaatkan aplikasi JMO dan aplikasi Lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA015":
		p_pesan_tidak_layak = "Hasil :Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Data peserta sedang dalam proses penggabungan saldo JHT, dan saat ini tidak dapat melanjutkan proses klain JHT sampai proses penggabungan saldo JHT selesai dilakukan|Tindak Lanjut : Untuk informasi lebih lanjut  silakan  menghubungi perusahaan Anda atau Kantor Cabang BPJS Ketenagakerjaan terdekat, untuk proses pengajuan klaim anda dapat memanfaatkan aplikasi JMO dan aplikasi Lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA017":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Terdapat permasalahan dalam rincian saldo JHT peserta dan dan membutuhkan verifikasi lebih lanjut di kantor cabang BPJS Ketenagakerjaan|Tindak Lanjut : Silahkan melakukan proses klaim JHT melalui Kantor Cabang BPJS Ketenagakerjaan terdekat  atau melalui website lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA018":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Terdapat sisa rincian saldo JHT peserta pada pengajuan klaim sebelumnya dan untuk harus dilakukan penggabungan dan verifikasi lebih lanjut di kantor cabang BPJS Ketenagakerjaan|Tindak Lanjut :Silahkan melakukan proses klaim JHT melalui Kantor Cabang BPJS Ketenagakerjaan terdekat  atau melalui website lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA019":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Peserta masih dalam masa tunggu 1 bulan sejak dilakukan proses Non Aktif kepesertaan dan belum layan mengajukan klaim sampai selesai masa tunggu satu bulan |Tindak Lanjut : Untuk informasi lebih lanjut  silakan  menghubungi perusahaan Anda atau Kantor Cabang BPJS Ketenagakerjaan terdekat, untuk proses pengajuan klaim anda dapat memanfaatkan aplikasi JMO dan aplikasi Lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA020":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Data peserta sedang dalam proses pengkinian data kepesertaan, dan saat ini tidak dapat melanjutkan proses klain JHT sampai proses pengkinian data kepesertaan selesai dilakukan|Tindak Lanjut : Untuk informasi lebih lanjut  silakan  menghubungi perusahaan Anda atau Kantor Cabang BPJS Ketenagakerjaan terdekat, untuk proses pengajuan klaim anda dapat memanfaatkan aplikasi JMO dan aplikasi Lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA021":
		p_pesan_tidak_layak = "Hasil :Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Peserta sedang mengajukan klaim di kanal layanan ${it.data.branchOfficeChannelName}, |Tindak Lanjut : Silahkan menghubungi Kantor Cabang BPJS Ketenagakerjaan terdekat atau ke contact center 175"
	case "JHTA100":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Pengajuan sebab klaim JHT peserta membutuhkan verifikasi lebih lanjut di kantor cabang BPJS Ketenagakerjaan Tindak Lanjut : Silahkan melakukan proses klaim JHT melalui Kantor Cabang BPJS Ketenagakerjaan terdekat  atau melalui website lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA022":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Pengajuan sebab klaim JHT peserta membutuhkan verifikasi lebih lanjut di kantor cabang BPJS Ketenagakerjaan|Tindak Lanjut : Silahkan melakukan proses klaim JHT melalui Kantor Cabang BPJS Ketenagakerjaan terdekat  atau melalui website lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA023":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Pengajuan sebab klaim JHT peserta membutuhkan verifikasi lebih lanjut di kantor cabang BPJS Ketenagakerjaan|Tindak Lanjut : Silahkan melakukan proses klaim JHT melalui Kantor Cabang BPJS Ketenagakerjaan terdekat  atau melalui website lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
	case "JHTA024":
		p_pesan_tidak_layak = "Hasil : Peserta TIDAK LAYAK mengajukan klaim via PT POS|Sebab : Pengajuan sebab klaim JHT peserta membutuhkan verifikasi lebih lanjut di kantor cabang BPJS Ketenagakerjaan| Tindak Lanjut : Silahkan melakukan proses klaim JHT melalui Kantor Cabang BPJS Ketenagakerjaan terdekat  atau melalui website lapakasik online di lapakasik.bpjsketenagakerjaan.go.id"
		// default:
		// 	return MessageEligibleClaim{
		// 		StatusCode: "120",
		// 		Message:    "Silahkan datang ke kantor cabang",
		// 	}
	}
	var objectw = []response.CheckJHTEligible{{
		StatusKelayakan:     "Y",
		KodeKelayakan:       p_kode_pesan,
		KeteranganKelayakan: p_pesan_tidak_layak,
	}}

	resultEligible := &response.CheckJHTEligibleResponse{
		StatusCode: 200,
		StatusDesc: "Ok",
		Data:       objectw,
	}

	return resultEligible, nil
}
