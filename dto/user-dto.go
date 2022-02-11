package dto

//UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}

// type UserCreateDTO struct {
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
// 	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
// }

type UserCheckDTO struct {
	Email string `json:"email" form:"email" binding:"required" validate:"email"`
}

type UserWebVerify struct {
	Email string `json:"email" form:"email" binding:"required,email"`
	OTP   string `json:"otp" form:"email" `
}

type UserWebRegisterDTO struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}

type SendWebOTPDTO struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}
