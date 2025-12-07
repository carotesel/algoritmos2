/*
2) Implementar una función eliminarRepetidos(arreglo []int) []int que, dado un arreglo de números,
   devuelva otro en el que estén los elementos del original sin repetidos.
   La primera aparición de cada número debe mantenerse, y las demás deben ser descartadas.

   Indicar y justificar la complejidad del algoritmo implementado.
*/


// Tal como se puede observar la complejidad total del algoritmo es de O(n),
// dado que se realizan ciclos recorriendo a lo sumo n elementos
// aplicando operaciones constantes (O(1)).

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

