package claves

func (hash *hashCerrado[K, V]) Claves() Lista[K] {
    
	lista := CrearListaEnlazada[K]()

	for _, item := range hash.tabla{
		if item.estado == OCUPADO{
			lista.InsertarUltimo(item.clave)
		}
	}

	return lista
}