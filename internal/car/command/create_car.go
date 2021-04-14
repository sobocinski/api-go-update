package command

type CreateCarCommand struct {
	Model string `json:"model" validate:"required"`
	UserId uint `json:"userId" validate:"required"`
}
