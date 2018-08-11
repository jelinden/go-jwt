package app

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Error    string `json:"error,omitempty"`
}

type Username struct {
	Username string `json:"username"`
}

type JwtToken struct {
	Token string `json:"token"`
}
