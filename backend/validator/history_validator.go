package validator

import (
	"backend/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IHistoryValidation interface {
	HistoryValidate(history model.History) error
}

type historyValidator struct{}

func NewHistoryValidator() IHistoryValidation {
	return &historyValidator{}
}

// Description: Validate 関数でバリデーション実行する。第一引数に対象の変数を指定し、第二引数以降にルールを指定する。バリデーションを全て成功すればnilを返し、1つでも失敗すればエラーを返す。
func (hv *historyValidator) HistoryValidate(history model.History) error {
	return validation.ValidateStruct(&history,
		validation.Field(&history.Win, validation.Required),
		validation.Field(&history.Lose, validation.Required),
		validation.Field(&history.Money, validation.Required),
	)
}
