/*
1. Implementar la primitiva Invertir() para el Heap, que invierta su forma de comparar los elementos (es decir, si era de máximos,
ahora sea de mínimos), sin modificar las funciones y primitivas previamente implementadas (simplemente contener todo el cambio
dentro de la primitiva). Indicar y justificar la complejidad del algoritmo. 

Indicar qué consecuencias podría tener esta forma de implementación si se invoca a la primitiva Invertir una cantidad k 
de veces, y cómo podría resolverse si se permitiera modificar otras funciones y/o primitivas 
(y/o la estructura del heap en sí).
*/

func (h *heap[T]) Invertir() {
	cmpVieja := h.cmp

	h.cmp = func (a, b T) int{
		return -cmpVieja(a, b)
	}

	heapify(h.datos, h.cmp)
}

// O(n)

// Consecuencias de llamar k veces: o(k*n) -> costo elevado

/*
2. Implementar el algoritmo de Dijsktra para encontrar el camino mínimo de un vértice hacia los demás vértices del grafo, con la
siguiente modificación: si para llegar a un vértice tengo dos (o más) opciones de misma distancia mínima, elegir aquella cuyo primer
vértice en el camino sea de menor distancia (en caso que la distancia sea la misma, por la razón que fuere, entonces se puede elegir
cualquiera de aquellos). Indicar y justificar la complejidad del algoritmo implementado. 
*/

def dijkstra_modif_2(grafo, origen):
	padres = {}
	distancias = {}
	primer_paso = {}
	heap = Heap()

	for v in grafo.vertices():
		distancias[v] = float("inf")
	
	distancias[origen] = 0
	padres[origen] = None
	heap.append((0, origen))
	
	while len(cola) > 0:
		_, v = heap.popleft()

		for w in grafo.adyacentes(v):
			dist_w = distancias[v] + grafo.peso_arista(v, w)
			if distancias[w] > dist_w:
				distancias[w] = dist_w
				padres[w] = V
				heap.append((distancias[w], w))
	
	return padres

/*
3. Analizar y explicar qué características debe tener un algoritmo de División y Conquista para que su 
complejidad sea O(log n).
*/

//T(n) = A * T(n/2) + O(n ^ c)
//logb(A) = c -> O(n ^c log n)
// c = 0
// logb(a) = 0
// b ^ 0 = 1
// por prop logaritmo: b > 1

// Cualquier algoritmo que lo cumpla sera o(log n)

/*
4. Implementar para la lista enlazada la primitiva Downsample(k int). Esta debe eliminar todos los elementos que se encuentren en
posiciones múltiplos de k (k > 1). La primera posición es la posición 0. Indicar y justificar la complejidad del algoritmo implementado.
*/

func (lista *listaEnlazada[T]) Downsample(k int){
	contador := 0
	actual := lista.primero
	var anterior *nodoLista[T] = nil

	while actual != nil{
		if contador % k == 0{
			if anterior == nil{
				lista.primero = actual.siguiente
			} else if actual.siguiente == nil{
				anterior.siguiente = nil
				lista.ultimo = anterior
			} else{
				anterior.siguiente = actual.siguiente
			}
			lista.largo--
			actual = actual.siguiente // actual avanza pero anterior no porque lo borre xd
		} else{
			anterior = actual
			actual = actual.siguiente
		}
		contador++
	}
}

// Complejidad: o(n)

/*
5. Implementar la primitiva Interseccion(otro *abb[K, V]) Lista[K] para el ABB que nos devuelva una lista ordenada con la
intersección entre el árbol y el recibido por parámetro, que estén ocupando el mismo lugar en el árbol. Indicar y justificar la
complejidad del algoritmo implementado. En el ejemplo a continuación, la intersección sería [4, 10, 18, 20].
*/

func (abb *abb[K, V]) Interseccion(otro *abb[K, V]) Lista[K] {
	lista := CrearListaEnlazada[K]()
	interseccion(abb.raiz, otro.raiz, abb.cmp, lista)
	return lista
}

func interseccion[K comparable, V any](nodo1, nodo2 *nodoABB[K, V], lista Lista[K], cmp func(a, b int) int){
	if nodo1 == nil || nodo2 == nil{
		return
	}

	interseccion(nodo1.izq, nodo2.izq, lista, cmp)

	if cmp(nodo1.clave, nodo2.clave) == 0{
		lista.AgregarUltimo(nodo1.clave)
	}

	interseccion(nodo1.der, nodo2.der, lista, cmp)

}