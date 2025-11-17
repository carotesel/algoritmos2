package cola

type nodoCola[T any] struct {
	dato    T
	proximo *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func crearNodo[T any](dato T) *nodoCola[T] {
	return &nodoCola[T]{dato: dato, proximo: nil}
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{primero: nil, ultimo: nil}
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

// Cuando el que desencolo es el ultimo que me quedaba, no se si deberia tambien hacer c.ultimo = nil.
// En ese caso c.primero se carga con nil por la linea 41, pero c.ultimo sigue apuntando a ese nodo que luego se pisa cuando vuelvo a encolar.
// Por las dudas lo hice por el recolector de basura, pero no se si talvez sea irrelevante
func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	retorno := c.primero.dato
	c.primero = c.primero.proximo

	if c.primero == nil {
		c.ultimo = nil
	}

	return retorno
}

func (c *colaEnlazada[T]) Encolar(dato T) {
	nuevo_nodo := crearNodo(dato)
	if c.EstaVacia() {
		c.primero = nuevo_nodo
	} else {
		c.ultimo.proximo = nuevo_nodo
	}
	c.ultimo = nuevo_nodo
}
