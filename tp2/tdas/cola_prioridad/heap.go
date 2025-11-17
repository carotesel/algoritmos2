package cola_prioridad

const (
	_TAMANIO_INICIAL       = 2
	_FACTOR_DE_REDIMENCION = 2
	_CARGA_MINIMA          = 4
	_MENSAJE_ERROR         = "La cola esta vacia"
)

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(a, b T) int
}

func CrearHeap[T any](funcion_cmp func(a, b T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: make([]T, _TAMANIO_INICIAL),
		cant:  0,
		cmp:   funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	largo := len(arreglo)
	if largo < _TAMANIO_INICIAL {
		largo = _TAMANIO_INICIAL
	}
	nuevo := make([]T, largo)
	copy(nuevo, arreglo)

	heapify(nuevo, funcion_cmp)

	heapMax := &colaConPrioridad[T]{
		datos: nuevo,
		cant:  len(arreglo),
		cmp:   funcion_cmp,
	}

	return heapMax
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {

	heapify(elementos, funcion_cmp)

	for i := len(elementos) - 1; i > 0; i-- {
		swap(&elementos[0], &elementos[i])
		downHeap(elementos, 0, i, funcion_cmp)
	}
}

func (heapMax *colaConPrioridad[T]) EstaVacia() bool {
	return heapMax.cant == 0
}

func (heapMax *colaConPrioridad[T]) Cantidad() int {
	return heapMax.cant
}

func (heapMax *colaConPrioridad[T]) VerMax() T {
	if heapMax.EstaVacia() {
		panic(_MENSAJE_ERROR)
	}
	return heapMax.datos[0]
}

func (heapMax *colaConPrioridad[T]) Encolar(dato T) {
	if cap(heapMax.datos) == heapMax.cant {
		heapMax.redimencion(_FACTOR_DE_REDIMENCION * cap(heapMax.datos))
	}
	heapMax.datos[heapMax.cant] = dato
	upHeap(heapMax.datos, heapMax.cant, heapMax.cmp)
	heapMax.cant++

}

func (heapMax *colaConPrioridad[T]) Desencolar() T {
	if heapMax.EstaVacia() {
		panic(_MENSAJE_ERROR)
	}

	retorno := heapMax.datos[0]
	heapMax.cant--
	heapMax.datos[0] = heapMax.datos[heapMax.cant]
	downHeap(heapMax.datos, 0, heapMax.cant, heapMax.cmp)

	if cap(heapMax.datos) >= _CARGA_MINIMA*heapMax.cant && heapMax.cant > _TAMANIO_INICIAL {
		heapMax.redimencion(cap(heapMax.datos) / _FACTOR_DE_REDIMENCION)
	}

	return retorno
}

// primitiva auxiliar privada
func (heapMax *colaConPrioridad[T]) redimencion(nuevo_tam int) {
	nuevo := make([]T, nuevo_tam)
	copy(nuevo, heapMax.datos[:heapMax.cant])
	heapMax.datos = nuevo
}

// auxiliares privadas
func indicePadre(i int) int {
	return (i - 1) / 2
}

func indiceHijoIzq(i int) int {
	return (2 * i) + 1
}

func indiceHijoDer(i int) int {
	return (2 * i) + 2
}

func swap[T any](padre, hijo *T) {
	*padre, *hijo = *hijo, *padre
}

func downHeap[T any](arr []T, i, cant int, cmp func(a, b T) int) {
	hijoIzq := indiceHijoIzq(i)
	hijoDer := indiceHijoDer(i)

	hijoMax := i
	if hijoIzq < cant && cmp(arr[hijoIzq], arr[i]) > 0 {
		hijoMax = hijoIzq
	}
	if hijoDer < cant && cmp(arr[hijoDer], arr[hijoMax]) > 0 {
		hijoMax = hijoDer
	}

	if hijoMax != i {
		swap(&arr[i], &arr[hijoMax])
		downHeap(arr, hijoMax, cant, cmp)
	}
}

func upHeap[T any](arr []T, i int, cmp func(a, b T) int) {
	if i == 0 {
		return
	}
	padre := indicePadre(i)
	if cmp(arr[padre], arr[i]) < 0 {
		swap(&arr[i], &arr[padre])
		upHeap(arr, padre, cmp)
	}
}

func heapify[T any](arr []T, cmp func(a, b T) int) {
	largo := len(arr)
	if largo <= 1 {
		return
	}

	for i := (largo / 2) - 1; i >= 0; i-- {
		downHeap(arr, i, largo, cmp)
	}
}
