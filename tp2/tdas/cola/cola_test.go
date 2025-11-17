package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN_TEST = 50000

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
}

func TestDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.Desencolar() })

	cola.Encolar(1)
	cola.Encolar(2)
	require.Equal(t, 1, cola.Desencolar())
	require.Equal(t, 2, cola.Desencolar())
	require.Panics(t, func() { cola.Desencolar() })
	require.True(t, cola.EstaVacia())
}

func TestVerPrimero(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.Panics(t, func() { cola.VerPrimero() })

	cola.Encolar(1)
	cola.Encolar(2)
	require.Equal(t, 1, cola.VerPrimero())
	require.Equal(t, 1, cola.VerPrimero())
	require.Equal(t, 1, cola.Desencolar())
	require.Equal(t, 2, cola.VerPrimero())
	require.Equal(t, 2, cola.Desencolar())
	require.Panics(t, func() { cola.VerPrimero() })
}

func TestEncolarDesencolarYEncolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.Panics(t, func() { cola.Desencolar() })

	cola.Encolar(1)
	cola.Encolar(2)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.Desencolar())
	require.Equal(t, 2, cola.Desencolar())
	require.True(t, cola.EstaVacia())

	cola.Encolar(3)
	cola.Encolar(4)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 3, cola.Desencolar())
	require.Equal(t, 4, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestColaVaciaIgualColaVaciada(t *testing.T) {
	cola_vaciada := TDACola.CrearColaEnlazada[int]()

	cola_vaciada.Encolar(1)
	cola_vaciada.Encolar(2)
	cola_vaciada.Encolar(3)

	require.Equal(t, 1, cola_vaciada.Desencolar())
	require.Equal(t, 2, cola_vaciada.Desencolar())
	require.Equal(t, 3, cola_vaciada.Desencolar())

	require.True(t, cola_vaciada.EstaVacia())

	require.Panics(t, func() { cola_vaciada.Desencolar() })

	require.Panics(t, func() { cola_vaciada.VerPrimero() })
}

func TestIntegrandoEnteros(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	require.Equal(t, 1, cola.VerPrimero())
	require.Equal(t, 1, cola.VerPrimero())
	require.Equal(t, 1, cola.Desencolar())
	require.Panics(t, func() { cola.Desencolar() })

	cola.Encolar(1)
	cola.Encolar(2)
	require.Equal(t, 1, cola.Desencolar())
	require.Equal(t, 2, cola.Desencolar())
	require.Panics(t, func() { cola.Desencolar() })
	require.Panics(t, func() { cola.VerPrimero() })
}

func TestIntegrandoStrings(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()

	dias := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes"}

	for i := 0; i < 3; i++ {
		cola.Encolar(dias[i])
	}
	require.False(t, cola.EstaVacia())

	for i := 0; i < 3; i++ {
		require.Equal(t, dias[i], cola.VerPrimero())
		require.Equal(t, dias[i], cola.Desencolar())
	}

	require.Panics(t, func() { cola.Desencolar() })
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.Desencolar() })
	require.Panics(t, func() { cola.VerPrimero() })
}

func TestIntegrandoFloats(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float32]()

	var i float32
	for i = 0; i <= 4; i++ {
		cola.Encolar(i - (i / 2.0))
	}

	require.False(t, cola.EstaVacia())
	require.Equal(t, float32(0), cola.Desencolar())
	require.Equal(t, float32(0.5), cola.Desencolar())
	require.Equal(t, float32(1), cola.Desencolar())
	require.Equal(t, float32(1.5), cola.VerPrimero())
	require.Equal(t, float32(1.5), cola.Desencolar())
	require.Equal(t, float32(2), cola.VerPrimero())
	require.Equal(t, float32(2), cola.Desencolar())
	require.Panics(t, func() { cola.Desencolar() })
	require.Panics(t, func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
}

func TestIntegrandoVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	for i := 0; i <= _VOLUMEN_TEST; i++ {
		cola.Encolar(i)
	}

	for i := 0; i < _VOLUMEN_TEST-1; i++ {
		require.Equal(t, i, cola.Desencolar())
	}

	require.Equal(t, _VOLUMEN_TEST-1, cola.VerPrimero())
	require.Equal(t, _VOLUMEN_TEST-1, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, _VOLUMEN_TEST, cola.Desencolar())
	require.Panics(t, func() { cola.Desencolar() })
	require.True(t, cola.EstaVacia())
}
