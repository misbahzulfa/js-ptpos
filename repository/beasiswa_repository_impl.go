package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"js-ptpos/config"
	db "js-ptpos/config"
	"js-ptpos/entity"
	"js-ptpos/exception"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
	"net/http"
	"strconv"
	"strings"
)

func NewBeasiswaRepository() BeasiswaRepository {
	return &BeasiswaRepositoryImpl{}
}

type BeasiswaRepositoryImpl struct {
}

func (repository *BeasiswaRepositoryImpl) CheckEligibleBeasiswa(params *request.CheckEligibleBeasiswaRequest) (*response.CheckEligibleBeasiswaResponse, error) {
	var p_status_pencarian string
	var p_keterangan_pencarian string
	var p_sukses string
	var p_mess string

	session := db.EngineEcha.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return nil, err
	}

	result, err := db.EngineEcha.Query(`select kode, keterangan from BPJSTKU.ASIK_MS_LOOKUP 
                                                                     where tIPE= 'KODE_DOK_PELAPOR_BEASISWA'
    `)

	if err != nil {
		exception.PanicIfNeeded(err)
	}

	dokumens := make([]response.DataDokumenBeasiswa, 0)
	for _, item := range result {
		dokumen := response.DataDokumenBeasiswa{
			KodeDokumen: string(item["KODE"]),
			NamaDokumen: string(item["KETERANGAN"]),
		}
		dokumens = append(dokumens, dokumen)
	}

	resultProbing, err := db.EngineEcha.Query(`
    select
    distinct
    kode_probing, 
    nama_probing, 
    kategori, 
    keterangan,
    respon_probing
    from
    (
    select 
    a.kode_probing, 
    a.nama_probing, 
    a.kategori, 
    a.keterangan,
    case 
        when a.kode_probing = 'PROBING001' then
            (select nama_tk from kn.vw_kn_tk@to_kn
            where nomor_identitas = ? 
            and rownum = 1)
        when a.kode_probing = 'PROBING002' then
            (select tempat_lahir||' / '||to_char(nvl(tgl_lahir,'31-DEC-3000'),'DD-MM-YYYY') from kn.vw_kn_tk@to_kn
            where nomor_identitas = ? 
            and rownum = 1)  
        when a.kode_probing = 'PROBING003' then
            (select nama_ibu_kandung from kn.vw_kn_tk@to_kn
            where nomor_identitas = ? 
            and rownum = 1)  
        when a.kode_probing = 'PROBING004' then
            (select alamat from kn.vw_kn_tk@to_kn
            where nomor_identitas = ? 
            and rownum = 1)  
        else null
    end respon_probing  
    from BPJSTKU.ASIK_KODE_PROBING a inner join BPJSTKU.ASIK_KODE_PROBING_PENERIMA b on b.KODE_PROBING = a.KODE_PROBING
    where substr(b.KODE_TIPE_KLAIM,1,3) in ('JKK','JKM')
            and nvl(b.STATUS_NONAKTIF,'X') = 'T'
            order by a.no_urut
            )
    `, params.NikPeserta,
		params.NikPeserta,
		params.NikPeserta,
		params.NikPeserta)

	if err != nil {
		exception.PanicIfNeeded(err)
	}

	probings := make([]response.DataProbingBeasiswa, 0)
	for _, item := range resultProbing {
		probing := response.DataProbingBeasiswa{
			KodeProbing:   string(item["KODE_PROBING"]),
			NoUrut:        string(item["NAMA_PROBING"]),
			ResponProbing: string(item["RESPON_PROBING"]),
			Kategori:      string(item["KATEGORI"]),
		}
		probings = append(probings, probing)
	}

	// return dataSet, nil

	query := `begin
                pn.p_pn_ptpos2pn.x_cek_kelayakan_konf_beasiswa(?,
                                                                 ?,
                                                                 ?,
                                                                 ?,                                                                                                                              
                                                                 :p_status_pencarian,
                                                                 :p_keterangan_pencarian,
                                                                 :p_sukses,
                                                                 :p_mess                                                                 
                                                                 );
             end;`
	if _, err := db.EngineOltp.Exec(query,
		params.NikPelapor,
		params.NamaPelapor,
		params.NikPeserta,
		params.ChannelID,
		sql.Named("p_status_pencarian", sql.Out{Dest: &p_status_pencarian}),
		sql.Named("p_keterangan_pencarian", sql.Out{Dest: &p_keterangan_pencarian}),
		sql.Named("p_sukses", sql.Out{Dest: &p_sukses}),
		sql.Named("p_mess", sql.Out{Dest: &p_mess})); err != nil {
		return nil, err
	}

	var CheckEligibleBeasiswa = []response.CheckEligibleBeasiswa{{
		StatusPencarian:     p_status_pencarian,
		KeteranganPencarian: p_keterangan_pencarian,
		DataDokumenBeasiswa: dokumens,
		DataProbingBeasiswa: probings,
	}}

	resultCheckEligibleBeasiswa := &response.CheckEligibleBeasiswaResponse{
		StatusCode:            "200",
		StatusDesc:            "Data Ditemukan",
		CheckEligibleBeasiswa: CheckEligibleBeasiswa,
	}

	return resultCheckEligibleBeasiswa, nil
}

func (repository *BeasiswaRepositoryImpl) GetKodePengajuan() string {
	var claimCode string

	var qry = `select bpjstku.p_ptpos_asik_klaim.f_gen_kode_pengajuan from dual`
	if _, err := db.EngineEcha.SQL(qry).Get(&claimCode); err != nil {
		fmt.Println(err)
	}

	return claimCode
}

