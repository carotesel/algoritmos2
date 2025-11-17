package cola_prioridad_test

import (
	"math/rand"
	"strconv"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN_TEST int = 50000

func cmpEnteros(a, b int) int {
	return a - b
}

func TestHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
}

func TestHeapEncolarUnElemento(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)
	heap.Encolar(10)
	require.False(t, heap.EstaVacia())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 10, heap.VerMax())
}

func TestHeapEncolarVariosElementos(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)
	elementos := []int{3, 7, 5, 10, 1}

	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 10, heap.VerMax())
}

func TestHeapDesencolar(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)
	elementos := []int{4, 9, 2, 7}
	for _, e := range elementos {
		heap.Encolar(e)
	}

	require.Equal(t, 9, heap.Desencolar())
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())
}

func TestHeapDesencolarHastaVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)
	require.Panics(t, func() { heap.Desencolar() })

	elementos := []int{1, 2, 3}
	for _, e := range elementos {
		heap.Encolar(e)
	}

	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 1, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.Panics(t, func() { heap.Desencolar() })
}

func TestHeapCrearDesdeArreglo(t *testing.T) {
	arr := []int{5, 1, 8, 3, 7}
	heap := TDAHeap.CrearHeapArr[int](arr, cmpEnteros)

	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 8, heap.VerMax())

	heap.Encolar(9)
	require.Equal(t, 6, heap.Cantidad())
	require.Equal(t, 9, heap.VerMax())

	require.Equal(t, 9, heap.Desencolar())
	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 8, heap.VerMax())
}

func TestHeapCrearDesdeArregloVacio(t *testing.T) {
	arr := []int{}
	heap := TDAHeap.CrearHeapArr[int](arr, cmpEnteros)

	require.Equal(t, 0, heap.Cantidad())
	require.Panics(t, func() { heap.Desencolar() })
	require.Panics(t, func() { heap.VerMax() })

	heap.Encolar(1)
	heap.Encolar(0)
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 1, heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 0, heap.VerMax())
}

func TestHeapSortOrdenaCorrectamente(t *testing.T) {
	arr := []int{5, 3, 8, 1, 2}
	TDAHeap.HeapSort(arr, cmpEnteros)

	require.Equal(t, []int{1, 2, 3, 5, 8}, arr)
}

func TestHeapVerMaxVacioPanic(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)
	require.Panics(t, func() { heap.VerMax() })
}

func TestHeapDeStrings(t *testing.T) {
	comparoStrings := func(s1, s2 string) int {
		return strings.Compare(s1, s2)
	}

	heap := TDAHeap.CrearHeap[string](comparoStrings)
	require.True(t, heap.EstaVacia())
	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })

	arr := []string{"2", "3", "5", "1", "4"}
	for _, e := range arr {
		heap.Encolar(e)
	}

	heapArr := TDAHeap.CrearHeapArr(arr, comparoStrings)

	require.False(t, heap.EstaVacia())
	require.False(t, heap.EstaVacia())

	TDAHeap.HeapSort(arr, comparoStrings)

	for i := 4; i >= 0; i-- {
		require.Equal(t, arr[i], heap.VerMax())
		require.Equal(t, arr[i], heapArr.VerMax())
		require.Equal(t, arr[i], heap.Desencolar())
		require.Equal(t, arr[i], heapArr.Desencolar())
	}

	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })
}

