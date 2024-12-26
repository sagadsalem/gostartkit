package request

type LoginRequest struct {
	Email    string `json:"email" binding:"required" validate:"email,required"`
	Password string `json:"password" binding:"required" validate:"required"`
}
