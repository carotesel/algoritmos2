// 1) Implementar una primitiva eliminarColisiones(clave K) []K para el hash,
//    que elimine del hash todas las claves que colisionen con la clave pasada
//    por parámetro en el estado actual (sin eliminar dicha clave del
//    diccionario, si se encuentra) y devuelva dichas claves.
//
//    Implementar tanto para el hash abierto como para el hash cerrado. Si no
//    se implementa para alguno, el ejercicio no estará aprobable.
//
//    Indicar y justificar la complejidad de la primitiva para ambos casos.



// 3) Implementar una función mejoresPromedios(alumnos []Alumno, k int) Lista[Alumno]
//    que, dado un arreglo de Alumnos y un valor entero k, nos devuelva una lista de
//    los k alumnos de mayor promedio (ordenada de mayor a menor).
//    Indicar y justificar la complejidad del algoritmo implementado.
//
//    Considerar que la estructura del alumno es:
//
/*type Alumno struct {
       nombre string
       padron int
       notas  []int
}*/

func cmp(a1, a2 Alumno) int{
	suma1 := 0
	suma2 := 0

	for _, nota := range a1.notas{
		suma1 += nota
	}

	for _, nota := range a2.notas{
		suma2 += nota
	}

	prom1 := suma1 / len(a1.notas)
	prom2 := suma2 / len(a2.notas)

	return prom2 - prom1
}

func mejoresPromedios(alumnos []Alumno, k int) Lista[Alumno]{
	heapMax := CrearHeap(cmp)
	lista := CrearListaEnlazada[Alumno]()

	for _, alumno := range alumnos{ // O(n log n)
		heapMax.Encolar(alumno)
	}

	for i:= 0; i < k; i++{ // O(k log n)
		lista.AgregarUltimo(heapMax.Desencolar())
	}
	return lista
}

// complejidad: O((n+k)log n)
