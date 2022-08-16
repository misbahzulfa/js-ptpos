package repository

import (
	// "database/sql"
	// "encoding/json"
	// "errors"

	"database/sql"
	"fmt"
	"js-ptpos/exception"

	// "js-ptpos/model"

	// "js-ptpos/model"

	// "js-ptpos/config"
	db "js-ptpos/config"
	"js-ptpos/entity"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
	// "js-ptpos/model/response"
	// "net/http"
	// "strings"
)

func NewJHTClaimRepository() ClaimJHTRepository {
	return &claimJHTRepositoryImpl{}
}

type claimJHTRepositoryImpl struct {
}

func (repository *claimJHTRepositoryImpl) GenKodeKlaim() string {
	var KodeKlaim string

	var qry = `select bpjstku.p_bpjstku_asik_klaim.f_gen_kode_pengajuan from dual`
	if _, err := db.EngineEcha.SQL(qry).Get(&KodeKlaim); err != nil {
		fmt.Println(err)
	}

	return KodeKlaim
}
func (repository *claimJHTRepositoryImpl) CheckTotalKodeTK(kpj string) int {
	var total int

	var qry = `
	select count(kode_tk) cnt_kode_tk 
	from (
		select count(a.kode_tk),a.kode_tk
		from kn.vw_kn_tk a where kpj =?
        group by a.kode_tk
	)xx 
	group by xx.kode_tk
	`
	if _, err := db.EngineOltp.SQL(qry, kpj).Get(&total); err != nil {
		fmt.Println(err)
	}

	return total
}

// func (repository *claimJHTRepositoryImpl) GetKodeTK(kpj string) string {
// 	var KodeTK string

// 	var qry = `
// 	select kode_tk
// 	from (
// 		select a.kode_tk,a.nomor_identitas,a.nama_tk
// 		from kn.vw_kn_tk a where kpj =?
// 	)xx
// 	group by xx.kode_tk
// 	`
// 	if _, err := db.EngineOltp.SQL(qry, kpj).Get(&KodeTK); err != nil {
// 		fmt.Println(err)
// 	}

// 	return KodeTK
// }

