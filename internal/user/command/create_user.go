package command

type CreateUserCommand struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Lang string `json:"password" validate:"required"`
}
