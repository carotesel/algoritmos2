package diccionario_test

import (
	"math/rand"
	"strings"
	TDAABB "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN_TEST int = 50000

func funcionCmpEnteros(a, b int) int {
	return a - b
}

func TestABBVacio(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, int](funcionCmpEnteros)
	require.Equal(t, 0, diccionario.Cantidad())
}

func TestGuardarYObtener(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	diccionario.Guardar(5, "cinco")
	diccionario.Guardar(3, "tres")
	diccionario.Guardar(7, "siete")

	require.Equal(t, 3, diccionario.Cantidad())
	require.True(t, diccionario.Pertenece(5))
	require.True(t, diccionario.Pertenece(3))
	require.True(t, diccionario.Pertenece(7))
	require.False(t, diccionario.Pertenece(10))

	require.Equal(t, "cinco", diccionario.Obtener(5))
	require.Equal(t, "tres", diccionario.Obtener(3))
	require.Equal(t, "siete", diccionario.Obtener(7))
}

func TestActualizar(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	diccionario.Guardar(1, "a")
	require.Equal(t, "a", diccionario.Obtener(1))
	require.Equal(t, 1, diccionario.Cantidad())

	diccionario.Guardar(1, "b")
	require.Equal(t, 1, diccionario.Cantidad())
	require.Equal(t, "b", diccionario.Obtener(1))
}

func TestObtenerYBorrarInexistente(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)
	require.Panics(t, func() { diccionario.Obtener(5) })
	require.Panics(t, func() { diccionario.Borrar(5) })
}

func TestBorrarRaiz(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	diccionario.Guardar(1, "raiz")
	require.Equal(t, 1, diccionario.Cantidad())

	require.Equal(t, "raiz", diccionario.Borrar(1))
	require.Equal(t, 0, diccionario.Cantidad())
	require.False(t, diccionario.Pertenece(1))
}

func TestBorrarHojas(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	diccionario.Guardar(5, "raiz")
	diccionario.Guardar(3, "izquierda")
	diccionario.Guardar(7, "derecha")

	require.Equal(t, "izquierda", diccionario.Borrar(3))
	require.Equal(t, 2, diccionario.Cantidad())
	require.False(t, diccionario.Pertenece(3))

	require.Equal(t, "derecha", diccionario.Borrar(7))
	require.Equal(t, 1, diccionario.Cantidad())
	require.False(t, diccionario.Pertenece(7))
}

func TestBorrarNodoConUnHijo(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	diccionario.Guardar(10, "raiz")
	diccionario.Guardar(5, "izquierda")
	diccionario.Guardar(2, "izq-izq")

	require.Equal(t, "izquierda", diccionario.Borrar(5))
	require.True(t, diccionario.Pertenece(2))
	require.False(t, diccionario.Pertenece(5))
	require.Equal(t, 2, diccionario.Cantidad())
}

func TestBorrarNodoConDosHijos(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	diccionario.Guardar(10, "raiz")
	diccionario.Guardar(5, "izquierda")
	diccionario.Guardar(15, "derecha")
	diccionario.Guardar(12, "izq-der")
	diccionario.Guardar(17, "der-der")

	require.Equal(t, "raiz", diccionario.Borrar(10))
	require.False(t, diccionario.Pertenece(10))
	require.Equal(t, 4, diccionario.Cantidad())
}

func TestABBQueSeVacio(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	diccionario.Guardar(5, "cinco")
	diccionario.Guardar(3, "tres")
	diccionario.Guardar(7, "siete")
	require.Equal(t, 3, diccionario.Cantidad())

	require.Equal(t, "cinco", diccionario.Borrar(5))
	require.Equal(t, "tres", diccionario.Borrar(3))
	require.Equal(t, "siete", diccionario.Borrar(7))
	require.Equal(t, 0, diccionario.Cantidad())

	require.False(t, diccionario.Pertenece(5))

	diccionario.Guardar(5, "cinco")
	require.Equal(t, 1, diccionario.Cantidad())
	require.True(t, diccionario.Pertenece(5))

}