func (repository *claimJHTRepositoryImpl) CheckJHTEligible(params *request.CheckEligibleJHTRequest) (*response.CheckJHTEligibleResponse, error) {
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

	var qryCnt = `
	select count(kode_tk) cnt_kode_tk 
	from (
		select count(a.kode_tk),a.kode_tk
		from kn.vw_kn_tk a where kpj =?
        group by a.kode_tk
	)xx 
	group by xx.kode_tk
	`

	resultCnt, err := db.EngineOltp.Query(qryCnt, params.Kpj)
	if err != nil {
		return nil, err
	}

	if len(resultCnt) > 0 {
		kodeTK = string(resultCnt[0]["CNT_KODE_TK"])
	}

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
		tglLahir = string(result[0]["TGL_LAHIR"])
		email = string(result[0]["EMAIL"])
	}

	// get data probing
	var probingQry = `select 
						a.KODE_PROBING,a.no_urut,
						(
							select b.KATEGORI from BPJSTKU.ASIK_KODE_PROBING b
							where b.KODE_PROBING = a.KODE_PROBING
							and rownum = 1
							) KATEGORI,
						(
						select b.NAMA_PROBING from BPJSTKU.ASIK_KODE_PROBING b
						where b.KODE_PROBING = a.KODE_PROBING
						and rownum = 1
						) NAMA_PROBING,
						a.FLAG_TAMBAHAN,
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
											(SELECT to_char(count(kpj))
												FROM (SELECT *
														FROM (SELECT npp,
																	kode_perusahaan,
																	nama_perusahaan,
																	kode_divisi,
																	kode_segmen,
																	kpj,
																	tgl_aktif,
																	kode_na,
																	NVL (tgl_na, TO_DATE ('31/12/3000', 'dd/mm/yyyy'))
																		tgl_na,
																	RANK ()
																	OVER (
																		PARTITION BY a.kpj, a.nomor_identitas, a.nama_tk
																		ORDER BY
																		NVL (a.tgl_na,
																				TO_DATE ('31/12/3000', 'dd/mm/yyyy')) DESC)
																		RANK
																FROM kn.vw_kn_tk@to_kn a
															WHERE     1 = 1
																	AND a.nomor_identitas = ?
																	AND NVL (a.kode_na, 'XXXXX') NOT IN ('AS', 'AG')
																	))
											WHERE RANK = 1 AND ROWNUM = 1)
										when a.kode_probing = 'PROBING005' then
											(SELECT nama_perusahaan
												FROM (SELECT *
														FROM (SELECT npp,
																	kode_perusahaan,
																	nama_perusahaan,
																	kode_divisi,
																	kode_segmen,
																	kpj,
																	tgl_aktif,
																	kode_na,
																	NVL (tgl_na, TO_DATE ('31/12/3000', 'dd/mm/yyyy'))
																		tgl_na,
																	RANK ()
																	OVER (
																		PARTITION BY a.kpj, a.nomor_identitas, a.nama_tk
																		ORDER BY
																		NVL (a.tgl_na,
																				TO_DATE ('31/12/3000', 'dd/mm/yyyy')) DESC)
																		RANK
																FROM kn.vw_kn_tk@to_kn a
															WHERE     1 = 1
																	AND a.nomor_identitas = ?
																	AND NVL (a.kode_na, 'XXXXX') NOT IN ('AS', 'AG')
																	))
											WHERE RANK = 1 AND ROWNUM = 1)  
										when a.kode_probing = 'PROBING006' then
											(SELECT (select alamat_perusahaan from kn.kn_perusahaan@to_kn b where b.kode_perusahaan = kd_prs)
												FROM (SELECT *
														FROM (SELECT npp,
																	kode_perusahaan kd_prs,
																	nama_perusahaan,
																	kode_divisi,
																	kode_segmen,
																	kpj,
																	tgl_aktif,
																	kode_na,
																	NVL (tgl_na, TO_DATE ('31/12/3000', 'dd/mm/yyyy'))
																		tgl_na,
																	RANK ()
																	OVER (
																		PARTITION BY a.kpj, a.nomor_identitas, a.nama_tk
																		ORDER BY
																		NVL (a.tgl_na,
																				TO_DATE ('31/12/3000', 'dd/mm/yyyy')) DESC)
																		RANK
																FROM kn.vw_kn_tk@to_kn a
															WHERE     1 = 1
																	AND a.nomor_identitas = ?
																	AND NVL (a.kode_na, 'XXXXX') NOT IN ('AS', 'AG')
																	))
											WHERE RANK = 1 AND ROWNUM = 1) 
										when a.kode_probing = 'PROBING007' then
											(SELECT to_char(tgl_aktif,'dd-mm-yyyy')
												FROM (SELECT *
														FROM (SELECT npp,
																	kode_perusahaan,
																	nama_perusahaan,
																	kode_divisi,
																	kode_segmen,
																	kpj,
																	tgl_aktif,
																	kode_na,
																	NVL (tgl_na, TO_DATE ('31/12/3000', 'dd/mm/yyyy'))
																		tgl_na,
																	RANK ()
																	OVER (
																		PARTITION BY a.kpj, a.nomor_identitas, a.nama_tk
																		ORDER BY
																		NVL (a.tgl_na,
																				TO_DATE ('31/12/3000', 'dd/mm/yyyy')) DESC)
																		RANK
																FROM kn.vw_kn_tk@to_kn a
															WHERE     1 = 1
																	AND a.nomor_identitas = ?
																	AND NVL (a.kode_na, 'XXXXX') NOT IN ('AS', 'AG')
																	))
											WHERE RANK = 1 AND ROWNUM = 1) 
										when a.kode_probing = 'PROBING008' then
											(SELECT to_char(tgl_na,'dd-mm-yyyy')
												FROM (SELECT *
														FROM (SELECT npp,
																	kode_perusahaan,
																	nama_perusahaan,
																	kode_divisi,
																	kode_segmen,
																	kpj,
																	tgl_aktif,
																	kode_na,
																	NVL (tgl_na, TO_DATE ('31/12/3000', 'dd/mm/yyyy'))
																		tgl_na,
																	RANK ()
																	OVER (
																		PARTITION BY a.kpj, a.nomor_identitas, a.nama_tk
																		ORDER BY
																		NVL (a.tgl_na,
																				TO_DATE ('31/12/3000', 'dd/mm/yyyy')) DESC)
																		RANK
																FROM kn.vw_kn_tk@to_kn a
															WHERE     1 = 1
																	AND a.nomor_identitas = ?
																	AND NVL (a.kode_na, 'XXXXX') NOT IN ('AS', 'AG')
																	))
											WHERE RANK = 1 AND ROWNUM = 1)   
										else null
									end respon_probing 
						from BPJSTKU.ASIK_KODE_PROBING_PENERIMA a
						where a.KODE_TIPE_KLAIM = 'JHT01'
						and a.KODE_TIPE_PENERIMA = 'TK'
						and nvl(a.STATUS_NONAKTIF,'X') = 'T'
						order by a.NO_URUT asc`
	// fmt.Println(qry, "2")

	probingQryRes, err := db.EngineEcha.Query(
		probingQry,
		params.Nik,
		params.Nik,
		params.Nik,
		params.Nik,
		params.Nik,
		params.Nik,
		params.Nik,
		params.Nik,
	)
	if err != nil {
		return nil, err
	}

	daftarProbings := make([]response.DataProbing, 0)
	for _, item := range probingQryRes {
		daftarProbing := response.DataProbing{
			KodeProbing:   string(item["KODE_PROBING"]),
			NoUrut:        string(item["NO_URUT"]),
			NamaProbing:   string(item["NAMA_PROBING"]) + " " + string(item["RESPON_PROBING"]),
			ResponProbing: string(item["RESPON_PROBING"]),
			Kategori:      string(item["KATEGORI"]),
		}
		daftarProbings = append(daftarProbings, daftarProbing)
	}

	// get data sebab Klaim
	var sebabKlaimQry = `select * from table (bpjstku.p_ptpos_asik_klaim.f_get_list_sebab_klaim(?,?,'JHT01',to_date(?,'dd-mm-yyyy')))`

	// fmt.Println(tglLahir, "3")

	sebabKlaimQryRes, err := db.EngineEcha.Query(
		sebabKlaimQry,
		kodeSegmen,
		kodeTK,
		tglLahir,
	)
	if err != nil {
		return nil, err
	}
	fmt.Println(sebabKlaimQry, kodeSegmen, tglLahir)
	daftarSebabKlaim := make([]response.DataSebabKlaim, 0)
	for _, item := range sebabKlaimQryRes {
		dataDaftarSebabKlaim := response.DataSebabKlaim{
			KodeSebabKlaim: string(item["KODE_SEBAB_KLAIM"]),
			NamaSebabKlaim: string(item["NAMA_SEBAB_KLAIM"]),
		}
		daftarSebabKlaim = append(daftarSebabKlaim, dataDaftarSebabKlaim)
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

	fmt.Println(kodeSegmen, params.Kpj, kodeTK, params.Nik, params.Fullname, tglLahir, email)
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
	// fmt.Println("aaa", p_param_maksimal_saldo_jht, p_klaim_sebagian, p_nama_kanal_kantor_cabang, p_pesan_tidak_layak, params.Fullname, tglLahir)
	if p_kode_pesan == "JHTA000" {
		var objectw = []response.CheckJHTEligible{{
			StatusKelayakan:     "Y",
			KodeKelayakan:       p_kode_pesan,
			KeteranganKelayakan: p_pesan_tidak_layak,
			SebabKlaim:          daftarSebabKlaim,
			DataProbingDetil:    daftarProbings}}

		resultEligible := &response.CheckJHTEligibleResponse{
			StatusCode: 200,
			StatusDesc: "Ok",
			Data:       objectw,
		}

		return resultEligible, nil
	} else {
		var objectw2 = []response.CheckJHTEligible{{
			StatusKelayakan:     "T",
			KodeKelayakan:       p_kode_pesan,
			KeteranganKelayakan: p_pesan_tidak_layak,
			SebabKlaim:          []response.DataSebabKlaim{},
			DataProbingDetil:    []response.DataProbing{},
		}}

		resultEligible := &response.CheckJHTEligibleResponse{
			StatusCode: 200,
			StatusDesc: "Ok",
			Data:       objectw2,
		}
		return resultEligible, nil

	}
}

