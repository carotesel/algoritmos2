/*Ejercicio 6: Heap 
Implementá una primitiva para el Heap func (heap *heap[T]) KesimoMayor(k int) T que devuelva el k-ésimo elemento más grande del heap (siendo 1 el máximo). El heap NO debe modificarse. Indicá y justificá la complejidad.
Importante: Pensá en usar estructuras auxiliares.*/

// para no modificar el heap, uso la prop de que el proximo maximo esta en uno de sus hijos (2 * i + 1 y 2 * i + 2)

type Item[T any] struct{
	dato T
	idx int
}

func (heap *heap[T]) KesimoMayor(k int) T{
	aux := CrearHeap[T](heap.cmp)
	
	aux.Encolar(Item{dato: heap.datos[0], idx: 0})

	var actual Item[T]

	for i:=0; i<k; i++{ // O(k)
		actual = aux.Desencolar() // O(log m)

		izq := 2 * actual.idx + 1
		der := 2 * actual.idx + 2

		if izq < len(heap.datos){
			aux.Encolar(Item{dato: heap.datos[izq], idx: izq})
		}

		if der < len(heap.datos){
			aux.Encolar(Item{dato: heap.datos[der], idx: der})
		}
	}
	return actual
}

// m = len(heap)

// en el peor de los casos el heap tiene k cosas

// O(k log k)

/*Ejercicio 9: Lista Enlazada (siempre cae una primitiva)
Implementá la primitiva func (lista *listaEnlazada[T]) Particionar(predicado func(T) bool) (Lista[T], Lista[T]) 
que devuelva dos listas: una con los elementos que cumplen el predicado y otra con los que no. 
La lista original debe quedar vacía. No crear nodos nuevos, solo reorganizar los existentes. 
Indicá y justificá la complejidad.*/

func (lista *listaEnlazada[T]) Particionar(predicado func(T) bool) (Lista[T], Lista[T]){
	l1 := CrearListaEnlazada[T]()
	l2 := CrearListaEnlazada[T]()

	actual := lista.primero

	for actual != nil{
		siguiente := actual.siguiente
		actual.siguiente = nil

		if predicado(actual.dato){
			l1.Agregar(actual.dato)
		} else{
			l1.Agregar(actual.dato)
		}
		actual = siguiente
	}
	lista.primero = nil
	lista.ultimo = nil
	lista.largo = 0
	return l1, l2
}

func (lista *listaEnlazada[T]) Agregar (nodo *nodo[T]){
	if lista.primero == nil{
		lista.primero = nodo
		lista.ultimo = nodo
	} else{
		lista.ultimo.siguiente = nodo
		lista.ultimo = nodo
	}
	lista.largo++
}

// prim lista. agregar sacar etc con punteros y editar lista.largo

/*Ejercicio 18: Heap - Top K elementos (MUY COMÚN)
Implementá una función func TopK(arreglo []int, k int) []int que devuelva los k elementos más grandes de un arreglo desordenado. Debe ejecutar en O(n log k). Justificá por qué tu solución cumple esa complejidad.
Ejemplo: [3, 1, 4, 1, 5, 9, 2, 6], k=3 → [9, 6, 5] (en cualquier orden)*/

func cmp(a, b int) int{
	return a - b // minimos
}

func TopK(arreglo []int, k int) []int{
	heap := CrearHeapArr[int](arreglo[:k], cmp) // O(k)
	
	for i:= k; i<len(arreglo); i++{  // O(n - k)
		if arr[i] > heap.VerMinimo(){
			heap.Desencolar() // o(log k)
			heap.Encolar(arr[i])
		}
	}

	res := make([]int, 0, k)

	for !heap.EstaVacia(){ // O(k)
		res = append(res, heap.Desencolar())
	}
	return res
}

// O(k + [(n - k) log k])
// O(k + n log k - k log k)
// como k << n -> o(n log k)

/*Ejercicio 7: Grafos - Caminos (patrón recurrente en 2025)
Tenemos un grafo dirigido y acíclico (DAG) y dos vértices v y w. 
Implementá un algoritmo que cuente la cantidad de caminos diferentes que hay desde v hasta w. 
Debe ser en tiempo lineal O(V + E). Justificá la complejidad.*/

def orden_top(grafo):

	g_ent = grados_entrada(grafo)
	cola = Cola()
	orden = []
	
	for v in grafo:
		if g_ent[v] == 0:
			cola.Encolar(v)
	
	while cola:
		x = cola.Desencolar()
		orden.append(x)

		for w in grafo.adyacentes(x):
			g_ent[w] -= 1
			if g_ent[w] == 0:
				cola.Encolar(w)
	
	return orden

