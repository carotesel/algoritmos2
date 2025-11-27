/*
3. Implementar una función mejoresPromedios(alumnos []Alumno, k int) Lista[Alumno] que, dado un arreglo
de Alumnos y un valor entero k, nos devuelva una lista de los k alumnos de mayor promedio (ordenada de mayor a menor). 

Indicar y justificar la complejidad del algoritmo implementado.

Considerar que la estructura del alumno es:

type Alumno struct {
nombre string
padron int
notas []int
}

TOP K

*/

func cmp(a1, a2 Alumno) int{
	var sumaNotas1, sumaNotas2 float64

	for _, x := range a1.notas{
		sumaNotas1 += float64(x)
	}

	for _, x := range a2.notas{
		sumaNotas2 += float64(x)
	}

	promedio1 := sumaNotas1 / float64(len(a1.notas))
	promedio2 := sumaNotas2 / float64(len(a2.notas))

	if promedio1 > promedio2{
		return 1
	} else if promedio1 < promedio2{
		return -1
	}
	return 0
}

func mejoresPromedios(alumnos []Alumno, k int) Lista[Alumno]{
	lista := CrearListaEnlazada[Alumno]()

	if k <= 0 || k > len(alumnos) {
        return lista
    }

	heap := CrearHeapArr(alumnos, cmp)

	for i:=0; i < k; i++{
		mejor := heap.Desencolar()
		lista.AgregarUltimo(mejor)
	}
	return lista
}

// O(n + k log n)