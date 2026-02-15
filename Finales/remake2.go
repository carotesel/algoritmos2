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

/*1. Implementar un algoritmo que reciba un arreglo de n enteros (con n ≥ 3) en el que todos sus elementos son iguales
salvo 1, y determine (utilizando división y conquista) cual es dicho elemento no repetido. Indicar y justificar la
complejidad del algoritmo implementado.*/

func distintaa(arr []int) int{
    return laDis(arr, 0, len(arr)-1, arr[0])
}

func laDis(arr []int, ini, fin, repet int) int{
	if ini == fin{
		return arr[ini]
	}

	medio := (ini + fin)/2

	if arr[medio] != repet{
		return arr[medio]
	}
	if arr[medio-1] != repet{
		return arr[medio-1]
	}
	if arr[medio+1] != repet{
		return arr[medio+1]
	}

	cantIzq := medio - ini
	cantDer := fin - medio

	if cantIzq % 2 == 1{ // impar = algo raro
		return laDis(arr, ini, medio-1, repet)
	} else{
		return laDis(arr, medio+1, fin, repet)
	}
}

// O(n) si o si porque no hay forma de asegurar que este a la izq o a la der sin recorrerlos.

/*5. Implementar una primitiva de árbol binario de búsqueda que devuelva un diccionario en el cual las claves sean los
niveles (int) y los datos sean listas de todos las claves del ABB que se encuentran en dicho nivel. Indicar y justificar la
complejidad del algoritmo implementado.*/

type Elem[K comparable, V any] struct {
	nodo  *nodoABB[K, V]
	nivel int
}

func (abb *abb[K, V]) Niveles() Diccionario[int, Lista[K]] {

	dicc := CrearHash[int, Lista[K]]()
	if abb.raiz == nil {
		return dicc
	}

	cola := CrearColaEnlazada[Elem[K,V]]()
	cola.Encolar(Elem[K,V]{nodo: abb.raiz, nivel: 0})

	for !cola.EstaVacia() {

		v := cola.Desencolar()

		if !dicc.Pertenece(v.nivel){
			lista := CrearListaEnlazada[K]()
			dicc.Guardar(v.nivel, lista)
		} else{
			lista := dicc.Obtener(v.nivel)
		}

		lista.AgregarUltimo(v.nodo.clave)
		

		if v.nodo.izq != nil {
			cola.Encolar(Elem[K,V]{v.nodo.izq, v.nivel+1})
		}

		if v.nodo.der != nil {
			cola.Encolar(Elem[K,V]{v.nodo.der, v.nivel+1})
		}
	}

	return dicc
}
