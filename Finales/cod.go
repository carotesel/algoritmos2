func noRepetido(arr []int) int {
	return noRep(arr, 0, len(arr)-1)
}

func noRep(arr []int, ini, fin int) int {
	// Caso base
	if ini == fin {
		return arr[ini]
	}

	medio := (ini + fin) / 2

	// Si hay al menos 3 elementos, intentamos detectar el distinto en el medio
	if medio > ini && medio < fin {
		if arr[medio] == arr[medio+1] && arr[medio-1] != arr[medio] {
			return arr[medio-1]
		}

		if arr[medio] == arr[medio-1] && arr[medio+1] != arr[medio] {
			return arr[medio+1]
		}

		if arr[medio+1] == arr[medio-1] && arr[medio] != arr[medio+1] {
			return arr[medio]
		}
	}

	// División y conquista: buscar en ambas mitades
	izq := noRep(arr, ini, medio)
	der := noRep(arr, medio+1, fin)

	// Combinar resultados
	if izq != der {
		// El distinto aparece una sola vez
		if count(arr, ini, fin, izq) == 1 {
			return izq
		}
		return der
	}

	return izq
}

func count(arr []int, ini, fin, val int) int {
	c := 0
	for i := ini; i <= fin; i++ {
		if arr[i] == val {
			c++
		}
	}
	return c
}

// Complejidad: O(n)

/*
2. Implementar una función que reciba un arreglo A de n enteros y un número k y devuelva un nuevo arreglo en el que
para cada posición i de dicho arreglo, contenga el resultado de la multiplicación de los primeros k máximos del arreglo A
entre las posición [0;i] (incluyendo a i). Las primeras k − 1 posiciones del arreglo a devolver deben tener como valor -1.
Por ejemplo, para el arreglo [1, 5, 3, 4, 2, 8] y k = 3, el resultado debe ser [-1, -1, 15, 60, 60, 160]. Indicar
y justificar la complejidad del algoritmo implementado.
*/

func primerosKmax (a []int, k int) []int{
	n := len(a)
	nuevo := make([]int, 0, n)

	for i:=0; i<n; i++{ // O(n)
		if i < k-1{
			nuevo = append(nuevo, -1)
		} else{ // O(i + k log i)
			heap := CrearHeapArr[int](cmp, a[:i+1]) // O(i)
			nro := 1

			for j:=0; j<k; j++{ // O(k log i)
				nro = nro * heap.Desencolar()
			}
			nuevo = append(nuevo, nro)
		}
	}
	return nuevo
}

/////////////////////////////


func (abb *ABB[K, V]) niveles() Diccionario[int, Lista[K]]{
	hash := CrearHash[int, Lista[K]]()
	niv(abb.raiz, hash)
	return hash
}

func niv (nodo *nodoABB[K, V], dicc Diccionario[int, Lista[K]]){
	cola := CrearColaEnlazada[*nodoABB[K, V]]()
	cola.Encolar(nodo)
	nivel := 0

	for !cola.EstaVacia(){
		nodo := Cola.Desencolar()
		lista := dicc.Obtener(nivel)
		lista.AgregarUltimo(nodo.clave)
		dicc.Guardar(nivel, lista)

		if nodo.izq != nil{
			cola.Encolar(nodo.izq)
		}

		if nodo.der != nil{
			cola.Encolar(nodo.der)
		}
	}
}

/*Definimos a la reducción de una secuencia de números como la operación de sacar los dos elementos más pequeños de la misma, volver a guardar el resultado de **2 * mínimo + segundo_mínimo**. Esto sólo puede hacerse si hay al menos dos elementos.

Ejemplo:

`[1, 7, 2, 3] → [7, 3, 4]`

De esta forma la secuencia queda con un elemento menos.

Implementar un algoritmo que reciba un arreglo y devuelva el único valor que quedaría en el arreglo si aplicamos dicha reducción hasta que quede un único elemento en el arreglo.

Indicar y justificar la complejidad del algoritmo.*/

func cmp(a, b int) int{
	return a - b
}

func reducir(arr []int) int{
	heap_min := CrearHeapArr(cmp, arr)

	for heap_min.Cantidad() > 1{
		min := heap_min.Desencolar()
		min_2 := heap_min.Desencolar()
		nro := (2 * min) + min_2
		
		heap_min.Encolar(nro)
		pila.Apilar(nro)
	}
	return heap_min.Desencolar()
}
// O(n + n log n)