func (repository *BeasiswaRepositoryImpl) InsertKonfirmasiBeasiswa(params *request.InsertKonfirmasiBeasiswaRequest) (*response.InsertKonfirmasiBeasiswaResponse, error) {

	configuration := config.New()
	client := &http.Client{}

	session := db.EngineEcha.NewSession()
	defer session.Close()

	GetKodePengajuan := repository.GetKodePengajuan()
	var KodeKlaimInduk string
	var KodeKantorPengajuan string
	var KodeManfaat string

	err := session.Begin()
	if err != nil {
		return nil, err
	}

	resultKodeKlaimInduk, err := db.EngineOltp.Query(`
    select
    kode_klaim_induk,
    kode_klaim,
    kode_manfaat,
    kode_kantor
    from
    (
        select
        a.kode_klaim,
        a.kode_klaim_induk,
        d.kode_manfaat,
        a.tgl_klaim,
        a.tgl_ubah,
        a.kode_kantor,
        b.nomor_identitas nik_pelapor,
        b.nama_penerima nama_pelapor,
        a.nomor_identitas nik_peserta,
        rank() over (order by a.tgl_klaim,a.tgl_ubah desc) ranktglubah
        from pn.pn_klaim a
        inner join pn.pn_klaim_penerima_manfaat b on b.kode_klaim = a.kode_klaim
        inner join pn.pn_klaim_manfaat_detil d on d.kode_klaim = a.kode_klaim
        inner join pn.pn_klaim_manfaat_detil_beasis c on c.kode_klaim = b.kode_klaim
        where
            a.kode_klaim in
                        (select kode_klaim from pn.pn_klaim
                            start with kode_klaim = (
                                                        select
                                                        distinct
                                                        a.kode_klaim
                                                        from pn.pn_klaim a
                                                        inner join pn.pn_klaim_penerima_manfaat b on b.kode_klaim = a.kode_klaim
                                                        inner join pn.pn_klaim_manfaat_detil_beasis c on c.kode_klaim = b.kode_klaim
                                                        where
                                                        -- a.kode_klaim = 'KL18032202568009'
                                                            b.nomor_identitas = ?
                                                           -- and b.nama_penerima = ?
                                                           and UTL_MATCH.EDIT_DISTANCE_SIMILARITY (upper(B.NAMA_PENERIMA), UPPER(?)) > 75 
                                                            and a.nomor_identitas = ?
                                                            and nvl(status_batal,'X') = 'T'
                                                            and rownum = 1 --nanti akan di prior lagi
                                                    )
                                connect by prior kode_klaim = kode_klaim_induk
                        )
    )y
    where ranktglubah = 1
    `, params.NikPelapor, params.NamaPelapor, params.NikPeserta)

	if err != nil {
		exception.PanicIfNeeded(err)
	}

	if len(resultKodeKlaimInduk) > 0 {
		KodeKlaimInduk = string(resultKodeKlaimInduk[0]["KODE_KLAIM_INDUK"])
		KodeKantorPengajuan = string(resultKodeKlaimInduk[0]["KODE_KANTOR"])
		KodeManfaat = string(resultKodeKlaimInduk[0]["KODE_MANFAAT"])
	}

	params.ChId = KodeKlaimInduk

	query := `insert into BPJSTKU.ASIK_KONFIRMASI(
                                                    kode_pengajuan,
                                                    kode_klaim_induk,
                                                    kode_kantor_pengajuan,
                                                    kode_tipe_manfaat,
                                                    kode_pointer_asal,
                                                    id_pointer_asal,
                                                    petugas_rekam,
                                                    tgl_rekam,
                                                    nik_pelapor,
                                                    nama_pelapor,
                                                    nik_peserta,
                                                    tgl_lahir_pelapor,
                                                    email_pelapor,
                                                    handphone_pelapor,
                                                    tgl_pengajuan,
                                                    kode_billing,
                                                    skor_face,
                                                    kemiripan_nama_pelapor,
                                                    keterangan_approval,
                                                    kanal_pelayanan
                                                    ) values (
                                                    ?, --kode_pengajuan
                                                    ?, --KodeKlaimInduk
                                                    ?, --KodeKantorPengajuan
                                                    ?, --kode_tipe_manfaat
                                                    ?,--kode_pointer_asal
                                                    ?,--id_pointer_asal
                                                    ?,--petugas_rekam
                                                    sysdate,--tgl_rekam 
                                                    ?, --nik_pelapor
                                                    ?, --nama_pelapor
                                                    ?,--nik_peserta
                                                    to_date(?,'dd-mm-yyyy'),--tgl_lahir_pelapor
                                                    ?,--email_pelapor
                                                    ?,--handphone_pelapor
                                                    to_date(?,'dd-mm-yyyy'),--tgl_pengajuan
                                                    ?,--kode_billing
                                                    ?,--skor_face
                                                    ?,--kemiripan_nama_pelapor
                                                    ?,--keterangan_approval
                                                    ?--kanal_pelayanan
                                                    )
                    `

	if _, err = session.Exec(
		query,
		GetKodePengajuan,
		KodeKlaimInduk,
		KodeKantorPengajuan,
		"F002",
		"PTPOS",
		GetKodePengajuan,
		params.PetugasRekamPTPOS,
		//time.Now().Format("01-02-2006"),
		params.NikPelapor,
		params.NamaPelapor,
		params.NikPeserta,
		params.TglLahirPelapor,
		params.EmailPelapor,
		params.HandphonePelapor,
		params.TglPengajuan,
		params.KodeBilling,
		params.SkorFace,
		params.KemiripanNamaPelapor,
		params.KeteranganApproval,
		"43",
	); err != nil {
		session.Rollback()
		return nil, err
	}

	for _, PenerimaBeasiswa := range params.PenerimaBeasiswa {

		//hitung manfaat beasiswa

		KodeManfaatInt, _ := strconv.Atoi(KodeManfaat)
		NoUrutInt, _ := strconv.Atoi(PenerimaBeasiswa.NoUrutPenerima)

		ParamHitungManfaatBeasiswa := &entity.HitungManfaatBeasiswaEntity{
			ChId:              "PTPOS",
			ReqId:             "PTPOS",
			KodeManfaat:       KodeManfaatInt,
			NikPenerima:       params.NikPeserta,
			KodeKlaim:         KodeKlaimInduk,
			NoUrut:            NoUrutInt,
			Tahun:             PenerimaBeasiswa.TahunBeasiswa,
			BeasiswaJenis:     PenerimaBeasiswa.KodeJenisBeasiswa,
			JenjangPendidikan: PenerimaBeasiswa.JenjangPendidikan,
		}

		jsonData, err := json.Marshal(ParamHitungManfaatBeasiswa)
		if err != nil {
			fmt.Println(err)
		}

		payload := strings.NewReader(string(jsonData))

		// fmt.Println(jsonData)
		fmt.Println(payload)

		req, err := http.NewRequest("POST", configuration.Get("JSPN5041_URL")+"/JSPN5041/HitungMnfBeasiswapp82", payload)

		if err != nil {
			fmt.Println(err)
		}

		HitungManfaatBeasiswaResponse := &response.HitungManfaatBeasiswaResponse{}
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		json.NewDecoder(res.Body).Decode(&HitungManfaatBeasiswaResponse)
		defer res.Body.Close()

		fmt.Println(HitungManfaatBeasiswaResponse.PNomDisetujui)

		//end hitung manfaaat

		querybeasiswadetil := `insert into bpjstku.asik_konfirmasi_beasiswa_detil(
            kode_pengajuan,
            nik_penerima,
            nama_penerima,
            no_urut,
            tgl_lahir_penerima,
            flag_masih_sekolah,
            tahun,
            jenis,
            tingkat,
            jenjang,
            lembaga,
            nom_manfaat,
            flag_dok_lengkap,
            petugas_rekam,
            tgl_rekam
        )values(
            ?, --kode_pengajuan
            ?, --nik_penerima
            ?, --nama_penerima
            ?, --no_urut
            to_date(?,'dd-mm-yyyy'),--tgl_lahir_penerima
            ?,--flag_masih_sekolah
            ?,--tahun
            ?,--jenis
            ?,--tingkat
            ?,--jenjang
            ?,--lembaga
            ?,--nom_manfaat
            ?,--flag_dok_lengkap
            ?,--petugas_rekam
            sysdate--tgl_--tgl_rekam
        )
        `
		if _, err = session.Exec(
			querybeasiswadetil,
			GetKodePengajuan,
			PenerimaBeasiswa.NikPenerimaBeasiswa,
			PenerimaBeasiswa.NamaPenerimaBeasiswa,
			PenerimaBeasiswa.NoUrutPenerima,
			PenerimaBeasiswa.TglLahirPenerimaBeasiswa,
			PenerimaBeasiswa.FlagMasihSekolah,
			PenerimaBeasiswa.TahunBeasiswa,
			PenerimaBeasiswa.KodeJenisBeasiswa,
			PenerimaBeasiswa.TingkatPendidikan,
			PenerimaBeasiswa.JenjangPendidikan,
			PenerimaBeasiswa.LembagaPendidikan,
			//1000000,
			HitungManfaatBeasiswaResponse.PNomDisetujui,
			PenerimaBeasiswa.FlagDokLengkap,
			params.PetugasRekamPTPOS,
			//time.Now().Format("01-02-2006"),
		); err != nil {
			session.Rollback()
			return nil, err
		}

		for _, DataDokumen := range PenerimaBeasiswa.DataDokumenBeasiswa {
			querydokumenbeasiswa := `insert into BPJSTKU.ASIK_KONFIRMASI_BEASISWA_DOK (
                kode_pengajuan,
                nama_dokumen,
                path_url,
                tgl_rekam,
                petugas_rekam
                )
                values
                (
                ?, --kode_pengajuan
                ?, --kode_dokumen
                ?, --path_url
                sysdate,--tgl_rekam
                ?--petugas_rekam
            )
        `
			if _, err = session.Exec(
				querydokumenbeasiswa,
				GetKodePengajuan,
				DataDokumen.KodeDokumen,
				DataDokumen.PathUrl,
				//time.Now().Format("01-02-2006"),
				"PETUGASPOS",
			); err != nil {
				session.Rollback()
				return nil, err
			}
		}
	}

	for _, DataDokumen := range params.DataDokumenPelapor {
		querydokumen := `insert into BPJSTKU.ASIK_KONFIRMASI_DOKUMEN (
                kode_pengajuan,
                kode_dokumen,
                path_url,
                tgl_rekam,
                petugas_rekam
                )
                values
                (
                ?, --kode_pengajuan
                ?, --kode_dokumen
                ?, --path_url
                sysdate,--tgl_rekam
                ?--petugas_rekam
            )
        `
		if _, err = session.Exec(
			querydokumen,
			GetKodePengajuan,
			DataDokumen.KodeDokumen,
			DataDokumen.PathUrl,
			"PETUGASPOS",
		); err != nil {
			session.Rollback()
			return nil, err
		}
	}

	for _, dataProbing := range params.DataProbingInsertKonfBeasiswa {
		queryprobing := `insert into bpjstku.asik_probing
        (
            kode_pengajuan,
            kode_probing,
            no_urut,
            respon_probing,
            jawaban_probing,
            tgl_rekam,
            petugas_rekam
        )values(
            ?, --kode_pengajuan
            ?, --kode_probing
            ?, --no_urut
            ?,--respon_probing
            ?,--jawaban_probing
            sysdate,--tgl_rekam
            ?--petugas_rekam
        )
    `
		if _, err = session.Exec(
			queryprobing,
			GetKodePengajuan,
			dataProbing.KodeProbing,
			dataProbing.NoUrut,
			dataProbing.ResponProbing,
			dataProbing.JawabanProbing,
			//time.Now().Format("01-02-2006"),
			"PETUGASPOS",
		); err != nil {
			session.Rollback()
			return nil, err
		}
	}

	err = session.Commit()
	if err != nil {
		fmt.Println(err)
	}

	//send Email
	// qryEmail := `select * from kit.vas_email_content where vas_email_content_id=?`
	// resultEmail, err := db.EngineOltp.Query(qryEmail, "54")
	// if err != nil {
	// 	return nil, err
	// }

	// var tgl_rekam string

	// var qry = `select to_char(tgl_rekam,'DD-MM-YYYY HH24:mm:ss')||' WIB' tgl_rekam from bpjstku.asik_konfirmasi where kode_pengajuan='` + GetKodePengajuan + `'`
	// if _, err := db.EngineEcha.SQL(qry).Get(&tgl_rekam); err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(tgl_rekam)
	// var kpj string
	// var qryKpj = `
	// select
	// kpj
	// from
	// pn.pn_klaim
	// where kode_klaim = '` + KodeKlaimInduk + `'
	// `
	// if _, err := db.EngineOltp.SQL(qryKpj).Get(&kpj); err != nil {
	// 	fmt.Println(err)
	// }

	// stringSubject := strings.Replace(string(resultEmail[0]["SUBJECT"]), ":0:", kpj, 1)
	// kanalLayanan := strings.Split(params.KodeKantorPTPOS, "|")

	// string1 := strings.Replace(string(resultEmail[0]["HTML_CONTENT"]), ":1:", strings.ToUpper(params.NamaPelapor), 1)
	// string2 := strings.Replace(string1, ":2:", strings.ToUpper(kanalLayanan[1]), 1)
	// string3 := strings.Replace(string2, ":3:", GetKodePengajuan, 1)
	// string4 := strings.Replace(string3, ":4:", tgl_rekam, 1)
	// //string5 := strings.Replace(string4, ":5:", "PENGAJUAN", 1)

	// paramSendEmail := &request.SendEmailRequest{
	// 	Subject: stringSubject,
	// 	Body:    stringSubject,
	// 	Message: string4,
	// 	Email:   params.EmailPelapor,
	// }

	// jsonDataEmail, err := json.Marshal(paramSendEmail)
	// if err != nil {
	// 	return nil, err
	// }

	// payload := strings.NewReader(string(jsonDataEmail))
	// req, err := http.NewRequest("POST", configuration.Get("JS_PTPOS_URL")+"/JSPTPOS/SendEmail", payload)
	// if err != nil {
	// 	return nil, err
	// }

	// sendEmailResponse := &response.SendEmailResponse{}
	// req.Header.Add("Content-Type", "application/json")
	// res, err := client.Do(req)
	// if err != nil {
	// 	return nil, err
	// }
	// json.NewDecoder(res.Body).Decode(&sendEmailResponse)
	// defer res.Body.Close()
	//end of send email

	// DatainsertKonfirmationJPBerkalaRes := []response.DataInsertKonfirmasiJP{{
	//  StatusKirim:                      "SUKSES",
	//  Keterangan:                       "SUKSES",
	//  KodePengajuanKonfirmasiJPBerkala: kodePengajuan,
	// }}

	// insertKonfirmationJPBerkalaRes := &response.InsertKonfirmasiJPResponse{
	//  StatusCode: "200",
	//  StatusDesc: "Berhasil Submit Data",
	//  Data:       DatainsertKonfirmationJPBerkalaRes,
	// }

	// if res.StatusCode == 200 {
	//  return insertKonfirmationJPBerkalaRes, nil
	// } else {
	//  return nil, errors.New(sendEmailResponse.Message)
	// }

	err = session.Commit()
	if err != nil {
		fmt.Println(err)
	}

	var insertKonfirmasiBeasiswa = []response.InsertKonfirmasiBeasiswa{{
		KodePengajuanKonfirmasi: GetKodePengajuan,
		StatusKirim:             "SUKSES",
		KeteranganSukses:        "rest",
	}}

	resultinsertKonfirmasiBeasiswa := &response.InsertKonfirmasiBeasiswaResponse{
		StatusCode:               "200",
		StatusDesc:               "Data Ditemukan",
		InsertKonfirmasiBeasiswa: insertKonfirmasiBeasiswa,
	}

	return resultinsertKonfirmasiBeasiswa, nil
	// return GetKodePengajuan, err

}

