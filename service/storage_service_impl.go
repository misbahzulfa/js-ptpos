package service

import (
	"encoding/base64"
	"fmt"
	"js-ptpos/exception"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
	"js-ptpos/repository"
	"js-ptpos/util"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NewStorageService(storageRepository *repository.StorageRepository) StorageService {
	return &storageServiceImpl{
		StorageRepository: *storageRepository,
	}
}

type storageServiceImpl struct {
	StorageRepository repository.StorageRepository
}

func (service *storageServiceImpl) UploadDokumen(ctx *fiber.Ctx, params *request.UploadDocRequest) (*response.UploadDokumenResponse, error) {
	//var result *response.UploadDokumenResponse

	dec, err := base64.StdEncoding.DecodeString(params.File[0].Data)
	if err != nil {
		//panic(err)
	}

	mimeFile := strings.Split(params.File[0].Mime, "/")
	ext := mimeFile[1]

	if ext == "pdf" || ext == "jpg" || ext == "jpeg" || ext == "png" {
		filename := params.KodePengajuan + "_" + params.KodeDokumen + "." + ext
		pathFileUpload := fmt.Sprintf("temp_folder/%s", filename)

		fmt.Println("####" + filename)
		fmt.Println("####" + pathFileUpload)

		f, err := os.Create(pathFileUpload)
		if err != nil {
			//panic(err)
		}
		defer f.Close()
		f.Write(dec)

		//get file information
		fi, err := f.Stat()
		if err != nil {
			// Could not obtain stat, handle error
		}

		//check file size
		if fi.Size() > 2097152 {
			//Hapus File
			errDelFile := os.Remove(pathFileUpload)
			if errDelFile != nil {
				fmt.Println(errDelFile)
			}

			message := "File Size lebih besar dari 2 MB"
			logStop := util.LogResponse(ctx, message, params.ReqId)
			fmt.Println(logStop)
			panic(exception.GeneralError{
				Message: message,
			})
		} else {
			if params.File[0].Data != "" {
				fmt.Println("putObjectRemoteRequest")
				fmt.Println("####" + filename)
				fmt.Println("####" + pathFileUpload)
				putObjectRemoteRequest := &request.PutObjectRemoteRequest{
					FilePath:   pathFileUpload,
					NamaBucket: "ptpos",
					NamaFolder: params.KodePengajuan + "/" + params.KodeDokumen,
				}

				storage, err := service.StorageRepository.PutObject(putObjectRemoteRequest)
				if err != nil {
					fmt.Println(err)
				}

				if storage.Ret == "0" {
					fmt.Println("CEPHFILE->", storage.Path)
					mimeFile_ := strings.Replace(params.File[0].Mime, "@", "", 0)
					reqUpdatedok := &request.UpdateDokumenRequest{
						KodePengajuan:     params.KodePengajuan,
						JenisLayanan:      params.JenisLayanan,
						KodeDokumen:       params.KodeDokumen,
						PetugasRekamPtPos: params.PetugasRekamPtPos,
						Mimetype:          mimeFile_,
						PathURL:           storage.Path,
					}

					updateDokumen, err := service.StorageRepository.UpdateDokumenEC(reqUpdatedok)
					if err != nil {
						fmt.Println(err)
					}

					if updateDokumen.StatusDesc == "0" {
						errDelFile := os.Remove(pathFileUpload)
						if errDelFile != nil {
							fmt.Println(errDelFile)
						}
					}
				} else {
					errDelFile := os.Remove(pathFileUpload)
					if errDelFile != nil {
						fmt.Println(errDelFile)
					}

					message := "Gagal upload ke storage file, silahkan cek kembali file stream dan mime type"
					logStop := util.LogResponse(ctx, message, params.ReqId)
					fmt.Println(logStop)
					panic(exception.GeneralError{
						Message: message,
					})
				}

				result := &response.UploadDokumenResponse{
					StatusUpload: 200,
					Message:      "Sukses Upload File",
				}
				return result, nil
			} else {
				//Hapus File
				errDelFile := os.Remove(pathFileUpload)
				if errDelFile != nil {
					fmt.Println(errDelFile)
				}

				message := "File Stream Kosong"
				logStop := util.LogResponse(ctx, message, params.ReqId)
				fmt.Println(logStop)
				panic(exception.GeneralError{
					Message: message,
				})
			}
		}
	} else {
		message := "File Extension Tidak diperbolehkan"
		logStop := util.LogResponse(ctx, message, params.ReqId)
		fmt.Println(logStop)
		panic(exception.GeneralError{
			Message: message,
		})
	}
}
