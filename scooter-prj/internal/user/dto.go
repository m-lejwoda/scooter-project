package user

import "time"

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegister struct {
	Username string `json:"username"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID                     int    `json:"id"`
	Username               string `json:"username"`
	AccessToken            string `json:"accesstoken"`
	RefreshToken           string `json:"refreshtoken"`
	AccessTokenExpiration  int64  `json:"access_token_expiration"`
	RefreshTokenExpiration int64  `json:"refresh_token_expiration"`
}

type TokenResponse struct {
	AccessToken            string `json:"accesstoken"`
	RefreshToken           string `json:"refreshtoken"`
	AccessTokenExpiration  int64  `json:"access_token_expiration"`
	RefreshTokenExpiration int64  `json:"refresh_token_expiration"`
}

type TokensTTL struct {
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}
