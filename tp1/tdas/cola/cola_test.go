package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVaciaRecienCreada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	require.True(t, cola.EstaVacia())
}

func TestColaRecienCreada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.Desencolar() })
	require.Panics(t, func() { cola.VerPrimero() })
}

func TestDesencolarRespetaFifo(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()

	item1 := "palabra1"
	item2 := "palabra2"
	item3 := "palabra3"

	cola.Encolar(item1)
	cola.Encolar(item2)
	cola.Encolar(item3)

	res1 := cola.Desencolar()
	res2 := cola.Desencolar()
	res3 := cola.Desencolar()

	require.Equal(t, item1, res1)
	require.Equal(t, item2, res2)
	require.Equal(t, item3, res3)
	require.True(t, cola.EstaVacia())
}

func TestColaQueSeVacio(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[bool]()

	item1 := true
	item2 := false
	item3 := false

	cola.Encolar(item1)
	cola.Encolar(item2)
	cola.Encolar(item3)

	cola.Desencolar()
	cola.Desencolar()
	cola.Desencolar()

	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.Desencolar() })
	require.Panics(t, func() { cola.VerPrimero() })
}

func TestColaIntercalado(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[bool]()

	cola.Encolar(true)
	cola.Encolar(false)
	cola.Encolar(false)

	require.False(t, cola.EstaVacia())
	require.Equal(t, true, cola.VerPrimero())

	cola.Encolar(true)
	require.Equal(t, true, cola.VerPrimero())
	cola.Encolar(true)

	require.False(t, cola.EstaVacia())
	require.Equal(t, true, cola.VerPrimero())

	require.Equal(t, true, cola.Desencolar()) // primero en salir
	require.Equal(t, false, cola.Desencolar())
	require.Equal(t, false, cola.Desencolar())
	require.Equal(t, true, cola.Desencolar())
	resultado := cola.Desencolar()
	require.Equal(t, true, resultado)

	require.True(t, cola.EstaVacia())
}

func TestColaDistintosTipos(t *testing.T) {
	// Cola de enteros
	colaInt := TDACola.CrearColaEnlazada[int]()
	require.True(t, colaInt.EstaVacia())
	colaInt.Encolar(10)
	colaInt.Encolar(20)
	require.Equal(t, 10, colaInt.VerPrimero())
	require.Equal(t, 10, colaInt.Desencolar())
	require.Equal(t, 20, colaInt.Desencolar())
	require.True(t, colaInt.EstaVacia())

	// Cola de cadenas
	colaStr := TDACola.CrearColaEnlazada[string]()
	colaStr.Encolar("hola")
	colaStr.Encolar("mundo")
	require.Equal(t, "hola", colaStr.VerPrimero())
	require.Equal(t, "hola", colaStr.Desencolar())
	require.Equal(t, "mundo", colaStr.Desencolar())
	require.True(t, colaStr.EstaVacia())

	// Cola de floats
	colaFloat := TDACola.CrearColaEnlazada[float64]()
	colaFloat.Encolar(3.14)
	colaFloat.Encolar(2.71)
	require.Equal(t, 3.14, colaFloat.Desencolar())
	require.Equal(t, 2.71, colaFloat.VerPrimero())

	// Cola de structs
	type Persona struct {
		Nombre string
		Edad   int
	}
	colaStruct := TDACola.CrearColaEnlazada[Persona]()
	juan := Persona{"Juan", 30}
	ana := Persona{"Ana", 25}

	colaStruct.Encolar(juan)
	colaStruct.Encolar(ana)
	require.Equal(t, juan, colaStruct.Desencolar())
	require.Equal(t, ana, colaStruct.VerPrimero())
}

// CORRECCION: AUMENTO VOLUMEN DEL TEST
func TestColaVolumen1000(t *testing.T) {
	const n = 1000
	cola := TDACola.CrearColaEnlazada[int]()

	// Encolar muchos elementos y verificar el primero
	for i := 0; i < n; i++ {
		cola.Encolar(i)
		require.False(t, cola.EstaVacia())
		require.Equal(t, 0, cola.VerPrimero()) // siempre debe ser el primer encolado (FIFO)
	}

	// Desencolar y verificar orden correcto
	for i := 0; i < n; i++ {
		require.Equal(t, i, cola.VerPrimero())
		valor := cola.Desencolar()
		require.Equal(t, i, valor) // debe salir en orden
	}

	require.True(t, cola.EstaVacia())

	// VerPrimero y Desencolar deben tirar pánico
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestColaVolumen10000(t *testing.T) {
	const n = 10000
	cola := TDACola.CrearColaEnlazada[int]()

	// Encolar muchos elementos y verificar el primero
	for i := 0; i < n; i++ {
		cola.Encolar(i)
		require.False(t, cola.EstaVacia())
		require.Equal(t, 0, cola.VerPrimero()) // siempre debe ser el primer encolado (FIFO)
	}

	// Desencolar y verificar orden correcto
	for i := 0; i < n; i++ {
		require.Equal(t, i, cola.VerPrimero())
		valor := cola.Desencolar()
		require.Equal(t, i, valor) // debe salir en orden
	}

	require.True(t, cola.EstaVacia())

	// VerPrimero y Desencolar deben tirar pánico
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}
