package lista

const (
	_MENSAJE_ERROR_LISTA    = "La lista esta vacia"
	_MENSAJE_ERROR_ITERADOR = "El iterador termino de iterar"
)

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func crearNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato: dato, siguiente: nil}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevo_nodo := crearNodo(dato)

	nuevo_nodo.siguiente = lista.primero

	if lista.EstaVacia() {
		lista.ultimo = nuevo_nodo
	}

	lista.largo++
	lista.primero = nuevo_nodo
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevo_nodo := crearNodo(dato)

	if lista.EstaVacia() {
		lista.primero = nuevo_nodo
	} else {
		lista.ultimo.siguiente = nuevo_nodo
	}

	lista.largo++
	lista.ultimo = nuevo_nodo
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_ERROR_LISTA)
	}
	dato := lista.primero.dato

	if lista.primero == lista.ultimo {
		lista.ultimo = nil
	}

	lista.primero = lista.primero.siguiente
	lista.largo--

	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_ERROR_LISTA)
	}

	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_ERROR_LISTA)
	}

	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	resultado := true
	actual := lista.primero

	for actual != nil && resultado {
		resultado = visitar(actual.dato)
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return creariteradorlistaEnlazada(lista)
}

func creariteradorlistaEnlazada[T any](lista *listaEnlazada[T]) IteradorLista[T] {
	return &iterListaEnlazada[T]{actual: lista.primero, anterior: nil, lista: lista}
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_ERROR_ITERADOR)
	}
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_ERROR_ITERADOR)
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iterListaEnlazada[T]) Insertar(dato T) {
	nuevo_nodo := crearNodo(dato)

	nuevo_nodo.siguiente = iter.actual

	if iter.anterior == nil {
		iter.lista.primero = nuevo_nodo
	} else {
		iter.anterior.siguiente = nuevo_nodo
	}

	if iter.actual == nil {
		iter.lista.ultimo = nuevo_nodo
	}

	iter.lista.largo++
	iter.actual = nuevo_nodo
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_ERROR_ITERADOR)
	}

	dato := iter.actual.dato

	if iter.actual != iter.lista.primero {
		iter.anterior.siguiente = iter.actual.siguiente
	} else {
		iter.lista.primero = iter.actual.siguiente
	}

	if iter.actual == iter.lista.ultimo {
		iter.lista.ultimo = iter.anterior
	}

	iter.actual = iter.actual.siguiente
	iter.lista.largo--

	return dato
}
