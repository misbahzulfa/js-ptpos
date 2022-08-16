package response

type DaftarPenerimaBeasiswaResponse struct {
	StatusCode             string                   `json:"statusCode"`
	StatusDesc             string                   `json:"statusDesc"`
	DaftarPenerimaBeasiswa []DaftarPenerimaBeasiswa `json:"dataDaftarPenerimaBeasiswa"`
}

type DaftarPenerimaBeasiswa struct {
	NikPenerimaBeasiswa  string `json:"nikPenerimaBeasiswa"`
	NamaPenerimaBeasiswa string `json:"namaPenerimaBeasiswa"`
	//KodeKlaim                     string                          `json:"kodeKlaim"`
	NoUrut                        string                          `json:"noUrut"`
	TanggalLahirPenerimaBeasiswa  string                          `json:"tanggalLahirPenerimaBeasiswa"`
	DataTahunPenerimaBeasiswa     []DataTahunPenerimaBeasiswa     `json:"dataTahunPenerimaBeasiswa"`
	DataJenjangPendidikanTerakhir []DataJenjangPendidikanTerakhir `json:"dataJenjangPendidikanTerakhirLOV"`
	//DataDokumenPenerimaBeasiswa   []DataDokumenPenerimaBeasiswa   `json:"dataDokumenPenerimaBeasiswa"`
}

type DataTahunPenerimaBeasiswa struct {
	TahunBeasiswa               string                        `json:"tahunBeasiswa"`
	JenisBeasiswa               string                        `json:"jenisBeasiswa"`
	JenjangPendidikanTerakhir   string                        `json:"jenjangPendidikanTerakhir"`
	DataDokumenPenerimaBeasiswa []DataDokumenPenerimaBeasiswa `json:"dataDokumenPenerimaBeasiswa"`
}

type DataJenjangPendidikanTerakhir struct {
	Keterangan                    string                          `json:"jenjangPendidikan"`
	DataTingkatPendidikanTerakhir []DataTingkatPendidikanTerakhir `json:"dataTingkatPendidikanTerakhirLOV"`
}

type DataTingkatPendidikanTerakhir struct {
	Keterangan string `json:"tingkatPendidikan"`
}

type DataDokumenPenerimaBeasiswa struct {
	KodeDokumen string `json:"kodeDokumen"`
	NamaDokumen string `json:"namaDokumen"`
}