func (repository *BeasiswaRepositoryImpl) DaftarJenisBeasiswa(params *request.DaftarJenisBeasiswaRequest) (*response.DaftarJenisBeasiswaResponse, error) {
	//var DaftarJenisBeasiswaHasil entity.DaftarJenisBeasiswaEntity

	var qry = `
            select kode kode_jenis_beasiswa, keterangan nama_jenis_beasiswa from ms.ms_lookup where tipe= 'KLMJNSBEAS'  
    `
	result, err := db.EngineOltp.Query(qry)

	if err != nil {
		return nil, err
	}

	JenisBeasiswas := make([]response.DataJenisBeasiswa, 0)
	for _, item := range result {
		JenisBeasiswa := response.DataJenisBeasiswa{
			KodeJenisBeasiswa: string(item["KODE_JENIS_BEASISWA"]),
			NamaJenisBeasiswa: string(item["NAMA_JENIS_BEASISWA"]),
		}
		JenisBeasiswas = append(JenisBeasiswas, JenisBeasiswa)
	}

	resultDaftarJenisBeasiswa := &response.DaftarJenisBeasiswaResponse{
		StatusCode:        "200",
		StatusDesc:        "Data Ditemukan",
		DataJenisBeasiswa: JenisBeasiswas,
	}

	// objects := make([]entity.DaftarJenisBeasiswaEntity, 0)
	// for _, item := range result {
	//  object := entity.DaftarJenisBeasiswaEntity{
	//      KodeJenisBeasiswa: string(item["KODE_JENIS_BEASISWA"]),
	//      NamaJenisBeasiswa: string(item["NAMA_JENIS_BEASISWA"]),
	//  }
	//  objects = append(objects, object)
	// }

	return resultDaftarJenisBeasiswa, nil
}

