package entity

type Mail struct {
	Subject  string `json:"subject"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Body     string `json:"body"`
}
