/*
3. Implementar un algoritmo que, dado un grafo no dirigido, conexo, y sin puentes (es decir, sin ninguna arista que al
quitarla formaría más de una componente conexa), determine una dirección para cada arista, para que el grafo dirigido
resultante sea fuertemente conexo (es decir, haya una única componente fuertemente conexa). Indicar y justificar la
complejidad del algoritmo.
*/

def determinar_direcc(grafo):
    visitados = set()
    origen = grafo.vertice_aleatorio()
    padres = {origen: None}
    aristas_visitadas = set()

    g_dir = Grafo(es_dirigido=True, vertices=grafo.obtener_vertices())

    _dfs(grafo, origen, visitados, padres, g_dir, aristas_visitadas)
    return g_dir


def _dfs(grafo, v, visitados, padres, g_dir, aristas_visitadas):
    visitados.add(v)

    for w in grafo.adyacentes(v):

        # no procesar dos veces la misma arista
        if (v, w) in aristas_visitadas or (w, v) in aristas_visitadas:
            continue

        aristas_visitadas.add(arista)

        if w not in visitados:
            # arista de árbol
            padres[w] = v
            g_dir.agregarArista(v, w)
            _dfs(grafo, w, visitados, padres, g_dir, aristas_visitadas)
        elif w != padres[v]:
            # arista de retorno (cierra ciclo)
            g_dir.agregarArista(v, w)



/*
4. Tenemos un arreglo de n elementos en el que cada elemento se encuentra a lo sumo k posiciones de la que le correspondería
si el arreglo estuviera ordenado (2 ≤ k ≤ n). Implementar un algoritmo de ordenamiento que funcione en O(n log k).
*/

func cmp(a, b int) int{
	return a - b
}

func k_pos(arr []int, k int) []int{
	n := len(arr)
	res := make([]int, n, 0)
	heap := CrearHeapArr[int](cmp, arr[:k+1])
	actual := k+1

	if k >= n{
		k = n -1
	}

	for actual < n{
		res = append(res, heap.Desencolar())
		heap.Encolar(arr[actual])
		actual++
	}

	// Vaciar el heap
	for !heap.EstaVacia() {
		res = append(res, heap.Desencolar())
	}

	return res
}