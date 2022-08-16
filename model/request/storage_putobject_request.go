package request

type PutObjectRemoteRequest struct {
	FilePath   string `json:"file"`
	NamaBucket string `json:"namaBucket"`
	NamaFolder string `json:"namaFolder"`
}