func (repository *BeasiswaRepositoryImpl) DaftarJenjangPendidikan(params *request.DaftarJenjangPendidikanRequest) ([]entity.DaftarJenjangPendidikanEntity, error) {
	//var DaftarJenisBeasiswaHasil entity.DaftarJenisBeasiswaEntity

	session := db.EngineOltp.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return nil, err
	}

	var qry = `
    select 
    (
        select
        distinct c.jenis
        from pn.pn_klaim a
        inner join pn.pn_klaim_penerima_manfaat b on b.kode_klaim = a.kode_klaim
        inner join pn.pn_klaim_manfaat_detil_beasis c on c.kode_klaim = b.kode_klaim
        where
        -- a.kode_klaim = 'KL21041801209352'
        -- and rownum = 1   
         b.nomor_identitas = ?  
         --and b.nama_penerima = ?
         and UTL_MATCH.EDIT_DISTANCE_SIMILARITY (upper(B.NAMA_PENERIMA), UPPER(?)) > 75 
         and a.nomor_identitas = ?
         and  nvl(status_batal,'X') = 'T' 
         and rownum = 1 
    )jenis_beasiswa,
 kode,   
 keterangan
 from
 ms.ms_lookup where tipe = 'TKSKLHPP82'
 and seq >= (
              select
              seq
              from
              (      
                    select
                    y.*,
                    --y.seq,
                    rank() over (order by y.seq desc) rankseq
                    from
                    (
                        select
                        a.kode_klaim,
                        a.kode_klaim_induk,
                        a.tgl_klaim,
                        a.tgl_ubah,
                        b.nomor_identitas nik_pelapor,
                        b.nama_penerima nama_pelapor,
                        a.nomor_identitas nik_peserta,
                        c.jenis,
                        c.jenjang,
                        (select seq from ms.ms_lookup where kode = c.jenjang and tipe = 'TKSKLHPP82') seq,
                        rank() over (order by a.tgl_klaim,a.tgl_ubah desc) ranktglubah
                        from pn.pn_klaim a 
                        inner join pn.pn_klaim_penerima_manfaat b on b.kode_klaim = a.kode_klaim    
                        inner join pn.pn_klaim_manfaat_detil_beasis c on c.kode_klaim = b.kode_klaim
                        where
                            a.kode_klaim in 
                                        (select kode_klaim from pn.pn_klaim           
                                            start with kode_klaim = (
                                                                        select
                                                                        distinct a.kode_klaim
                                                                        from pn.pn_klaim a
                                                                        inner join pn.pn_klaim_penerima_manfaat b on b.kode_klaim = a.kode_klaim
                                                                        inner join pn.pn_klaim_manfaat_detil_beasis c on c.kode_klaim = b.kode_klaim
                                                                        where
                                                                        -- a.kode_klaim = 'KL21041801209352'
                                                                        -- and rownum = 1
                                                                         b.nomor_identitas = ?  
                                                                         --and b.nama_penerima = ?
                                                                         and UTL_MATCH.EDIT_DISTANCE_SIMILARITY (upper(B.NAMA_PENERIMA), UPPER(?)) > 75 
                                                                         and a.nomor_identitas = ?
                                                                         and  nvl(status_batal,'X') = 'T'
                                                                         and rownum = 1 --nanti akan di prior lagi  
                                                                    )                                                 
                                                connect by prior kode_klaim = kode_klaim_induk
                                        )
                    )y     
                    where ranktglubah = 1    
              )x
              where rankseq = 1   
              and rownum = 1   
  )     
    `
	result, err := db.EngineOltp.Query(qry,
		params.NikPelapor,
		params.NamaPelapor,
		params.NikPeserta,
		params.NikPelapor,
		params.NamaPelapor,
		params.NikPeserta,
	)

	if err != nil {
		return nil, err
	}

	objects := make([]entity.DaftarJenjangPendidikanEntity, 0)
	for _, item := range result {
		object := entity.DaftarJenjangPendidikanEntity{
			KodeJenisBeasiswa:     string(item["JENIS_BEASISWA"]),
			KodeJenjangPendidikan: string(item["KODE"]),
			NamaJenjangPendidikan: string(item["KETERANGAN"]),
		}
		objects = append(objects, object)
	}

	return objects, nil

}

