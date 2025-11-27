/*(★) Implementar una primitiva que devuelva la suma de todos los datos (números) 
de un árbol binario. Indicar y justificar el orden de la primitiva.*/

func (arbol *ab) SumaDatos(){
	if arbol == nil{
		return 0
	}

	sumaIzq := arbol.izq.SumaDatos()
	sumaDer := arbol.der.SumaDatos()

	return arbol.dato + sumaIzq + sumaDer
}

// Orden:
// T(n) = 2 T(n/2) + O(1)
// Log 2 (2) = 1 < C
// Compl: O(n^c) = O(n)