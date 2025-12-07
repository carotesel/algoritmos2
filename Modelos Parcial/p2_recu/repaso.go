/*
Implementar en Go una primitiva para el árbol binario:

    func (ab *Arbol[int]) EsSuma() bool

que determine si un árbol es suma o no.
Un árbol es suma cuando el valor de todo nodo padre es igual a la suma de sus hijos.
Se considera que si un nodo es inexistente su valor es 0.
Si un nodo es hoja, entonces se define que es suma.

Indicar y justificar la complejidad de la primitiva.
Al dorso del examen pueden ver tanto la estructura del árbol como algunos ejemplos.


type Arbol[T any] struct {
    dato T
    izq  *Arbol
    der  *Arbol
}

Ejemplo de árbol que **es suma**:

        10
       /  \
      5    5
     / \  / \
    2  3 4   1

(5 + 5 = 10)
(2 + 3 = 4 + 1 = 5)


Ejemplo de árbol que **no** es suma:

        20
       /  \
     12    21

(12 + 21 = 33 != 20)
*/


// -----------------------------------------------------------------------------
// Implementar para el Hash Cerrado la primitiva Limpieza(), la cual se encarga
// de eliminar todos los borrados, asegurándose de dejar al Hash en un estado
// correcto.
//
// Pista: pensar en las búsquedas de los elementos que efectivamente se
// encuentran en el hash.
//
// Indicar y justificar la complejidad de la primitiva implementada.
// -----------------------------------------------------------------------------


// -----------------------------------------------------------------------------
// 3. Implementar en Go una primitiva para el árbol binario:
//
//     func (ab *Arbol[T]) EsCompleto() bool
//
// que determine si el árbol es completo.
//
// Indicar y justificar la complejidad de la primitiva.
//
// A fines del ejercicio, considerar que el árbol está definido como:
//
//     type Arbol[T any] struct {
//         dato T
//         izq  *Arbol[T]
//         der  *Arbol[T]
//     }
//
// Completo:
//
//     10
//    /  \
//   5    5
//  / \  / \
// 2  3 4   1
//
// -----------------------------------------------------------------------------



/*
1) Implementar una función que reciba un arreglo A de n enteros y un número k,
   y devuelva un nuevo arreglo R tal que, para cada posición i de R, el valor
   sea la máxima suma obtenible de EXACTAMENTE k elementos dentro del rango
   A[0..i] (incluyendo a i). Si para una posición i no hay suficientes
   elementos (i < k-1), entonces R[i] debe valer -1.

   Ejemplo:
     A = [1, 5, 3, 4, 2, 8], k = 3
     Resultado esperado: R = [-1, -1, 9, 12, 12, 17]

   Requisito de eficiencia:
     La complejidad del algoritmo debe ser MEJOR que O(n * k).

   Indicar y justificar la complejidad del algoritmo implementado.
*/


// Como se puede apreciar la complejidad maxima en el algoritmos desarrollado es
// O(n.log k), dado que a lo sumo se recorren los n elementos del arreglo, aplicando
// a lo sumo operaciones O(log k), al Encolar o Desencolar en el heap

/*
2) Implementar una función que reciba un hash (diccionario) de claves genéricas K
   y como dato una lista cuyos elementos sean genéricos T. La función debe devolver
   un nuevo diccionario donde las claves sean las mismas que las del diccionario
   recibido, pero el valor asociado a cada clave sea el elemento del medio de la
   lista correspondiente del diccionario original.

   El diccionario original (ni sus listas) no deben modificarse.

   Se sabe que cada lista tiene a lo sumo M elementos (M no es constante).

   Ejemplo:
     Si el diccionario original tiene la clave "boquita" con la lista [1977, 1978, 2000, 2001, 2003, 2007],
     el nuevo diccionario debe tener la clave "boquita" con el dato 2000 (elemento del medio).

   La firma de la función debe ser:
     func DictMedio[K comparable, T any](dict Diccionario[K, Lista[T]]) Diccionario[K, T]

   Indicar y justificar la complejidad del algoritmo implementado.
*/


// Como se puede apreciar la complejidad maxima en el algoritmos desarrollado es
// O(n.m), dado que a lo sumo se recorren los n elementos del arreglo, aplicando
// a lo sumo operaciones O(m), si se recorre toda una lista interna.

/*
3) Implementar una primitiva del ABB que, dado un valor entero M, una clave inicial "inicio"
   y una clave final "fin", devuelva una lista con todos los datos cuyas claves estén
   dentro del rango [inicio, fin] y que además se encuentren dentro de los primeros M niveles
   del árbol (considerando a la raíz en el nivel 1).

   Ejemplo:
     Si se tiene el siguiente ABB:

              10
             /  \
            5    15
           / \   / \
          3   8 12 20
             /     \
            7       14

     y se llama con M = 3, inicio = 5 y fin = 15,
     el resultado debe contener los datos de las claves 10, 5, 8, 12, 15 (en cualquier orden).

   Se debe indicar y justificar la complejidad temporal del algoritmo implementado.
*/


// Como se puede apreciar la complejidad maxima en el algoritmos desarrollado es
// O(n), dado que a lo sumo se recorren los nodos del arbol, aplicando
// a lo sumo operaciones O(1).

/*
Implementar una primitiva para el ABB:

   func (arbol *Arbol[K, V]) AncestroComun(clave1, clave2 K) K

   que reciba dos claves y devuelva la **clave del último ancestro en común** (Lowest Common Ancestor, LCA)
   entre ambas en el ABB. El ancestro común puede coincidir con alguna de las claves pasadas.
   Si **alguna** de las claves no se encuentra en el árbol, finalizar con **panic**.

   Ejemplos (para el árbol del enunciado):
     arbol.AncestroComun(1, 4) --> 2
     arbol.AncestroComun(2, 4) --> 2
     arbol.AncestroComun(9, 1) --> 5

   Indicar y justificar la complejidad: O(h) tiempo y O(h) espacio (por la recursión),
   donde h es la altura del árbol.
*/


