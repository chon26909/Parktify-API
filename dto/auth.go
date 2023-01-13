package dto

type SignInRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type SignUpRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
