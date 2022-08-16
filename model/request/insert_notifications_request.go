package request

type InsertNotificationRequest struct {
	Tipe        string                       `json:"tipe"`
	Subject     string                       `json:"subject"`
	Message     string                       `json:"message"`
	Email       string                       `json:"email"`
	UserCreated string                       `json:"userCreated"`
	Receiver    []InsertNotificationToDetail `json:"receiver"`
}

type InsertNotificationToDetail struct {
	Email string `json:"email"`
}
