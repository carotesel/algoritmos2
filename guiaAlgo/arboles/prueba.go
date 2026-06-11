/*1. Implementar una primitiva para el ABB func (arbol *abb[K, V]) AncestroComun(clave1, clave2 K) K que reciba
2 claves y devuelva el último ancestro en común entre ambas claves. Dicho ancestro en común podría ser incluso alguna
de estas claves. Si alguna clave pasada no se encuentra en el árbol, finalizar con panic. Indicar y justificar la complejidad
de la primitiva implementada.
Mostramos ejemplos de resultados esperados de invocar la primitiva al árbol del dorso:
arbol.AncestroComun(1, 4) --> 2
arbol.AncestroComun(2, 4) --> 2
arbol.AncestroComun(9, 1) --> 5*/

func (arbol *abb[K, V]) AncestroComun(clave1, clave2 K) K {
	if !arbol.Pertenece(clave1) || !arbol.Pertenece(clave2) {
		panic("Una de las claves no se encuentra en el arbol")
	}

	return ancestroCom(arbol.raiz, arbol.cmp, clave1, clave2)
}

func ancestroCom(nodo *nodoABB[K, V], cmp func(K, K) int, clave1, clave2 K) K {
	if nodo == nil {
		return panic("No hay ancestro comun de claves en un arbol vacio")
	}

	if cmp(nodo.clave, clave1) > 0 && cmp(nodo.clave, clave2) > 0 {
		return ancestroCom(nodo.izq, cmp, clave1, clave2)
	}

	if cmp(nodo.clave, clave1) < 0 && cmp(nodo.clave, clave2) < 0 {
		return ancestroCom(nodo.der, cmp, clave1, clave2)
	}

	caso1 := cmp(nodo.clave, clave1) <= 0 && cmp(nodo.clave, clave2) >= 0
	caso2 := cmp(nodo.clave, clave1) >= 0 && cmp(nodo.clave, clave2) <= 0

	if caso1 || caso2 {
		return nodo.clave
	}

}

// O(log n)

/*4. Implementar una primitiva para el árbol binario esABB(comparacion func(T, T) int) bool que determine si el mismo cumple
propiedad de ABB dada la función de comparación pasada por parámetro. Indicar y justificar la complejidad del algoritmo
implementado. La estructura del árbol es:
type Arbol[T any] struct {
izq *Arbol[T]
der *Arbol[T]
dato T
}
*/

func (ab *Arbol[T]) esABB(comparacion func(T, T) int) bool{
	return esABBrec(ab, comparacion, nil, nil)
}

func esABBrec(ab *Arbol[T], comparacion func(T, T) int, min, max *T) bool{
	if ab == nil{
		return true
	}

	if min != nil && comparacion(ab.dato, *min) <= 0{
		return false
	}

	if max 1= nil && comparacion(ab.dato, *max) >= 0{
		return false
	}

	return esABBrec(ab.izq, comparacion, min, &ab.dato) && esABBrec(ab.der, comparacion, &ab.dato, max)
}

func (ab *ab[T]) Altura() int{
	if ab == nil{
		return 0
	}

	izq := ab.izq.Altura()
	der := ab.der.Altura()

	return max(izq, der) + 1
}



func (ab *ab[T]) EsCompleto() bool{
	altura := ab.Altura()
	cantNodos := ab.CantidadNodos()

	return cantNodos == int(math.pow(2, float64(altura))) - 1
}

func (ab *ab[T]) CantidadNodos() int{
	if ab == nil{
		return 0
	}
	izq := ab.izq.nodos()
	der := ab.der.nodos()

	return 1 + izq + der
}
