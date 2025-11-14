package request

type UpdateAccount struct {
	Name     string `validate:"required,min=2" json:"name"`
	Password string `validate:"omitempty,min=8" json:"password"`
	Age      int64  `validate:"omitempty,gt=0" json:"age"`
	Phone    int64  `validate:"omitempty,gt=0" json:"phone"`
	District string `validate:"omitempty" json:"district"`
	City     string `validate:"omitempty" json:"city"`
	Country  string `validate:"omitempty" json:"country"`
}

type UpdateClient struct {
	Name     string  `validate:"required,min=2" json:"name"`
	Email    string  `validate:"omitempty,email" json:"email"`
	Password string  `validate:"omitempty,min=8" json:"password"`
	Role     []int64 `validate:"omitempty,dive,gt=0" json:"role_id"`
}