def cant_caminos_diferentes(grafo, v, w):
	orden = orden_top(grafo) // O(V + E)
	caminos = {}

	for x in grafo: // O(v)
		caminos[x] = 0

	for x in orden: // O (V + E)
		for y in grafo.adyacentes(x):
			caminos[y] += caminos[x]
	return caminos[w]

/*Ejercicio 11: División y Conquista (siempre hay uno)
Tenés un arreglo ordenado que fue rotado en algún punto desconocido. 
Por ejemplo: [4, 5, 6, 7, 0, 1, 2] (originalmente era [0, 1, 2, 4, 5, 6, 7]).
Implementá un algoritmo que encuentre el mínimo elemento en O(log n). 
Justificá la complejidad con el Teorema Maestro.*/

//ENCONTRAR EL PRIMER CORTE! -> COMPARO MEDIO CON FIN!!!!

func minimo_dyc(arr []int) int{
	return min(arr, 0, len(arr)-1)
}

func min(arr []int, ini, fin int) int{
	if ini == fin{
		return arr[ini]
	}

	medio := (ini + fin)/2

	if arr[medio] > arr[fin]{
		return min(arr, medio+1, fin)
	} else{
		return min(arr, ini, medio)
	}
}

/*Ejercicio 16: Heap - Merge con condición (patrón diciembre 2025)
Implementá una primitiva func (heap *heap[T]) MergeSinDuplicados(otro *heap[T], igualdad func(T,T) bool) ColaPrioridad[T] 
que devuelva un nuevo heap con los elementos de ambos heaps, pero sin elementos duplicados 
(si un elemento está en ambos, incluirlo solo una vez). 

La función de comparación es la del primer heap. Indicá y justificá la complejidad.
*/

/*Ejercicio 17: Lista - Intercalar (patrón julio 2025)
Implementá la primitiva func (lista *listaEnlazada[T]) Intercalar(otra *listaEnlazada[T]) que intercale 
los elementos de la otra lista en la lista actual. Por ejemplo:

Lista 1: [1, 3, 5, 7]
Lista 2: [2, 4, 6]
Resultado en Lista 1: [1, 2, 3, 4, 5, 6, 7]

Ambas listas deben quedar vacías al final excepto la lista resultado. No crear nodos nuevos. 
Indicá y justificá la complejidad.*/

/*
Ejercicio 32: Cola - Filtrar manteniendo orden
Implementá func (cola *colaEnlazada[T]) FiltrarConOrden(mantener func(T) bool) Cola[T] que devuelva una nueva cola con los elementos que cumplen la condición, manteniendo el orden relativo. La cola original debe quedar intacta. Indicá y justificá la complejidad.

Ejercicio 33: Pila - Elemento mínimo en O(1)
Diseñá una implementación de Pila que además de las operaciones normales (apilar, desapilar, ver_tope), soporte obtener_minimo() que devuelva el elemento mínimo de la pila, todas en O(1).
Explicá la estructura interna y justificá por qué cada operación es O(1).
Pista: Necesitás una estructura auxiliar.

Ejercicio 35: ABB - Sucesor Inorder
Implementá func (abb *abb[K,V]) Sucesor(clave K, cmp func(K,K) int) (K, bool) que devuelva el sucesor inorder de una clave (el siguiente elemento en un recorrido inorder). Si no existe, devolver false. Debe ser O(h) donde h es la altura. Justificá.

Ejercicio 36: Lista - Swap en pares
Implementá func (lista *listaEnlazada[T]) SwapPares() que intercambie cada par de nodos adyacentes. No intercambiar datos, intercambiar nodos (mover punteros).
Ejemplo: [1, 2, 3, 4, 5] → [2, 1, 4, 3, 5]
Complejidad: O(n). Justificá.

Ejercicio 37: Heap - Reemplazar raíz
Implementá func (heap *heap[T]) ReemplazarRaiz(nuevo T) que reemplace la raíz del heap con un nuevo elemento y mantenga la propiedad de heap. Debe ser O(log n). Justificá.
Nota: Esto es más eficiente que desencolar() + encolar()

Ejercicio 44: Lista - Eliminar N-ésimo desde el final en una pasada (DIFÍCIL)
Implementá func (lista *listaEnlazada[T]) EliminarNesimoDesdeAtras(n int) que elimine el n-ésimo nodo desde el final, recorriendo la lista una sola vez (sin calcular el largo primero). Complejidad: O(longitud). Justificá.

Ejercicio 45: Grafos - Ciclo de longitud exacta K (DIFÍCIL)
Implementá un algoritmo que determine si existe un ciclo de exactamente K vértices en un grafo dirigido. Complejidad esperada: O(V+E) usando DFS con backtracking. Justificá.
*/