func (repository *BeasiswaRepositoryImpl) DaftarPenerimaBeasiswa(params *request.DaftarPenerimaBeasiswaRequest) (*response.DaftarPenerimaBeasiswaResponse, error) {
	//var DaftarJenisBeasiswaHasil entity.DaftarJenisBeasiswaEntity

	session := db.EngineOltp.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return nil, err
	}

	var qry = `
    select
    distinct
    kode_klaim,
    no_urut,
    nik_penerima,
    NAMA_PENERIMA,
    TGL_LAHIR
    from
    (
        select
        a.kode_klaim,
        c.no_urut,
        --a.kode_klaim_induk,
        --a.tgl_klaim,
        --a.tgl_ubah,
        --b.nomor_identitas nik_pelapor,
        --b.nama_penerima nama_pelapor,
        --a.nomor_identitas nik_peserta,
        e.nik_penerima,
        E.NAMA_PENERIMA,
        E.TGL_LAHIR,
        --(select nama_hubungan from kn.kn_kode_hubungan_tk where kode_hubungan = b.kode_hubungan) nama_hubungan,
        --c.jenis,
        --c.jenjang,
        rank() over (order by a.tgl_klaim,a.tgl_ubah desc) ranktglubah
        from pn.pn_klaim a 
        inner join pn.pn_klaim_penerima_manfaat b on b.kode_klaim = a.kode_klaim    
        inner join pn.pn_klaim_manfaat_detil_beasis c on c.kode_klaim = b.kode_klaim and c.no_urut = b.no_urut_keluarga
        inner join pn.pn_klaim_manfaat_detil d on d.kode_klaim = c.kode_klaim and d.no_urut = c.no_urut
        inner join pn.pn_penerima_beasiswa e on e.nik_penerima = d.beasiswa_nik_penerima
        where
            a.kode_klaim in 
                        (select kode_klaim from pn.pn_klaim           
                            start with kode_klaim = (
                                                        select 
                                                        distinct a.kode_klaim
                                                        from pn.pn_klaim a
                                                        inner join pn.pn_klaim_penerima_manfaat b on b.kode_klaim = a.kode_klaim
                                                        inner join pn.pn_klaim_manfaat_detil_beasis c on c.kode_klaim = b.kode_klaim
                                                        inner join pn.pn_klaim_manfaat_detil d on d.kode_klaim = c.kode_klaim
                                                        inner join pn.pn_penerima_beasiswa e on e.nik_penerima = d.beasiswa_nik_penerima
                                                        where
                                                        --a.kode_klaim = 'KL21050601375865'
                                                        --and rownum = 1                                                        
                                                            b.nomor_identitas = ?--:p_nik_pelapor 3401121108850002
                                                            --and b.nama_penerima = ?--:p_nama_pelapor AGUS RIYANTA
                                                            and UTL_MATCH.EDIT_DISTANCE_SIMILARITY (upper(B.NAMA_PENERIMA), UPPER(?)) > 75 
                                                            and a.nomor_identitas = ?--:p_nik_peserta 3401121108850002
                                                            and nvl(a.status_batal,'X') = 'T'
                                                            and rownum = 1 --nanti akan di prior lagi 
                                                    )                                                 
                                connect by prior kode_klaim = kode_klaim_induk
                        )
    )y     
    where ranktglubah = 1
    order by no_urut asc
    `
	result1, err := db.EngineOltp.Query(qry,
		params.NikPelapor,
		params.NamaPelapor,
		params.NikPeserta,
	)

	if err != nil {
		return nil, err
	}

	objects1 := make([]response.DaftarPenerimaBeasiswa, 0)
	for _, item := range result1 {

		kodeklaim := string(item["KODE_KLAIM"])
		object := response.DaftarPenerimaBeasiswa{
			//KodeKlaim:                    string(item["KODE_KLAIM"]),
			NoUrut:                       string(item["NO_URUT"]),
			NikPenerimaBeasiswa:          string(item["NIK_PENERIMA"]),
			NamaPenerimaBeasiswa:         string(item["NAMA_PENERIMA"]),
			TanggalLahirPenerimaBeasiswa: string(item["TGL_LAHIR"]),
			//HubunganPenerimaBeasiswa:     string(item["NAMA_HUBUNGAN"]),
			//JenisBeasiswa:             string(item["JENIS"]),
			//JenjangPendidikanTerakhir: string(item["JENJANG"]),
		}

		// var qrytahunpenerima = `  select
		// distinct
		// nik_penerima,
		// NAMA_PENERIMA,
		// TGL_LAHIR,
		//    tahun,
		//  jenis,
		//  jenjang
		// from
		// (
		//  select
		//  --a.kode_klaim,
		//  --a.kode_klaim_induk,
		//  --a.tgl_klaim,
		//  --a.tgl_ubah,
		//  --b.nomor_identitas nik_pelapor,
		//  --b.nama_penerima nama_pelapor,
		//  --a.nomor_identitas nik_peserta,
		//  e.nik_penerima,
		//  E.NAMA_PENERIMA,
		//  E.TGL_LAHIR,
		//  --(select nama_hubungan from kn.kn_kode_hubungan_tk where kode_hubungan = b.kode_hubungan) nama_hubungan,
		//  c.tahun,
		//  c.jenis,
		//  c.jenjang,
		//  rank() over (order by a.tgl_klaim,a.tgl_ubah desc) ranktglubah
		//  from pn.pn_klaim a
		//  inner join pn.pn_klaim_penerima_manfaat b on b.kode_klaim = a.kode_klaim
		//  inner join pn.pn_klaim_manfaat_detil_beasis c on c.kode_klaim = b.kode_klaim
		//  inner join pn.pn_klaim_manfaat_detil d on d.kode_klaim = c.kode_klaim
		//  inner join pn.pn_penerima_beasiswa e on e.nik_penerima = d.beasiswa_nik_penerima
		//  where
		//      a.kode_klaim in
		//                  (select kode_klaim from pn.pn_klaim
		//                      start with kode_klaim = (
		//                                                  select
		//                                                  distinct a.kode_klaim
		//                                                  from pn.pn_klaim a
		//                                                  inner join pn.pn_klaim_penerima_manfaat b on b.kode_klaim = a.kode_klaim
		//                                                  inner join pn.pn_klaim_manfaat_detil_beasis c on c.kode_klaim = b.kode_klaim
		//                                                  inner join pn.pn_klaim_manfaat_detil d on d.kode_klaim = c.kode_klaim
		//                                                  inner join pn.pn_penerima_beasiswa e on e.nik_penerima = d.beasiswa_nik_penerima
		//                                                  where
		//                                                  --a.kode_klaim = 'KL21050601375865'
		//                                                  --and rownum = 1
		//                                                      b.nomor_identitas = ?--:p_nik_pelapor 3401121108850002
		//                                                      --and b.nama_penerima = ?--:p_nama_pelapor AGUS RIYANTA
		//                                                      and UTL_MATCH.EDIT_DISTANCE_SIMILARITY (upper(B.NAMA_PENERIMA), UPPER(?)) > 75
		//                                                      and a.nomor_identitas = ?--:p_nik_peserta 3401121108850002
		//                                                      and nvl(a.status_batal,'X') = 'T'
		//                                                      and rownum = 1 --nanti akan di prior lagi
		//                                              )
		//                          connect by prior kode_klaim = kode_klaim_induk
		//                  )
		// )y
		// where ranktglubah = 1
		// `

		var qrytahunpenerima = `select tahun, jenis,jenjang from  pn.pn_klaim_manfaat_detil_beasis where kode_klaim = ? and no_urut = ?`

		resulttahunpenerima, err := db.EngineOltp.Query(qrytahunpenerima,
			kodeklaim,
			object.NoUrut,
		)

		if err != nil {
			return nil, err
		}

		tahuns := make([]response.DataTahunPenerimaBeasiswa, 0)
		for _, item := range resulttahunpenerima {
			tahun := response.DataTahunPenerimaBeasiswa{
				TahunBeasiswa:             string(item["TAHUN"]),
				JenisBeasiswa:             string(item["JENIS"]),
				JenjangPendidikanTerakhir: string(item["JENJANG"]),
			}

			result2, err := db.EngineEcha.Query(`select kode, keterangan from BPJSTKU.ASIK_MS_LOOKUP 
            where TIPE = decode(?, 'BEASISWA','KODE_DOK_BEA_PENDIDIKAN','PELATIHAN','KODE_DOK_BEA_PELATIHAN','')
            `, tahun.JenisBeasiswa)

			if err != nil {
				exception.PanicIfNeeded(err)
			}

			objects2 := make([]response.DataDokumenPenerimaBeasiswa, 0)
			for _, item := range result2 {
				data := response.DataDokumenPenerimaBeasiswa{
					KodeDokumen: string(item["KODE"]),
					NamaDokumen: string(item["KETERANGAN"]),
				}
				objects2 = append(objects2, data)
			}

			// object2 := entity.DaftarPenerimaBeasiswaEntity{
			//  DataDokumen: objects2,
			// }
			tahun.DataDokumenPenerimaBeasiswa = objects2

			tahuns = append(tahuns, tahun)

			result3, err := db.EngineOltp.Query(` select kode from MS.MS_LOOKUP 
                                                            where tipe = 'TKSKLHPP82' 
                                                                and seq >=  
                                                                    (select seq from MS.MS_LOOKUP 
                                                                        where tipe = 'TKSKLHPP82' 
                                                                        and kode = ?) 
            `, tahun.JenjangPendidikanTerakhir)

			if err != nil {
				exception.PanicIfNeeded(err)
			}

			objects3 := make([]response.DataJenjangPendidikanTerakhir, 0)
			for _, item := range result3 {
				DataJenjangPendidikanTerakhir := response.DataJenjangPendidikanTerakhir{
					Keterangan: string(item["KODE"]),
				}

				result4, err := db.EngineEcha.Query(` select keterangan from ANTRIAN.ATR_MS_LOOKUP 
                                                                            WHERE   trim(substr(TIPE,12,6)) = ?
                                                                            order by seq asc
                    `, DataJenjangPendidikanTerakhir.Keterangan)

				if err != nil {
					exception.PanicIfNeeded(err)
				}

				objects4 := make([]response.DataTingkatPendidikanTerakhir, 0)
				for _, item := range result4 {
					data := response.DataTingkatPendidikanTerakhir{
						Keterangan: string(item["KETERANGAN"]),
					}

					objects4 = append(objects4, data)
				}

				DataJenjangPendidikanTerakhir.DataTingkatPendidikanTerakhir = objects4

				objects3 = append(objects3, DataJenjangPendidikanTerakhir)
			}

			object.DataJenjangPendidikanTerakhir = objects3
			object.DataTahunPenerimaBeasiswa = tahuns
		}

		objects1 = append(objects1, object)

	}

	resultDaftarPenerimaBeasiswa := &response.DaftarPenerimaBeasiswaResponse{
		StatusCode:             "200",
		StatusDesc:             "Data Ditemukan",
		DaftarPenerimaBeasiswa: objects1,
	}

	return resultDaftarPenerimaBeasiswa, nil

}

