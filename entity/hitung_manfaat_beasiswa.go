package entity

type HitungManfaatBeasiswaEntity struct {
	ChId              string `json:"chId"`
	ReqId             string `json:"reqId"`
	KodeManfaat       int    `json:"KODE_MANFAAT"`
	NikPenerima       string `json:"NIK_PENERIMA"`
	KodeKlaim         string `json:"KODE_KLAIM"`
	NoUrut            int    `json:"NO_URUT"`
	Tahun             string `json:"TAHUN"`
	BeasiswaJenis     string `json:"BEASISWA_JENIS"`
	JenjangPendidikan string `json:"JENJANG_PENDIDIKAN"`
}
