/*
2. Implementar un algoritmo que reciba un grafo no dirigido y determine si el mismo tiene forma de estrella. Es decir,
si todos los vértices, salvo 1, se conectan al mismo vértice, mientras ese único vértice se conecta con todos los demás.
Indicar y justificar la complejidad del algoritmo si se implementara el grafo con una lista de adyacencia (diccionario de
diccionarios), y también si se hiciera con una matriz de adyacencia.
*/

def es_estrella(grafo):
	contador_n_adyacentes = 0
	cant_vertices = len(grafo.vertices())

	for v in grafo.vertices():
		if len(grafo.adyacentes(v)) == cant_vertices - 1:
			contador_n_adyacentes+=1
		elif len(grafo.adyacentes(v)) != 1: // ya descarto que no cumple porque el vertice no tiene n - 1 ady ni 1 exacto!
			return false
		
	return contador_n_adyacentes == 1

// Con matriz de ady la complejidad es O(v^2) dado que por cada v recorro todos los vertices 
// Con lista de ady es O(V) dado que en su implementacion para cada vertice recorro sus adyacentes. En este caso simplemente recorro cada 
// vertice 1 vez.

/*
3. Tenemos un arreglo de números de 1 a n, ordenado. A dicho arreglo se le quita un elemento. Implementar un algoritmo
que determine qué elemento falta en el arreglo. Indicar y justificar la complejidad del algoritmo implementado.
*/

func arrOrdenado(arr []int) int{
	faltante := buscarFaltante(arr, 0, len(arr)-1)
	return faltante
}

func buscarFaltante(arr []int, ini, fin int) int{

	if ini > fin{
		return ini + 1
	}

	medio := (ini+fin)/2

	if arr[medio] != medio+1{
		return buscarFaltante(arr, ini, medio-1)
	}
	else if arr[medio] == medio + 1{
		return buscarFaltante(arr, medio+1, fin)
	}

}

/*
4. Implementar en Go una primitiva que reciba un árbol binario que representa un heap (árbol binario izquierdista, que
cumple la propiedad de heap), y devuelva la representación en arreglo del heap. La firma de la primitiva debe ser
RepresentacionArreglo() []T. Indicar y justificar la complejidad de la primitiva. 

La estructura del árbol binario es:
type ab[T any] struct {
izquierda *ab[T]
derecha *ab[T]
dato T
}
*/

// prop izquierdista: RECORRIDO POR NIVELES -> "BFS"

func (ab *ab[T]) RepresentacionArreglo() []T{

	if ab == nil{
		return nil
	}

	res := make([]T, 0)
	cola := CrearColaEnlazada[*ab[T]]()
	cola.Encolar(ab)

	for !cola.EstaVacia(){
		nodo := cola.Desencolar()
		res = append(res, nodo.dato)
		
		if nodo.izquierda != nil{
			cola.Encolar(nodo.izquierda)
		}
		
		if nodo.derecha != nil{
			cola.Encolar(nodo.derecha)
		}
	}

	return res
}

/*
5. Implementar un algoritmo que permita ordenar cronológicamente un arreglo de cadenas que representan horarios en
formato HH:MM:SS. Indicar y justificar la complejidad del algoritmo implementado.
*/

// posiciones: 01:34:56

// bucket sort falopa

// Rango de valores: [1; 60] -> FINITO (1 bucket por cada seg). BUCKET SORT!

func ordenarCronologico(arr []string) []string{
	res := make([]string, 0, len(arr))
	buckets := make([][]string, 24*3600)

	for _, h := range arr{
		horas := atoi(h[0:2])
		mins := atoi(h[3:5])
		segs := atoi(h[6:8])
		total := horas*3600 + mins*60 + segs

		buckets[total] = append(buckets[total], h)
	}

	for i:=0; i < len(buckets); i++{
		res = append(res, buckets[i])
	}

	return res
}