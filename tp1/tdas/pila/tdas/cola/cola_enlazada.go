package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

// extras
// CORRECCION: TEST EN 1 LINEA
func CrearColaEnlazada[T any]() Cola[T] { return &colaEnlazada[T]{} }

// nuevoNodo crea un nodo de cola con el dato dado.
func nuevoNodo[T any](dato T) *nodoCola[T] {
	return &nodoCola[T]{dato: dato}
}

// primitivas
func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

// Agrega un nuevo elemento a la cola, al final de la misma.
func (c *colaEnlazada[T]) Encolar(item T) {
	nodo := nuevoNodo(item)
	if c.EstaVacia() {
		c.primero = nodo
	} else {
		c.ultimo.prox = nodo // nodo ya tiene prox == nil
	}
	c.ultimo = nodo
}

// saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
//
//	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := c.primero.dato
	c.primero = c.primero.prox

	// si la cola se vacio
	if c.primero == nil {
		c.ultimo = nil
	}

	return dato
}

// primitivas
func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}
