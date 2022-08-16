package entity

type DaftarPenerimaBeasiswaEntity struct {
	NikPenerimaBeasiswa           string                          `json:"nikPenerimaBeasiswa"`
	NamaPenerimaBeasiswa          string                          `json:"namaPenerimaBeasiswa"`
	TanggalLahirPenerimaBeasiswa  string                          `json:"tanggalLahirPenerimaBeasiswa"`
	JenisBeasiswa                 string                          `json:"jenisBeasiswa"`
	JenjangPendidikanTerakhir     string                          `json:"jenjangPendidikanTerakhir"`
	DataJenjangPendidikanTerakhir []DataJenjangPendidikanTerakhir `json:"dataJenjangPendidikanTerakhir"`
	DataDokumen                   []DataDokumen                   `json:"dataDokumen"`
}

type DataDokumen struct {
	KodeDokumen string `json:"kodeDokumen"`
	NamaDokumen string `json:"namaDokumen"`
}

type DataJenjangPendidikanTerakhir struct {
	Keterangan                    string                          `json:"jenjangPendidikan"`
	DataTingkatPendidikanTerakhir []DataTingkatPendidikanTerakhir `json:"dataTingkatPendidikanTerakhir"`
}

type DataTingkatPendidikanTerakhir struct {
	Keterangan string `json:"tingkatPendidikan"`
}
