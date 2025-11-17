package calculadora

import (
	"strconv"
	"strings"
)

// Funciones

// - tokenizarLinea(linea string)[] string: Devuelve la linea pero separada por tokens (caracteres)
//y sin espacios raros. DONE

// - esNumero (token string) bool: me dice si mi caracter es numero o no (para poder decir si es un operador). DONE

// - cantidadOperandos(op string) int: recibe un operador y me dice cuantos numeros necesita para realizar la operacion. DONE

// - aplicarOperacion(op string, operandos []intg4)(int64, error):
// 		* Ejecuta la operación matemática.
//		* Valida errores como: división por 0, raíz negativa, log inválido, etc.
//		* Devuelve el resultado y/o error.
//DONE

// -evaluarLinea(linea string) (int64, error):
// 		* Tokenizar la línea (split de espacios)
//		* Pila Vacia para operandos
//
//		*Recorro arr tokens:
//			- Si es nro convierto y apilo
// 			- Si es operador proceso op
//			- Verifico que op sea valido
// 			- Verifico que haya suficientes operandos en la pila
//			- Desapilo operandos necesarios
//			- Calculo op
//			- Apilo res
//
//		* Debe quedar exactamente 1 resultado en la pila -> Lo devuelvo

func TokenizarLinea(linea string) []string {
	tokens := strings.Fields(linea) // es como el split pero no guarda los " " en el array
	return tokens
}

func EsNumero(token string) bool {
	_, err := strconv.ParseInt(token, 10, 64) // convierte a int64
	return err == nil
}
