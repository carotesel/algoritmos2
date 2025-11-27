/*(★) Implementar en lenguaje Go una función recursiva con la firma func esHeap(arr []int). 
Esta función debe devolver true o false de acuerdo a si el arreglo que recibe como parámetro 
cumple la propiedad de heap (de mínimos).

Hacer el seguimiento de la función para el arreglo [ 1, 7, 2, 8, 7, 6, 3, 3, 9, 10 ].*/

// En heap min: todo indice es < que su hijo izq (2*i +1) y der (2*i + 2)

func esHeap(arr []int) bool {
	return arrEsheap(arr, 0, len(arr))
}

func arrEsheap(arr []int, pos_actual, largo int) bool {
	if pos_actual >= largo {
		return true
	}

	izq := 2 * pos_actual + 1
	der := 2 * pos_actual + 2

	if izq < largo && arr[izq] < arr[pos_actual]{
		return false
	}

	if der < largo && arr[der] < arr[pos_actual]{
		return false
	}

	return arrEsheap(arr, izq, largo) && arrEsheap(arr, der, largo)

}

// la recursion revisa los hijos izq y der