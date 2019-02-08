package json_models

type UserJSON struct {
	Email    string `validate:"required, email" json:"email" binding:"required"`
	Password string `validate:"required" json:"password" binding:"required"`
}

type RegInfo struct {
	Email    string `validate:"required, email" json:"email" binding:"required"`
	Password string `validate:"required" json:"password" binding:"required"`
	VeriCode string `validate:"required" json:"code" binding:"required"`
}

type SendCodeRecv struct {
	Email string `validate:"required, email" json:"email" binding:"required"`
}

