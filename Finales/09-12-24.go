/*
1. Queremos hacer una estructura de mensajería para producir y consumir. Tiene las primitivas de Producir y Consumir. Cuando se
manda a Producir, se crea un elemento (esto puede hacerse recibiendo una función de producción de elementos en la función que
crea nuestro TDA). Al Consumir, se consume (devuelve) el primer elemento producido (luego el segundo, etc..). ¿Qué estructura
de datos utilizarías internamente para manejar eso? 

Ahora, se le quiere agregar una primitiva Deshacer. Dicha primitiva debe
reincorporar el último elemento eliminado, pero dicha operación se debe poder utilizar hasta K veces, siendo K la cantidad de elemento
que ya fueron consumidos hasta ese momento (luego, si se consumen más elementos, se podrá seguir deshaciendo). Primero debe
reincorporarse el último en ser consumido, luego el anteúltimo, etc. . . 

Una vez reincorporado, el elemento reincorporado debe ser
el siguiente en consumirse. ¿Cómo podemos hacer para poder tener nuestra primitiva deshacer, y que el resto siga funcionando
correctamente? Justificar, basándote en los temas vistos en la materia.
*/

/*
Para la mensajería se utiliza una cola, ya que permite consumir los elementos en el orden en que fueron producidos (FIFO).

Para implementar la primitiva Deshacer, se utiliza una pila donde se almacenan los elementos a medida que son consumidos. Dado que la pila es LIFO, permite reincorporar primero el último elemento consumido, cumpliendo con el orden requerido.

Al deshacer, el elemento se vuelve a insertar de manera que sea el próximo en consumirse, por ejemplo utilizando una cola con inserción en ambos extremos o una estructura equivalente. De esta forma, el resto de las primitivas continúa funcionando correctamente.
*/

/*
2. Implementar un algoritmo que reciba un grafo no dirigido y pesado (todos pesos positivos), un vértice de inicio y una función
estaHabilitada(org, dst, tiempo) bool, y obtenga los caminos mínimos hacia los demás vértices, con una modificación: las
aristas no están habilitadas en todo momento y, por lo tanto, no pueden utilizarse cuando no estén habilitadas. Sí pueden volver
a habilitarse en otro tiempo posterior (considerar siempre tiempos discretos). Indicar y justificar la complejidad del algoritmo
implementado. Para esto, suponer que los los valores de los pesos son pequeños. . . si no lo fueran, ¿Cuál podría ser el problema?
*/

ddef dijkstra(grafo, origen):

    padres = {origen: None}
    distancias = {}

    for v in grafo:
        distancias[v] = float("inf")
    
    distancias[origen] = 0
    heap_min = Heap()
    heap_min.Encolar((0, origen))

    while not heap_min.EstaVacia():
        dist_v, v = heap_min.Desencolar()

        if dist_v > distancias[v]:
            continue
        
        for w in grafo.adyacentes(v):

            tiempo_actual = dist_v

            //buscar primer tiempo habilitado >= tiempo_actual
            tiempo_salida = tiempo_actual
            while not estaHabilitada(v, w, tiempo_salida):
                tiempo_salida += 1   # esperamos

            nuevo_tiempo = tiempo_salida + grafo.peso_arista(v, w)

            if nuevo_tiempo < distancias[w]:
                distancias[w] = nuevo_tiempo
                padres[w] = v
                heap_min.Encolar((nuevo_tiempo, w))
    
    return distancias, padres


	/*
	4. Implementar una función que reciba un arreglo de números y determine cuáles aparecen una única vez. Indicar y justificar la
complejidad del algoritmo implementado.
	*/

	func unaVez(arr []int) []int{
		dicc := CrearHash[int, int]()

		for _, elem := range arr{ // O(n)
			if !dicc.Pertenece(elem){
				dicc.Guardar(elem, 1)
			} else{
				dicc.Guardar(elem, dicc.Obtener(elem)+1)
			}
		}

		iter := dicc.iterador()
		res := make([]int, 0)

		for iter.HaySiguiente(){ // O(n)
			clave, valor := iter.VerActual()
			if valor == 1{
				res = append(res, clave)
			}
			iter.Siguiente()
		}
		return res
	}

	/*
5. Implementar una primitiva para el árbol binario EsABB(func(T, T) int) bool que reciba una función de comparación y determine
si el árbol cumple con la propiedad de ABB para dicha función de comparación. Indicar y justificar la complejidad del algoritmo
implementado.
A fines del ejercicio, considerar que la estructura del árbol es la indicada en el dorso a este examen.

type arbol[T any] struct {
izq *arbol
der *arbol
clave T
}
*/

// ESABB SIEMPRE CON RANGOS!
// A LA IZQ EL MAXIMO ES LA RAIZ, SI ALGO LA PASA NO CUMPLE
// A LA DER EL MINIMO ES LA RAIZ, SI HAY ALHO MAS CHICO NO CUMPLE

func (abb *abb[T]) EsABB(cmp func(T, T) int) bool {
	return abbRec(abb, cmp, nil, nil)
}

func abbRec(abb *abb[T], cmp func(T, T) int, min, max *T) bool{
	if abb == nil{
		return true
	}

	if min != nil && cmp(abb.clave, min) <= 0{
		return false
	}

	if max != nil && cmp(abb.clave, max) >= 0{
		return false
	}

	return abbRec(abb.izq, cmp, min, &abb.clave) && abbRec(abb.der, cmp, &abb.clave, max)
}