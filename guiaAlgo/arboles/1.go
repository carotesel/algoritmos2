/*
(★) Dado un árbol binario, escribir una primitiva recursiva que determine la altura del mismo. 
Indicar y justificar el orden de la primitiva.
*/

func Max (a, b int) int{
	if a > b{
		return a
	} else{
		return b
	}
}

func (arbol *ab) Altura() int {
    if arbol == nil{
		return 0
	}

	alturaIzq := arbol.izq.Altura()
	alturaDer := arbol.der.Altura()

	return Max(alturaIzq, alturaDer) + 1
}

// Orden:
// T(n) = 2 T(n/2) + O(1)
// Log 2 (2) = 1 < C
// Compl: O(n^c) = O(n)