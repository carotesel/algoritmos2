package diccionario

import TDAPila "tdas/pila"

// estructuras

type nodoABB[K any, V any] struct {
	clave         K
	dato          V
	hijoIzquierdo *nodoABB[K, V]
	hijoDerecho   *nodoABB[K, V]
}

type abb[K any, V any] struct {
	raiz           *nodoABB[K, V]
	compararClaves func(K, K) int
	cantElementos  int
}

type iteradorABB[K any, V any] struct {
	pila           TDAPila.Pila[*nodoABB[K, V]]
	compararClaves func(K, K) int
	desde          *K
	hasta          *K
}

func crearNodo[K any, V any](clave K, dato V) *nodoABB[K, V] {
	return &nodoABB[K, V]{
		clave:         clave,
		dato:          dato,
		hijoIzquierdo: nil,
		hijoDerecho:   nil,
	}
}

func CrearABB[K any, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		raiz:           nil,
		compararClaves: funcionCmp,
		cantElementos:  0,
	}
}

// primitivas del abb

func (diccionario *abb[K, V]) Guardar(clave K, dato V) {
	nodo := buscarABB(diccionario.compararClaves, &diccionario.raiz, clave)

	if *nodo == nil {
		diccionario.cantElementos++
		*nodo = crearNodo(clave, dato)

	} else {
		(*nodo).dato = dato
	}
}

func (diccionario *abb[K, V]) Pertenece(clave K) bool {
	nodo := buscarABB(diccionario.compararClaves, &diccionario.raiz, clave)
	return *nodo != nil
}

func (diccionario *abb[K, V]) Obtener(clave K) V {
	nodo := buscarABB(diccionario.compararClaves, &diccionario.raiz, clave)

	if *nodo == nil {
		panic(_MENSAJE_ERROR_DICCIONARIO)
	}
	return (*nodo).dato
}

func (diccionario *abb[K, V]) Borrar(clave K) V {
	nodo := buscarABB(diccionario.compararClaves, &diccionario.raiz, clave)

	if *nodo == nil {
		panic(_MENSAJE_ERROR_DICCIONARIO)
	}

	dato := (*nodo).dato

	if (*nodo).hijoIzquierdo != nil && (*nodo).hijoDerecho != nil {
		reemplazo := buscarMasIzquierdo(&(*nodo).hijoDerecho)
		(*nodo).clave = (*reemplazo).clave
		(*nodo).dato = (*reemplazo).dato

		*reemplazo = (*reemplazo).hijoDerecho
	} else {
		if (*nodo).hijoIzquierdo == nil {
			*nodo = (*nodo).hijoDerecho
		} else {
			*nodo = (*nodo).hijoIzquierdo
		}
	}

	diccionario.cantElementos--
	return dato
}

func (diccionario *abb[K, V]) Cantidad() int {
	return diccionario.cantElementos
}

func (diccionario *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	diccionario.raiz.iterarRango(nil, nil, diccionario.compararClaves, visitar)
}

func (diccionario *abb[K, V]) Iterador() IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoABB[K, V]]()
	apilarMasIzquierdo(diccionario.raiz, &pila)
	return &iteradorABB[K, V]{pila: pila, compararClaves: diccionario.compararClaves, desde: nil, hasta: nil}
}

func (diccionario *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(K, V) bool) {
	if diccionario.raiz == nil {
		return
	}

	diccionario.raiz.iterarRango(desde, hasta, diccionario.compararClaves, visitar)
}

func (diccionario *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

	if desde == nil && hasta == nil {
		return diccionario.Iterador()
	}

	pila := TDAPila.CrearPilaDinamica[*nodoABB[K, V]]()
	inicializarPilaConCondicion(diccionario.raiz, &pila, desde, hasta, diccionario.compararClaves)
	return &iteradorABB[K, V]{pila: pila, desde: desde, hasta: hasta, compararClaves: diccionario.compararClaves}
}

// primitivas de iterador

func (iter *iteradorABB[K, V]) HaySiguiente() bool {
	if iter.pila.EstaVacia() || (iter.hasta != nil && iter.compararClaves(iter.pila.VerTope().clave, *iter.hasta) > 0) {
		return false
	}
	return true
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

// primitivas de nodo

func (nodo *nodoABB[K, V]) iterarRango(desde *K, hasta *K, compararClaves func(K, K) int, visitar func(K, V) bool) bool {
	if nodo == nil {
		return true
	}

	comparacion_desde := desde == nil || compararClaves(*desde, nodo.clave) <= 0
	comparacion_hasta := hasta == nil || compararClaves(*hasta, nodo.clave) >= 0

	if comparacion_desde {
		if !nodo.hijoIzquierdo.iterarRango(desde, hasta, compararClaves, visitar) {
			return false
		}
	}

	if comparacion_desde && comparacion_hasta {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}

	if comparacion_hasta {
		if !nodo.hijoDerecho.iterarRango(desde, hasta, compararClaves, visitar) {
			return false
		}
	}

	return true
}

// auxiliares

func buscarABB[K any, V any](compararClaves func(K, K) int, nodo **nodoABB[K, V], clave K) **nodoABB[K, V] {
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

func buscarMasIzquierdo[K any, V any](nodo **nodoABB[K, V]) **nodoABB[K, V] {
	if (*nodo).hijoIzquierdo == nil {
		return nodo
	}

	return buscarMasIzquierdo(&(*nodo).hijoIzquierdo)
}

func apilarMasIzquierdo[K any, V any](nodo *nodoABB[K, V], pila *TDAPila.Pila[*nodoABB[K, V]]) {
	if nodo == nil {
		return
	}

	(*pila).Apilar(nodo)

	apilarMasIzquierdo(nodo.hijoIzquierdo, pila)
}

func inicializarPilaConCondicion[K any, V any](nodo *nodoABB[K, V], pila *TDAPila.Pila[*nodoABB[K, V]], desde *K, hasta *K, compararClaves func(K, K) int) {
	if nodo == nil {
		return
	}

	comparacion_desde := desde == nil || compararClaves(*desde, nodo.clave) <= 0
	comparacion_hasta := hasta == nil || compararClaves(*hasta, nodo.clave) >= 0

	if comparacion_desde && comparacion_hasta {
		(*pila).Apilar(nodo)
	}

	if comparacion_desde {
		inicializarPilaConCondicion(nodo.hijoIzquierdo, pila, desde, hasta, compararClaves)
	}

	if comparacion_hasta && !comparacion_desde {
		inicializarPilaConCondicion(nodo.hijoDerecho, pila, desde, hasta, compararClaves)
	}
}
