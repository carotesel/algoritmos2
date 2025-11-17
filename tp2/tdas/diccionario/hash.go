package diccionario

import (
	"fmt"
)

const (
	_TAMAÑO_INICIAL                    = 3
	_CARGA_MAXIMA_SOBRE_TOTAL  float32 = 0.7
	_CARGA_MINIMA_SOBRE_TOTAL  float32 = 0.2
	_FACTOR_DE_REDIMENCION     int     = 2
	_MENSAJE_ERROR_DICCIONARIO         = "La clave no pertenece al diccionario"
	_MENSAJE_ERROR_ITERADOR            = "El iterador termino de iterar"
)

type estado int

const (
	VACIO estado = iota
	OCUPADO
	BORRADO
)

type celdaHashCerrado[K any, V any] struct {
	clave  K
	dato   V
	estado estado
}

type hashCerrado[K any, V any] struct {
	tabla                 []celdaHashCerrado[K, V]
	funcComparacionClaves func(K, K) bool
	cantOcupados          int
	cantBorrados          int
}

type iteradorHashCerrado[K any, V any] struct {
	dicc             *hashCerrado[K, V]
	indice_actual    int
	cantidad_iterada int
}

func convertirABytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func fnvHashing[K any](clave K, largo int) int {
	clave_conv := convertirABytes(clave)
	var h uint64

	h = 14695981039346656037
	for _, c := range clave_conv {
		h *= 1099511628211
		h ^= uint64(c)
	}
	return int(h % uint64(largo))
}

func buscar[K any, V any](clave K, tabla []celdaHashCerrado[K, V], f func(K, K) bool) int {
	pos := fnvHashing(clave, len(tabla))
	for tabla[pos].estado != VACIO && !(tabla[pos].estado == OCUPADO && f(tabla[pos].clave, clave)) {
		pos = (pos + 1) % len(tabla)
	}
	return pos
}

func crearCelda[K any, V any](clave K, dato V) celdaHashCerrado[K, V] {
	return celdaHashCerrado[K, V]{clave: clave, dato: dato, estado: OCUPADO}
}

func CrearHash[K any, V any](funcionComparacionClaves func(K, K) bool) Diccionario[K, V] {
	var diccionario hashCerrado[K, V]

	diccionario.tabla = make([]celdaHashCerrado[K, V], _TAMAÑO_INICIAL)
	diccionario.funcComparacionClaves = funcionComparacionClaves
	diccionario.cantBorrados = 0
	diccionario.cantOcupados = 0

	return &diccionario
}

func (diccionario *hashCerrado[K, V]) redimension(nuevo_largo int) {
	contador := 0
	nueva_tabla := make([]celdaHashCerrado[K, V], nuevo_largo)

	for i := 0; i < len(diccionario.tabla) && contador < diccionario.Cantidad(); i++ {
		if diccionario.tabla[i].estado == OCUPADO {
			contador++
			indice := buscar(diccionario.tabla[i].clave, nueva_tabla, diccionario.funcComparacionClaves)
			nueva_tabla[indice] = diccionario.tabla[i]
		}
	}
	diccionario.cantBorrados = 0
	diccionario.tabla = nueva_tabla
}

func (diccionario *hashCerrado[K, V]) Cantidad() int {
	return diccionario.cantOcupados
}

func (diccionario *hashCerrado[K, V]) Pertenece(clave K) bool {
	indice := buscar(clave, diccionario.tabla, diccionario.funcComparacionClaves)

	return diccionario.tabla[indice].estado != VACIO
}

func (diccionario *hashCerrado[K, V]) Guardar(clave K, dato V) {
	indice := buscar(clave, diccionario.tabla, diccionario.funcComparacionClaves)

	if indice != -1 && diccionario.tabla[indice].estado != VACIO {
		diccionario.tabla[indice].dato = dato
	} else {
		diccionario.cantOcupados++

		cant_total := diccionario.cantBorrados + diccionario.cantOcupados
		factor_carga := float32(cant_total) / float32(len(diccionario.tabla))

		if factor_carga >= _CARGA_MAXIMA_SOBRE_TOTAL || indice == -1 {
			diccionario.redimension(cant_total * _FACTOR_DE_REDIMENCION)
			indice = buscar(clave, diccionario.tabla, diccionario.funcComparacionClaves)
		}

		diccionario.tabla[indice] = crearCelda(clave, dato)
	}
}

func (diccionario *hashCerrado[K, V]) Obtener(clave K) V {
	indice := buscar(clave, diccionario.tabla, diccionario.funcComparacionClaves)

	if diccionario.tabla[indice].estado != OCUPADO {
		panic(_MENSAJE_ERROR_DICCIONARIO)
	}

	return diccionario.tabla[indice].dato
}

func (diccionario *hashCerrado[K, V]) Borrar(clave K) V {
	indice := buscar(clave, diccionario.tabla, diccionario.funcComparacionClaves)

	if diccionario.tabla[indice].estado != OCUPADO {
		panic(_MENSAJE_ERROR_DICCIONARIO)
	}

	diccionario.cantBorrados++
	diccionario.cantOcupados--

	diccionario.tabla[indice].estado = BORRADO
	borrado_dato := diccionario.tabla[indice].dato

	factor_carga := float32(diccionario.cantOcupados) / float32(len(diccionario.tabla))

	if factor_carga < _CARGA_MINIMA_SOBRE_TOTAL {
		nuevo_tamanio := len(diccionario.tabla) / _FACTOR_DE_REDIMENCION
		if nuevo_tamanio < _TAMAÑO_INICIAL {
			nuevo_tamanio = _TAMAÑO_INICIAL
		}
		diccionario.redimension(nuevo_tamanio)
	}

	return borrado_dato
}

func (diccionario *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, i := range diccionario.tabla {
		if i.estado == OCUPADO {
			if !visitar(i.clave, i.dato) {
				return
			}
		}
	}
}

func (diccionario *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	return crearIteradorHashCerrado(diccionario)
}

func crearIteradorHashCerrado[K any, V any](dicc *hashCerrado[K, V]) IterDiccionario[K, V] {

	pos := 0
	if dicc.cantOcupados != 0 {
		pos, _ = buscarSiguiente(0, dicc.tabla)
	}

	return &iteradorHashCerrado[K, V]{dicc: dicc, indice_actual: pos, cantidad_iterada: 0}
}

func (iter *iteradorHashCerrado[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_ERROR_ITERADOR)
	}

	actual := iter.dicc.tabla[iter.indice_actual]
	return actual.clave, actual.dato
}

func (iter *iteradorHashCerrado[K, V]) HaySiguiente() bool {
	return iter.cantidad_iterada < iter.dicc.cantOcupados
}

func (iter *iteradorHashCerrado[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_ERROR_ITERADOR)
	}

	pos, encontre_siguiente := buscarSiguiente(iter.indice_actual+1, iter.dicc.tabla)

	if encontre_siguiente {
		iter.indice_actual = pos
	}

	iter.cantidad_iterada++
}

func buscarSiguiente[K any, V any](indiceEmpiezo int, tabla []celdaHashCerrado[K, V]) (int, bool) {

	for i := indiceEmpiezo; i < len(tabla); i++ {
		if tabla[i].estado == OCUPADO {
			return i, true
		}
	}

	return 0, false
}
