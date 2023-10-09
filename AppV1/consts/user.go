package consts

import "time"

const (
	AccessTokenHeader  = "access_token"
	RefreshTokenHeader = "refresh_token"
	MaxAge             = 3600 * 24
)

const (
	AccessTokenExpireDuration  = 24 * time.Hour
	RefreshTokenExpireDuration = 10 * 24 * time.Hour
)