func (repository *BeasiswaRepositoryImpl) NominalPerjenjangBeasiswa(params *request.NominalPerjenjangBeasiswaRequest) (*response.NominalPerjenjangBeasiswaResponse, error) {
	//var DaftarJenisBeasiswaHasil entity.DaftarJenisBeasiswaEntity

	var KodeKlaim string
	var NoUrut string
	var KodeManfaat string
	var KdPrg string
	var JenisBeasiswa string
	var JenjangPendidikan string
	var p_nom_disetujui string
	var p_sukses string
	var p_mess string

	var qry = `
    select
    *
    from
    (
        select
        a.kode_klaim,
        a.kode_klaim_induk,
        a.tgl_klaim,
        a.tgl_ubah,
        b.nomor_identitas nik_pelapor,
        b.nama_penerima nama_pelapor,
        a.nomor_identitas nik_peserta,
        e.nik_penerima,
        E.NAMA_PENERIMA,
        E.TGL_LAHIR,
        --(select nama_hubungan from kn.kn_kode_hubungan_tk where kode_hubungan = b.kode_hubungan) nama_hubungan,
        d.no_urut,
        d.kode_manfaat,
        d.kd_prg,
        c.jenis,
        c.jenjang,
        rank() over (order by a.tgl_klaim,a.tgl_ubah desc) ranktglubah
        from pn.pn_klaim a 
        inner join pn.pn_klaim_penerima_manfaat b on b.kode_klaim = a.kode_klaim    
        inner join pn.pn_klaim_manfaat_detil_beasis c on c.kode_klaim = b.kode_klaim
        inner join pn.pn_klaim_manfaat_detil d on d.kode_klaim = c.kode_klaim
        inner join pn.pn_penerima_beasiswa e on e.nik_penerima = d.beasiswa_nik_penerima
        where
            a.kode_klaim in 
                        (select kode_klaim from pn.pn_klaim           
                            start with kode_klaim = (
                                                        select 
                                                        distinct
                                                        a.kode_klaim
                                                        from pn.pn_klaim a
                                                        inner join pn.pn_klaim_penerima_manfaat b on b.kode_klaim = a.kode_klaim
                                                        inner join pn.pn_klaim_manfaat_detil_beasis c on c.kode_klaim = b.kode_klaim
                                                        inner join pn.pn_klaim_manfaat_detil d on d.kode_klaim = c.kode_klaim
                                                        inner join pn.pn_penerima_beasiswa e on e.nik_penerima = d.beasiswa_nik_penerima
                                                        where
                                                        --a.kode_klaim = 'KL21050601375865'
                                                        --and rownum = 1                                                        
                                                            b.nomor_identitas = ?--:p_nik_pelapor 3401121108850002
                                                            --and b.nama_penerima = ?--:p_nama_pelapor AGUS RIYANTA
                                                            and UTL_MATCH.EDIT_DISTANCE_SIMILARITY (upper(B.NAMA_PENERIMA), UPPER(?)) > 75 
                                                            and a.nomor_identitas = ?--:p_nik_peserta 3401121108850002
                                                            and nvl(a.status_batal,'X') = 'T'
                                                            and rownum = 1 --nanti akan di prior lagi 
                                                    )                                                 
                                connect by prior kode_klaim = kode_klaim_induk
                        )
    )y     
    where ranktglubah = 1
    `

	result, err := db.EngineOltp.Query(qry, params.NikPelapor, params.NamaPelapor, params.NikPeserta)
	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		KodeKlaim = string(result[0]["KODE_KLAIM"])
		NoUrut = string(result[0]["NO_URUT"])
		KodeManfaat = string(result[0]["KODE_MANFAAT"])
		KdPrg = string(result[0]["KD_PRG"])
		JenisBeasiswa = string(result[0]["JENIS"])
		JenjangPendidikan = string(result[0]["JENJANG"])
	}

	var QryManfaatBeas = `  

    begin
    pn.p_pn_pn5040.x_hitung_mnf_beasiswa(?,
                                                     ?,
                                                     ?,
                                                     ?,
                                                     ?,
                                                     ?,                                                                                                                                                                                                   
                                                     :p_nom_disetujui,
                                                     :p_sukses,
                                                     :p_mess                                                                 
                                                     );             
 end;
 
    `

	if _, err := db.EngineOltp.Exec(QryManfaatBeas,
		KodeKlaim,
		NoUrut,
		KodeManfaat,
		KdPrg,
		JenisBeasiswa,
		JenjangPendidikan,
		sql.Named("p_nom_disetujui", sql.Out{Dest: &p_nom_disetujui}),
		sql.Named("p_sukses", sql.Out{Dest: &p_sukses}),
		sql.Named("p_mess", sql.Out{Dest: &p_mess})); err != nil {
		return nil, err
	}

	var result2 = []response.NominalPerjenjangBeasiswa{{
		NominalManfaat: p_nom_disetujui,
		Keterangan:     p_mess,
	}}

	resultNominalPerjenjangBeasiswa := &response.NominalPerjenjangBeasiswaResponse{
		StatusCode:                "200",
		StatusDesc:                "Data Ditemukan",
		NominalPerjenjangBeasiswa: result2,
	}

	return resultNominalPerjenjangBeasiswa, nil

	// NominalPerjenjangBeasiswaEntity = entity.NominalPerjenjangBeasiswaEntity{
	//  NominalManfaat: string(result[0]["NOMINAL"]),
	//  Keterangan:     string(result[0]["KETERANGAN"]),
	// }

	// return &NominalPerjenjangBeasiswaEntity, nil
}

