package models

type RefreshTokenMetadata struct {
	Jti     string
	Expires int64
	Iat     int64
	Nbf     int64
	Sub     string
}

type AccessTokenMetadata struct {
	Expires int64
	Iat     int64
	Nbf     int64
	Sub     string
}
