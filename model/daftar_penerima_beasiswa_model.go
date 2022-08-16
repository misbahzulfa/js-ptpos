package model

type DaftarPenerimaBeasiswa struct {
	NikPenerimaBeasiswa          string                        `json:"nikPenerimaBeasiswa"`
	NamaPenerimaBeasiswa         string                        `json:"namaPenerimaBeasiswa"`
	TanggalLahirPenerimaBeasiswa string                        `json:"tanggalLahirPenerimaBeasiswa"`
	JenisBeasiswa                string                        `json:"jenisBeasiswa"`
	JenjangPendidikanTerakhir    string                        `json:"jenjangPendidikanTerakhir"`
	DataDokumen                  []DaftarPenerimaBeasiswaDetil `json:"dataDokumen"`
}

type DaftarPenerimaBeasiswaDetil struct {
	KodeDokumen string `json:"KodeDokumen"`
	NamaDokumen string `json:"NamaDokumen"`
}
