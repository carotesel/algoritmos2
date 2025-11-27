/*
1. Implementar una primitiva eliminarColisiones(clave K) []K para el hash, que elimine del hash todas las
claves que colisionen con la clave pasada por parámetro en el estado actual (sin eliminar dicha clave del
diccionario, si se encuentra), y devuelva dichas claves. Implementar tanto para el hash abierto como para el hash
cerrado. Si no se implementa para alguno, el ejercicio no estará aprobable. Indicar y justificar la complejidad de
la primitiva para ambos casos.
*/

func (dicc *hashAbierto[K, V])eliminarColisionesAbierto(clave K) []K{
	h := hash(clave)
	bucket := dicc.tabla[h]

	// Lista de claves que colisionan (las borradas)
	res := make([]K, 0)

	// Nuevo bucket donde solo queda la clave original
	    nuevoBucket := make([]K, 0)


	for _, x := range bucket{
		if x == clave{
			nuevoBucket = append(nuevoBucket, x)
		} else{
			res = append(res, x)
			dicc.cantidad--
		}
	}

	dicc.tabla[h] = nuevoBucket

	if (dicc.cantidad/len(dicc.tabla)) < FACTOR_MINIMO {
		nuevoTam := len(dicc.tabla) / 2
        dicc.redimensionar(nuevoTam)
    }
    if float64(dicc.cantidad)/float64(len(dicc.tabla)) > FACTOR_MAXIMO {
		nuevoTam := len(dicc.tabla) * 2
        dicc.redimensionar(nuevoTam)
    }

	return res
}

func (dicc *hashCerrado[K, V])eliminarColisionesCerrado(clave K) []K{
	inicio := hash(clave)
	tabla := dicc.tabla
	res := make([]K, 0)

	i := inicio

	for {
		celda := tabla[i]

		if celda.estado == VACIO {
			break
		} 
	}

	
}