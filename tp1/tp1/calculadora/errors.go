package calculadora

import "errors"

var (
	ErrExpresionInvalida = errors.New("expresión inválida")
	ErrDivisionPorCero   = errors.New("división por cero")
)
