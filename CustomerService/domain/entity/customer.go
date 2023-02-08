package entity

type Customer struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	City     City   `json:"city"`
}
