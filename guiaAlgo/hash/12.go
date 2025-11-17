package proveedores

func MergeProveedores(prov1, prov2 Diccionario[string, int]) Diccionario[string, int] {
    merge := CrearHash[string, int]()

	it1 := prov1.Iterador()
	it2 := prov2.Iterador()

	for it1.HaySiguiente() {
		k, v:= it1.VerActual()

		if !prov2.Pertenece(k){
			merge.Guardar(k, v)
		} else{
			v2 := prov2.Obtener(k)
			if v < v2{
				merge.Guardar(k, v)
			} else{
				merge.Guardar(k, v2)
			}
		}

		it1.Siguiente()
	}

	for it2.HaySiguiente() {
		k, v := it2.VerActual()

		if !prov1.Pertenece(k){
			merge.Guardar(k, v)
		} else{
			v1 := prov1.Obtener(k)
			if v < v1{
				merge.Guardar(k, v)
			} else{
				merge.Guardar(k, v1)
			}
		}

		it2.Siguiente()
	}


	return merge
}

// Orden:
// Pertenece(), Obtener(), Guardar() son o(1)
// Compl: O(n + m)
// n = cant en prov1 
// m = cant en prov 2