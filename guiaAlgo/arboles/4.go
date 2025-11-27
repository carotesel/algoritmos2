/*(★★) Dado un árbol binario, escriba una primitiva recursiva que cuente la cantidad 
de nodos que tienen exactamente dos hijos directos. 
¿Qué orden de complejidad tiene la función implementada?*/

func (ab *arbol) CantidadDosHijos() int{
	if ab == nil{
		return 0
	}

	izq := ab.izq.CantidadDosHijos()
	der := ab.der.CantidadDosHijos()

	if ab.izq != nil && ab.der != nil{
		return 1 + izq + der
	}

	return izq + der
}

// Orden:
// T(n) = 2 T(n/2) + O(1)
// Log 2 (2) = 1 < C
// Compl: O(n^c) = O(n)