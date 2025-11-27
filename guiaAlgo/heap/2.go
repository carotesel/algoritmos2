/*(★) Implementar una primitiva para el heap (de máximos) que obtenga
 los 3 elementos más grandes del heap en O(1).*/

func (h heap[T]) Max3 () []T{
	res := make([]T, 3)

	if h.Cantidad() == 0{
		return res
	}
	
	// primer maximo
	res = append(res, h.datos[0])

	if h.Cantidad() > 1{
		res = append(res, h.datos[1])
	}

	if h.Cantidad() > 2{
		res = append(res, h.datos[2])
	}

	return res
}

// Primitiva en O(1) es SIN DESENCOLAR (O LOG N), sino que usando los indices del array de datos.