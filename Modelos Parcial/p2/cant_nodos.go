// Cantidad Nodos 

// arbol es nodo
func (ab *Arbol) CantidadNodos() int{
	if ab == nil{
		return 0
	}

	izq := ab.izq.CantidadNodos()
	der := ab.der.CantidadNodos()

	return 1 + izq + der
}

// arbol no es nodo
func (ab *Arbol) CantidadNodos() int{
	return cantidad(ab.raiz)
}

func cantidad(nodo *nodo) int{
	if nodo == nil{
		return 0
	}

	return 1 + cantidad(nodo.izq) + cantidad(nodo.der)
}