package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"js-ptpos/config"
	db "js-ptpos/config"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
	"net/http"
	"strconv"
	"strings"
)

func NewClaimJPRepository() ClaimJPRepository {
	return &claimJPRepositoryImpl{}
}

type claimJPRepositoryImpl struct {
}

func (repository *claimJPRepositoryImpl) JPCheckEligible(params *request.CheckEligibleJPBerkalaRequest) (*response.CheckEligibleJPBerkalaResponse, error) {
	//cek data penerima berkala
	var qryCnt = `select 
				count(*) CNT_PELAPOR
				from
				(
					select a.*,
					rank() over (partition by a.kode_klaim order by a.NO_KONFIRMASI desc, a.TGL_KONFIRMASI desc) rank
					from PN.PN_KLAIM_BERKALA a 
					where  
						nvl(a.STATUS_BATAL,'X') = 'T'
						and nvl(a.STATUS_SUBMIT,'T') = 'Y'
						and exists
						(
						select * from PN.PN_KLAIM_PENERIMA_BERKALA b
						where b.KODE_KLAIM = a.KODE_KLAIM
						and b.KODE_PENERIMA_BERKALA = a.KODE_PENERIMA_BERKALA
						and b.NOMOR_IDENTITAS = ? 
						and UTL_MATCH.EDIT_DISTANCE_SIMILARITY (UPPER(B.NAMA_LENGKAP), UPPER(?)) > 75 
						)
				)
				where rank = 1`

	cekPelapor, err := db.EngineOltp.Query(
		qryCnt,
		params.NikPenerimaManfaat,
		params.NamaPenerimaManfaat,
	)
	if err != nil {
		return nil, err
	}

	cnt := string(cekPelapor[0]["CNT_PELAPOR"])

	if cnt == "0" {
		var object = []response.DataPencarianJPBerkala{{
			StatusPencarian:     "T",
			KeteranganPencarian: "TIDAK LAYAK | Saat ini anda tidak terdaftar sebagai ahli waris | Silahkan menghubungi Kantor Cabang BPJS Ketenagakerjaan terdekat atau ke contact center 175",
			DaftarDokumen:       []response.DokumenCekBerkala{},
			DaftarProbing:       []response.DaftarProbingCekBerkala{},
		}}
		result := &response.CheckEligibleJPBerkalaResponse{
			StatusCode:       200,
			StatusDesc:       "OK",
			DataPenJPBerkala: object,
		}
		return result, nil

	} else {
		var qryPenerima = `select z.*,
								to_char(z.blth_akhir,'dd/mm/yyyy') blthakhir,
								(select nama_lengkap from PN.PN_KLAIM_PENERIMA_BERKALA where kode_klaim = z.kode_klaim and kode_penerima_berkala = z.kode_penerima_berkala) nama_penerima,
								(select nomor_identitas from PN.PN_KLAIM_PENERIMA_BERKALA where kode_klaim = z.kode_klaim and kode_penerima_berkala = z.kode_penerima_berkala) nik_penerima,
								(select tgl_lahir from PN.PN_KLAIM_PENERIMA_BERKALA where kode_klaim = z.kode_klaim and kode_penerima_berkala = z.kode_penerima_berkala) tgl_lahir_penerima
							from
							(
								select a.*,
								rank() over (partition by a.kode_klaim order by a.NO_KONFIRMASI desc, a.TGL_KONFIRMASI desc) rank
								from PN.PN_KLAIM_BERKALA a
								where
									nvl(a.STATUS_BATAL,'X') = 'T'
									and nvl(a.STATUS_SUBMIT,'T') = 'Y'
									and exists
									(
									select * from PN.PN_KLAIM_PENERIMA_BERKALA b
									where b.KODE_KLAIM = a.KODE_KLAIM
									and b.KODE_PENERIMA_BERKALA = a.KODE_PENERIMA_BERKALA
									and b.NOMOR_IDENTITAS = ?
									and UTL_MATCH.EDIT_DISTANCE_SIMILARITY (UPPER(B.NAMA_LENGKAP), UPPER(?)) > 75
									)
							) z
							where z.rank = 1`

		dataPenerimaBerkala, err := db.EngineOltp.Query(
			qryPenerima,
			params.NikPenerimaManfaat,
			params.NamaPenerimaManfaat,
		)
		if err != nil {
			return nil, err
		}

		kodeKlaim := string(dataPenerimaBerkala[0]["KODE_KLAIM"])
		namaPenerima := string(dataPenerimaBerkala[0]["NAMA_PENERIMA"])
		nikPenerima := string(dataPenerimaBerkala[0]["NIK_PENERIMA"])
		tglLahirPenerima := string(dataPenerimaBerkala[0]["TGL_LAHIR_PENERIMA"])

		var qryCekJPBelumdiBayarkan = `select count (*) v_cnt_blmbyr, to_char(last_day(max(blth_proses)),'DD/MM/YYYY') next_conf
							from pn.pn_klaim_berkala_rekap x
							where x.kode_klaim = ?
										and nvl (nom_berkala, 0) > 0
										and nvl (status_lunas, 'T') = 'T'`
		dataBelumdiBayarkan, err := db.EngineOltp.Query(
			qryCekJPBelumdiBayarkan,
			kodeKlaim,
		)
		if err != nil {
			return nil, err
		}

		cntBelumDibayarkan, err := strconv.Atoi(string(dataBelumdiBayarkan[0]["V_CNT_BLMBYR"]))

		if cntBelumDibayarkan > 0 {
			var object = []response.DataPencarianJPBerkala{{
				StatusPencarian: "T",
				KeteranganPencarian: "TIDAK LAYAK | Saat ini anda belum memasuki jadwal konfirmasi JP Berkala, anda dapat mengajukan konfirmasi JP Berkala setelah tanggal " + string(dataBelumdiBayarkan[0]["NEXT_CONF"]) +
					" Silahkan datang ke kembali ke PT POS atau ke Kantor Cabang BPJS Ketenagakerjaan terdekat sesuai dengan jadwal konfirmasi JP Berkala setelah tanggal " + string(dataBelumdiBayarkan[0]["NEXT_CONF"]),
				DaftarDokumen: []response.DokumenCekBerkala{},
				DaftarProbing: []response.DaftarProbingCekBerkala{},
			}}
			result := &response.CheckEligibleJPBerkalaResponse{
				StatusCode:       200,
				StatusDesc:       "OK",
				DataPenJPBerkala: object,
			}
			return result, nil
		}

		if string(dataPenerimaBerkala[0]["KODE_PENERIMA_BERKALA"]) == "A1" || string(dataPenerimaBerkala[0]["KODE_PENERIMA_BERKALA"]) == "A2" {
			//cek sudah bekerja atau belum
			var qryCntKpj = `SELECT NVL (kpj, 0) CNT_KPJ
							FROM (SELECT COUNT (kpj)     kpj
									FROM (SELECT a.*,
												RANK ()
													OVER (
														PARTITION BY a.kpj,
																	a.nomor_identitas,
																	a.nama_tk
														ORDER BY
															NVL (a.tgl_na,
																TO_DATE ('31/12/3000', 'dd/mm/yyyy')) DESC)
													RANK
											FROM kn.vw_kn_tk a
										WHERE     1 = 1
												AND a.nomor_identitas = ?
												AND a.nama_tk = ?
												AND a.tgl_lahir = TO_DATE (?, 'dd/mm/yyyy')
												AND NVL (a.kode_na, 'XXXXX') NOT IN ('AS', 'AG')
												AND EXISTS
														(SELECT *
															FROM kn.kn_kepesertaan_tk_prg c
														WHERE     c.kode_kepesertaan =
																	a.kode_kepesertaan
																AND c.kode_tk = a.kode_tk)
												AND EXISTS
														(SELECT *
															FROM kn.kn_iuran_tk b
														WHERE     b.kode_tk = a.kode_tk
																AND EXISTS
																		(SELECT *
																			FROM kn.kn_iuran_tk_prg c
																		WHERE c.kode_iuran =
																				b.kode_iuran)))
								WHERE RANK = 1 AND ROWNUM = 1)`

			dataCntKPJ, err := db.EngineOltp.Query(
				qryCntKpj,
				nikPenerima,
				namaPenerima,
				tglLahirPenerima,
			)
			if err != nil {
				return nil, err
			}

			var qryCek23 = `select case when 
									add_months(trunc(tgl_lahir,'mm'),12*23) <= trunc(add_months(to_date(?,'dd/mm/yyyy'),1),'MM')
										then 1
										else 0
									end validU23
							from pn.pn_klaim_penerima_berkala
							where kode_klaim = ?
							and kode_penerima_berkala = ?
							and nvl(status_layak,'T')='Y'
							and rownum = 1`
			dataqryCek23, err := db.EngineOltp.Query(
				qryCek23,
				string(dataPenerimaBerkala[0]["BLTHAKHIR"]),
				string(dataPenerimaBerkala[0]["KODE_KLAIM"]),
				string(dataPenerimaBerkala[0]["KODE_PENERIMA_BERKALA"]),
			)
			if err != nil {
				return nil, err
			}

			cntKpj, err := strconv.Atoi(string(dataCntKPJ[0]["CNT_KPJ"]))
			cntU23, err := strconv.Atoi(string(dataqryCek23[0]["VALIDU23"]))

			fmt.Println(cntBelumDibayarkan)

			if cntKpj > 0 {
				var object = []response.DataPencarianJPBerkala{{
					StatusPencarian:     "T",
					KeteranganPencarian: "TIDAK LAYAK | Saat ini anda sudah bekerja kembali, dan sudah tidak berhak atas JP Berkala BPJS Ketenagakerjaan | Silahkan menghubungi Kantor Cabang BPJS Ketenagakerjaan terdekat atau ke contact center 175",
					DaftarDokumen:       []response.DokumenCekBerkala{},
					DaftarProbing:       []response.DaftarProbingCekBerkala{},
				}}
				result := &response.CheckEligibleJPBerkalaResponse{
					StatusCode:       200,
					StatusDesc:       "OK",
					DataPenJPBerkala: object,
				}
				return result, nil
			} else {
				if cntU23 > 0 {
					var object = []response.DataPencarianJPBerkala{{
						StatusPencarian:     "T",
						KeteranganPencarian: "TIDAK LAYAK, Saat ini anda sudah memasuki usia 23 tahun, dan sudah tidak berhak atas JP Berkala BPJS Ketenagakerjaan | Silahkan menghubungi Kantor Cabang BPJS Ketenagakerjaan terdekat atau ke contact center 175",
						DaftarDokumen:       []response.DokumenCekBerkala{},
						DaftarProbing:       []response.DaftarProbingCekBerkala{},
					}}
					result := &response.CheckEligibleJPBerkalaResponse{
						StatusCode:       200,
						StatusDesc:       "OK",
						DataPenJPBerkala: object,
					}
					return result, nil
				}
			}
		}

		//get data dokumen
		dokumenQry, err := db.EngineEcha.Query(
			`select KODE, KETERANGAN from bpjstku.asik_ms_lookup where tipe = 'KODE_DOK_PELAPOR_BERKALA'`,
		)
		if err != nil {
			return nil, err
		}

		dokumens := make([]response.DokumenCekBerkala, 0)
		for _, item := range dokumenQry {
			dokumen := response.DokumenCekBerkala{
				KodeDokumen: string(item["KODE"]),
				NamaDokumen: string(item["KETERANGAN"]),
			}
			dokumens = append(dokumens, dokumen)
		}

		//get data probing
		var probingQry = `select kode_probing, no_urut, 
					(select b.nama_probing 
						from bpjstku.asik_kode_probing b 
					where b.kode_probing = a.kode_probing and rownum = 1) nama_probing, 
					(select b.kategori 
						from bpjstku.asik_kode_probing b 
					where b.kode_probing = a.kode_probing and rownum = 1) kategori,
					case 
						when kode_probing = 'PROBING001' then
							(select nama_tk from kn.vw_kn_tk@to_kn
								where nomor_identitas = ?
								and rownum = 1)
						when kode_probing = 'PROBING002' then
							(select tempat_lahir|| ' / ' ||to_char(nvl(tgl_lahir,'31-DEC-3000'),'DD-MM-YYYY') from kn.vw_kn_tk@to_kn
								where nomor_identitas = ?
								and rownum = 1)  
						when kode_probing = 'PROBING009' then
							(select nama_lengkap from pn.pn_klaim_penerima_berkala@to_kn
								where nomor_identitas = ?
								and rownum = 1)  
						when kode_probing = 'PROBING0010' then
							(select tempat_lahir|| ' / ' ||to_char(nvl(tgl_lahir,'31-DEC-3000'),'DD-MM-YYYY') from pn.pn_klaim_penerima_berkala@to_kn
								where nomor_identitas = ?
								and rownum = 1)  
						else null
					end respon_probing
				from bpjstku.asik_kode_probing_penerima a
				where  a.kode_tipe_klaim = 'JPN01' and a.kode_tipe_penerima = 'AW'
				order by a.no_urut`

		probingQryRes, err := db.EngineEcha.Query(
			probingQry,
			params.NikPeserta,
			params.NikPeserta,
			params.NikPenerimaManfaat,
			params.NikPenerimaManfaat,
		)
		if err != nil {
			return nil, err
		}

		daftarProbings := make([]response.DaftarProbingCekBerkala, 0)
		for _, item := range probingQryRes {
			daftarProbing := response.DaftarProbingCekBerkala{
				KodeProbing:     string(item["KODE_PROBING"]),
				NoUrut:          string(item["NO_URUT"]),
				NamaProbing:     string(item["NAMA_PROBING"]) + " " + string(item["RESPON_PROBING"]),
				ResponProbing:   string(item["RESPON_PROBING"]),
				KategoriProbing: string(item["KATEGORI"]),
			}
			daftarProbings = append(daftarProbings, daftarProbing)
		}

		var object = []response.DataPencarianJPBerkala{{
			StatusPencarian:     "Y",
			KeteranganPencarian: "Data Ditemukan",
			DaftarDokumen:       dokumens,
			DaftarProbing:       daftarProbings,
		}}

		result := &response.CheckEligibleJPBerkalaResponse{
			StatusCode:       200,
			StatusDesc:       "OK",
			DataPenJPBerkala: object,
		}

		return result, nil

	}

}

