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