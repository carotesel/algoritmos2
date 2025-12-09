/*
1. Implementar una función que reciba un arreglo A de n enteros y un número k,
   y devuelva un nuevo arreglo R tal que, para cada posición i de R, el valor
   sea la máxima suma obtenible de EXACTAMENTE k elementos dentro del rango
   A[0..i] (incluyendo a i). Si para una posición i no hay suficientes
   elementos (i < k-1), entonces R[i] debe valer -1.

   Ejemplo:
     A = [1, 5, 3, 4, 2, 8], k = 3
     Resultado esperado: R = [-1, -1, 9, 12, 12, 17]

   Requisito de eficiencia:
     La complejidad del algoritmo debe ser MEJOR que O(n * k).

   Indicar y justificar la complejidad del algoritmo implementado.
*/

func cmp(a, b int) int{
	return a - b
}

func SumasMaximasDeKElementos(arr []int, k int) []int {

	heapMin := CrearHeap(cmp)
	res := make([]int, 0)
	suma := 0

	for i, elem := range arr{
		
		if i < k{
			heapMin.Encolar(elem)
			suma += elem

			if i < k - 1{
				res = append(res, -1)
			} else{
				res = append(res, suma)
			}
		}

		if i >= k-1{
			min := heapMin.VerMin()

			if elem > min{
				eliminado := heapMin.Desencolar()
				suma -= eliminado
				heapMin.Encolar(elem)
				suma += elem
			}

			res = append(res, suma)
		}
	}
	return res
}