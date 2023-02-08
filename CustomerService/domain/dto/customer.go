package dto

type Customer struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Country  string `json:"country"`
	City     string `json:"city"`
}
