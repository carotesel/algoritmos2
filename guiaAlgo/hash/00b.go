package claves

func (hash *hashAbierto[K, V]) Claves() Lista[K] {
    
	lista := CrearListaEnlazada[K]()

	for _, item := range hash.tabla {
		item.Iterar(func (pareja parClaveValor[K, V]) bool{
			lista.InsertarUltimo(pareja.clave)
			return true
		})
	}

	return lista
}

// uso el iterador externo de la lista. itero y por cada elemento inserto la clave. devuelvo true para cortar iteracion