func TestIterarVacio(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	contadorVueltas := 0
	contarVueltas := func(clave int, dato string) bool {
		contadorVueltas++
		return true
	}

	diccionario.Iterar(contarVueltas)

	require.Equal(t, 0, contadorVueltas)
}

func TestIteradorVacio(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, int](funcionCmpEnteros)
	iter := diccionario.Iterador()
	require.False(t, iter.HaySiguiente())

	require.Panics(t, func() { iter.VerActual() })
	require.Panics(t, func() { iter.Siguiente() })
}

func TestIterarInOrder(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	claves := []int{5, 3, 7, 6}
	valores := []string{"a", "b", "c", "d"}

	for i, k := range claves {
		diccionario.Guardar(k, valores[i])
	}

	recorrido := []int{}
	guardarRecorrido := func(clave int, dato string) bool {
		recorrido = append(recorrido, clave)
		return true
	}

	diccionario.Iterar(guardarRecorrido)

	require.Equal(t, []int{3, 5, 6, 7}, recorrido)
}

func TestIterarInOrderCortaAntes(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	claves := []int{5, 3, 7, 6}
	valores := []string{"a", "b", "c", "d"}

	for i, k := range claves {
		diccionario.Guardar(k, valores[i])
	}

	contadorVueltas := 0
	suma := 0
	contarVueltas := func(clave int, dato string) bool {
		suma += clave
		contadorVueltas++
		return contadorVueltas < 2
	}

	diccionario.Iterar(contarVueltas)

	require.Equal(t, 2, contadorVueltas)
	require.Equal(t, 8, suma)
}

func TestIteradorInOrder(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)
	diccionario.Guardar(5, "a")
	diccionario.Guardar(3, "b")
	diccionario.Guardar(7, "c")

	iter := diccionario.Iterador()
	claves := []int{}

	for iter.HaySiguiente() {
		k, _ := iter.VerActual()
		claves = append(claves, k)
		iter.Siguiente()
	}

	require.Equal(t, []int{3, 5, 7}, claves)
}

func TestIterarRangoBasico(t *testing.T) {
	dicc := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	claves := []int{10, 5, 15, 2, 7, 12, 20}
	valores := []string{"A", "B", "C", "D", "E", "F", "G"}

	for i, k := range claves {
		dicc.Guardar(k, valores[i])
	}

	visitadas := []int{}
	visitar := func(k int, v string) bool {
		visitadas = append(visitadas, k)
		return true
	}

	dicc.IterarRango(&claves[1], &claves[2], visitar)

	require.Equal(t, []int{5, 7, 10, 12, 15}, visitadas)
}

func TestIterarRangoSinDesde(t *testing.T) {
	dicc := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	claves := []int{1, 3, 5, 7, 9}
	valores := []string{"A", "B", "C", "D", "E"}

	for i, k := range claves {
		dicc.Guardar(k, valores[i])
	}

	visitadas := []int{}
	visitar := func(k int, v string) bool {
		visitadas = append(visitadas, k)
		return true
	}

	dicc.IterarRango(nil, &claves[2], visitar)

	require.Equal(t, []int{1, 3, 5}, visitadas)
}

func TestIterarRangoSinHasta(t *testing.T) {
	dicc := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	claves := []int{1, 3, 5, 7, 9}
	valores := []string{"A", "B", "C", "D", "E"}

	for i, k := range claves {
		dicc.Guardar(k, valores[i])
	}

	visitadas := []int{}
	visitar := func(k int, v string) bool {
		visitadas = append(visitadas, k)
		return true
	}

	dicc.IterarRango(&claves[2], nil, visitar)

	require.Equal(t, []int{5, 7, 9}, visitadas)
}

func TestIterarRangoFueraDeRango(t *testing.T) {
	dicc := TDAABB.CrearABB[int, string](funcionCmpEnteros)
	claves := []int{2, 1, 3}
	valores := []string{"A", "B", "C"}
	for i, k := range claves {
		dicc.Guardar(k, valores[i])
	}

	visitadas := []int{}
	visitar := func(k int, v string) bool {
		visitadas = append(visitadas, k)
		return true
	}

	desde := 10
	hasta := 15

	dicc.IterarRango(&desde, &hasta, visitar)

	require.Equal(t, 0, len(visitadas))
}

