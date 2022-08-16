package request

type InsertKonfirmasiBeasiswaRequest struct {
	NikPelapor                    string                          `json:"nikPelapor"`
	NamaPelapor                   string                          `json:"namaPelapor"`
	NikPeserta                    string                          `json:"nikPeserta"`
	TglLahirPelapor               string                          `json:"tglLahirPelapor"`
	EmailPelapor                  string                          `json:"emailPelapor"`
	HandphonePelapor              string                          `json:"handphonePelapor"`
	TglPengajuan                  string                          `json:"tglPengajuan"`
	KodeBilling                   string                          `json:"kodeBilling"`
	SkorFace                      string                          `json:"skorFace"`
	KemiripanNamaPelapor          string                          `json:"kemiripanNamaPelapor"`
	KeteranganApproval            string                          `json:"keteranganApproval"`
	ChannelID                     string                          `json:"channelID"`
	ChId                          string                          `json:"chId"`
	ReqId                         string                          `json:"reqId"`
	KodeKantorPTPOS               string                          `json:"kodeKantorPTPOS"`
	PetugasRekamPTPOS             string                          `json:"petugasRekamPTPOS"`
	PenerimaBeasiswa              []PenerimaBeasiswa              `json:"penerimaBeasiswa"`
	DataProbingInsertKonfBeasiswa []DataProbingInsertKonfBeasiswa `json:"dataProbing"`
	DataDokumenPelapor            []DataDokumenPelapor            `json:"dataDokumenPelapor"`
}

type PenerimaBeasiswa struct {
	NikPenerimaBeasiswa      string                `json:"nikPenerimaBeasiswa"`
	NamaPenerimaBeasiswa     string                `json:"namaPenerimaBeasiswa"`
	NoUrutPenerima           string                `json:"noUrutPenerima"`
	TglLahirPenerimaBeasiswa string                `json:"tglLahirPenerimaBeasiswa"`
	FlagMasihSekolah         string                `json:"flagMasihSekolah"`
	TahunBeasiswa            string                `json:"tahunBeasiswa"`
	KodeJenisBeasiswa        string                `json:"kodeJenisBeasiswa"`
	TingkatPendidikan        string                `json:"tingkatPendidikan"`
	JenjangPendidikan        string                `json:"jenjangPendidikan"`
	LembagaPendidikan        string                `json:"lembagaPendidikan"`
	FlagDokLengkap           string                `json:"flagDokLengkap"`
	DataDokumenBeasiswa      []DataDokumenBeasiswa `json:"dataDokumenBeasiswa"`
}

// type PenerimaBeasiswaDetil struct {
// 	TahunBeasiswa     string `json:"tahunBeasiswa"`
// 	KodeJenisBeasiswa string `json:"kodeJenisBeasiswa"`
// 	TingkatPendidikan string `json:"tingkatPendidikan"`
// 	JenjangPendidikan string `json:"jenjangPendidikan"`
// 	LembagaPendidikan string `json:"lembagaPendidikan"`
// 	FlagDokLengkap    string `json:"flagDokLengkap"`
// }

type DataDokumenBeasiswa struct {
	KodeDokumen string `json:"kodeDokumen"`
	PathUrl     string `json:"pathUrl"`
}

type DataDokumenPelapor struct {
	KodeDokumen string `json:"kodeDokumen"`
	PathUrl     string `json:"pathUrl"`
}

type DataProbingInsertKonfBeasiswa struct {
	KodeProbing    string `json:"kodeProbing"`
	NoUrut         string `json:"noUrut"`
	ResponProbing  string `json:"responProbing"`
	JawabanProbing string `json:"jawabanProbing"`
}