func (repository *claimJHTRepositoryImpl) GetPengajuanJHT(params *request.GetPengajuanJHTRequest) ([]entity.GetPengajuanJHTEntity, error) {
	session := db.EngineEcha.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return nil, err
	}

	result, err := db.EngineEcha.Query(`select kode_segmen,kode_perusahaan,kode_tk,kode_divisi,kode_kepesertaan,tgl_aktif,tgl_nonaktif,kode_na,kode_sebab_klaim,kode_kantor 
	from bpjstku.asik_klaim`)

	if err != nil {
		exception.PanicIfNeeded(err)
	}

	dataSet := make([]entity.GetPengajuanJHTEntity, 0)
	for _, item := range result {
		data := entity.GetPengajuanJHTEntity{
			KodeSegmen: string(item["KODE_SEGMEN"]),
			NPP:        string(item["KODE_PERUSAHAAN"]),
			KodeTK:     string(item["KODE_TK"]),
		}
		dataSet = append(dataSet, data)
	}

	return dataSet, nil

}

func (repository *claimJHTRepositoryImpl) DaftarSegmen(params *request.DaftarSegmenRequest) ([]entity.DaftarKodeSegmenEntity, error) {
	session := db.EngineEcha.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return nil, err
	}
	result, err := db.EngineEcha.Query(`select kode,keterangan from BPJSTKU.ASIK_MS_LOOKUP where tipe ='KODE_SEGMEN'`)

	if err != nil {
		exception.PanicIfNeeded(err)
	}

	dataSet := make([]entity.DaftarKodeSegmenEntity, 0)
	for _, item := range result {
		data := entity.DaftarKodeSegmenEntity{
			Kode:       string(item["KODE"]),
			Keterangan: string(item["KETERANGAN"]),
		}
		dataSet = append(dataSet, data)
	}

	return dataSet, nil
}

