/*2. Implementar una función func minimoExcluido(arr []int) int que dado un arreglo de valores enteros 
(mayores o iguales a 0), obtenga el mínimo valor que no se encuentre en el arreglo. I
ndicar y justificar la complejidad del algoritmo
(explicar en detalle este paso, porque es fácil que se te puedan pasar detalles importantes a explicar).

Por ejemplo:
minimoExcluido([]int{0, 5, 1}) --> 2
minimoExcluido([]int{3, 5, 1}) --> 0
minimoExcluido([]int{0, 5, 1, 3, 4, 1, 2}) --> 6
minimoExcluido([]int{0, 5, 1, 3, 4, 1, 2, 12345675433221345}) --> 6*/

// Cuando actual < esperado, significa que ese número ya lo procesé antes y no sirve para encontrar el mínimo excluido. 
// El candidato real sigue siendo esperado.

func cmp(a, b int) int{
	return a - b
}

func minimoExcluido(arr []int) int{
	heap:= CrearHeapArr(arr, cmp) // O(n)

	esperado := 0

	for !heap.EstaVacia{ // O(n)
		actual := heap.Desencolar() // O(Log(n))

		if actual == esperado{
			esperado += 1
			continue
		} else if actual < esperado{
			continue
		} else{
			return esperado
		}
	}
	return  // devuelve esperado + 1 xd
}

// O(n log n) porque en el peor de los casos el heap desencola todos los numeros del array y el minimo excluido es len(arr)