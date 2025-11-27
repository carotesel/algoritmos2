/*2. Implementar una función eliminarRepetidos(arreglo []int) []int que dado un arreglo de números, nos devuelva
otro en el que estén los elementos del original sin repetidos. La primera aparición debe mantenerse, y las demás no ser
consideradas. Indicar y justificar la complejidad del algoritmo implementado.*/

func eliminarRepetidos(arr []int) []int{
	hash := CrearHash[int, bool]()
	res := make([]int, 0)

	for _, elem := range arr{
		if !hash.Pertenece(elem){
			hash.Guardar(elem, true)
			res = append(res, elem)
		}	
	}
	return res
}

// Tal como se puede observar la complejidad total del algoritmo es de O(n),
// dado que se realizan ciclos recorriendo a lo sumo n elementos
// aplicando operaciones constantes (O(1)).