func (repository *claimJHTRepositoryImpl) DaftarSebabKlaim(params *request.DaftarSebabKlaimRequest) ([]entity.DaftarKodeSebabKlaimEntity, error) {

	var filterSebabKlaim string

	session := db.EngineOltp.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return nil, err
	}

	if params.KodeSegmen == "PU" {
		filterSebabKlaim = `select 
		b.KODE_SEGMEN, a.KODE_TIPE_KLAIM, a.KODE_SEBAB_KLAIM, 
		a.NAMA_SEBAB_KLAIM,  'TK' KODE_TIPE_PENERIMA, a.NO_URUT
		from PN.PN_KODE_SEBAB_KLAIM a, PN.PN_KODE_SEBAB_SEGMEN b
		where a.KODE_SEBAB_KLAIM = b.KODE_SEBAB_KLAIM
		and a.KODE_TIPE_KLAIM = 'JHT01'
		and b.KODE_SEGMEN = ?
		and b.KODE_SEBAB_KLAIM in ('SKJ01','SKJ21','SKJ22','SKJ23','SKJ07','SKJ04','SKJ06')
		and nvl(b.STATUS_NONAKTIF,'X') = 'T'`
	} else {
		filterSebabKlaim = `select 
		b.KODE_SEGMEN, a.KODE_TIPE_KLAIM, a.KODE_SEBAB_KLAIM, 
		a.NAMA_SEBAB_KLAIM,  'TK' KODE_TIPE_PENERIMA, a.NO_URUT
		from PN.PN_KODE_SEBAB_KLAIM a, PN.PN_KODE_SEBAB_SEGMEN b
		where a.KODE_SEBAB_KLAIM = b.KODE_SEBAB_KLAIM
		and a.KODE_TIPE_KLAIM = 'JHT01'
		and b.KODE_SEGMEN = ?
		and b.KODE_SEBAB_KLAIM in ('SKJ06')
		and nvl(b.STATUS_NONAKTIF,'X') = 'T'`
	}
	result, err := db.EngineOltp.Query(filterSebabKlaim, params.KodeSegmen)

	if err != nil {
		exception.PanicIfNeeded(err)
	}

	dataSet := make([]entity.DaftarKodeSebabKlaimEntity, 0)
	for _, item := range result {
		data := entity.DaftarKodeSebabKlaimEntity{
			KodeSebabKlaim: string(item["KODE_SEBAB_KLAIM"]),
			NamaSebabKlaim: string(item["NAMA_SEBAB_KLAIM"]),
		}
		dataSet = append(dataSet, data)
	}

	return dataSet, nil
}
func (repository *claimJHTRepositoryImpl) DaftarDokumenSebabKlaim(params *request.DaftarDokumenSebabKlaimRequest) ([]entity.DaftarKodeDokumenSebabKlaimEntity, error) {

	session := db.EngineEcha.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return nil, err
	}

	result, err := db.EngineEcha.Query(`select a.kode_dokumen, a.nama_dokumen,b.flag_mandatory from BPJSTKU.asik_kode_dokumen a, bpjstku.asik_kode_dokumen_segmen b
	where a.kode_dokumen = b.kode_dokumen
	and b.kode_sebab_klaim = ?
	and nvl(b.status_nonaktif,'X') = 'T'
	order by b.no_urut asc
	`, params.KodeSebabKlaim)

	if err != nil {
		exception.PanicIfNeeded(err)
	}

	dataSet := make([]entity.DaftarKodeDokumenSebabKlaimEntity, 0)
	for _, item := range result {
		data := entity.DaftarKodeDokumenSebabKlaimEntity{
			KodeDokumen:   string(item["KODE_DOKUMEN"]),
			NamaDokumen:   string(item["NAMA_DOKUMEN"]),
			FlagMandatory: string(item["FLAG_MANDATORY"]),
		}
		dataSet = append(dataSet, data)
	}

	return dataSet, nil
}

