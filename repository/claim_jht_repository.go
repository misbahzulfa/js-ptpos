package repository

import (
	"js-ptpos/entity"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
)

type ClaimJHTRepository interface {
	CheckJHTEligible(params *request.CheckEligibleJHTRequest) (*response.CheckJHTEligibleResponse, error)
	GetPengajuanJHT(params *request.GetPengajuanJHTRequest) ([]entity.GetPengajuanJHTEntity, error)
	DaftarSegmen(params *request.DaftarSegmenRequest) ([]entity.DaftarKodeSegmenEntity, error)
	DaftarSebabKlaim(params *request.DaftarSebabKlaimRequest) ([]entity.DaftarKodeSebabKlaimEntity, error)
	DaftarDokumenSebabKlaim(params *request.DaftarDokumenSebabKlaimRequest) ([]entity.DaftarKodeDokumenSebabKlaimEntity, error)
	InsertPengajuanJHT(params *request.InsertPengajuanJHTRequest) (*entity.InsertPengajuanJHTEntity, error)
	UpdatePengajuanJHT(params *request.UpdatePengajuanJHTRequest) (*entity.UpdatePengajuanJHTEntity, error)
}
