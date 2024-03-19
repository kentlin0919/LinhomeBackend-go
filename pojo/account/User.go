package pojo

type Account struct {
	Id       int    `json:"userid"`
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
