package lista

type Lista[T any] interface {
	// devuelve un booleano indicando si la lista tiene elementos o no
	EstaVacia() bool

	// inserta en la primera "posicion" de las lista
	InsertarPrimero(T)

	// inserta en la ultima "posicion" de las lista, si la lista esta vacia, lo inserta al principio.
	InsertarUltimo(T)

	// borra el primer elemento de la lista, si esta vacia entra en pánico
	BorrarPrimero() T

	// ver el primer elemento de la lista, si esta vacia entra en pánico
	VerPrimero() T

	// ver el ultimo elemento de la lista, si esta vacia entra en pánico
	VerUltimo() T

	// devuelve el largo de la lista
	Largo() int

	// itera internamente la lista, las funciones que se le pasen por parametro deben respetar la firma func(any) bool
	Iterar(visitar func(T) bool)

	// devuelve una interfaz de iterador externo
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	// devuelve el dato del elemento donde esta parado el iterador (elemento actual)
	VerActual() T

	// devuelve un booleano que indica si termine de iterar la lista
	HaySiguiente() bool

	// pasa el iterador al siguiente elemento
	Siguiente()

	// inserta un elemento luego del elemento actual y el nuevo elemento pasa a ser el actual
	Insertar(T)

	// borra el elemento actual
	Borrar() T
}