// func (repository *BeasiswaRepositoryImpl) SendEmail(params *request.SendEmailRequest) (*response.SendEmailResponse, error) {
//  configuration := config.New()
//  client := &http.Client{}

//  //send Email
//  qryEmail := `select * from kit.vas_email_content where vas_email_content_id=?`
//  resultEmail, err := db.EngineOltp.Query(qryEmail, "54")
//  if err != nil {
//      return nil, err
//  }

//  ContentEmail := string(resultEmail[0]["HTML_CONTENT"])
//  string1 := strings.Replace(ContentEmail, ":0:", params.OfficeCode, 1)
//  string2 := strings.Replace(string1, ":1:", params.FullName, 1)
//  string3 := strings.Replace(string2, ":2:", params.KodeKantorPPTOS, 1)
//  string4 := strings.Replace(string3, ":3:", params.KodePengajuan, 1)
//  string5 := strings.Replace(string4, ":4:", params.TanggalPengajuan, 1)
//  string6 := strings.Replace(string5, ":5:", "PENGAJUAN", 1)

//  SubjectContent := string(resultEmail[0]["SUBJECT"])
//  SubjectContent1 := strings.Replace(SubjectContent, ":0:", params.Kpj, 1)

//  paramSendEmail := &request.SendEmailRequest{
//      Subject: SubjectContent1,
//      Body:    SubjectContent1,
//      Message: string6,
//      Email:   params.Email,
//  }

//  jsonDataEmail, err := json.Marshal(paramSendEmail)
//  if err != nil {
//      return nil, err
//  }

//  payload := strings.NewReader(string(jsonDataEmail))
//  req, err := http.NewRequest("POST", configuration.Get("JS_PTPOS_URL")+"/JSPTPOS/SendEmail", payload)
//  if err != nil {
//      return nil, err
//  }

//  sendEmailResponse := &response.SendEmailResponse{}
//  req.Header.Add("Content-Type", "application/json")
//  res, err := client.Do(req)
//  if err != nil {
//      return nil, err
//  }
//  json.NewDecoder(res.Body).Decode(&sendEmailResponse)
//  defer res.Body.Close()
//  //end of send email

//  if res.StatusCode == 200 {
//      return sendEmailResponse, nil
//  } else {
//      return nil, errors.New(sendEmailResponse.Message)
//  }
// }
