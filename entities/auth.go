package entities

type AuthRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type CustomerAuthResponse struct {
	Token string           `json:"token"`
	User  CustomerResponse `json:"user"`
}

type InternalAuthResponse struct {
	Token string           `json:"token"`
	User  InternalResponse `json:"user"`
}
