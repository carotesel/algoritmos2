/*(★★★) 
♠ Implementar un algoritmo que reciba un arreglo desordenado de enteros, su largo (n) y un número K y determinar 
en O(n) si existe un par de elementos en el arreglo que sumen exactamente K.*/

func HaySumaK (arr []int, n, k int) bool{
	dicc := CrearDiccionario[int, bool]()

	for i := 0; i < n; i++ { // O(n)
		complemento := k - arr[i]
		if dicc.Pertenece(complemento){ // O(1)
			return true
		} else{
			dicc.Guardar(arr[i], true)
		}
	}
	return false
}

// Complejidad: O(n)