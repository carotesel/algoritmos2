/*(★★★) Implementar una primitiva del ABB que dado un valor entero M, una clave inicial inicio 
y una clave final fin, se devuelva una lista con todos los datos cuyas claves estén entre inicio y fin, 
que estén dentro de los primeros M niveles del árbol (considerando a la raíz en nivel 1). 

Indicar y justificar la complejidad temporal.

Por ejemplo, si tenemos el siguiente ABB con M = 3, inicio = 5 y fin = 15:

      10
    /    \
   5      15                    Un resultado final serían los datos de las 
  / \    /  \                   claves 10, 5, 8, 15, 12 (en cualquier orden).
 3   8  12   20
    /     \
   7       14

   type abb[K comparable, V any] struct {
    izq *abb[K, V]
    der *abb[K, V]
    clave K
    dato V
}
*/

func[] (ab *abb[K, V]) datosRango(M int, ini, fin K) Lista[K]{
	lista := CrearListaEnlazada[K]()

	recorrido(ab, 1, M, ini, fin, lista)

	return lista
}

func recorrido(nodo *nodo[K, v], nivel, M int, ini, fin K, lista Lista[V]){
	if nodo == nil{
		return
	}

	if nivel > M{
		return 
	}

	if nodo.clave > ini{
		recorrido(nodo.izq, nivel+1, M, ini, fin, lista)
	}

	if nodo.clave >= ini && nodo.clave <= fin{
		lista.AgregarUltimo(nodo.dato)
	}

	if nodo.clave < fin{
		recorrido(nodo.der, nivel+1, M, ini, fin, lista)
	}
}
