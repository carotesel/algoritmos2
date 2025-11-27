/*(★★) En un diccionario todas las claves tienen que ser diferentes, no así sus valores. 
Escribir en Go una primitiva para el hash cerrado 
func (dicc *hashCerrado[K, V]) CantidadValoresDistintos(cmp func (V, V) bool) int 
que, sin usar el iterador interno, dado un hash devuelva la cantidad de valores diferentes que almacena. 

La función pasada por parámetro determina si dos valores son iguales, o no. 
Indicar y justificar el orden del algoritmo.*/

func (dicc *hashCerrado[K, V]) CantidadValoresDistintos(cmp func (V, V) bool) int{
	iter := dicc.Iterador()
	diccAux := CrearDiccionario[V, bool]() // O(1)

	for iter.HaySiguiente(){ // O(n)
		existe := false

		_, valor := iter.VerActual() // O(1)

		// reiniciamos iterador del diccionario auxiliar. la idea es revisar cada valor del diccionario inicial
		iter2 := diccAux.Iterador()

		for iter2.HaySiguiente(){
			actual, _ := iter2.VerActual()

			if cmp(valor, actual){
				existe = true
				break
			}
			iter2.Siguiente()
		}

		if !existe{
			diccAux.Guardar(valor, true)
		}
		iter.Siguiente()
	}

	return diccAux.Cantidad() // O(1)
}

// Complejidad: O(n^2)