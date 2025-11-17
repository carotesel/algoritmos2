package calculadora

import (
	"fmt"
	"math"
)

const (
	OpSuma     = "+"
	OpResta    = "-"
	OpMult     = "*"
	OpDiv      = "/"
	OpPotencia = "^"
	OpRaiz     = "sqrt"
	OpLog      = "log"
	OpTernario = "?"
)

func CantidadOperandos(op string) int {
	switch op {
	case OpSuma, OpResta, OpMult, OpDiv, OpPotencia, OpLog:
		return 2
	case OpRaiz:
		return 1
	case OpTernario:
		return 3
	default:
		return -1
	}
}

func AplicarOperacion(op string, operandos []int64) (int64, error) {
	var res int64

	switch op {
	case OpSuma:
		res = operandos[0] + operandos[1]
	case OpResta:
		res = operandos[0] - operandos[1]
	case OpMult:
		res = operandos[0] * operandos[1]
	case OpDiv:
		if operandos[1] == 0 {
			return 0, fmt.Errorf("división por cero")
		}
		res = operandos[0] / operandos[1]
	case OpPotencia:
		if operandos[1] < 0 {
			return 0, fmt.Errorf("exponente negativo")
		}
		res = int64(math.Pow(float64(operandos[0]), float64(operandos[1])))
	case OpRaiz:
		if operandos[0] < 0 {
			return 0, fmt.Errorf("raíz negativa")
		}
		res = int64(math.Sqrt(float64(operandos[0])))
	case OpLog:
		base := operandos[1]
		valor := operandos[0]

		if base < 2 {
			return 0, fmt.Errorf("base inválida para log")
		}
		if valor <= 0 {
			return 0, fmt.Errorf("argumento inválido para log")
		}

		// Casos especiales
		if valor == 1 {
			res = 0 // log_base(1) = 0 para cualquier base válida
		} else if valor == base {
			res = 1 // log_base(base) = 1
		} else {
			// Uso fórmula matemática para logaritmo
			// log_base(valor) = ln(valor) / ln(base)

			logBase := math.Log(float64(base))
			logValor := math.Log(float64(valor))
			resultado := logValor / logBase

			// Truncar hacia cero como especifica la consigna
			res = int64(resultado)
		}

	case OpTernario:
		condicion := operandos[0]
		if condicion != 0 {
			res = operandos[1]
		} else {
			res = operandos[2]
		}
	default:
		return 0, fmt.Errorf("operador desconocido: %s", op)
	}

	return res, nil
}
