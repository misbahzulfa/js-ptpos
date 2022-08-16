package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"js-ptpos/config"
	db "js-ptpos/config"
	"js-ptpos/exception"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// func NewStorageRepository() StorageRepository {
// 	return &storageRepositoryImpl{}
// }

// type storageRepositoryImpl struct {
// 	Configuration config.Config
// }

func NewStorageRepository(configuration *config.Config) StorageRepository {
	return &storageRepositoryImpl{
		Configuration: *configuration,
	}
}

type storageRepositoryImpl struct {
	Configuration config.Config
}

func (repository *storageRepositoryImpl) PutObject(params *request.PutObjectRemoteRequest) (*response.PutObjectRemoteResponse, error) {
	fmt.Println("Start PutObject")
	client := &http.Client{}
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(params.FilePath)
	if errFile1 != nil {
		return nil, errFile1
	}

	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("file", filepath.Base(params.FilePath))
	if errFile1 != nil {
		return nil, errFile1
	}

	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		return nil, errFile1
	}

	_ = writer.WriteField("namaBucket", params.NamaBucket)
	_ = writer.WriteField("namaFolder", params.NamaFolder)

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	fmt.Println("Call put objectt")
	fmt.Println(repository.Configuration.Get("STORAGE_URL"))

	req, err := http.NewRequest("POST", repository.Configuration.Get("STORAGE_URL")+"/put-object", payload)
	if err != nil {
		fmt.Println(err)
		return nil, exception.CallApiExceptionMessage
	}

	fmt.Println("After Call put objectt")

	putObjectRemoteResponse := &response.PutObjectRemoteResponse{}
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, exception.ClientDoExceptionMessage
	}
	json.NewDecoder(res.Body).Decode(&putObjectRemoteResponse)
	defer res.Body.Close()

	os.Remove(params.FilePath)
	return putObjectRemoteResponse, nil
}

func (repository *storageRepositoryImpl) DeleteObject(params *request.StorageDeleteObjectRemoteRequest) (*response.StorageDeleteObjectRemoteResponse, error) {
	client := &http.Client{}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, exception.JSONParseExceptionMessage
	}

	payload := strings.NewReader(string(jsonData))
	req, err := http.NewRequest("POST", repository.Configuration.Get("STORAGE_URL")+"/object/delete", payload)

	if err != nil {
		return nil, exception.CallApiExceptionMessage
	}

	deleteObjectRemoteResponse := &response.StorageDeleteObjectRemoteResponse{}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, exception.ClientDoExceptionMessage
	}
	json.NewDecoder(res.Body).Decode(&deleteObjectRemoteResponse)
	defer res.Body.Close()

	return deleteObjectRemoteResponse, nil
}

func (repository *storageRepositoryImpl) UpdateDokumenEC(params *request.UpdateDokumenRequest) (*response.UpdateDokumenResponse, error) {
	var tabel string
	var tabel2 string
	var kode_dokumen string
	session := db.EngineEcha.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return nil, err
	}

	jenisLayanan := params.JenisLayanan

	if jenisLayanan == "JHT" {
		tabel = `bpjstku.asik_klaim_dokumen`
		tabel2 = `bpjstku.asik_klaim`
		kode_dokumen = `kode_dokumen`
	} else if jenisLayanan == "JP" {
		tabel = `bpjstku.asik_konfirmasi_dokumen`
		tabel2 = `bpjstku.asik_konfirmasi`
		kode_dokumen = `kode_dokumen`
	} else if jenisLayanan == "BEASISWA" {
		tabel = `bpjstku.asik_konfirmasi_beasiswa_dok`
		tabel2 = `bpjstku.asik_konfirmasi`
		kode_dokumen = `nama_dokumen`
	}

	qryUpdatePathUrl := `UPDATE ` + tabel + ` 
								SET
									path_url = ?,
									mime_type = ?,
									tgl_ubah = sysdate,
									petugas_ubah = ?
								where kode_pengajuan = ?
									and ` + kode_dokumen + ` = ?`

	if _, err = session.Exec(
		qryUpdatePathUrl,
		params.PathURL,
		strings.Replace(params.Mimetype, "@", "", 1),
		params.PetugasRekamPtPos,
		params.KodePengajuan,
		params.KodeDokumen,
	); err != nil {
		session.Rollback()
		return nil, err
	}

	qryUpdateSubmitDokument := `UPDATE ` + tabel2 + `
								SET 
									TGL_SUBMIT_DOKUMEN = SYSDATE,
									STATUS_SUBMIT_DOKUMEN = 'Y'
								where kode_pengajuan = ?`

	if _, err = session.Exec(
		qryUpdateSubmitDokument,
		params.KodePengajuan,
	); err != nil {
		session.Rollback()
		return nil, err
	}

	err = session.Commit()

	var respon *response.UpdateDokumenResponse

	if err != nil {
		fmt.Println(err)
	} else {
		respon = &response.UpdateDokumenResponse{
			StatusUpdate: "0",
			StatusDesc:   "Update pathURL Dokumen berhasil",
		}
	}
	return respon, nil
}
