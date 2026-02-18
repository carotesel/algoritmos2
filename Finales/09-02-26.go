/*
2. Se tiene un arreglo de enteros de tamaño conocido n, inicialmente ordenado de menor a mayor y sin elementos repetidos. Al mismo
se le aplica un corrimiento de una cierta cantidad k de elementos. Esto es, todos los elementos que están a partir del índice k se
los correrá hacia la izquierda k posiciones. Los elementos entre las posiciones 0 y k − 1, estarán ahora al final del arreglo (cada
sub-segmento mantendrá el orden original de los elementos). Por ejemplo: v = [0, 1, 3, 5, 7, 8, 9] luego de correrlo con k = 3
resulta en v = [5, 7, 8, 9, 0, 1, 3].
Implementar en Go una función que devuelva el valor de k (1 ≤ k ≤ n − 1) para un arreglo ya corrido en dicha cantidad k
desconocida, utilizando un algoritmo de complejidad O(log(n)).
La firma de la función es: func buscarK(v int[], ini int, fin int) int, y será llamada inicialmente con: buscarK(v, 0,
n-1). Justificar la complejidad del algoritmo propuesto.
*/

func buscarK(v []int, ini int, fin int) int {
    if ini == fin {
        return ini
    }

    medio := (ini + fin) / 2

    if v[medio] > v[fin] {
        return buscarK(v, medio+1, fin)
    } else {
        return buscarK(v, ini, medio)
    }
}
