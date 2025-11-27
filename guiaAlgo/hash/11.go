/*11.(★★★) 
♠ La diferencia simétrica entre dos conjuntos A y B es un conjunto que contiene todos los elementos que 
se encuentran en A y no en B, y viceversa.

Implementar una función DiferenciaSimetricaDict[K comparable, V any](d1, d2 Diccionario[K, V]) Diccionario[K, V] 
que devuelva un nuevo Diccionario (puede ser el hash que hayas implementado) 
con la diferencia simétrica entre los dos recibidos por parámetro. 

La diferencia tiene que ser calculada teniendo en cuenta las claves, 
y los datos asociados a las claves deben ser los mismos que estaban en cada uno de los hashes originales.*/

func DiferenciaSimetricaDict[K comparable, V any](d1, d2 Diccionario[K, V]) Diccionario[K, V] {
	dicc := CrearDiccionario[K, V]()

	iter := d1.Iterador()

	for iter.HaySiguiente(){ // O(n1)
		actual, valor := iter.VerActual()

		if !d2.Pertenece(actual){
			dicc.Guardar(actual, valor)
		}
		iter.Siguiente()
	}

	iter2 := d2.Iterador()


	for iter2.HaySiguiente(){ // O(n2)
		actual, valor := iter2.VerActual()

		if !d1.Pertenece(actual){
			dicc.Guardar(actual, valor)
		}
		iter2.Siguiente()
	}

	return dicc
}

// Complejidad: O(n1 + n2) = O(n)