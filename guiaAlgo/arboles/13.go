/*(★★★★) Implementar una función que reciba un arreglo ordenado y devuelva un arreglo 
o lista con los elementos en orden para ser insertados en un ABB, 
de tal forma que al insertarlos en dicho orden se asegure que el ABB quede balanceado. 

¿Cómo cambiarías tu resolución si en vez de querer guardarlos en un ABB se fueran a insertar en un AVL?*/

// CADA MITAD DEBE ESTAR BALANCEADA

// balanceado -> preorder -> medio, izq, der

func balancear[K comparable](claves []K) Lista[K] {

	lista := CrearListaEnlazada[K]()

	if len(claves)==0{
		return lista
	}

	medio := len(claves) / 2

	lista.AgregarUltimo(claves[medio])

	izq := balancear(claves[:medio])
	der := balancear(claves[medio+1:])

	// itero lista izq y der y agrego a la lista a devolver
	izq.Iterar(func (elem K) bool {
		lista.InsertarUltimo(elem)
		return true
	})

	der.Iterar(func (elem K) bool {
		lista.InsertarUltimo(elem)
		return true
	})

	return lista
}

// ⭐ Complejidad

//Cada elemento es procesado exactamente una vez → O(n).

//Las llamadas recursivas hacen slicing que cuesta O(1), ya que son views, no copias.

//La misma solución sirve para AVL. Sin embargo, un AVL se balancearía solo incluso si el orden no fuera el óptimo.
