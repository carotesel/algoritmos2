/*
1. Tenemos un mapa de caminos rurales que conectan diferentes ciudades, donde algunos de estos caminos se encuentran
bloqueados por alguna razón (un árbol caido, una piedra que cayó desde una montaña, los festejos de boquita campeón
quemando un local de comida rápida, etc. . . ). Cada bloqueo cuesta diferente de remover. Sabemos que algunos de
estos bloqueos no nos impiden llegar desde una ciudad a otra, pero podría ser el caso que sí. Queremos implementar un
algoritmo que nos determine qué bloqueos deben ser removidos para que se pueda llegar de cualquier ciudad a cualquier
otra. Para esto, primero modelar el problema con grafos, y luego implementar un algoritmo que reciba dicho grafo y nos
devuelva los bloqueos a eliminar. Indicar y justificar la complejidad del algoritmo implementado.
*/

def prim_modif(grafo):
	origen = random.choice(grafo)
	visitados = set()
	visitados.add(origen)
	bloqueos = []
	heap = Heap()

	for w in grafo.adyacentes(origen):
		heap.Encolar((origen, w, grafo.peso(origen, w)))
	
	while not heap.EstaVacia():
		v, w, peso = heap.Desencolar()

		if w in visitados:
			continue
		
		visitados.add(w)
		
		if peso > 0:
			bloqueos.append((v, w, peso))
		
		for x in grafo.adyacentes(w):
			if x not in visitados:
				heap.Encolar((w, x, grafo.peso(w, x)))
	
	return bloqueos

// O(E log V)

/*
2. Implementar un algoritmo que obtenga la parte entera de la raíz de un número n entero en O (log n). Justificar la
complejidad de la primitiva implementada.
*/

func obtenerRaiz(n int) int{
	return raiz(0, n, n)
}

func raiz(ini, fin, n int) int{
	if ini > fin{
		return fin // valor mas grande que NO se paso del rango
	}

	medio := (ini + fin) / 2

	if medio * medio == n{
		return medio
	}
	else if medio * medio > n{
		return raiz(ini, medio-1, n)
	} else{
		return raiz(medio+1, fin, n)
	}
}

/*
3. Implementar una primitiva para una Cola implementada como una estructura en arreglo (como la vista en clase),
Filtrar[T](func condicion(T) bool) Cola[T] que devuelva una nueva cola para la cual los elementos de la cola
original dan true en la función condicion pasada por parámetro. La cola original debe quedar intacta, y los elementos
de la final deben tener el orden relativo que tenían en la original. Indicar y justificar la complejidad del algoritmo
implementado.
*/

func Filtrar(cola *ColaEnlazada[T])(func condicion(T) bool) Cola[T]{
	actual := cola.primero
	nueva := CrearColaEnlazada[T]()

	for actual != nil{ 
		valor := condicion(actual.dato)
		if valor{
			nueva.Encolar(actual.dato)
		}
		actual = actual.siguiente
	}

	return nueva
}

// Complejidad: O(n)

/*
4. Implementar un algoritmo que dado un texto, devuelva cuál es la palabra más frecuente del mismo. Indicar y justificar
la complejidad del algoritmo implementado. Nota: recordar que existe la función split(cadena, separador), que
funciona en O(m), siendo m el largo de la cadena.
*/

func palabraFrecuente(texto string) string{
	palabras := strings.Split(texto, " ")
	dicc := CrearHash[string, int]()

	for _, pal := range palabras{
		if dicc.Pertenece(pal){
			dicc.Guardar(pal, dicc.Obtener(pal)+1)
		} else{
			dicc.Guardar(pal, 1)
		}
	}

	max := 0
	palabra_max := ""

	iter := dicc.Iterador()

	for iter.HaySiguiente(){
		actual := iter.VerActual()
		if dicc.Obtener(actual) > max{
			max = dicc.Obtener(actual)
			palabra_max = actual
		}
		iter.Siguiente()
	}
	return palabra_max
}

// Complejidad O(k + m) ~= O(m)
// k = cantidad de palabras 
// m = cantidad de letras TOTALES
// k <= m (considerablemente mas chico)

/*5. Implementar un algoritmo que reciba un grafo dirigido, acíclico y pesado, un vértice v y otro w, y devuelva la longitud
del camino máximo. Indicar y justificar la complejidad del algoritmo implementado.*/

def camino_max(grafo, origen, fin):
	camino = []
	visitados = set()
	dfs(grafo, camino, visitados, origen)
	camino.reverse()

	dist = {}

	for v in grafo:
		dist[v] = float("inf")
	
	dist[origen] = 0
	padres = {origen: None}

	for v in camino:
		for w in grafo.adyacentes(v):
			peso = grafo.peso(v, w)

			if dist[v] + peso > dist[w]:
				padres[w] = v 
				dist[w] = dist[v] + peso
	
	return padres, distancias 

def dfs(grafo, camino, visitados, origen):
	visitados.add(origen)
	camino.append(origen)

	for w in grafo.adyacentes(origen):
		if w not in visitados:
			dfs(grafo, camino, visitados, w)