func (repository *claimJHTRepositoryImpl) InsertPengajuanJHT(params *request.InsertPengajuanJHTRequest) (*entity.InsertPengajuanJHTEntity, error) {
	session := db.EngineEcha.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return nil, err
	}

	KodePengajuanKlaim := repository.GenKodeKlaim()

	// var p_param_maksimal_saldo_jht string
	// var p_klaim_sebagian string
	// var p_kode_pesan string
	// var p_pesan_tidak_layak string
	// var p_nama_kanal_kantor_cabang string
	// var p_sukses string
	// var p_mess string

	// Nama
	// NIK
	// KPJ
	// Email
	// NoHP
	// TempatLahir
	// TanggalLahir
	// NamaIbuKandung
	// Alamat
	// KodePos
	// KodeKelurahan
	// KodeKecamatan
	// KodeKabupaten
	// KodeProvinsi
	// KodeBank
	// NomorRekening
	// NamaRekening
	// NPWP
	// PathURLDokumen --asikklaimdokumen?
	// TanggalPengajuan
	// KodeBillingPTPos --belumada
	// ScoreFaceMatch
	// ScoreSimilarityNama --blmada
	// JawabanProbing --blmada
	// ChannelID --kanalpelayanan?
	// KeteranganPengajuan --blmada

	query := `insert into bpjstku.asik_klaim(
		kode_pengajuan,
		kanal_pelayanan,
		nama_tk,
		nomor_identitas,
		kpj,
		email,
		no_hp,
		tempat_lahir,
		tgl_lahir,
		nama_ibu_kandung,
		alamat,
		kode_pos,
		kode_kelurahan,
		kode_kecamatan,
		kode_kabupaten,
		kode_propinsi,
		kode_bank,
		no_rekening,
		nama_rekening,
		npwp,
		tgl_pengajuan,
		kode_billing,
		kemiripan_nama_pelapor,
		score_face,
		status_pengajuan,
		status_submit_pengajuan,
		tgl_submit_pengajuan,
		tgl_rekam,
		petugas_rekam,
		status_submit_dokumen,
		tgl_submit_dokumen
) values (
		?, --kode_pengajuan,
		'43', --kanal_pelayanan,
		?, --nama_tk,
		?, --nomor_identitas,
		?, --kpj,
		?, --email,
		?, --no_hp,
		?, --tempat_lahir,
		to_date(?,'dd-mm-yyyy'), --tgl_lahir,
		?, --nama_ibu_kandung,
		?, --alamat,
		?, --kode_pos,
		?, --kode_kelurahan,
		?, --kode_kecamatan,
		?, --kode_kabupaten,
		?, --kode_propinsi,
		?, --kode_bank,
		?, --no_rekening,
		?, --nama_rekening,
		?, --npwp,
		sysdate, --tgl_pengajuan,
		?, --kodebillingpos
		?, --kemiripannamapelapor
		?, --score_face
		'KLA1',--status_pengajuan
		'Y',--status_submit_pengajuan,
		sysdate,--tgl_submit_pengajuan
		sysdate,--tgl_rekam
		?,--petugas_rekam
		'Y',--status_submit_dokumen
		sysdate --tgl_submit_dokumen
	)
`
	// scoreFace, _ := strconv.Atoi(params.ScoreFaceMatch)
	// scoreFaceLiveness, _ := strconv.Atoi(params.ScoreFaceMatch)
	if _, err = session.Exec(
		query,
		KodePengajuanKlaim,
		params.Nama,
		params.NIK,
		params.KPJ,
		params.Email,
		params.NoHP,
		params.TempatLahir,
		params.TanggalLahir,
		params.NamaIbuKandung,
		params.Alamat,
		params.KodePos,
		params.KodeKelurahan,
		params.KodeKecamatan,
		params.KodeKabupaten,
		params.KodeProvinsi,
		params.KodeBank,
		params.NomorRekening,
		params.NamaRekening,
		params.NPWP,
		params.KodeBillingPTPos,
		params.ScoreSimilarityNama,
		params.ScoreFaceMatch,
		params.PetugasRekamPTPOS,
	); err != nil {
		session.Rollback()
		return nil, err
	}

	for _, itemDokumen := range params.Dokumen {
		query := `insert /*+append*/ into bpjstku.asik_klaim_dokumen(kode_pengajuan, kode_dokumen, path_url)
						 values(?, ?, ?)`
		if _, err = session.Exec(query, KodePengajuanKlaim, itemDokumen.KodeDokumen, itemDokumen.PathUrlDokumen); err != nil {
			session.Rollback()
			return nil, err
		}
	}

	for _, itemProbing := range params.DataProbing {
		query := `insert /*+append*/ into bpjstku.asik_probing(kode_pengajuan, kode_probing, no_urut,respon_probing,jawaban_probing,tgl_rekam,petugas_rekam)
						 values(?, ?, ?,?,?,sysdate,'PT POS')`
		if _, err = session.Exec(
			query,
			KodePengajuanKlaim,
			itemProbing.KodeProbing,
			itemProbing.NomorUrut,
			itemProbing.ResponseProbing,
			itemProbing.JawabanProbing,
		); err != nil {
			session.Rollback()
			return nil, err
		}
	}

	err = session.Commit()
	if err != nil {
		fmt.Println(err)
	}

	result := &entity.InsertPengajuanJHTEntity{
		StatusCode:         200,
		StatusKirim:        "SUKSES",
		KeteranganKirim:    "SUKSES",
		KodePengajuanKlaim: KodePengajuanKlaim,
	}

	return result, nil
}

