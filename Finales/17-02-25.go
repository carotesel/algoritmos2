/*
2. Implementar para la cola enlazada la primitiva Consumir(accion func (T)) que aplique la función accion a todos
los elementos de la cola. Al terminar la ejecución, la cola debe quedar vacía. Se espera que se implemente sin utilizar
otras primitivas, para demostrar el conocimiento sobre estructuras enlazadas. Indicar y justificar la complejidad de la
primitiva.
*/

func (cola *ColaEnlazada[T]) Consumir(accion func (T)){
	for cola.primero != nil(){
		accion(cola.primero.valor)
		cola.primero = cola.primero.siguiente
	}
	cola.ultimo = nil
}

/*
4. Tenemos un arreglo de n elementos en el que cada elemento se encuentra a lo sumo k posiciones de la que le correspondería
si el arreglo estuviera ordenado (2 ≤ k ≤ n). Implementar un algoritmo de ordenamiento que funcione en O(n log k).
*/

func cmp(a, b int) int{
	return a - b
}

func aKposiciones(arr []int, k int) []int{
	res := make([]int, len(arr))

	heap_min := CrearHeapArr(arr[:k+1], cmp) // O(k)

	indice := 0

	for j:= k+1; j < len(arr); j++{ // O(n-(k+1)) * O(log k) ≈ O(n log k)
		res[indice] = heap_min.Desencolar()
		indice++
		heap_min.Encolar(arr[j])
	}

	for !heap_min.EstaVacia(){ //O(k+1) * log k
		res[indice] = heap_min.Desencolar()
		indice++
	}
	return res
}

// O(n log k) = encolo los primeros k, desp recorro el arr y mientras ordeno los que desencolo encolo los que me faltan del array.
// Al final desencolo los q quedan y estan todos ordenaditos.

/*
3. Implementar un algoritmo que, dado un grafo no dirigido, conexo, y sin puentes (es decir, sin ninguna arista que al
quitarla formaría más de una componente conexa), determine una dirección para cada arista, para que el grafo dirigido
resultante sea fuertemente conexo (es decir, haya una única componente fuertemente conexa). Indicar y justificar la
complejidad del algoritmo.
*/

// componentes conexas etc dfs

def obtener_dirigido(grafoND):
	visitados = set()
	aristas_visitadas = set()
	grafoD = Grafxo(es_dirigido = true, vertices = grafoND.obtener_vertices())
	origen = grafo.vertice_aleatorio()
	dfs(grafoND, grafoD, origen, visitados, aristas_visitadas)
	return grafoD


def dfs(grafoND, grafoD, origen, visitados, aristas_visitadas):
	visitados.add(origen)

	for w in grafoND.adyacentes(origen):
		if (origen, w) in aristas_visitadas or (w, origen) in aristas_visitadas:
			continue
		
		aristas_visitadas.add(origen, w)

		if w not in visitados:
			grafoD.agregarArista(origen, w)
			dfs(grafoND, grafoD, w, visitados, aristas_visitadas)
		else:
			 grafoD.agregarArista(origen, w)