func TestHeapDeStructs(t *testing.T) {
	type elefante struct {
		nombre string
		edad   int
	}

	comparoElefantes := func(e1, e2 elefante) int {
		comp := -1 * strings.Compare(e1.nombre, e2.nombre)
		if comp == 0 {
			comp = e1.edad - e2.edad
		}
		return comp
	}

	heap := TDAHeap.CrearHeap[elefante](comparoElefantes)
	require.True(t, heap.EstaVacia())
	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })

	arr := []elefante{
		elefante{nombre: "Alan", edad: 52},
		elefante{nombre: "Messi", edad: 38},
		elefante{nombre: "Barbara", edad: 52},
		elefante{nombre: "Alan", edad: 90},
		elefante{nombre: "Cristian", edad: 27},
	}

	arrOrdenado := []elefante{
		elefante{nombre: "Messi", edad: 38},
		elefante{nombre: "Cristian", edad: 27},
		elefante{nombre: "Barbara", edad: 52},
		elefante{nombre: "Alan", edad: 52},
		elefante{nombre: "Alan", edad: 90},
	}

	for _, e := range arr {
		heap.Encolar(e)
	}

	heapArr := TDAHeap.CrearHeapArr(arr, comparoElefantes)

	require.False(t, heap.EstaVacia())
	require.False(t, heap.EstaVacia())

	TDAHeap.HeapSort(arr, comparoElefantes)

	require.Equal(t, arr, arrOrdenado)

	for i := 4; i >= 0; i-- {
		require.Equal(t, arr[i], heap.VerMax())
		require.Equal(t, arr[i], heapArr.VerMax())
		require.Equal(t, arr[i], heap.Desencolar())
		require.Equal(t, arr[i], heapArr.Desencolar())
	}

	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })
}

func TestHeapArrVolumenStrings(t *testing.T) {
	comparoStrings := func(s1, s2 string) int {
		return strings.Compare(s1, s2)
	}

	heap := TDAHeap.CrearHeap[string](comparoStrings)
	require.True(t, heap.EstaVacia())
	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })

	arr := make([]string, _VOLUMEN_TEST)
	for i := 0; i < _VOLUMEN_TEST; i++ {
		arr[i] = strconv.Itoa(i)
		heap.Encolar(strconv.Itoa(i))
	}

	heapArr := TDAHeap.CrearHeapArr(arr, comparoStrings)

	require.False(t, heap.EstaVacia())
	require.False(t, heap.EstaVacia())

	TDAHeap.HeapSort(arr, comparoStrings)

	for i := _VOLUMEN_TEST - 1; i >= 0; i-- {
		require.Equal(t, arr[i], heap.VerMax())
		require.Equal(t, arr[i], heapArr.VerMax())
		require.Equal(t, arr[i], heap.Desencolar())
		require.Equal(t, arr[i], heapArr.Desencolar())
	}

	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })
}

var random = rand.New(rand.NewSource(5454545))

func desordenar(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func TestHeapCrearDesdeArregloVolumen(t *testing.T) {
	arr := make([]int, _VOLUMEN_TEST)
	for i := 0; i < _VOLUMEN_TEST; i++ {
		arr[i] = i
	}

	desordenar(arr)

	heapArr := TDAHeap.CrearHeapArr(arr, cmpEnteros)

	require.False(t, heapArr.EstaVacia())

	TDAHeap.HeapSort(arr, cmpEnteros)

	for i := _VOLUMEN_TEST - 1; i >= 0; i-- {
		require.Equal(t, arr[i], heapArr.VerMax())
		require.Equal(t, arr[i], heapArr.Desencolar())
	}

	require.Panics(t, func() { heapArr.VerMax() })
	require.Panics(t, func() { heapArr.Desencolar() })
}

func TestHeapCrearDesdeArregloDeSringsVolumen(t *testing.T) {
	comparoStrings := func(s1, s2 string) int {
		return strings.Compare(s1, s2)
	}

	arrStings := make([]string, _VOLUMEN_TEST)
	arrInt := make([]int, _VOLUMEN_TEST)
	for i := 0; i < _VOLUMEN_TEST; i++ {
		arrInt[i] = i
	}

	desordenar(arrInt)

	for i := 0; i < _VOLUMEN_TEST; i++ {
		arrStings[i] = strconv.Itoa(arrInt[i])
	}

	heapArr := TDAHeap.CrearHeapArr(arrStings, comparoStrings)

	require.False(t, heapArr.EstaVacia())

	TDAHeap.HeapSort(arrStings, comparoStrings)

	for i := _VOLUMEN_TEST - 1; i >= 0; i-- {
		require.Equal(t, arrStings[i], heapArr.VerMax())
		require.Equal(t, arrStings[i], heapArr.Desencolar())
	}

	require.Panics(t, func() { heapArr.VerMax() })
	require.Panics(t, func() { heapArr.Desencolar() })
}
