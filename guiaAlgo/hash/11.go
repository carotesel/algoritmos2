package difsimetrica

func DiferenciaSimetricaDict[K comparable, V any](d1, d2 Diccionario[K, V]) Diccionario[K, V] {
    diferencia := CrearHash[K, V]()

	it1 := d1.Iterador()
	it2 := d2.Iterador()

	for it1.HaySiguiente() {
		k_Actual_1, v_Actual_1 := it1.VerActual()

		if !d2.Pertenece(k_Actual_1){
			diferencia.Guardar(k_Actual_1, v_Actual_1)
		}

		it1.Siguiente()
	}

	for it2.HaySiguiente() {
		k_Actual_2, v_Actual_2 := it2.VerActual()

		if !d1.Pertenece(k_Actual_2){
			diferencia.Guardar(k_Actual_2, v_Actual_2)
		}

		it2.Siguiente()
	}

	return diferencia
}

// Pertenece() y Guardar() son o(1)
// complejidad: O(n+m)
// n = cant en d1
// m = cant en d2