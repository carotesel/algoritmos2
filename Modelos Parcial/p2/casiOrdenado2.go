/*3. Tenemos un arreglo de n números en el que cada elemento se encuentra a lo sumo k
posiciones de la que le correspondería si estuviera ordenado(2 ≤ k < n). Implementar una
función que reciba el arreglo y el valor de k y ordene el arreglo en O(n log k). Justificar la
complejidad del algoritmo implementado.*/

/*Traducido:

El elemento que debería ir en la posición 0 está entre las posiciones 0 y k

El elemento que debería ir en la posición 1 está entre las posiciones 1 y k+1

El que debería ir en la posición i está entre i y (i+k)
*/

func cmp (a, b int) int{
	return a - b
}

func arrCasiOrdenado(arr[]int, k int) []int{
	n := len(arr)
	heapMin := CrearHeapArr(arr[:k+1], cmp)
	res := make([]int, n)

	i := 0

	// desencolo los k+1 minimos y encolo los q faltan del arr
	for j := k+1; j < n; j++{
		res[i] = heapMin.Desencolar()
		i++
		heapMin.Encolar(arr[j])
	}

	// guardo los q quedan del heap hasta vaciarlo
	for !heapMin.EstaVacia(){
		res[i] = heapMin.Desencolar()
		i++
	}
	return res

}