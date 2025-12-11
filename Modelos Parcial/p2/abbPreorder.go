/*Implementar una primitiva de ABB, que devuelva una lista con las claves del mismo tal que si las insertamos en un ABB
vacio, tendria la misma estructura q el original.*/

func (ab *ABB[K, V]) Claves() Lista[K]{
	return ab.raiz.claves()
}

func (nodo *NodoABB[K, V]) claves() Lista[K] {

	lista := CrearListaEnlazada[K]()
	nodo.clavesAux(lista)
	return lista
}

func (nodo *NodoABB[K, V]) clavesPreorder(lista Lista[K]) {

	if nodo == nil{
		return
	}

	lista.AgregarUltimo(nodo.clave)

	nodo.izq.clavesAux(lista)
	nodo.der.clavesAux(lista)
}

/*Implementar una primitiva para el ABB, que devuelva una lista con las claves del mismo, ordenadas tal que si 
insertáramos las claves en un ABB vacío, dicho ABB tendría la misma estructura que el árbol original, pero solamente 
queremos las claves que estén en el sub-árbol de un elemento E.*/

func buscar(nodo *NodoABB[K, V], clave K) *NodoABB[K, V] {
	if nodo == nil{
		return nil
	}
	if nodo.clave > clave{
		return buscar(nodo.izq, clave)
	}

	if nodo.clave < clave{
		return buscar(nodo.der, clave)
	}

	return nodo
}

func (arbol *ABB[K, V]) buscarNodo(clave K) *NodoABB[K, V] {
    nodo := buscar(arbol.raiz, clave)
	return nodo
}

func (abb *ABB[K, V]) ClavesSubarbol(elem K) Lista[K] {
    nodo := abb.buscarNodo(elem)
    if nodo == nil {
        return CrearListaEnlazada[K]() // no existe el elemento
    }
    return nodo.Claves()
}


func (nodo *NodoABB[K, V]) Claves() Lista[K] {

	lista := CrearListaEnlazada[K]()
	nodo.clavesPreorder(lista)
	return lista
}

func (nodo *NodoABB[K, V]) clavesPreorder(lista Lista[K]) {

	if nodo == nil{
		return
	}

	lista.AgregarUltimo(nodo.clave)

	nodo.izq.clavesAux(lista)
	nodo.der.clavesAux(lista)
}


