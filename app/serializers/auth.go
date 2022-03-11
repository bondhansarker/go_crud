package serializers

type LoginResponse struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	User         *UserResponse `json:"user"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtToken struct {
	UserID        uint   `json:"uid"`
	AccessToken   string `json:"act"`
	RefreshToken  string `json:"rft"`
	AccessUuid    string `json:"aid"`
	RefreshUuid   string `json:"rid"`
	AccessExpiry  int64  `json:"axp"`
	RefreshExpiry int64  `json:"rxp"`
}
