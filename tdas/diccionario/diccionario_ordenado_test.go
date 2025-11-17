package diccionario_test

import (
	TDAABB "algo2/tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestABBReciencreado(t *testing.T) {
	abb := TDAABB.CrearABB[int, int](func(a, b int) int {
		if a == b {
			return 0
		}
		if a < b {
			return -1
		}
		return 1
	})

	require.Equal(t, 0, abb.Cantidad())
}
