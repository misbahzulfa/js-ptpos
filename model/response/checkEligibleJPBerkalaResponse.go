package response

type CheckEligibleJPBerkalaResponse struct {
	StatusCode       int                      `json:"StatusCode"`
	StatusDesc       string                   `json:"StatusDesc"`
	DataPenJPBerkala []DataPencarianJPBerkala `json:"Data"`
}

type DataPencarianJPBerkala struct {
	StatusPencarian     string                    `json:"statusPencarian"`
	KeteranganPencarian string                    `json:"keteranganPencarian"`
	DaftarDokumen       []DokumenCekBerkala       `json:"dataDokumen"`
	DaftarProbing       []DaftarProbingCekBerkala `json:"dataProbing"`
}

type DokumenCekBerkala struct {
	KodeDokumen string `json:"kodeDokumen"`
	NamaDokumen string `json:"namaDokumen"`
}

type DaftarProbingCekBerkala struct {
	KodeProbing     string `json:"kodeProbing"`
	NoUrut          string `json:"noUrut"`
	NamaProbing     string `json:"namaProbing"`
	ResponProbing   string `json:"responProbing"`
	KategoriProbing string `json:"kategori"`
}
