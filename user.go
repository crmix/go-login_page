package loginapi

type SignUp struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
