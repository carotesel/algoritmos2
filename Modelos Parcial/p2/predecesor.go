/*
1. Implementar una primitiva para el ABB Predecesor(clave K) K que reciba una clave
(que puede estar en el árbol, o no) y devuelva la clave inmediatamente anterior a esta en el
recorrido inorder. Si no hay ninguna anterior, finalizar con un panic. Indicar y justificar
la complejidad de la primitiva.
*/

func (ab *Abb[K, V]) Predecesor(clave K) K{
	// busco nodo
	nodo := buscar(ab.raiz, clave)

	// nodo vacio: no existe clave
	if nodo == nil{
		panic("Clave no existe")
	}

	// hay nodo.izq -> mayor del subarbol
	if nodo.izq != nil{
		return max(nodo.izq)
	}

	// no hay nodo.izq -> mayor desde raiz
	return ultimoMenor(ab.raiz, clave)
}

func buscar(nodo *nodoABB[K, V], clave K) *nodoABB[K, V]{
	if nodo == nil{
		return nil
	}

	if clave < nodo.clave{
		return buscar(nodo.izq, clave)
	}

	if clave > nodo.clave{
		return buscar(nodo.der, clave)
	}
	return nodo
}

func max(nodo *nodoABB[K, V]) *nodoABB[K, V]{
	for nodo.der != nil{
		nodo = nodo.der
	}
	return nodo.clave
}

func ultimoMenor(nodo *nodoABB[K, V], clave K) K{
	var candidato **nodoABB[K, V]
	
	for nodo != nil{
		if clave > nodo.clave{
			candidato = nodo
			nodo = nodo.der
		} else {
			nodo = nodo.izq
		}
	}

	if candidato == nil{
		panic("no tiene predecesor")
    }
    return candidato.clave
}

