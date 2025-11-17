package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const (
	_TAMAÑO_INICIAL        = 2
	_FACTOR_DE_REDIMENCION = 2
	_CARGA_MINIMA          = 4
)

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, _TAMAÑO_INICIAL),
		cantidad: 0,
	}
}

func (p *pilaDinamica[T]) redimencion(nuevo_tam int) {
	nuevo := make([]T, nuevo_tam)
	copy(nuevo, p.datos[:p.cantidad])
	p.datos = nuevo
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) Apilar(dato T) {
	if cap(p.datos) == p.cantidad {
		p.redimencion(_FACTOR_DE_REDIMENCION * cap(p.datos))
	}
	p.datos[p.cantidad] = dato
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	p.cantidad--
	retorno := p.datos[p.cantidad]

	if cap(p.datos) >= _CARGA_MINIMA*p.cantidad && p.cantidad > _TAMAÑO_INICIAL {
		p.redimencion(cap(p.datos) / _FACTOR_DE_REDIMENCION)
	}

	return retorno
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}
