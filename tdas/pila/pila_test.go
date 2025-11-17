package pila_test

import (
	TDAPila "algo2/tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVaciaRecienCreada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
}

// Test 5 - Desapilar
func TestDesapilarInvalidaPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.Desapilar()
	})
}

// Test 5 - Ver Tope
func TestVerTopeInvalidoPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.VerTope()
	})
}

// Test 1
func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[bool]()

	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })

	// ver q no modifico el estado
	require.True(t, pila.EstaVacia())
}

// Test 2
func TestApilarRespetaLifo(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()

	item1 := "palabra1"
	item2 := "palabra2"
	item3 := "palabra3"

	pila.Apilar(item1)
	pila.Apilar(item2)
	pila.Apilar(item3)

	res1 := pila.Desapilar()
	res2 := pila.Desapilar()
	res3 := pila.Desapilar()

	require.Equal(t, item3, res1)
	require.Equal(t, item2, res2)
	require.Equal(t, item1, res3)
	require.True(t, pila.EstaVacia())
}

// Test 3

func TestVolumen(t *testing.T) {
	const n = 100
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i < n; i++ {
		pila.Apilar(i)
		require.False(t, pila.EstaVacia())
		require.Equal(t, i, pila.VerTope()) // tope == i
	}

	for i := n - 1; i >= 0; i-- {
		require.Equal(t, i, pila.VerTope())
		valor := pila.Desapilar()
		require.Equal(t, i, valor) // valor desapilado == i
	}

	require.True(t, pila.EstaVacia())

	// Desapilar y ver tope deben tirar panick
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

// Test 4 y 7 comprimidos
func TestPilaQueSeVacio(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()

	item1 := 3.14
	item2 := 123456.789

	pila.Apilar(item1)
	pila.Apilar(item2)

	pila.Desapilar()
	pila.Desapilar()

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

// Test 8
func TestPilaDistintosTipos(t *testing.T) {
	// Pila de enteros
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pilaInt.EstaVacia())
	pilaInt.Apilar(10)
	pilaInt.Apilar(20)
	require.Equal(t, 20, pilaInt.VerTope())
	require.Equal(t, 20, pilaInt.Desapilar())
	require.Equal(t, 10, pilaInt.Desapilar())
	require.True(t, pilaInt.EstaVacia())

	// Pila de cadenas
	pilaStr := TDAPila.CrearPilaDinamica[string]()
	pilaStr.Apilar("hola")
	pilaStr.Apilar("mundo")
	require.Equal(t, "mundo", pilaStr.VerTope())
	require.Equal(t, "mundo", pilaStr.Desapilar())
	require.Equal(t, "hola", pilaStr.Desapilar())
	require.True(t, pilaStr.EstaVacia())

	// Pila de floats
	pilaFloat := TDAPila.CrearPilaDinamica[float64]()
	pilaFloat.Apilar(3.14)
	pilaFloat.Apilar(2.71)
	require.Equal(t, 2.71, pilaFloat.Desapilar())
	require.Equal(t, 3.14, pilaFloat.VerTope())

	// Pila de structs
	type Persona struct {
		Nombre string
		Edad   int
	}
	pilaStruct := TDAPila.CrearPilaDinamica[Persona]()
	juan := Persona{"Juan", 30}
	ana := Persona{"Ana", 25}

	pilaStruct.Apilar(juan)
	pilaStruct.Apilar(ana)
	require.Equal(t, ana, pilaStruct.Desapilar())
	require.Equal(t, juan, pilaStruct.VerTope())
}

// CORRECCION: AGREGO TEST INTERCALADO
func TestIntercalado(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[bool]()

	pila.Apilar(true)
	pila.Apilar(false)
	pila.Apilar(false)

	require.False(t, pila.EstaVacia())
	require.Equal(t, false, pila.VerTope())

	pila.Apilar(true)
	require.Equal(t, true, pila.VerTope())
	pila.Apilar(true)

	require.False(t, pila.EstaVacia())
	require.Equal(t, true, pila.VerTope())

	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()
	resultado := pila.Desapilar()
	require.Equal(t, true, resultado)
	require.True(t, pila.EstaVacia())
}