func TestIterarRangoCortaAntes(t *testing.T) {
	dicc := TDAABB.CrearABB[int, string](funcionCmpEnteros)

	claves := []int{1, 3, 5, 7, 9}
	valores := []string{"A", "B", "C", "D", "E"}

	for i, k := range claves {
		dicc.Guardar(k, valores[i])
	}

	visitadas := []int{}
	visitar := func(k int, v string) bool {
		if k <= 5 {
			visitadas = append(visitadas, k)

		}
		return k < 5
	}

	dicc.IterarRango(&claves[0], &claves[4], visitar)

	require.Equal(t, []int{1, 3, 5}, visitadas)
}

func TestIteradorRangoCompleto(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)
	claves := []int{5, 3, 7, 1, 4, 6, 9}
	valores := []string{"e", "c", "g", "a", "d", "f", "i"}
	for i, clave := range claves {
		diccionario.Guardar(clave, valores[i])
	}

	desde := 1
	hasta := 9
	resultado := []int{}

	iter := diccionario.IteradorRango(&desde, &hasta)

	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		resultado = append(resultado, clave)
		iter.Siguiente()
	}

	require.Equal(t, []int{1, 3, 4, 5, 6, 7, 9}, resultado)
}

func TestIteradorRangoSinDesde(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)
	claves := []int{5, 3, 7, 1, 4, 6, 9}
	valores := []string{"e", "c", "g", "a", "d", "f", "i"}
	for i, clave := range claves {
		diccionario.Guardar(clave, valores[i])
	}

	hasta := 5

	iter := diccionario.IteradorRango(nil, &hasta)

	resultado := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		resultado = append(resultado, clave)
		iter.Siguiente()
	}
	require.Equal(t, []int{1, 3, 4, 5}, resultado)
}

func TestIteradorRangoSinHasta(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)
	claves := []int{5, 3, 7, 1, 4, 6, 9}
	valores := []string{"e", "c", "g", "a", "d", "f", "i"}
	for i, clave := range claves {
		diccionario.Guardar(clave, valores[i])
	}

	desde := 5

	iter := diccionario.IteradorRango(&desde, nil)

	resultado := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		resultado = append(resultado, clave)
		iter.Siguiente()
	}
	require.Equal(t, []int{5, 6, 7, 9}, resultado)
}

func TestIteradorRangoFueraDeRango(t *testing.T) {
	diccionario := TDAABB.CrearABB[int, string](funcionCmpEnteros)
	claves := []int{5, 3, 7, 1, 4, 6, 9}
	valores := []string{"e", "c", "g", "a", "d", "f", "i"}
	for i, clave := range claves {
		diccionario.Guardar(clave, valores[i])
	}

	desde := 10
	hasta := 12

	iter := diccionario.IteradorRango(&desde, &hasta)

	resultado := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		resultado = append(resultado, clave)
		iter.Siguiente()
	}
	require.Equal(t, []int{}, resultado)
}

func TestGenearlClavesStrings(t *testing.T) {
	comparoStrings := func(s1, s2 string) int {
		return strings.Compare(s1, s2)
	}

	diccionario := TDAABB.CrearABB[string, int](comparoStrings)
	valores := []int{5, 3, 7, 1, 4, 6, 9}
	claves := []string{"e", "c", "g", "a", "d", "f", "i"}
	for i, clave := range claves {
		diccionario.Guardar(clave, valores[i])
	}

	require.Equal(t, 7, diccionario.Cantidad())
	require.Equal(t, 5, diccionario.Borrar("e"))
	require.Equal(t, 6, diccionario.Cantidad())
	require.False(t, diccionario.Pertenece("e"))
	require.Panics(t, func() { diccionario.Borrar("e") })
	require.Panics(t, func() { diccionario.Obtener("e") })

	require.Equal(t, 7, diccionario.Obtener("g"))

	desde := "a"
	hasta := "c"

	iter := diccionario.IteradorRango(&desde, &hasta)

	resultado := []string{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		resultado = append(resultado, clave)
		iter.Siguiente()
	}
	require.Equal(t, []string{"a", "c"}, resultado)
}

