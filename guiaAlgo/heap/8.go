/*(★★★) 
♠ Escribir una función en Go que, dado un arreglo de 
n cadenas y un entero positivo k, devuelva una lista con las 
k cadenas más largas. Se espera que el orden del algoritmo sea O(n+klogn). 
Justificar el orden.*/

func cmp(s1, s2 string) int{
	return len(s1) - len(s2)
}

func kMasLargas(arr []string, k int) []string{
	res := make([]string, k)

	if k <= 0 || k > len(arr) {
        return nil
    }

	heapAux := CrearHeapArr(arr, cmp) // O(n). USA HEAPIFY.

	for i:=0; i < k; i++{ // k * O(log n)
		res[i] = heapAux.Desencolar() // NO APPENDEAR PORQUE AGREGA ESPACIO
	}

	return res
}

// O(n+klogn)