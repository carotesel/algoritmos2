package pila

/* Definición del struct pila proporcionado por la cátedra. */

const CAPACIDAD_INICIAL = 6
const AUMENTO = 2
const DECREMENTO = 2

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

// agrego el puntero a la fc cuando modifica la pila

// extras
func (p *pilaDinamica[T]) redimensionar(nueva_capacidad int) {
	nuevo_slice := make([]T, nueva_capacidad)
	copy(nuevo_slice, p.datos)
	p.datos = nuevo_slice
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, CAPACIDAD_INICIAL)
	pila.cantidad = 0
	return pila
}

// primitivas
func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	} else {
		return p.datos[p.cantidad-1]
	}
}

func (p *pilaDinamica[T]) Apilar(item T) {
	// si pila no esta vacia:
	// si cant == len(datos)
	// 	creo una nueva pila con capacidad = cap inicial * 2
	if p.cantidad == cap(p.datos) {
		// redimensionar(capacidad_pila * FACTOR AUMENTO)
		p.redimensionar(cap(p.datos) * AUMENTO)
	}
	p.datos[p.cantidad] = item
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	// pila esta vacia: panic
	// pila no esta vacia:
	// 	tope = verTope
	//  cantidad --
	//  si cant * 4 <= capacidad:
	// 	nueva_cap = cap / 2
	//  nueva_cap >= cap_inicial

	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	elemento := p.datos[p.cantidad-1]

	//CORRECCION DEL IF
	if p.cantidad > 0 && p.cantidad*4 <= cap(p.datos) && (cap(p.datos)/DECREMENTO) >= CAPACIDAD_INICIAL {
		p.redimensionar(cap(p.datos) / DECREMENTO)
	}
	p.cantidad--

	return elemento
}
