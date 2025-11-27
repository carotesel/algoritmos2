/*3. Tenemos un arreglo de n números en el que cada elemento se encuentra a lo sumo k
posiciones de la que le correspondería si estuviera ordenado(2 ≤ k < n). Implementar una
función que reciba el arreglo y el valor de k y ordene el arreglo en O(n log k). Justificar la
complejidad del algoritmo implementado.*/

/*Traducido:

El elemento que debería ir en la posición 0 está entre las posiciones 0 y k

El elemento que debería ir en la posición 1 está entre las posiciones 1 y k+1

El que debería ir en la posición i está entre i y (i+k)
*/

func cmp(a, b int) int{
	return a - b
}

func arrCasiOrdenado(arr[]int, k int) []int{
	n := len(arr)
	heapMin := CrearHeapArr(arr[:k+1], cmp) // O(k+1) = O(k)
	ordenado := make([]int, n)

	index := 0

	// saco min meto k+1, k+2 etc en el heap
	for j:= k+1; j < n; j++{ // O (n−(k+1)) * O(log k) ≈ O(n log k)
		ordenado[index] = heapMin.Desencolar()
		index++
		heapMin.Encolar(arr[j])
	}

	for !heapMin.EstaVacia(){ // O(k+1) * log k
		ordenado[index] = heapMin.Desencolar()
		index++
	}
	
	return ordenado
}