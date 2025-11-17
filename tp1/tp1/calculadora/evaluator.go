package calculadora

import (
	"strconv"
	"tdas/pila"
)

func EvaluarLinea(linea string) (int64, error) {
	tokens := TokenizarLinea(linea)
	p := pila.CrearPilaDinamica[int64]()
	var res int64
	count := 0

	for _, token := range tokens {
		if EsNumero(token) {
			num, err := strconv.ParseInt(token, 10, 64)
			if err != nil {
				return 0, ErrExpresionInvalida
			}
			p.Apilar(num)
			count++
		} else {
			n := CantidadOperandos(token)
			if n == -1 {
				return 0, ErrExpresionInvalida
			}

			if count < n {
				return 0, ErrExpresionInvalida
			}

			operandos := make([]int64, n)
			for i := n - 1; i >= 0; i-- {
				operandos[i] = p.Desapilar()
			}
			count -= n

			var err error
			res, err = AplicarOperacion(token, operandos)
			if err != nil {
				return 0, err
			}

			p.Apilar(res)
			count++
		}
	}

	if count != 1 {
		return 0, ErrExpresionInvalida
	}
	return p.Desapilar(), nil
}