func (repository *claimJPRepositoryImpl) JPCheckJumlahKlaimBerkalaByNIKPelapor(params *request.CheckJumlahKlaimJPBerkalaNikPelaporRequest) (*response.CheckJumlahKlaimJPBerkalaNikPelaporResponse, error) {

	result := &response.CheckJumlahKlaimJPBerkalaNikPelaporResponse{
		StatusLebihDariSatu: "Y",
	}

	return result, nil
}

func (repository *claimJPRepositoryImpl) GetSubmissionCodeJP() string {
	var claimCode string

	var qry = `select bpjstku.p_ptpos_asik_klaim.f_gen_kode_pengajuan from dual`
	if _, err := db.EngineEcha.SQL(qry).Get(&claimCode); err != nil {
		fmt.Println(err)
	}

	return claimCode
}

func (repository *claimJPRepositoryImpl) InsertConfirmationJPBerkala(params *request.InsertJPConfirmationRequest) (*response.InsertKonfirmasiJPResponse, error) {
	configuration := config.New()
	client := &http.Client{}

	session := db.EngineEcha.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return nil, err
	}

	//Start Check Layak
	//Cek Apakah Ada Data dengan Status KLA1
	qry1 := `select count(*) cnt_pengajuan from bpjstku.asik_konfirmasi a
			where nik_penerima_manfaat = ? and status_pengajuan in ('KLA1','KLA2')`
	dataQry1, err := db.EngineEcha.Query(
		qry1,
		params.NikPenerimaManfaat,
	)
	if err != nil {
		return nil, err
	}

	cntPengajuan, err := strconv.Atoi(string(dataQry1[0]["CNT_PENGAJUAN"]))

	if cntPengajuan > 0 {
		DatainsertKonfirmationJPBerkalaRes := []response.DataInsertKonfirmasiJP{{
			StatusKirim:                      "T",
			Keterangan:                       "Pengajuan Konfirmasi JP Berkala dengan NIK " + params.NikPenerimaManfaat + " sudah terdapat didalam sistem dengan status KLA1-Pengajuan",
			KodePengajuanKonfirmasiJPBerkala: "",
		}}

		insertKonfirmationJPBerkalaRes := &response.InsertKonfirmasiJPResponse{
			StatusCode: 200,
			StatusDesc: "OK",
			Data:       DatainsertKonfirmationJPBerkalaRes,
		}
		return insertKonfirmationJPBerkalaRes, nil
	}

	var qryPenerima = `select z.*,
								to_char(z.blth_akhir,'dd/mm/yyyy') blthakhir,
								(select nama_lengkap from PN.PN_KLAIM_PENERIMA_BERKALA where kode_klaim = z.kode_klaim and kode_penerima_berkala = z.kode_penerima_berkala) nama_penerima,
								(select nomor_identitas from PN.PN_KLAIM_PENERIMA_BERKALA where kode_klaim = z.kode_klaim and kode_penerima_berkala = z.kode_penerima_berkala) nik_penerima,
								(select tgl_lahir from PN.PN_KLAIM_PENERIMA_BERKALA where kode_klaim = z.kode_klaim and kode_penerima_berkala = z.kode_penerima_berkala) tgl_lahir_penerima
							from
							(
								select a.*,
								rank() over (partition by a.kode_klaim order by a.NO_KONFIRMASI desc, a.TGL_KONFIRMASI desc) rank
								from PN.PN_KLAIM_BERKALA a
								where
									nvl(a.STATUS_BATAL,'X') = 'T'
									and nvl(a.STATUS_SUBMIT,'T') = 'Y'
									and exists
									(
									select * from PN.PN_KLAIM_PENERIMA_BERKALA b
									where b.KODE_KLAIM = a.KODE_KLAIM
									and b.KODE_PENERIMA_BERKALA = a.KODE_PENERIMA_BERKALA
									and b.NOMOR_IDENTITAS = ?
									and UTL_MATCH.EDIT_DISTANCE_SIMILARITY (UPPER(B.NAMA_LENGKAP), UPPER(?)) > 75
									)
							) z
							where z.rank = 1`

	dataPenerimaBerkala, err := db.EngineOltp.Query(
		qryPenerima,
		params.NikPenerimaManfaat,
		params.NamaPenerimaManfaat,
	)
	if err != nil {
		fmt.Println("## err dataPenerimaBerkala ")
		return nil, err
	}

	kodeKlaim := string(dataPenerimaBerkala[0]["KODE_KLAIM"])
	namaPenerima := string(dataPenerimaBerkala[0]["NAMA_PENERIMA"])
	nikPenerima := string(dataPenerimaBerkala[0]["NIK_PENERIMA"])
	tglLahirPenerima := string(dataPenerimaBerkala[0]["TGL_LAHIR_PENERIMA"])

	var qryCekJPBelumdiBayarkan = `select count (*) v_cnt_blmbyr, to_char(last_day(max(blth_proses)),'DD/MM/YYYY') next_conf
						from pn.pn_klaim_berkala_rekap x
						where x.kode_klaim = ?
									and nvl (nom_berkala, 0) > 0
									and nvl (status_lunas, 'T') = 'T'`
	dataBelumdiBayarkan, err := db.EngineOltp.Query(
		qryCekJPBelumdiBayarkan,
		kodeKlaim,
	)
	if err != nil {
		return nil, err
	}

	cntBelumDibayarkan, err := strconv.Atoi(string(dataBelumdiBayarkan[0]["V_CNT_BLMBYR"]))

	if cntBelumDibayarkan > 0 {
		DatainsertKonfirmationJPBerkalaRes := []response.DataInsertKonfirmasiJP{{
			StatusKirim:                      "T",
			Keterangan:                       "Pengajuan Konfirmasi JP Berkala Tidak layak, Saat ini anda belum memasuki jadwal konfirmasi JP Berkala, anda dapat mengajukan konfirmasi JP Berkala setelah tanggal " + string(dataBelumdiBayarkan[0]["NEXT_CONF"]),
			KodePengajuanKonfirmasiJPBerkala: "",
		}}

		insertKonfirmationJPBerkalaRes := &response.InsertKonfirmasiJPResponse{
			StatusCode: 200,
			StatusDesc: "OK",
			Data:       DatainsertKonfirmationJPBerkalaRes,
		}
		return insertKonfirmationJPBerkalaRes, nil
	}

	if string(dataPenerimaBerkala[0]["KODE_PENERIMA_BERKALA"]) == "A1" || string(dataPenerimaBerkala[0]["KODE_PENERIMA_BERKALA"]) == "A2" {
		//cek sudah bekerja atau belum
		var qryCntKpj = `SELECT NVL (kpj, 0) CNT_KPJ
						FROM (SELECT COUNT (kpj)     kpj
								FROM (SELECT a.*,
											RANK ()
												OVER (
													PARTITION BY a.kpj,
																a.nomor_identitas,
																a.nama_tk
													ORDER BY
														NVL (a.tgl_na,
															TO_DATE ('31/12/3000', 'dd/mm/yyyy')) DESC)
												RANK
										FROM kn.vw_kn_tk a
									WHERE     1 = 1
											AND a.nomor_identitas = ?
											AND a.nama_tk = ?
											AND a.tgl_lahir = TO_DATE (?, 'dd/mm/yyyy')
											AND NVL (a.kode_na, 'XXXXX') NOT IN ('AS', 'AG')
											AND EXISTS
													(SELECT *
														FROM kn.kn_kepesertaan_tk_prg c
													WHERE     c.kode_kepesertaan =
																a.kode_kepesertaan
															AND c.kode_tk = a.kode_tk)
											AND EXISTS
													(SELECT *
														FROM kn.kn_iuran_tk b
													WHERE     b.kode_tk = a.kode_tk
															AND EXISTS
																	(SELECT *
																		FROM kn.kn_iuran_tk_prg c
																	WHERE c.kode_iuran =
																			b.kode_iuran)))
							WHERE RANK = 1 AND ROWNUM = 1)`

		dataCntKPJ, err := db.EngineOltp.Query(
			qryCntKpj,
			nikPenerima,
			namaPenerima,
			tglLahirPenerima,
		)
		if err != nil {
			return nil, err
		}

		var qryCek23 = `select case when 
								add_months(trunc(tgl_lahir,'mm'),12*23) <= trunc(add_months(to_date(?,'dd/mm/yyyy'),1),'MM')
									then 1
									else 0
								end validU23
						from pn.pn_klaim_penerima_berkala
						where kode_klaim = ?
						and kode_penerima_berkala = ?
						and nvl(status_layak,'T')='Y'
						and rownum = 1`
		dataqryCek23, err := db.EngineOltp.Query(
			qryCek23,
			string(dataPenerimaBerkala[0]["BLTHAKHIR"]),
			string(dataPenerimaBerkala[0]["KODE_KLAIM"]),
			string(dataPenerimaBerkala[0]["KODE_PENERIMA_BERKALA"]),
		)
		if err != nil {
			return nil, err
		}

		cntKpj, err := strconv.Atoi(string(dataCntKPJ[0]["CNT_KPJ"]))
		cntU23, err := strconv.Atoi(string(dataqryCek23[0]["VALIDU23"]))

		if cntKpj > 0 {
			DatainsertKonfirmationJPBerkalaRes := []response.DataInsertKonfirmasiJP{{
				StatusKirim:                      "T",
				Keterangan:                       "Pengajuan Konfirmasi JP Berkala Tidak layak, Saat ini anda sudah bekerja kembali, dan sudah tidak berhak atas JP Berkala BPJS Ketenagakerjaan",
				KodePengajuanKonfirmasiJPBerkala: "",
			}}

			insertKonfirmationJPBerkalaRes := &response.InsertKonfirmasiJPResponse{
				StatusCode: 200,
				StatusDesc: "OK",
				Data:       DatainsertKonfirmationJPBerkalaRes,
			}
			return insertKonfirmationJPBerkalaRes, nil
		} else {
			if cntU23 > 0 {
				DatainsertKonfirmationJPBerkalaRes := []response.DataInsertKonfirmasiJP{{
					StatusKirim:                      "T",
					Keterangan:                       "Pengajuan Konfirmasi JP Berkala Tidak layak, Saat ini anda sudah memasuki usia 23 tahun, dan sudah tidak berhak atas JP Berkala BPJS Ketenagakerjaan",
					KodePengajuanKonfirmasiJPBerkala: "",
				}}

				insertKonfirmationJPBerkalaRes := &response.InsertKonfirmasiJPResponse{
					StatusCode: 200,
					StatusDesc: "OK",
					Data:       DatainsertKonfirmationJPBerkalaRes,
				}
				return insertKonfirmationJPBerkalaRes, nil
			}
		}
	}
	//End Of Pengecekan
	//####

	kodePengajuan := repository.GetSubmissionCodeJP()
	kodeTipeManfaat := "F001"

	query := `insert  /*+append*/ into bpjstku.asik_konfirmasi(
									KODE_PENGAJUAN,
									KODE_TIPE_MANFAAT,
									KODE_TIPE_KLAIM,
									NIK_PELAPOR,
									NAMA_PELAPOR,
									TGL_LAHIR_PELAPOR,
									NIK_PESERTA,
									EMAIL_PELAPOR,
									HANDPHONE_PELAPOR,
									KODE_BILLING, 
									KODE_POINTER_ASAL,
									ID_POINTER_ASAL,
									SKOR_FACE, 
									KEMIRIPAN_NAMA_PELAPOR,
									KANAL_PELAYANAN,
									TGL_PENGAJUAN,
									TGL_REKAM,
									PETUGAS_REKAM,
									NIK_PENERIMA_MANFAAT,
									NAMA_PENERIMA_MANFAAT,
									TGL_LAHIR_PENERIMA_MANFAAT 
								) values (
									?,--KODE_PENGAJUAN
									?,--KODE_TIPE_MANFAAT
									?,--KODE_TIPE_KLAIM
									?,--NIK_PELAPOR
									?,--NAMA_PELAPOR
									to_date(?,'dd/mm/yyyy'),--TGL_LAHIR_PELAPOR
									?,--NIK_PESERTA
									?,--EMAIL_PELAPOR
									?,--HANDPHONE_PELAPOR
									?,--KODE_BILLING 
									?,--KODE_POINTER_ASAL
									?,--ID_POINTER_ASAL
									?,--SKOR_FACE 
									?,--KEMIRIPAN_NAMA_PELAPOR
									'43', 
									sysdate,--TGL_PENGAJUAN
									sysdate,--TGL_REKAM
									?,--PETUGAS_REKAM 
									?,--NIK_PENERIMA_MANFAAT
									?,--NAMA_PENERIMA_MANFAAT
									to_date(?,'dd/mm/yyyy')--TGL_LAHIR_PENERIMA_MANFAAT 
									)`

	scoreFace, _ := strconv.Atoi(params.ScoreFaceMatch)
	kemiripanNamaPelapor, _ := strconv.Atoi(params.SimilarityNamaPTPOSkeAdminduk)

	if _, err = session.Exec(
		query,
		kodePengajuan,
		kodeTipeManfaat,
		"JPN01",
		params.NikPelapor,
		params.NamaPelapor,
		params.TanggalLahirPelapor,
		params.NikPeserta,
		params.EmailPelapor,
		params.NoHPPelapor,
		params.KodeBillingPTPos,
		params.KodeKantorPtPos,
		kodePengajuan,
		scoreFace,
		kemiripanNamaPelapor,
		params.PetugasRekamPtPos,
		params.NikPenerimaManfaat,
		params.NamaPenerimaManfaat,
		params.TanggalLahirPenerimaManfaat,
	); err != nil {
		session.Rollback()
		return nil, err
	}

	err = session.Commit()
	if err != nil {
		fmt.Println(err)
	}

	queryDokument := `insert /*+append*/ into bpjstku.asik_konfirmasi_dokumen (
									KODE_PENGAJUAN,
									KODE_DOKUMEN,
									PATH_URL_MITRA,
									FLAG_MANDATORY,
									FLAG_PALSU,
									TGL_REKAM,
									PETUGAS_REKAM
								)values(
									?,--KODE_PENGAJUAN
									?,--KODE_DOKUMEN
									?,--PATH_URL_MITRA
									?,--FLAG_MANDATORY,
									?,--FLAG_PALSU,
									SYSDATE,--TGL_REKAM
									?--PETUGAS_REKAM
								)`

	for _, itemDokumen := range params.DataDokumen {
		if _, err = session.Exec(queryDokument, kodePengajuan, itemDokumen.KodeDokumen, itemDokumen.PathURL, "T", "T", params.PetugasRekamPtPos); err != nil {
			session.Rollback()
			return nil, err
		}
	}

	queryProbing := `insert /*+append*/ into bpjstku.asik_probing(
									KODE_PENGAJUAN,
									KODE_PROBING,
									NO_URUT,
									RESPON_PROBING,
									JAWABAN_PROBING,
									KETERANGAN,
									TGL_REKAM,
									PETUGAS_REKAM 
								)values(
									?,--KODE_PENGAJUAN,
									?,--KODE_PROBING,
									?,--NO_URUT,
									?,--RESPON_PROBING,
									?,--JAWABAN_PROBING,
									?,--KETERANGAN,
									SYSDATE,--TGL_REKAM,
									?--PETUGAS_REKAM 
								)`

	for _, itemProbing := range params.DataProbing {
		if _, err = session.Exec(
			queryProbing,
			kodePengajuan,
			itemProbing.KodeProbing,
			itemProbing.NoUrut,
			itemProbing.ResponProbing,
			itemProbing.JawabanProbing,
			"",
			params.PetugasRekamPtPos); err != nil {
			session.Rollback()
			return nil, err
		}
	}

	//send Email
	qryEmail := `select * from kit.vas_email_content where vas_email_content_id=?`
	resultEmail, err := db.EngineOltp.Query(qryEmail, "53")
	if err != nil {
		return nil, err
	}

	var tgl_rekam string
	var qry = `select to_char(tgl_rekam,'DD-MM-YYYY HH24:mm:ss')||' WIB' tgl_rekam from bpjstku.asik_konfirmasi where kode_pengajuan='` + kodePengajuan + `'`
	if _, err := db.EngineEcha.SQL(qry).Get(&tgl_rekam); err != nil {
		fmt.Println(err)
	}

	var kpj string
	var qryKpj = `select (select kpj from pn.pn_klaim where kode_klaim = z.kode_klaim ) kpj 
					from 
						(
						select a.*,
						rank() over (partition by a.kode_klaim order by a.NO_KONFIRMASI desc, a.TGL_KONFIRMASI desc) ranking
							from PN.PN_KLAIM_BERKALA a 
						where  
							nvl(a.STATUS_BATAL,'X') = 'T'
							and nvl(a.STATUS_SUBMIT,'T') = 'Y' 
							and exists
							(
								select * from PN.PN_KLAIM_PENERIMA_BERKALA b
								where b.KODE_KLAIM = a.KODE_KLAIM
								and b.KODE_PENERIMA_BERKALA = a.KODE_PENERIMA_BERKALA
								and b.NOMOR_IDENTITAS ='` + params.NikPenerimaManfaat + `' 
								and UTL_MATCH.EDIT_DISTANCE_SIMILARITY (UPPER(B.NAMA_LENGKAP), UPPER('` + params.NamaPenerimaManfaat + `')) > 75 
							)
						) z
				where z.ranking = '1' and rownum = 1`
	if _, err := db.EngineOltp.SQL(qryKpj).Get(&kpj); err != nil {
		fmt.Println(err)
	}

	stringSubject := strings.Replace(string(resultEmail[0]["SUBJECT"]), ":0:", kpj, 1)
	kanalLayanan := strings.Split(params.KodeKantorPtPos, "|")

	string1 := strings.Replace(string(resultEmail[0]["HTML_CONTENT"]), ":1:", strings.ToUpper(params.NamaPelapor), 1)
	string2 := strings.Replace(string1, ":2:", strings.ToUpper(kanalLayanan[1]), 1)
	string3 := strings.Replace(string2, ":3:", kodePengajuan, 1)
	string4 := strings.Replace(string3, ":4:", tgl_rekam, 1)
	string5 := strings.Replace(string4, ":5:", "PENGAJUAN", 1)
	string6 := strings.Replace(string5, ":0:", kpj, 1)

	paramSendEmail := &request.SendEmailRequest{
		Subject: stringSubject,
		Body:    stringSubject,
		Message: string6,
		Email:   params.EmailPelapor,
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

	DatainsertKonfirmationJPBerkalaRes := []response.DataInsertKonfirmasiJP{{
		StatusKirim:                      "Y",
		Keterangan:                       "Berhasil Submit Data Pengajuan",
		KodePengajuanKonfirmasiJPBerkala: kodePengajuan,
	}}

	insertKonfirmationJPBerkalaRes := &response.InsertKonfirmasiJPResponse{
		StatusCode: 200,
		StatusDesc: "OK",
		Data:       DatainsertKonfirmationJPBerkalaRes,
	}

	if res.StatusCode == 200 {
		return insertKonfirmationJPBerkalaRes, nil
	} else {
		return nil, errors.New(sendEmailResponse.Message)
	}
	// return insertKonfirmationJPBerkalaRes, nil
}

func (repository *claimJPRepositoryImpl) CheckStatusKonfirmasiExsist(workerCode string) int {
	var total int

	var qry = `
			select 	count(1) 
			from 		bpjstku.asik_konfirmasi x 
			where 	nik_pelapor = ? 
							and nvl(status_batal, 'X') = 'T' 
							and status_pengajuan not in ('KLA5','KLA6')
	`
	if _, err := db.EngineEcha.SQL(qry, workerCode).Get(&total); err != nil {
		fmt.Println(err)
	}

	return total
}
