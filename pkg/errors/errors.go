package errors

import "errors"

var (
	ErrDataNotFound         = errors.New("не найдено")
	ErrAlreadyHasUser       = errors.New("вы уже зарегестрированы")
	ErrWrongPassword        = errors.New("неправильный пароль")
	ErrProjectAlreadyExists = errors.New("такой проект уже существует")
	ErrTaskAlreadyExists    = errors.New("такая задача уже существует")
)
