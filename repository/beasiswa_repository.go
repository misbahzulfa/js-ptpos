package repository

import (
	"js-ptpos/entity"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
)

type BeasiswaRepository interface {
	CheckEligibleBeasiswa(params *request.CheckEligibleBeasiswaRequest) (*response.CheckEligibleBeasiswaResponse, error)
	InsertKonfirmasiBeasiswa(params *request.InsertKonfirmasiBeasiswaRequest) (*response.InsertKonfirmasiBeasiswaResponse, error)
	DaftarJenisBeasiswa(params *request.DaftarJenisBeasiswaRequest) (*response.DaftarJenisBeasiswaResponse, error)
	DaftarJenjangPendidikan(params *request.DaftarJenjangPendidikanRequest) ([]entity.DaftarJenjangPendidikanEntity, error)
	DaftarPenerimaBeasiswa(params *request.DaftarPenerimaBeasiswaRequest) (*response.DaftarPenerimaBeasiswaResponse, error)
	NominalPerjenjangBeasiswa(params *request.NominalPerjenjangBeasiswaRequest) (*response.NominalPerjenjangBeasiswaResponse, error)
}
