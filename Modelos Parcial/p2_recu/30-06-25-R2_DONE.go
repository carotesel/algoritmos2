/*
1. Implementar una primitiva Filter(f func(T) bool) para el Heap, que deje al heap únicamente con los elementos para
los que la función devuelva true. La primitiva debe funcionar en O(n), siendo n la cantidad de elementos inicialmente
en el heap. Por supuesto, luego de aplicar la operación, el heap debe quedar en un estado válido para poder seguir
operando. Justificar la complejidad de la primitiva implementada.
*/

func (heap *heap[T]) Filter(f func(T) bool) {

  if heap.EstaVacia(){
    return 
  }
  
  for i, elem := range heap.datos{
    modificado = f(elem)
    heap.datos[i] = modificado
  }

  heapify(heap, heap.cmp)
}

/*
2) Implementar una función eliminarRepetidos(arreglo []int) []int que, dado un arreglo de números,
   devuelva otro en el que estén los elementos del original sin repetidos.
   La primera aparición de cada número debe mantenerse, y las demás deben ser descartadas.

   Indicar y justificar la complejidad del algoritmo implementado.
*/

func eliminarRepetidos(arr []int) []int{
  dicc := CrearDiccionario[int, int]()
  res := make([]int, 0)

  for _, elem := range arr{
    if !dicc.Pertenece(elem){
      dicc.Guardar(elem, 1)
      res = append(res, elem)
    }
  }
  return res
}

// O(n)

/*
3) Se tiene un árbol binario que representa la fase eliminatoria del mundial.
   En cada nodo se guarda el nombre del país y la cantidad de goles que convirtió en dicha fase
   (incluyendo la tanda de penales, si fuera necesario).

   El padre del nodo debe tener como valor al país que ganó entre sus dos hijos
   (es decir, aquel que tuvo mayor cantidad de goles).

   Implementar una primitiva que, dado un árbol donde solamente están los nombres de los equipos
   en las hojas (no en los nodos internos), complete el árbol asignando a cada nodo interno
   el nombre del equipo ganador entre sus dos hijos.

   Se puede asumir que el árbol es completo, o que al menos todos los nodos internos
   tienen exactamente 2 hijos.

   La cantidad de goles en la raíz no es relevante.

   La estructura del árbol es:

     type Arbol struct {
         pais  string
         goles int
         izq   *Arbol
         der   *Arbol
     }

   Tomando el ejemplo del dorso, si invocamos la función para el árbol de la izquierda,
   debe quedar como el árbol de la derecha, donde cada nodo interno contiene el nombre
   del equipo ganador en su respectiva fase.
*/

func (ab *ab[K, V]) completarNombres() {
  if ab == nil{
    return
  }

  ab.izq.completarNombres()
  ab.der.completarNombres()

  if (ab.izq != nil && ab.der != nil) && ab.izq.goles < ab.der.goles{
    ab.pais = ab.izq.pais
  } else{
    ab.pais = ab.der.pais
  }
}

