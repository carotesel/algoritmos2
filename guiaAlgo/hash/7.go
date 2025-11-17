package multiconj

type multiConj[K comparable] struct {
    datos Diccionario[K, int]
}

func CrearMulticonjunto[K comparable]() MultiConjunto[K] {
    return &multiConj[K]{ // devolvemos un puntero para que los métodos modifiquen el struct
		datos: CrearHash[K, int](), // usás tu función para crear un diccionario vacío
	}
}

func (conj *multiConj[K]) Guardar(elem K) {
    if conj.datos.Pertenece(elem){
        cantidad := conj.datos.Obtener(elem)
        conj.datos.Guardar(elem, cantidad+1)
    } else{
        conj.datos.Guardar(elem, 1)
    }
}

func (conj *multiConj[K]) Pertenece(elem K) bool {
    return conj.datos.Pertenece(elem)
}

func (conj *multiConj[K]) Borrar(elem K) {

    if !conj.datos.Pertenece(elem){
        panic("Elemento no esta en el multiconjunto")
    }

    cantidad := conj.datos.Obtener(elem)

    if cantidad > 1 {
        conj.datos.Guardar(elem, cantidad-1)
    } else{
        conj.datos.Borrar(elem)
    }
}