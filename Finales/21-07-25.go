/*
1. Implementar un algoritmo que dado un arreglo de dígitos (0-9) determine cuál es el número más grande que se puede
formar con dichos dígitos. Indicar y justificar la complejidad del algoritmo implementado.
*/

// atoi

import (
	"fmt"
	"strconv"
)

func numeroBigger(arr []int) int{
	contador := make([]int, 10)

	for _, elem := range arr{
		contador[elem]++
	}

	var nro string

	for i:=9; i>0; i--{
		for j:= 0; j < contador[i]; j++{
			nro += strconv.Itoa(i)
		}
	}

	return strconv.Atoi(nro)
}

// Complejidad: O(n)

/*
2. Implementar un algoritmo que dado un arreglo de números, determine si hay un elemento dentro del mismo que
aparece al menos la mitad veces. La complejidad del algoritmo debe ser lineal. Justificar la complejidad del algoritmo
implementado.
*/

func masMitadveces(arr []int) int{
	dicc := CrearHash[int, int]()

	for _, elem := range arr{
		if dicc.Pertenece(elem){
			dicc.Guardar(elem, dicc.Obtener(elem)+1)
		} else{
			dicc.Guardar(elem, 1)
		}
	}

	iter := dicc.Iterador()

	for iter.HaySiguiente(){
		actual := iter.VerActual()
		if dicc.Obtener(actual) > len(arr) / 2{
			return actual
		}
		iter.Siguiente()
	}
}

/*
3. Edgar necesitaba hacer unos cambios en un TDA que ya había implementado Alan, con todo y pruebas (todas
funcionando correctamente, como es de esperar de Alan). Esto le llevó a Edgar n días. Cada día realizó un commit en el
sistema de control de versiones que utilizan. El problema es que no corrió las pruebas hasta el día n, y recién ahí notó
que fallaban. Edgar implementó una función func todoOkElDía(n int) bool que recibe un número de día y devuelve
true si estaba todo ok hasta el día n o false si ese día ya fallaban.
Implementar una función func buscarDiaFalla(diasTotales int) int que devuelva el número de día en el que
empezaron a fallar las pruebas. Indicar y justificar la complejidad del algoritmo implementado (la complejidad de
todoOkElDia es O(n)).

*/

func buscarDiaFalla(diasTotales int) int{
	return diaFalla(0, diasTotales)
}

func diaFalla(ini, fin int) int{

	if ini > fin{
		return ini
	}

	medio := (ini + fin) / 2

	if todoOkElDia(medio){
		return diaFalla(medio+1, fin)
	} else{
		return diaFalla(ini, medio)
	}
}

/*
4. Implementar un algoritmo que reciba un arreglo de n números, y un número k, y devuelva los k números dentro del
arreglo cuya suma sería la máxima (entre todos los posibles subconjuntos de k elementos de dicho arreglo). Indicar y
justificar la complejidad de la función implementada.
*/

func cmp(a, b int) int{
	return b - a
}

func Topk(arr []int, k int) []int{
	if k > len(arr){
		k = len(arr)
	}
	res := make([]int, 0)
	heapMax := CrearHeapArr(cmp)

	for i:=0; i<k; i++{ // O(k)
		res = append(res, heapMax.Desencolar()) // O(log n)
	}
	return res
}

// O(n + k log n)

