package notify

// Body 메세지 내용
type Body struct {
	Users     []User            `json:"users"`
	Message   Message           `json:"message"`
	ClientKey map[string]string `json:"client_key"`
}

// Message 메세지
type Message struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// User 회신자
type User struct {
	EMail        string            `json:"email"`
	Phone        string            `json:"phone"`
	Channel      string            `json:"channel"`
	StringFormat map[string]string `json:"string_format"`
}
