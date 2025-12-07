/*
1) Implementar una primitiva para el ABB:

   func (arbol *Arbol[K, V]) AncestroComun(clave1, clave2 K) K

   que reciba dos claves y devuelva la **clave del último ancestro en común** (Lowest Common Ancestor, LCA)
   entre ambas en el ABB. El ancestro común puede coincidir con alguna de las claves pasadas.
   Si **alguna** de las claves no se encuentra en el árbol, finalizar con **panic**.

   Ejemplos (para el árbol del enunciado):
     arbol.AncestroComun(1, 4) --> 2
     arbol.AncestroComun(2, 4) --> 2
     arbol.AncestroComun(9, 1) --> 5

*/

// Indicar y justificar la complejidad: O(h) tiempo y O(h) espacio (por la recursión),
//donde h es la altura del árbol.

   func (abb *abb[K, V]) AncestroComun(clave1, clave2 K) K{

	if arbol == nil || !arbol.Pertenece(clave1) || !arbol.Pertenece(clave2){
		return panic("Alguna de las claves ni pertenece o el arbol esta vacio xd")
	}

	return ancestroNodo(abb.raiz, abb.cmp, clave1, clave2)

   }

   func ancestroNodo[K, V] (nodo *NodoABB[K, V], func cmp(K, K) int, clave1, clave2 K) K{
	if nodo == nil{
		panic("Error interno: nodo nulo durante la búsqueda del LCA")
	}

	if cmp(clave1, nodo.clave) < 0 && cmp(clave2, nodo.clave) < 0{
		return ancestroNodo(nodo.izq, cmp, clave1, clave2)
	}

	if cmp(clave1, nodo.clave) > 0 && cmp(clave2, nodo.clave) > 0{
		return ancestroNodo(nodo.der, cmp, clave1, clave2)
   	}

	return nodo.clave
}


/*
2) Implementar una función:

   func minimoExcluido(arr []int) int

   que, dado un arreglo de valores enteros (mayores o iguales a 0),
   obtenga el mínimo valor que NO se encuentre en el arreglo.
   Indicar y justificar la complejidad del algoritmo
   (explicar en detalle este paso, porque es fácil pasar por alto
   detalles importantes).

   Ejemplos:
     minimoExcluido([]int{0, 5, 1})                         --> 2
     minimoExcluido([]int{3, 5, 1})                         --> 0
     minimoExcluido([]int{0, 5, 1, 3, 4, 1, 2})             --> 6
     minimoExcluido([]int{0, 5, 1, 3, 4, 1, 2, 12345675433221345}) --> 6
*/

// cmp min
func cmp(a, b int) int{
	return a - b
}

func minimoExcluido(arr []int) int{
	heap := CrearHeapArr(cmp, arr) // O(n)
	esperado := 0

	for range arr{ // O(n)
		actual := heap.Desencolar()

		if actual < esperado{
			esperado++
		}

		if actual > esperado{
			return esperado
		}
	}
	return esperado
}

/*La complejidad es O(n) ya que en el peor de los casos, recorro todo el array y el minimo excluido es el "siguiente" 
que no aparece en el.*/