func ordenarBalanceados(arr []int, ini, fin int, res *[]int) {
	if ini > fin {
		return
	}
	med := (ini + fin) / 2
	(*res) = append((*res), arr[med])

	ordenarBalanceados(arr, ini, med-1, res)
	ordenarBalanceados(arr, med+1, fin, res)
}

func arrayParaInsertarBalanceado(n int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	orden := make([]int, 0)

	ordenarBalanceados(arr, 0, n-1, &orden)

	return orden
}

func TestVolumen(t *testing.T) {
	arr := arrayParaInsertarBalanceado(_VOLUMEN_TEST)

	diccionario := TDAABB.CrearABB[int, int](funcionCmpEnteros)
	for _, clave := range arr {
		diccionario.Guardar(clave, clave)
	}

	require.Equal(t, _VOLUMEN_TEST, diccionario.Cantidad())

	for _, clave := range arr {
		diccionario.Borrar(clave)
	}
	require.Panics(t, func() { diccionario.Obtener(arr[0]) })
	require.Equal(t, 0, diccionario.Cantidad())

	arr = arrayParaInsertarBalanceado(_VOLUMEN_TEST * 2)
	for _, clave := range arr {
		diccionario.Guardar(clave, clave)
	}
	require.Equal(t, _VOLUMEN_TEST*2, diccionario.Cantidad())

	desde := int(float32(_VOLUMEN_TEST*2) * 0.25)
	hasta := int(float32(_VOLUMEN_TEST*2) * 0.75)

	for _, clave := range arr[desde:hasta] {
		diccionario.Borrar(clave)
	}

	require.Equal(t, _VOLUMEN_TEST*2-(hasta-desde), diccionario.Cantidad())
	require.Panics(t, func() { diccionario.Obtener(arr[hasta-1]) })
	require.Panics(t, func() { diccionario.Obtener(arr[desde]) })
	require.False(t, diccionario.Pertenece(arr[desde]))
}

var random = rand.New(rand.NewSource(5454545))