func (repository *claimJHTRepositoryImpl) UpdatePengajuanJHT(params *request.UpdatePengajuanJHTRequest) (*entity.UpdatePengajuanJHTEntity, error) {
	session := db.EngineEcha.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return nil, err
	}

	// var p_param_maksimal_saldo_jht string
	// var p_klaim_sebagian string
	// var p_kode_pesan string
	// var p_pesan_tidak_layak string
	// var p_nama_kanal_kantor_cabang string
	// var p_sukses string
	// var p_mess string

	// Nama
	// NIK
	// KPJ
	// Email
	// NoHP
	// TempatLahir
	// TanggalLahir
	// NamaIbuKandung
	// Alamat
	// KodePos
	// KodeKelurahan
	// KodeKecamatan
	// KodeKabupaten
	// KodeProvinsi
	// KodeBank
	// NomorRekening
	// NamaRekening
	// NPWP
	// PathURLDokumen --asikklaimdokumen?
	// TanggalPengajuan
	// KodeBillingPTPos --belumada
	// ScoreFaceMatch
	// ScoreSimilarityNama --blmada
	// JawabanProbing --blmada
	// ChannelID --kanalpelayanan?
	// KeteranganPengajuan --blmada

	// 	query := `insert into bpjstku.asik_klaim(
	// 		kode_pengajuan,
	// 		kanal_pelayanan,
	// 		nama_tk,
	// 		nomor_identitas,
	// 		kpj,
	// 		email,
	// 		no_hp,
	// 		tempat_lahir,
	// 		tgl_lahir,
	// 		nama_ibu_kandung,
	// 		alamat,
	// 		kode_pos,
	// 		kode_kelurahan,
	// 		kode_kecamatan,
	// 		kode_kabupaten,
	// 		kode_propinsi,
	// 		kode_bank,
	// 		no_rekening,
	// 		nama_rekening,
	// 		npwp,
	// 		tgl_pengajuan,
	// 		score_face,
	// 		status_pengajuan,
	// 		status_submit_pengajuan,
	// 		tgl_submit_pengajuan,
	// 		tgl_rekam,
	// 		petugas_rekam,
	// 		status_submit_dokumen,
	// 		tgl_submit_dokumen
	// ) values (
	// 		?, --kode_pengajuan,
	// 		'43', --kanal_pelayanan,
	// 		?, --nama_tk,
	// 		?, --nomor_identitas,
	// 		?, --kpj,
	// 		?, --email,
	// 		?, --no_hp,
	// 		?, --tempat_lahir,
	// 		to_date(?,'dd-mm-yyyy'), --tgl_lahir,
	// 		?, --nama_ibu_kandung,
	// 		?, --alamat,
	// 		?, --kode_pos,
	// 		?, --kode_kelurahan,
	// 		?, --kode_kecamatan,
	// 		?, --kode_kabupaten,
	// 		?, --kode_propinsi,
	// 		?, --kode_bank,
	// 		?, --no_rekening,
	// 		?, --nama_rekening,
	// 		?, --npwp,
	// 		sysdate, --tgl_pengajuan,
	// 		?, --score_face
	// 		'KLA1',--status_pengajuan
	// 		'Y',--status_submit_pengajuan,
	// 		sysdate,--tgl_submit_pengajuan
	// 		sysdate,--tgl_rekam
	// 		'PT POS',--petugas_rekam
	// 		'Y',--status_submit_dokumen
	// 		sysdate --tgl_submit_dokumen
	// 	)
	// `
	// 	// scoreFace, _ := strconv.Atoi(params.ScoreFaceMatch)
	// 	// scoreFaceLiveness, _ := strconv.Atoi(params.ScoreFaceMatch)
	// 	if _, err = session.Exec(
	// 		query,
	// 		KodePengajuanKlaim,
	// 		params.Nama,
	// 		params.NIK,
	// 		params.KPJ,
	// 		params.Email,
	// 		params.NoHP,
	// 		params.TempatLahir,
	// 		params.TanggalLahir,
	// 		params.NamaIbuKandung,
	// 		params.Alamat,
	// 		params.KodePos,
	// 		params.KodeKelurahan,
	// 		params.KodeKecamatan,
	// 		params.KodeKabupaten,
	// 		params.KodeProvinsi,
	// 		params.KodeBank,
	// 		params.NomorRekening,
	// 		params.NamaRekening,
	// 		params.NPWP,
	// 		params.ScoreFaceMatch,
	// 	); err != nil {
	// 		session.Rollback()
	// 		return nil, err
	// 	}

	// 	err = session.Commit()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	result := &entity.UpdatePengajuanJHTEntity{
		StatusKirim:     "SUKSES",
		KeteranganKirim: "SUKSES",
	}

	return result, nil
}
