/*2.*/

func arrOrdenado(arr []int) bool{
	return ordenado(arr, 0, len(arr)-1)
}

func ordenado(arr []int, ini, fin int) bool{
	if ini >= fin{
		return true
	}
	medio := (ini + fin)/2
	izq := ordenado(arr, ini, medio)
	der := ordenado(arr, medio+1, fin)
	frontera := arr[medio] < = arr[medio+1]
	return izq && der && frontera
}

// o(n)

/*
3. Explicar detalladamente cómo modificarías la implementación del ABB para poder tener una primitiva Maximo y Minimo que nos
devuelva las claves máximas y mínimas, y que se realice en tiempo constante.

Lo que haria para implementar las primitivas Maximo() y Minimo() en el ABB es lo siguiente:

Primero tendria 2 punteros, que sean referencias al nodo maximo y minimo del arbol. Dichos punteros se actualizarian
a medida que yo agrego o elimino elementos del arbol. 

Si yo agrego un elemento, lo compararia con el maximo y con el minimo, y si el mismo fuera mas grande que el maximo o mas chico
que el minimo, el mismo se actualizaria (osea reemplazaria la referencia) con el nuevo correspondiente. 

Si yo elimino un elemento y el mismo fuera el maximo o el minimo del arbol, lo actualizaria con la referencia al anterior (osea el
segundo maximo o el segundo minimo). Pregunta a chatgpt, que onda la complejidad?
*/

/*
4. Cuando programamos un módulo en Go, tenemos un archivo go.mod que nos indica las dependencias del proyecto. Asímismo, esas
dependencias tienen sus propios go.mod que nos indican sus propias dependencias. Para compilar nuestro proyecto, Go debe traer
(i.e. descargar) y compilar todas nuestras dependencias, así como las dependencias transitivas. Si el módulo A depende del B, es
necesario sí o sí compilar primero el módulo B antes que el A.
Explicar detalladamente cómo modelarías este problema con grafos, e implementar un algoritmo que reciba dicha grafo y nos devuelva
un orden correcto para compilar el proyecto entero (de forma correcta). Indicar y justificar la complejidad de la función implementada,
en función de las variables del problema.

ORDEN TOPOLOGICO. 
O -> O -> O

Grafo de dependencias. Dirigido. No pesado. 
*/

def grados_entrada(grafo):
	g_entrada = {}

	for v in grafo.vertices():
		g_entrada[v] = 0
	
	for v in grafo.vertices():
		for w in grafo.adyacentes(v):
			g_entrada[w] += 1
	
	return g_entrada

def orden_top(grafo):
	g_entrada = grados_entrada(grafo)
	cola = deque()
	res = []

	for v in grafo.vertices():
		if g_entrada[v] == 0:
			cola.append(v)
	
	while len(cola) > 0:
		actual = cola.popleft()
		res.append(actual)

		for w in grafo.adyacentes(actual):
			g_entrada[w] -= 1

			if g_entrada[w] == 0:
				cola.append(w)
	return res


// Complejidad: O(V + E)

/*
5. Implementar una primitiva para el hash cerrado filtro(func(V) bool) que elimine del hash todas las claves del mismo que
tengan asociado un valor para el cuál la función devuelva false. Indicar y justificar la complejidad de la primitiva 
implementada.
Ejemplo: si mi diccionario es de cadenas a números como el de arriba, e invocamos a la primitiva con una función que 
devuelve true para los números pares y false para los impares, el diccionario debe quedar como el de abajo:
{"koala": 3, "rana": 2, "gato": 2, "perro": 5, "canguro": 4 }

↓

{"rana": 2, "gato": 2, "canguro": 4}
*/

func (dicc *hashCerrado[K, V]) Filtro(f func(V) bool){
	tabla := dicc.tabla
	nuevaTabla := make([]CeldaHash[K, V], dicc.tam)
	dicc.tabla = nuevaTabla
	
	for _, celda := range tabla{
		if celda.estado == OCUPADO{
			if f(celda.valor){
				dicc.Guardar(celda.clave, celda.valor)
			}
		}
	}

}
