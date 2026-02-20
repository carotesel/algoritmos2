/*
1. Dados dos arreglos ordenados A y B, donde B tiene “un elemento menos que A”, implementar un algoritmo de división y
conquista que permita obtener el valor faltante de A en B. Ejemplo, si A = {2, 4, 6, 8, 9, 10, 12} y B = {2, 4,
6, 8, 10, 12}, entonces la salida del algoritmo debe ser o bien la posición 4, o el valor 9 (lo que decidan que devuelva).
Indicar y justificar adecuadamente la complejidad del algoritmo implementado.
*/

func encontrarFaltante(arr1, arr2 []int) int{
	return faltante(arr1, arr2, 0, len(arr2))
}

func faltante(arr1, arr2 []int, ini, fin int) int{
	if ini == fin{
		return ini
	}

	medio := (ini + fin)/2

	if arr1[medio] != arr2[medio]{
		return faltante(arr1, arr2, ini, medio)
	} else{
		return faltante(arr1, arr2, medio+1, fin)
	}
}

/*
2. Implementar un algoritmo que dado un grafo dirigido y clico, obtenga el camino mínimo de un vértice v a otro w,
en tiempo lineal (de vértices y aristas). Justificar la complejidad del algoritmo implementado. Pista: utilizar un
recorrido DFS.
*/

def camino_min(grafo, origen, fin):
	visitados = set()
	distancias = {}
	padres = {origen: None}

	for v in grafo:
		distancias[v] = float("inf")
	
	distancias[origen] = 0
	visitados.add(origen)

	for v in grafo.adyacentes(origen):
		if v not in visitados:
			_dfs(grafo, v, fin, padres, visitados, distancias)
	
	return padres, distancias

def _dfs(grafo, v, fin, padres, visitados, distancias):
	visitados.add(v)
	
	for w in grafo.adyacentes(v):
		if w not in visitados:
			_dfs(grafo, w, fin, padres, visitados, distancias)

/*
5. Implementar una primitiva para el heap func (heap *heap[T]) DiferenciaSimetrica(otro *heap[T])
ColaPrioridad[T], que reciba otro Heap y cree un nuevo Heap con los elementos del primero que no se
ecuentren en el segundo, y viceversa (es decir, la diferencia simétrica entre ambos). La función de comparación del
nuevo heap debe ser la del primer heap. Indicar y justificar la complejidad del algoritmo implementado.
*/

func (heap *heap[T]) DiferenciaSimetrica(otro *heap[T]) ColaPrioridad[T]{
	d1 := CrearHash[T, bool]()
	d2 := CrearHash[T, bool]()
	nuevoArr := make([]T, 0)

	for _, elem := range heap.datos{
		if !d1.Pertenece(elem){
			d1.Guardar(elem, true)
		}
	}

	for _, elem := range otro.datos{
		if !d2.Pertenece(elem){
			d2.Guardar(elem, true)
		}
	}

	for _, elem := range heap.datos{
		if !d2.Pertenece(elem){
			nuevoArr = append(nuevoArr, elem)
		}
	}
	
	for _, elem := range otro.datos{
		if !d1.Pertenece(elem){
			nuevoArr = append(nuevoArr, elem)
		}
	}

	nuevo := CrearHeapArr(nuevoArr, heap.cmp)
	return nuevo
}
