package claves

func Claves[K comparable, V any](d Diccionario[K, V]) Lista[K] {
    iter := d.Iterador()
	claves := CrearListaEnlazada[K]()

	for iter.HaySiguiente(){
		clave := iter.Siguiente()
		claves.InsertarUltimo(clave)
	}

	return claves
}