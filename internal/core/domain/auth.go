package domain

type Authentication struct {
	Token string `json:"token" db:"token"`
}

type Auth struct {
	Username string `json:"username" binding:"required,lowercase"`
	Password string `json:"password" binding:"required"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
