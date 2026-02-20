/*
1. Se define como puente local de un grafo a una arista que une dos vértices sin adyacentes en común. Implementar un
algoritmo que reciba un grafo y devuelva una lista con todos los puentes locales. Indicar y justificar la complejidad del
algoritmo implementado. Recomendación: no te la rebusques.
*/

def puentes_locales(grafo):
	res = []
	aristas_visitadas = set()

	for v in grafo:
		for w in grafo.adyacentes(v):
			if (v, w) in aristas_visitadas:
				continue
			
			puente = True
			
			for x in grafo.adyacentes(v):
				if x != w and x in grafo.adyacentes(w):
					puente = False
					break
			
			if puente:
				res = append(v, w)
			
			aristas_visitadas.add(v, w)
			aristas_visitadas.add(w, v)
	return res


// Complejidad: O(V*E) porque por cada vertice recorre las aristas.

/*
2. Implementar en Go una primitiva para un árbol binario izquierdista, que reciba la cantidad de nodos que tiene, y
devuelva el dato del elemento más a abajo y a la derecha del árbol. En los árboles de las figuras mostradas, se debe
devolver en ambos casos 4. Para que el ejercicio se pueda considerar como aprobable, debe resolverse en no más que
O(n), sin contar con otros errores. Para que se considere completamente bien, debe ejecutar en O(log n). Justificar la
complejidad del algoritmo implementado. A fines del ejercicio, considerar que la estructura del árbol binario es:
type ab[T any] struct {
izquierda *ab[T]
derecha *ab[T]
dato T
}
*/

// RECORRIDO POR NIVEL -> COLA. NO RECURSIVO.

func (ab *ab[T]) mbALaDer(n int) *ab[T]{

	if ab == nil{
		panic("El arbol esta vacio")
	}

	cola := CrearColaEnlazada[*ab [T]]()
	pila := CrearPilaDinamica[T]()

	cola.Encolar(ab)

	while not cola.EstaVacia(){
		nodo = cola.Desencolar()
		pila.Apilar(nodo.dato)

		if nodo.izquierda != nil{
			cola.Encolar(nodo.izquierda)
		}

		if nodo.derecha != nil{
			cola.Encolar(nodo.derecha)
		}
	}

	return pila.Desapilar()
}

// Version o(n)

/*
3. Implementar un algoritmo para obtener caminos mínimos sobre un grafo pesado, pero con una modificación: el algoritmo
debe recibir el grafo, los vértices origen (v) y destino (w), pero también una arista (x, y) por la que el camino deba
pasar si o si (los vértices x e y pueden o no coincidir con v y/o w, inclusive en cualquier orden). El camino resultante
debe ser el mínimo que permita ir de v a w pasando por la arista (x, y). Si es necesario/conveniente volver a pasar por
un vértice, es correcto en tanto el resultado sea mínimo en función de la restricción adicional. Indicar y justificar la
complejidad del algoritmo implementado.
*/

def dijkstra(grafo, origen):
	distancias = {}

	for v in grafo:
		distancias[v] = float("inf")
	
	distancias[origen] = 0
	heap_min = Heap()
	heap_min.Encolar((0, origen))

	while not heap_min.EstaVacia():
		dist_v, v = heap_min.Desencolar()

		if dist_v > distancias[v]:
			continue
		
		for w in grafo.adyacentes(v):
			dist_w = dist_v + grafo.peso_arista(v, w)

			if dist_w < distancias[w]:
				distancias[w] = dist_w
				heap_min.Encolar((dist_w, w))
	return distancias

def main_HDP(grfao, origen, destino, x, y):

	dist_v = dijkstra(grafo, origen)
	dist_x = dijkstra(grafo, x)
	dist_y = dijkstra(grafo, y)

	peso_xy = grafo.peso_arista(x,y)

	costo_1 = dist_v[x] + peso_xy + dist_y[destino]
	costo_2 = dist_v[y] + peso_xy + dist_x[destino]

	return min(costo_1, costo_2)

	// 3 O((v+e) log v)

/*
5. Implementar una primitiva para la lista enlazada func (lista *listaEnlanzada[T]) Reducir(reductor func(T,
T) T, valorInicial T) T que reciba una función que reduce la lista a un único valor. Inicialmente, la función
reductor debe recibir a valorInicial, y el primer elemento de la lista (en ese orden). Luego, el resultado debe usarse
para invocar a reductor nuevamente, pero con el segundo de la lista (en ese orden), y así, hasta que ya no queden
elementos en la lista y se devuelve el resultado final. Si la lista se encuentra vacía, se debe devolver directamente
valorInicial. Ejemplo:
lista := [1, 2, 3, 4, 5] // como lista enlazada, esto es en pseudocódigo
lista.Reducir(func(a, b int) int {
return a + b
}, 0)
En este caso, la primitiva debe devolver 15 (la suma de todos los elementos de la lista). La lista debe mantenerse en el
mismo estado. Indicar y justificar la complejidad de la primitiva implementada.
*/

func (lista *listaEnlanzada[T]) Reducir(reductor func(T,T) T, valorInicial T) T{

	if lista.largo == 0{
		return valorInicial
	}

	valor := valorInicial

	actual := lista.primero

	for actual != nil{
		valor = reductor(valor, actual.dato)
		actual = actual.siguiente
	}
	return valor
}

// Complejidad: O(n) porque recorre nodo a nodo 1 vez c/u.