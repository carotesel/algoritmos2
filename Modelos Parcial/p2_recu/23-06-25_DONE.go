/* 1) Implementar una primitiva eliminarColisiones(clave K) []K para el hash,
      que elimine del hash todas las claves que colisionen con la clave pasada
      por parámetro en el estado actual (sin eliminar dicha clave del
      diccionario, si se encuentra) y devuelva dichas claves.

      Implementar tanto para el hash abierto como para el hash cerrado. Si no
      se implementa para alguno, el ejercicio no estará aprobable.

      Indicar y justificar la complejidad de la primitiva para ambos casos.
*/

func (dicc *hashAbierto[K, V]) eliminarColisiones(clave K) []K{
	indice := hash(clave) % dicc.tam
	tabla := dicc.tabla[indice]

	nuevaTabla := make([]K, 0)
	res := make([]K, 0)

	for _, k := range tabla{
		if k == clave{
			nuevaTabla = append(nuevaTabla, k)
		} else{
			res = append(res, k)
			dicc.cantidad--
		}
	}

	dicc.tabla[indice] = nuevaTabla

	// Redimensiono post borrado

	if float64(dicc.cantidad)/float64(len(dicc.tabla)) < FACTOR_MINIMO{
		nuevoT := len(dicc.tabla)/2

		if nuevoT < FACTOR_MINIMO{
			nuevoT = FACTOR_MINIMO
		}
		dicc.redimensionar(nuevoT)
	}

	if float64(dicc.cantidad)/float64(len(hash.tabla)) > FACTOR_MAXIMO{
		nuevoT := len(dicc.tabla) * 2

		dicc.redimensionar(nuevoT)
	}

	return res
}

// es solo O(1) ya que recorro unicamente la lista o bucket de esa clave, y la redimension se da en casos particulares,
// por lo que en esencia resulta o(1)

func (dicc *hashCerrado[K, V]) eliminarColisiones(clave K) []K{
	cHasheada := hash(clave) % dicc.tam

	res := make([]K, 0)

	for i:=0; i < len(dicc.tabla); i++{
		celda := &dicc.tabla[i]

		if celda.estado != OCUPADO{
			continue
		}

		if celda.clave != clave && hash(celda.clave) == cHasheada{
			res = append(res, celda.clave)
			hash.tabla[i].estado = BORRADO
			hash.borrados++
			hash.cantidad--
		}
	}

	if float64(dicc.cantidad)/float64(len(dicc.tabla)) < FACTOR_MINIMO{
		nuevoT := len(dicc.tabla)/2

		if nuevoT < FACTOR_MINIMO{
			nuevoT = FACTOR_MINIMO
		}
		dicc.redimensionar(nuevoT)
	}

	if float64(dicc.cantidad)/float64(len(hash.tabla)) > FACTOR_MAXIMO{
		nuevoT := len(dicc.tabla) * 2

		dicc.redimensionar(nuevoT)
	}
	return res
}


/*
2. Sobre un AVL cuyo estado inicial puede reconstruirse a partir del preorder: 40 - 10 - 3 - 17 - 15 - 64 -
47 - 74 - 92, realizar un seguimiento de insertar los siguientes elementos (incluyendo rotaciones intermedias):
20, 23, 13, 14, 16, 12.

EN EL CUADERNO -> DONE.
*/

/* 3) Implementar una función mejoresPromedios(alumnos []Alumno, k int) Lista[Alumno]
      que, dado un arreglo de Alumnos y un valor entero k, nos devuelva una lista de
      los k alumnos de mayor promedio (ordenada de mayor a menor).
      Indicar y justificar la complejidad del algoritmo implementado.

      Considerar que la estructura del alumno es:

/*type Alumno struct {
       nombre string
       padron int
       notas  []int
}*/

//CORREGIDO:

type AlumnoProm struct {
	alumno Alumno
	promedio int
}

func cmp(a, b AlumnoProm) int{
	return b.promedio - a.promedio
}

func calcularPromedio (notas []int) int{
	suma := 0

	for _, x := range notas{
		suma += x
	}

	return suma / len(notas)
}

func mejoresPromedios(alumnos []Alumno, k int) Lista[Alumno] {
	lista := CrearListaEnlazada[AlumnoProm]()

	aux := make([]AlumnoProm, 0, len(alumnos))

	for _, a := range alumnos{
		aux = append(aux, AlumnoProm{alumno: a, promedio: calcularPromedio(a.notas)})
	}

	heap := CrearHeapArr(aux, cmp) // O(n)

	for i:=0; i<k; i++{ // O(k log n)
		lista.AgregarUltimo(heap.Desencolar())
	}
	return lista
}

// O(n + k log n)