func desordenar(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func TestVolumenDesbalanceado(t *testing.T) {
	arr := make([]int, _VOLUMEN_TEST)
	for i := 0; i < _VOLUMEN_TEST; i++ {
		arr[i] = i
	}

	desordenar(arr)

	diccionario := TDAABB.CrearABB[int, int](funcionCmpEnteros)
	for _, clave := range arr {
		diccionario.Guardar(clave, clave)
	}

	require.Equal(t, _VOLUMEN_TEST, diccionario.Cantidad())

	for _, clave := range arr {
		diccionario.Borrar(clave)
	}
	require.Panics(t, func() { diccionario.Obtener(arr[0]) })
	require.Equal(t, 0, diccionario.Cantidad())

	arr = arrayParaInsertarBalanceado(_VOLUMEN_TEST * 2)
	for _, clave := range arr {
		diccionario.Guardar(clave, clave)
	}
	require.Equal(t, _VOLUMEN_TEST*2, diccionario.Cantidad())

	desde := int(float32(_VOLUMEN_TEST*2) * 0.25)
	hasta := int(float32(_VOLUMEN_TEST*2) * 0.75)

	for _, clave := range arr[desde:hasta] {
		diccionario.Borrar(clave)
	}

	require.Equal(t, _VOLUMEN_TEST*2-(hasta-desde), diccionario.Cantidad())
	require.Panics(t, func() { diccionario.Obtener(arr[hasta-1]) })
	require.Panics(t, func() { diccionario.Obtener(arr[desde]) })
	require.False(t, diccionario.Pertenece(arr[desde]))
}

func TestVolumenDesbalanceadoIterar(t *testing.T) {
	arr := make([]int, _VOLUMEN_TEST)
	for i := 0; i < _VOLUMEN_TEST; i++ {
		arr[i] = i
	}

	desordenar(arr)

	diccionario := TDAABB.CrearABB[int, int](funcionCmpEnteros)
	for _, clave := range arr {
		diccionario.Guardar(clave, clave)
	}

	cont := 0
	visitarTodas := func(k int, v int) bool {
		cont++
		return true
	}

	diccionario.Iterar(visitarTodas)

	require.Equal(t, _VOLUMEN_TEST, cont)

	cont = 0
	visitarCortaAntes := func(k int, v int) bool {
		cont++
		return cont < _VOLUMEN_TEST/2
	}

	diccionario.Iterar(visitarCortaAntes)

	require.Equal(t, _VOLUMEN_TEST/2, cont)
}

func TestVolumenDesbalanceadoIterarRangoCompleto(t *testing.T) {
	arr := make([]int, _VOLUMEN_TEST)
	for i := 0; i < _VOLUMEN_TEST; i++ {
		arr[i] = i + 1
	}

	desordenar(arr)

	diccionario := TDAABB.CrearABB[int, int](funcionCmpEnteros)
	for _, clave := range arr {
		diccionario.Guardar(clave, clave)
	}

	// con desde y hasta
	desde := int(float32(_VOLUMEN_TEST) * 0.25)
	hasta := int(float32(_VOLUMEN_TEST)*0.75) - 1

	cont := 0
	visitarTodas := func(k int, v int) bool {
		cont++
		return true
	}

	diccionario.IterarRango(&desde, &hasta, visitarTodas)
	require.Equal(t, _VOLUMEN_TEST/2, cont)

	// solo con desde
	desde = int(float32(_VOLUMEN_TEST)*0.75) + 1

	cont = 0
	diccionario.IterarRango(&desde, nil, visitarTodas)

	require.Equal(t, _VOLUMEN_TEST/4, cont)

	// solo con hasta
	hasta = int(float32(_VOLUMEN_TEST) * 0.25)

	cont = 0
	diccionario.IterarRango(nil, &hasta, visitarTodas)

	require.Equal(t, _VOLUMEN_TEST/4, cont)
}

func TestVolumenDesbalanceadoIterarRangoConCorte(t *testing.T) {
	arr := make([]int, _VOLUMEN_TEST)
	for i := 0; i < _VOLUMEN_TEST; i++ {
		arr[i] = i + 1
	}

	desordenar(arr)

	diccionario := TDAABB.CrearABB[int, int](funcionCmpEnteros)
	for _, clave := range arr {
		diccionario.Guardar(clave, clave)
	}

	// con desde y hasta
	desde := int(float32(_VOLUMEN_TEST) * 0.25)
	hasta := int(float32(_VOLUMEN_TEST)*0.75) - 1

	cont := 0
	visitarAlgunas := func(k int, v int) bool {
		if k < _VOLUMEN_TEST/2 {
			cont++
			return true
		}
		return false
	}

	diccionario.IterarRango(&desde, &hasta, visitarAlgunas)
	require.Equal(t, _VOLUMEN_TEST/4, cont)

	// solo con desde
	desde = int(float32(_VOLUMEN_TEST)*0.75) + 1

	cont = 0
	diccionario.IterarRango(&desde, nil, visitarAlgunas)

	require.Equal(t, 0, cont)

	// solo con hasta
	hasta = int(float32(_VOLUMEN_TEST) * 0.25)

	cont = 0
	diccionario.IterarRango(nil, &hasta, visitarAlgunas)

	require.Equal(t, _VOLUMEN_TEST/4, cont)
}

func TestVolumenDesbalanceadoIterardor(t *testing.T) {
	arr := make([]int, _VOLUMEN_TEST)
	for i := 0; i < _VOLUMEN_TEST; i++ {
		arr[i] = i
	}

	desordenar(arr)

	diccionario := TDAABB.CrearABB[int, int](funcionCmpEnteros)
	for _, clave := range arr {
		diccionario.Guardar(clave, clave)
	}

	cont := 0
	iter := diccionario.Iterador()

	for iter.HaySiguiente() {
		cont++
		iter.Siguiente()
	}

	require.Equal(t, _VOLUMEN_TEST, cont)

	cont = 0
	iter = diccionario.Iterador()

	for iter.HaySiguiente() && cont < _VOLUMEN_TEST/2 {
		cont++
		iter.Siguiente()
	}

	require.Equal(t, _VOLUMEN_TEST/2, cont)
}

func TestVolumenDesbalanceadoIteradorRangoCompleto(t *testing.T) {
	arr := make([]int, _VOLUMEN_TEST)
	for i := 0; i < _VOLUMEN_TEST; i++ {
		arr[i] = i + 1
	}

	desordenar(arr)

	diccionario := TDAABB.CrearABB[int, int](funcionCmpEnteros)
	for _, clave := range arr {
		diccionario.Guardar(clave, clave)
	}

	// con desde y hasta
	desde := int(float32(_VOLUMEN_TEST) * 0.25)
	hasta := int(float32(_VOLUMEN_TEST)*0.75) - 1

	cont := 0
	iter := diccionario.IteradorRango(&desde, &hasta)

	for iter.HaySiguiente() {
		cont++
		iter.Siguiente()
	}

	require.Equal(t, _VOLUMEN_TEST/2, cont)

	// solo con desde
	desde = int(float32(_VOLUMEN_TEST)*0.75) + 1

	cont = 0
	iter = diccionario.IteradorRango(&desde, nil)

	for iter.HaySiguiente() {
		cont++
		iter.Siguiente()
	}

	require.Equal(t, _VOLUMEN_TEST/4, cont)

	// solo con hasta
	hasta = int(float32(_VOLUMEN_TEST) * 0.25)

	cont = 0
	iter = diccionario.IteradorRango(nil, &hasta)

	for iter.HaySiguiente() {
		cont++
		iter.Siguiente()
	}

	require.Equal(t, _VOLUMEN_TEST/4, cont)
}

func TestVolumenDesbalanceadoIteradorRangoConCorte(t *testing.T) {
	arr := make([]int, _VOLUMEN_TEST)
	for i := 0; i < _VOLUMEN_TEST; i++ {
		arr[i] = i + 1
	}

	desordenar(arr)

	diccionario := TDAABB.CrearABB[int, int](funcionCmpEnteros)
	for _, clave := range arr {
		diccionario.Guardar(clave, clave)
	}

	// con desde y hasta
	desde := int(float32(_VOLUMEN_TEST) * 0.25)
	hasta := int(float32(_VOLUMEN_TEST)*0.75) - 1

	cont := 0
	iter := diccionario.IteradorRango(&desde, &hasta)

	clave := 0
	for iter.HaySiguiente() && clave < _VOLUMEN_TEST/2 {
		clave, _ = iter.VerActual()

		if clave < _VOLUMEN_TEST/2 {
			cont++
		}

		iter.Siguiente()
	}

	require.Equal(t, _VOLUMEN_TEST/4, cont)

	// solo con desde
	desde = int(float32(_VOLUMEN_TEST)*0.75) + 1

	cont = 0
	iter = diccionario.IteradorRango(&desde, nil)

	clave = 0
	for iter.HaySiguiente() && clave < _VOLUMEN_TEST/2 {
		clave, _ = iter.VerActual()

		if clave < _VOLUMEN_TEST/2 {
			cont++
		}

		iter.Siguiente()
	}

	require.Equal(t, 0, cont)

	// solo con hasta
	hasta = int(float32(_VOLUMEN_TEST) * 0.25)

	cont = 0
	iter = diccionario.IteradorRango(nil, &hasta)

	clave = 0
	for iter.HaySiguiente() && clave < _VOLUMEN_TEST/2 {
		clave, _ = iter.VerActual()

		if clave < _VOLUMEN_TEST/2 {
			cont++
		}

		iter.Siguiente()
	}

	require.Equal(t, _VOLUMEN_TEST/4, cont)
}
