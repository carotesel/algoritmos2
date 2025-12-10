/*
15/07/24

1. Implementar en Go una primitiva para un árbol binario izquierdista, que reciba la cantidad de nodos que tiene, y
devuelva el dato del elemento más a abajo y a la derecha del árbol. En los árboles de las figuras mostradas, se debe
devolver en ambos casos 4. Para que el ejercicio se pueda considerar como aprobable, debe resolverse en no más que
O(n), sin contar con otros errores. Para que se considere completamente bien, debe ejecutar en O(log n). Justificar la
complejidad del algoritmo implementado. A fines del ejercicio, considerar que la estructura del árbol binario es:

type ab[T any] struct {
izquierda *ab[T]
derecha *ab[T]
dato T
}*/

/*
2. Implementar una primitiva para el árbol binario EsABB(func(T, T) int) bool que reciba una función de comparación y determine
si el árbol cumple con la propiedad de ABB para dicha función de comparación. Indicar y justificar la complejidad del algoritmo
implementado.
A fines del ejercicio, considerar que la estructura del árbol es la indicada en el dorso a este examen.

type ab[T any] struct {
izquierda *ab[T]
derecha *ab[T]
dato T
}
*/

/*
### 19/12/24

1. Implementar en Go una primitiva que reciba un árbol binario que representa un heap (árbol binario izquierdista, que
cumple la propiedad de heap), y devuelva la representación en arreglo del heap. La firma de la primitiva debe ser
RepresentacionArreglo() []T. Indicar y justificar la complejidad de la primitiva. La estructura del árbol binario es:
type ab[T any] struct {
izquierda *ab[T]
derecha *ab[T]
dato T
}
*/

/*
### 11/08/25

1. Explicar detalladamente cómo modificarías la implementación del ABB para poder tener una primitiva Maximo y Minimo que nos
devuelva las claves máximas y mínimas, y que se realice en tiempo constante.
*/