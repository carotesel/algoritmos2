package diccionario

import TDAPila "algo2/tdas/pila"

type nodoArbol[K any, V any] struct {
	clave         K
	dato          V
	hijoIzquierdo *nodoArbol[K, V]
	hijoDerecho   *nodoArbol[K, V]
}

type ArbolBinario[K any, V any] struct {
	raiz           *nodoArbol[K, V]
	compararClaves func(K, K) int // segun enunciado: 0 si son iguales, <0 si la primera es menor, >0 si la primera es mayor.
	cantElementos  int
}

func crearNodo[K any, V any](clave K, dato V) *nodoArbol[K, V] {
	return &nodoArbol[K, V]{
		clave:         clave,
		dato:          dato,
		hijoIzquierdo: nil,
		hijoDerecho:   nil,
	}
}

func CrearABB[K any, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &ArbolBinario[K, V]{
		raiz:           nil,
		compararClaves: funcionCmp,
		cantElementos:  0,
	}
}

func buscarABB[K any, V any](compararClaves func(K, K) int, nodo **nodoArbol[K, V], clave K) **nodoArbol[K, V] {
	if *nodo == nil {
		return nodo
	}

	comparacion := compararClaves((*nodo).clave, clave)

	if comparacion < 0 {
		return buscarABB(compararClaves, &((*nodo).hijoDerecho), clave)
	}
	if comparacion > 0 {
		return buscarABB(compararClaves, &((*nodo).hijoIzquierdo), clave)
	}

	return nodo
}

func buscarMasIzquierdo[K any, V any](nodo **nodoArbol[K, V]) **nodoArbol[K, V] {
	if *nodo == nil {
		return nodo
	}

	return buscarMasIzquierdo(&(*nodo).hijoIzquierdo)
}

func apilarMasIzquierdo[K any, V any](nodo *nodoArbol[K, V], pila *TDAPila.Pila[*nodoArbol[K, V]]) {
	if nodo == nil {
		return
	}

	(*pila).Apilar(nodo)

	apilarMasIzquierdo(nodo.hijoIzquierdo, pila)
}

func (diccionario *ArbolBinario[K, V]) Guardar(clave K, dato V) {
	nodo := buscarABB(diccionario.compararClaves, &diccionario.raiz, clave)

	if *nodo == nil {
		diccionario.cantElementos++
		*nodo = crearNodo(clave, dato)

	} else {
		(*nodo).dato = dato
	}
}

func (diccionario *ArbolBinario[K, V]) Pertenece(clave K) bool {
	nodo := buscarABB(diccionario.compararClaves, &diccionario.raiz, clave)
	return *nodo != nil
}

func (diccionario *ArbolBinario[K, V]) Obtener(clave K) V {
	nodo := buscarABB(diccionario.compararClaves, &diccionario.raiz, clave)

	if *nodo == nil {
		panic(_MENSAJE_ERROR_DICCIONARIO)
	}
	return (*nodo).dato
}

func (diccionario *ArbolBinario[K, V]) Borrar(clave K) V {
	nodo := buscarABB(diccionario.compararClaves, &diccionario.raiz, clave)

	if *nodo == nil {
		panic(_MENSAJE_ERROR_DICCIONARIO)
	}

	dato := (*nodo).dato

	if (*nodo).hijoIzquierdo != nil && (*nodo).hijoDerecho != nil {
		// nodo tiene 2 hijos
		reemplazo := buscarMasIzquierdo(&(*nodo).hijoDerecho)
		(*nodo).clave = (*reemplazo).clave
		(*nodo).dato = (*reemplazo).dato

		if (*reemplazo).hijoDerecho == nil {
			*reemplazo = nil

		} else {
			*reemplazo = (*reemplazo).hijoDerecho
		}
	} else {
		// nodo tiene 1 hijo o no tiene
		if (*nodo).hijoIzquierdo == nil {
			*nodo = (*nodo).hijoDerecho
		} else {
			*nodo = (*nodo).hijoIzquierdo
		}
	}

	diccionario.cantElementos--
	return dato
}

func (diccionario *ArbolBinario[K, V]) Cantidad() int {
	return diccionario.cantElementos
}

func (diccionario *ArbolBinario[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	diccionario.raiz.iterar(visitar)
}

func (nodo *nodoArbol[K, V]) iterar(visitar func(K, V) bool) bool {
	if nodo == nil {
		return true
	}
	nodo.hijoIzquierdo.iterar(visitar)

	if !visitar(nodo.clave, nodo.dato) {
		return false
	}
	nodo.hijoDerecho.iterar(visitar)

	return true
}

type iteradorABB[K any, V any] struct {
	pila TDAPila.Pila[*nodoArbol[K, V]]
}

func (diccionario *ArbolBinario[K, V]) Iterador() IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoArbol[K, V]]()
	apilarMasIzquierdo(diccionario.raiz, &pila)
	return &iteradorABB[K, V]{pila: pila}
}

func (iter *iteradorABB[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iteradorABB[K, V]) VerActual() (K, V) {

	if !iter.HaySiguiente() {
		panic(_MENSAJE_ERROR_ITERADOR)
	}

	nodo := iter.pila.VerTope()

	return nodo.clave, nodo.dato
}

func (iter *iteradorABB[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_ERROR_ITERADOR)
	}

	ultimo := iter.pila.Desapilar()

	apilarMasIzquierdo(ultimo.hijoDerecho, &iter.pila)

}
