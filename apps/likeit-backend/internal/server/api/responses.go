package api

type GetLikesResponse struct {
	Likes int `json:"likes"`
}

type SignInResponse struct {
	AccessToken string `json:"accessToken"`
}

type SignUpResponse struct {
	AccessToken string `json:"accessToken"`
}
