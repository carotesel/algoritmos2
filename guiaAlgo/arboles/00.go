package mayores

func (abb *abb[K, V]) Mayores(clave K) Lista[K] {
    lista := CrearListaEnlazada[K]()
	mayoresAux(abb.raiz, clave, abb.cmp, lista)
    return lista
}

func mayoresAux[K comparable, V any](nodo *nodoAbb[K, V], clave K, cmp funcCmp[K], lista Lista[K]){
	if nodo == nil{
		return
	}

	// si el nodo es mayor a la clave pasada, recorro todo porque podrian haber cosas en izq y der que sean mayores tmb
	// si nodo es <, si o si voy a la derecha
	if cmp(nodo.clave, clave) > 0{
		mayoresAux(nodo.izquierdo, clave, cmp, lista)
		lista.InsertarUltimo(nodo.clave)
		mayoresAux(nodo.derecho, clave, cmp, lista)
	} else{
		mayoresAux(nodo.derecho, clave, cmp, lista)
	}
}