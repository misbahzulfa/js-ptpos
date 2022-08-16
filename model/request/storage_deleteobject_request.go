package request

type StorageDeleteObjectRemoteRequest struct {
	File       string `json:"file"`
	NamaBucket string `json:"namaBucket"`
	NamaFolder string `json:"namaFolder"`
}
