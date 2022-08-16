package request

type UpdateKonfirmasiJPECtoPNRequest struct {
	KPJ                   string              `json:"kpj"`
	NikPeserta            string              `json:"nikPeserta"`
	NamaPeserta           string              `json:"namaPeserta"`
	KodeTK                string              `json:"kodeTk"`
	KodeKlaim             string              `json:"kodeKlaim"`
	HubunganDenganPelapor string              `json:"hubunganDenganPelapor"`
	NamaHubunganLainnya   string              `json:"namaHubunganLainnya"`
	NamaPelapor           string              `json:"namaPelapor"`
	NikPelapor            string              `json:"nikPelapor"`
	TempatLahirPelapor    string              `json:"tempatLahirPelapor"`
	TanggalLahirPelapor   string              `json:"tanggalLahirPelapor"`
	JenisKelaminPelapor   string              `json:"jenisKelaminPelapor"`
	GolonganDarahPelapor  string              `json:"golonganDarahPelapor"`
	AlamatDomisiliPelapor string              `json:"alamatDomisiliPelapor"`
	KodePosPelapor        string              `json:"kodePosPelapor"`
	KelurahanPelapor      string              `json:"kelurahanPelapor"`
	KecamatanPelapor      string              `json:"kecamatanPelapor"`
	KabupatenPelapor      string              `json:"kabupatenPelapor"`
	Email                 string              `json:"email"`
	NoHP                  string              `json:"noHP"`
	NamaBank              string              `json:"namaBank"`
	NomorRekening         string              `json:"nomorRekening"`
	NamaRekening          string              `json:"namaRekening"`
	FaceMatchScore        string              `json:"faceMatchScore"`
	DataDokumen           []DataDokumenECtoPN `json:"dataDokumen"`
}

type DataDokumenECtoPN struct {
	KodeDokumen string `json:"kodeDokumen"`
	PathURL     string `json:"pathUrl"`
}
