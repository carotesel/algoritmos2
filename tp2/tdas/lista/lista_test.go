package lista_test

import (
	"strconv"
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN_TEST = 50000

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
}

func TestVerPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	require.Panics(t, func() { lista.VerPrimero() })

	lista.InsertarPrimero(2)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.VerPrimero())
	require.False(t, lista.EstaVacia())

	lista.InsertarUltimo(3)
	require.Equal(t, 2, lista.VerPrimero())

	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerPrimero())
}

func TestVerPrimeroYUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	require.Panics(t, func() { lista.VerPrimero() })

	lista.InsertarUltimo(2)
	require.Equal(t, 2, lista.VerPrimero())
}

func TestVerUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	require.Panics(t, func() { lista.VerPrimero() })

	lista.InsertarPrimero(2)
	require.Equal(t, 2, lista.VerUltimo())

	lista.InsertarPrimero(1)
	require.Equal(t, 2, lista.VerUltimo())

	lista.InsertarUltimo(3)
	require.Equal(t, 3, lista.VerUltimo())
}

func TestBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	require.Panics(t, func() { lista.BorrarPrimero() })

	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.BorrarPrimero())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.True(t, lista.EstaVacia())

	lista.InsertarUltimo(3)
	require.Equal(t, 3, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 2, lista.VerPrimero())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 3, lista.BorrarPrimero())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.True(t, lista.EstaVacia())
}

func TestLargo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	require.Equal(t, 2, lista.Largo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 1, lista.Largo())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia())

	require.Panics(t, func() { lista.BorrarPrimero() })
}

func TestIteradorInternoCompleto(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	suma := 0
	lista.Iterar(func(i int) bool {
		suma += i
		return true
	})
	require.Equal(t, 0, suma)

	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	lista.Iterar(func(i int) bool {
		suma += i
		return true
	})
	require.Equal(t, 6, suma)
}

func TestIteradorInternoConCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	suma := 0
	cont := 0

	lista.Iterar(func(i int) bool {
		suma += i
		cont++
		return cont != 2
	})
	require.Equal(t, 0, suma)

	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	suma = 0
	lista.Iterar(func(i int) bool {
		suma += i
		cont++
		return cont != 2
	})
	require.Equal(t, 3, suma)

	suma = 0
	lista.Iterar(func(i int) bool {
		suma += i
		return i%2 == 0
	})
	require.Equal(t, 1, suma)
}

func TestIntegrandoVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i <= _VOLUMEN_TEST; i++ {
		lista.InsertarPrimero(i)
	}

	require.False(t, lista.EstaVacia())
	require.Equal(t, 0, lista.VerUltimo())
	require.Equal(t, _VOLUMEN_TEST, lista.VerPrimero())

	for i := _VOLUMEN_TEST; i > 0; i-- {
		require.Equal(t, i, lista.BorrarPrimero())
	}

	require.False(t, lista.EstaVacia())
	require.Equal(t, 0, lista.BorrarPrimero())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.True(t, lista.EstaVacia())

	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.BorrarPrimero())
	require.Panics(t, func() { lista.BorrarPrimero() })
}

func TestStrings(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	for i := 0; i <= _VOLUMEN_TEST; i++ {
		lista.InsertarPrimero(strconv.Itoa(i))
	}

	require.False(t, lista.EstaVacia())
	require.Equal(t, "0", lista.VerUltimo())
	require.Equal(t, strconv.Itoa(_VOLUMEN_TEST), lista.VerPrimero())

	for i := _VOLUMEN_TEST; i > 0; i-- {
		require.Equal(t, strconv.Itoa(i), lista.BorrarPrimero())
	}

	require.False(t, lista.EstaVacia())
	require.Equal(t, strconv.Itoa(0), lista.BorrarPrimero())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.True(t, lista.EstaVacia())

	lista.InsertarPrimero(strconv.Itoa(1))
	require.Equal(t, strconv.Itoa(1), lista.BorrarPrimero())
	require.Panics(t, func() { lista.BorrarPrimero() })
}

func TestListaVaciaItero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	iter := lista.Iterador()

	iter.Insertar(1)

	require.Equal(t, 1, iter.VerActual())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
}

func TestListaIteroHastaFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	iter := lista.Iterador()

	for iter.HaySiguiente() {
		iter.Siguiente()
	}

	iter.Insertar(1)
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.VerPrimero())

	iter.Insertar(2)
	iter.Insertar(3)
	iter.Insertar(4)
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
}

func TestListaIteroInsertoInicio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	iter := lista.Iterador()

	iter.Insertar(1)
	iter.Insertar(2)
	iter.Insertar(3)
	iter.Insertar(4)

	for iter.HaySiguiente() {
		iter.Siguiente()
	}

	require.Equal(t, 4, lista.BorrarPrimero())
	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 1, lista.BorrarPrimero())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.True(t, lista.EstaVacia())
}

func TestListaIteroInsertoMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	lista.InsertarPrimero(5)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(1)

	require.Equal(t, 4, lista.Largo())

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Insertar(2)
	require.Equal(t, 2, iter.VerActual())

	require.Equal(t, 5, lista.Largo())

	iter = lista.Iterador()
	iter.Siguiente()
	require.Equal(t, 2, iter.VerActual())
}

func TestListaIteroBorroPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	require.Equal(t, 4, lista.Largo())

	iter := lista.Iterador()
	require.Equal(t, 1, iter.Borrar())
	require.Equal(t, 2, lista.VerPrimero())
}

func TestListaIteroBorroUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	require.Equal(t, 4, lista.Largo())

	cont := 1
	iter := lista.Iterador()
	for iter.HaySiguiente() && cont < 4 {
		iter.Siguiente()
		cont++
	}

	require.Equal(t, 4, cont)
	require.Equal(t, 4, lista.VerUltimo())
	require.Equal(t, 4, iter.Borrar())
	require.Equal(t, 3, lista.VerUltimo())
}

func TestListaIteroBorroMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	iter := lista.Iterador()

	iter.Insertar(1)
	iter.Insertar(2)
	iter.Insertar(3)
	iter.Insertar(4)
	iter.Insertar(5)

	require.Equal(t, 5, lista.Largo())

	iter = lista.Iterador()
	for i := 0; i < lista.Largo()/2; i++ {
		iter.Siguiente()
	}

	require.Equal(t, 3, iter.Borrar())

	require.Equal(t, 4, lista.Largo())

	no_esta := true
	lista.Iterar(func(i int) bool {
		no_esta = i != 3
		return no_esta
	})
	require.Equal(t, true, no_esta)
}

func TestListaIteroBorroUnico(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	lista.InsertarPrimero(1)
	iter := lista.Iterador()

	require.Equal(t, 1, iter.Borrar())
	require.Panics(t, func() { iter.Borrar() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
}

func TestListaIteroTodo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	cantidad_vueltas := 0

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		cantidad_vueltas++
		iter.Siguiente()
	}

	require.Equal(t, 4, cantidad_vueltas)
}
