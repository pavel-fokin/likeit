package api

type UserCredentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInRequest struct {
	UserCredentials
}

type SignUpRequest struct {
	UserCredentials
}
