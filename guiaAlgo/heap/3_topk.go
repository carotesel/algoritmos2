/*(★★) Si en el ejercicio anterior en vez de quererse los 3 elementos más grandes, 
se quisieran los K elementos más grandes ¿cómo se debería proceder? 
¿Cuál terminaría siendo la complejidad del algoritmo?*/

func cmp(a, b int) int{
	return b - a
}

func (h heap[T]) TopK (k int) []T{
	res := make([]T, k)

	heapAux := CrearHeapArr(h.datos, cmp) 

	for i:=0; i < k; i++{ // k * O(log n)
		res[i] = heapAux.Desencolar()
	}

	return res
}

// O(n + k log n)