/*2. Implementar una primitiva para el Heap merge(otroHeap heap[T]) ColaPrioridad[T] que dado otro heap, nos devuelva un
nuevo heap (con la función de comparación del heap al que se le invoca la primitiva) que contengan los elementos de ambos heaps,
pero con un único elemento de misma prioridad (no importa cuál elemento se elige). Indicar y justificar la complejidad de la primitiva
implementada.*/

func (h heap[T]) merge(otroHeap heap[T]) ColaPrioridad[T] {
	dicc := CrearHash[T, bool](h.cmp)

	for _, dato := range h.datos{
		if !dicc.Pertenece(dato){
			dicc.Guardar(dato, true)
		}
	}

	for _, dato := range otroHeap.datos{
		if !dicc.Pertenece(dato){
			dicc.Guardar(dato, true)
		}
	}

	iter := dicc.Iterador()
	arr := make([]T, 0)

	for iter.HaySiguiente(){
		actual := iter.VerActual()
		arr = append(arr, actual)
	}

	heapNuevo := CrearHeapArr[T](arr, h.cmp) // aplica heapify solo xd
	return heapNuevo
}

/*
3. ¿Como implementar un TDA Cola que además de permitir encolar y desencolar en O(1), permita eliminar el elemento del medio 
de la cola, también en O(1)?
*/

/*
Para implementar un TDA Cola que permita eliminar el item del medio, lo que haria es utilizar una lista doblemente enlazada con punteros,
y a medida que voy agregando elementos mantengo una referencia al nodo del medio moviendo los punteros la distancia correspondiente (utilizando un contador del total de elementos).
Si quiero eliminar el elemento del medio, basta con modificar los punteros de los nodos previo y posterior al que quiero eliminar,
esto me permite conectarlos entre si y sacar el elemento que quiero. En el caso de que la cola tenga un solo elemento, eliminar el medio equivale a vaciar la cola, actualizando los punteros correspondientes.
*/

/*
4. Queremos resolver el problema de obtención de caminos mínimos en un grafo, pero con una variante. Nuestro grafo corresponde a la
red de calles de la ciudad. Como es bien sabido, el tránsito por las calles no es siempre el mismo, lo cual implica que el tiempo para
llegar a destino no es siempre el mismo. En nuestro grafo los pesos corresponden a cuándo (en tiempo) se llega de un vértice a otro a
partir de saber cuándo se comienza a transitar la arista, pero no es un valor constante sino una función (cada arista tiene su propia
función como peso), que dado el tiempo actual nos indica en qué tiempo pasamos por dicha arista hasta llegar al vértice destino (la
función es monótona creciente, así que no hay nunca ventaja en quedarse esperando en un lugar). Por ejemplo, si llego en tiempo t =
2 a un vértice v, veo la arista hacia w y aplico la función que se encuentra en la arista (v, w) con tiempo 2, y nos da como resultado
4, significa que en este caso llegaríamos a w en tiempo 4 (no que nos tomaría 4 para llegar, resultando en 6 en total). Implementar el
algoritmo de Dijkstra con las modificaciones necesarias para que funcione para un grafo de estas características, que permita obtener
el mínimo tiempo para llegar a todos los vértices desde un origen. Indicar y justificar la complejidad del algoritmo implementado
(considerar que las funciones de peso ejecutan en tiempo constante).
*/

def dijkstra_modif(grafo, origen):
	tiempos = {}

	for v in grafo.obtener_vertices():
		tiempos[v] = float("inf")
	
	padres = {}
	tiempos[origen] = 0
	padres[origen] = None
	heap_min = Heap()
	heap_min.Encolar((0, origen))

	while not heap_min.EstaVacia():
		tiempo_v, v = heap_min.Desencolar()
		if tiempo_v > tiempos[v]:
			continue // no me sirve xq no hace mas chico al camino
		
		for w in grafo.adyacentes(v):
			func = grafo.peso(v, w)
			tiempo_w = func(tiempo_v)

			if tiempos[w] > tiempo_w:
				tiempos[w] = tiempo_w
				padres[w] = v
				heap_min.Encolar((tiempo_w, w))
	
	return tiempos, padres
/*
5. Implementar un algoritmo de determine la cantidad de veces que aparece un elemento en un arreglo ordenado, en O(log n). Justificar
adecuadamente la complejidad del algoritmo implementado. Si te resulta muy difícil de resolver, podés considerar que el elemento
aparece a lo sumo O(log n) veces (en caso de usar esta presunción, el ejercicio estará para a lo sumo B=). Explicar dónde/cómo se
usaría esto, en caso de haberlo hecho.
*/

func contarArrOrdenado(arr, ini, fin, elem int) int{
	if len(arr) == 0{
		return 0
	}

}