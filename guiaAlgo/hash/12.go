
/*(★★) Una fábrica de pastas de Lanús le pide a alumnos de Algoritmos y Estructuras de Datos 
que le solucionen un problema: sus dos distribuidoras de materias primas le enviaron un hash cada una, 
dónde sus claves son los nombres de los productos, y sus valores asociados, sus precios. 

La fábrica de pastas le pide a los alumnos que le implementen una función que le devuelva un nuevo
 hash con la unión de todos esos productos, y en caso de que una misma materia prima se encuentre en ambos hashes, 
 elegir la que tenga el precio más barato. 
 
 Indicar y justificar el orden del algoritmo.*/

 func UnionProductos[K comparable, float](d1, d2 Diccionario[K, float]) Diccionario[K, float]{
	res := CrearDiccionario[K, float]()

	iter := d1.Iterador()

	// ESTRATEGIA: COPIAR TODO D1 A RES E ITERAR D2, SI ESTA EN RES METO EL MENOR Y SINO METO ACTUAL DE D2

	for iter.HaySiguiente(){
		actual, valor := d1.VerActual()
		res.Guardar(actual, valor)
		iter.Siguiente()
	}

	iter2 := d2.Iterador()

	for iter2.HaySiguiente(){
		actual, valor := d2.VerActual()

		if res.Pertenece(actual){
			v1 := d1.Obtener(actual)
			valor_min := min(valor, v1)
			res.Guardar(actual, valor_min)
		}
		res.Guardar(actual, valor)
		iter2.Siguiente()
	}

	
	return res
 }

 // Total: O(n₁ + n₂)