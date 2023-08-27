package validator

import (
	"backend/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IPlayerValidation interface {
	PlayerValidate(player model.Player) error
}

type playerValidator struct {}

func NewPlayerValidator() IPlayerValidation {
	return &playerValidator{}
}

func (pv *playerValidator) PlayerValidate(player model.Player) error {
	return validation.ValidateStruct(&player,
		validation.Field(&player.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&player.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("limited min 6 max 30 char"),
		),
	)
}