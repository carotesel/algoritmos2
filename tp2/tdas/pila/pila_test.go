package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN_TEST = 50000

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
}

func TestDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.Desapilar() })

	pila.Apilar(1)
	require.Equal(t, 1, pila.Desapilar())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.Desapilar() })
	require.True(t, pila.EstaVacia())
}

func TestVerTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.Panics(t, func() { pila.VerTope() })

	pila.Apilar(1)
	require.Equal(t, 1, pila.VerTope())

	pila.Apilar(2)
	require.Equal(t, 2, pila.VerTope())
	require.Equal(t, 2, pila.VerTope())
}

func TestApilarDesapilarApilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	require.Equal(t, 3, pila.Desapilar())
	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.Desapilar())

	require.True(t, pila.EstaVacia())

	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	require.False(t, pila.EstaVacia())
}

func TestPilaVaciaIgualPilaVaciada(t *testing.T) {
	pila_vaciada := TDAPila.CrearPilaDinamica[int]()

	pila_vaciada.Apilar(1)
	pila_vaciada.Apilar(2)
	pila_vaciada.Apilar(3)
	require.Equal(t, 3, pila_vaciada.Desapilar())
	require.Equal(t, 2, pila_vaciada.Desapilar())
	require.Equal(t, 1, pila_vaciada.Desapilar())

	require.True(t, pila_vaciada.EstaVacia())

	require.Panics(t, func() { pila_vaciada.Desapilar() })

	require.Panics(t, func() { pila_vaciada.VerTope() })
}

func TestIntegrandoEnteros(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.Equal(t, 1, pila.VerTope())
	require.Equal(t, 1, pila.VerTope())
	require.Equal(t, 1, pila.Desapilar())
	require.Panics(t, func() { pila.Desapilar() })

	pila.Apilar(1)
	pila.Apilar(2)
	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.Desapilar())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })

}
func TestIntegrandoStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()

	dias := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes"}

	for i := 0; i < 3; i++ {
		pila.Apilar(dias[i])
	}
	require.False(t, pila.EstaVacia())
	require.Equal(t, "Miércoles", pila.VerTope())
	require.Equal(t, "Miércoles", pila.Desapilar())
	require.Equal(t, "Martes", pila.VerTope())
	require.Equal(t, "Martes", pila.Desapilar())
	require.Equal(t, "Lunes", pila.Desapilar())
	require.Panics(t, func() { pila.Desapilar() })
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
}

func TestIntegrandoFloats(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float32]()

	var i float32
	for i = 0; i <= 4; i++ {
		pila.Apilar(i - (i / 2.0))
	}
	require.False(t, pila.EstaVacia())
	require.Equal(t, float32(2), pila.VerTope())
	require.Equal(t, float32(2), pila.Desapilar())
	require.Equal(t, float32(1.5), pila.VerTope())
	require.Equal(t, float32(1.5), pila.Desapilar())
	require.Equal(t, float32(1), pila.Desapilar())
	require.Equal(t, float32(0.5), pila.Desapilar())
	require.Equal(t, float32(0), pila.Desapilar())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
}

func TestIntegrandoVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i < _VOLUMEN_TEST; i++ {
		pila.Apilar(i)
	}

	for i := _VOLUMEN_TEST - 1; i > 1; i-- {
		require.Equal(t, i, pila.Desapilar())
	}
	require.Equal(t, 1, pila.Desapilar())
	require.Equal(t, 0, pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.Equal(t, 0, pila.Desapilar())
	require.Panics(t, func() { pila.Desapilar() })
	require.True(t, pila.EstaVacia())
}
