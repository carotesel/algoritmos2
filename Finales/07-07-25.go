/*
1. Implementar una función func minimoExcluido(arr []int) int que dado un arreglo de valores enteros (mayores o iguales a 0),
obtenga el mínimo valor que no se encuentre en el arreglo. Indicar y justificar la complejidad del algoritmo (explicar en detalle este
paso, porque es fácil que se te puedan pasar detalles importantes a explicar). ¿Es el mismo ejercicio del parcialito? Si.
Por ejemplo:
minimoExcluido([]int{0, 5, 1}) --> 2
minimoExcluido([]int{3, 5, 1}) --> 0
minimoExcluido([]int{0, 5, 1, 3, 4, 1, 2}) --> 6
minimoExcluido([]int{0, 5, 1, 3, 4, 1, 2, 12345675433221345}) --> 6
*/

func minimoExcluido(arr []int) int{
	dicc := CrearHash[int, bool]()

	for _, elem := range arr{
		dicc.Guardar(elem, true)
	}

	for i:=0; ; i++{
		if !dicc.Pertenece(i){
			return i
		}
	}
}

// Complejidad: O(n) porque en el peor de los casos recorri todo.

/*
2. Implementar un algoritmo que dado Grafo no dirigido nos devuelva su complemento. Es decir, un grafo en el que 
una arista (v, w) significa que v y w no son adyacentes en el grafo original. 
Indicar y justificar la complejidad del algoritmo implementado.
*/

// hayArista
// agregarArista

// Grafo(esDirigido = false, vertices = grafo.obtenerVertices())

def complemento(grafo):
	nuevo = Grafo(esDirigido = false, vertices = grafo.obtener_vertices())

	for v in grafo.vertices():
		for w in grafo.vertices():
			if v != w and (not grafo.hayArista(v, w)):
				nuevo.agregarArista(v, w)
	
	return nuevo

/*
3. Se tiene un arreglo rutas, donde cada posición rutas[i] nos da la la ruta del colectivo de la línea i, que se repite para siempre. Por
ejemplo, si rutas[0] = [4, 7, 9, 6], significa que la ruta de los colectivos de la línea 0 es 4 -> 7 -> 9 -> 6 -> 4 -> 7 -> 9
-> 6 -> 4 -> ..., donde cada número es un número de parada, definido de alguna manera no relevante a este enunciado. Sólamente
se puede trasladar entre paradas utilizando colectivos (es decir, nada de a pie ni por otros medios). Por ejemplo, si rutas[8] contiene
la parada 9, entonces podríamos hacer un trasbordo de la línea 0 a la 8 en dicha parada. Se quiere, dada una parada inicial y una
parada final, obtener la cantidad mínima de colectivos a tomar para llegar desde la parada inicial a la final. Es decir, queremos
minimizar la cantidad de trasbordos a realizar. Teniendo todas las rutas de todas las líneas de colectivo, modelar el problema con
grafos (explicando claramete qué son los vértice y aristas, y sus características), enunciar y explicar un algoritmo que nos permita
obtener lo deseado (no es necesario implementar). Indicar y justificar la complejidad en función de las variables del problema.
*/

// explicado en papel

/*
4. Implementar una función map[T any, V any](Lista[K], func(K) V) Lista[V] que dada una lista original, cree una nueva lista
con el resultado de aplicarle a cada elemento la función pasada por parámetro. Para que el ejercicio esté completamente bien, se
espera que se implemente utilizando el iterador interno de la lista. Indicar y justificar la complejidad de la función.
*/

func map[K any, V any](lista Lista[K], f func(K) V) Lista[V]{
	nuevaLista := CrearLista[V]()

	lista.Iterar(func (clave K) bool{
		nuevaLista.AgregarUltimo(f(clave))
		return true
	})

	return nuevaLista
}

// Complejidad: O(n)