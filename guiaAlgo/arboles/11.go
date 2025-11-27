/*(★★) Implementar una primitiva para el ABB, que reciba el ABB y devuelva una lista 
con las claves del mismo, ordenadas tal que si insertáramos las claves en un ABB vacío 
(con la misma función de comparación), dicho ABB tendría la misma estructura que el árbol original. 

¿Qué tipo de recorrido utilizaste? Indicar y justificar el orden de la primitiva.*/

// PREORDER -> CONSERVA LA ESTRUCTURA

func (abb *abb[K, V]) Claves() Lista[K] {
    lista := CrearListaEnlazada[K]()
	ClavesNodo(abb.raiz, lista)
	return lista
}

func (nodo *nodoAbb[K, V]) ClavesNodo(n *nodoAbb[K, V], lista Lista[K]) {
	if n == nil{
		return
	}

	// Raiz
	lista.AgregarUltimo(nodo.clave)

	// Subarbol izq
	ClavesNodo(n.izquierdo, lista)

	// subarbol der
	ClavesNodo(n.derecho, lista)
}

// Complejidad: O(n)