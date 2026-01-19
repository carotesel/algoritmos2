/*
1. Explicar en detalle cómo implementar una estructura de pila enlazada como implementación del TDA Pila tal que asegure que todas
las primitivas del TDA se ejecuten en tiempo constante. 

PILA -> PRIMERO
COLA -> ULTIMO
*/

/*
Para implementar una pila enlazada en tiempo constante se utiliza una lista enlazada simple manteniendo una referencia al tope 
de la pila. Cada nodo contiene el dato y un puntero al siguiente nodo.
La operación Apilar se implementa creando un nuevo nodo cuyo siguiente apunta al antiguo tope, y luego actualizando el puntero 
al tope.
La operación Desapilar se realiza accediendo al nodo apuntado por el tope, guardando su dato y actualizando el puntero al 
siguiente nodo.
VerTope accede directamente al dato del nodo tope, y EstaVacia se implementa verificando si el puntero al tope es nil.
Todas las primitivas se ejecutan en tiempo constante O(1) ya que NO requieren recorrer la estructura.
*/

/*
2. Se tiene un arreglo rutas, donde cada posición rutas[i] nos da la la ruta del colectivo de la línea i, que se repite para siempre. Por
ejemplo, si rutas[0] = [4, 7, 9, 6], significa que la ruta de los colectivos de la línea 0 es 4 -> 7 -> 9 -> 6 -> 4 -> 7 -> 9
-> 6 -> 4 -> ..., donde cada número es un número de parada, definido de alguna manera no relevante a este enunciado. Sólamente
se puede trasladar entre paradas utilizando colectivos (es decir, nada de a pie ni por otros medios). Por ejemplo, si rutas[8] contiene
la parada 9, entonces podríamos hacer un trasbordo de la línea 0 a la 8 en dicha parada. Se quiere, dada una parada inicial y una
parada final, obtener la cantidad mínima de paradas a realizar para llegar desde la parada inicial a la final. Teniendo todas
las rutas de todas las líneas de colectivo, modelar el problema con grafos (explicando claramente qué son los vértice y aristas, y
sus características), enunciar y explicar un algoritmo que nos permita obtener lo deseado (no es necesario implementar). Indicar y
justificar la complejidad en función de las variables del problema.

Se modela el problema mediante un grafo dirigido y no pesado donde cada vértice representa una parada.
Existe una arista dirigida desde una parada A hacia una parada B si en alguna ruta de colectivo la parada B aparece inmediatamente después de A.
Dado que las rutas se repiten indefinidamente, cada ruta genera un ciclo en el grafo.
Para obtener la mínima cantidad de paradas entre la parada inicial y la final, se realiza un recorrido BFS desde el vértice correspondiente a la parada inicial.
Al ser un grafo no pesado, BFS garantiza encontrar el camino con menor cantidad de aristas, que corresponde a la menor cantidad de paradas recorridas.

Complejidad: O(V + E)
*/

/*
4. Implementar un algoritmo que dado un grafo no dirigido determine si el mismo es conexo. Se pide implementar utilizando un recorrido
DFS. Indicar y justificar la complejidad del algoritmo.
*/

import random 

def es_conexo(grafo):
	if len(grafo.obtener_vertices()) == 0:
		return True
	visitados = set()
	origen = random.choice(grafo.obtener_vertices())
	_dfs(grafo, origen, visitados)
	return len(visitados) == len(grafo.obtener_vertices())
	

def _dfs(grafo, v, visitados):
	visitados.add(v)
	for w in grafo.adyacentes(v):
		if w not in visitados:
			_dfs(grafo, w, visitados)

// Complejidad: O(V + E)

/*
5. Implementar una primitiva para el árbol binario que determine si cumple la propiedad de Heap. 
Indicar y justificar la complejidad del algoritmo implementado.
ASUMO PROP HEAP MAXIMOS
*/

func (ab *Arbol) PropHeap() bool{
	if ab == nil{
		return true
	}

	if ab.izq != nil{
		if ab.valor < ab.izq.valor{
		return false
		}
	}

	if ab.der != nil{
		if ab.valor < ab.der.valor{
		return false
		}
	}

	return ab.izq.PropHeap() && ab.der.PropHeap()
}
