package request



type AuthRequest struct{
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthResponse struct{
	ID int `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}