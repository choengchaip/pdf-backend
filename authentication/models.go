package authentication

type ILoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
