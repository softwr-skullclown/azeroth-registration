package domain

type Account struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
