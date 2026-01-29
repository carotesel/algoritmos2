/*
1. Implementar un algoritmo que, dado un árbol binario, determine si el mismo es completo (es decir, que todos los niveles
que tenga estén completos). Indicar y justificar la complejidad del algoritmo implementado.
*/

// Completo: cant_nodos == 2 ^ altura - 1

func (ab *ab[T]) Altura() int{
	if ab == nil{
		return 0
	}

	izq := ab.izq.Altura()
	der := ab.der.Altura()
	
	return max(izq, der) + 1
}

func (ab *ab[T]) CantidadNodos() int{
	if ab == nil{
		return 0
	}

	izq := ab.izq.Altura()
	der := ab.der.Altura()
	
	return izq + der + 1
}

func (ab *ab[T]) EsCompleto() bool{
	altura := ab.Altura()
	cantNodos := ab.CantidadNodos()

	return cantNodos == math.pow(2, altura) - 1
}

// O(n)

/*
2. El 10/10 un nuevo algoritmo de ordenamiento ha sido inventado: el MessiSort. Así como no podemos entender cómo
hace Messi para jugar como lo hace, vamos a asumir que no podemos entender cómo hace este algoritmo para ordenar.
El creador del algoritmo (que nada tiene que ver con el astro argentino), declara que el mismo ordena en tiempo mejor
a O(n log n). ¿Tenés algo para decir sobre su afirmación? Si esto no fuera cierto, ¿podría utilizarse como algoritmo
auxiliar de RadixSort?
*/

/*
Lo que podria decir sobre su afirmación a partir de la información proporcionada es que la misma es falsa, 
dado que, por el contexto del ejercicio, debo asumir que el algoritmo es un ordenamiento comparativo, y matemáticamente
no existen algoritmos comparativos mejores que o(n log n) en el peor de los casos. Es por esto que tampoco se podría utilizar como algoritmo
auxiliar de radix sort dado que éste precisa un algoritmo lineal y estable y messi sort no lo es.
*/

/*
3. Implementar un algoritmo que reciba dos arreglos desordenados y determine si ambos arreglos tienen los mismos
elementos (y en mismas cantidades). Indicar y justificar la complejidad del algoritmo implementado.
*/

func mismosElementos(arr1, arr2 []int) bool{

	if len(arr1) != len(arr2){
		return false
	}

	dicc1 := CrearHash[int, int]()
	dicc2 := CrearHash[int, int]()

	for _, elem := range arr1{
		if dicc1.Pertenece(elem){
			dicc1.Guardar(elem, dicc1.Obtener(elem)+1)
		} else{
			dicc1.Guardar(elem, 1)
		}
	}

	for _, elem := range arr2{
		if dicc2.Pertenece(elem){
			dicc2.Guardar(elem, dicc2.Obtener(elem)+1)
		} else{
			dicc2.Guardar(elem, 1)
		}
	}

	iter1 := dicc1.Iterador()

	for iter1.HaySiguiente(){
		actual := iter1.VerActual()

		if !dicc2.Pertenece(actual){
			return false
		}

		if dicc2.Obtener(actual) != dicc1.Obtener(actual){
			return false
		}
		iter1.Siguiente()
	}

	iter2 := dicc2.Iterador()

	for iter2.HaySiguiente(){
		actual := iter2.VerActual()

		if !dicc1.Pertenece(actual){
			return false
		}

		if dicc1.Obtener(actual) != dicc2.Obtener(actual){
			return false
		}
		iter2.Siguiente()
	}

	return true
}

/*Otra opc. arr 1 suma y arr 2 resta, esta vacio = espejado*/

func mismosElementos2(arr1, arr2 []int) bool{

	if len(arr1) != len(arr2){
		return false
	}

	dicc := CrearHash[int, int]()

	for _, elem := range arr1{
		if dicc.Pertenece(elem){
			dicc.Guardar(elem, dicc.Obtener(elem)+1)
		} else{
			dicc.Guardar(elem, 1)
		}
	}

	for _, elem := range arr2{
		if !dicc.Pertenece(elem){
			return false
		}
		cant := dicc.Obtener(elem) - 1
		if cant < 0{
			return false
		}
		if cant == 0{
			dicc.Borrar(elem)
		} else{
			dicc.Guardar(elem, cant)
		}
	}

	return dicc.Cantidad() == 0
}

/*
4. Trabajamos para una escuela primaria muy estructurada. En dicha escuela hay k cursos, cada uno con m alumnos (es
decir, hay un total de n = k · m alumnos). Todas las mañanas hay que armar filas para cantar Aurora en el patio del
colegio. Primero los alumnos se ubican en una fila correspondiente a su curso, de menor a mayor altura para cantar.
Una vez terminado, proceden a entrar a la escuela de a un alumno por vez, pero deben hacerlo de menor a mayor altura.
Es decir, se debe ordenar a todos los alumnos de menor a mayor. Nosotros sabemos que esto es ineficiente (suelen usar
mergesort, así que es O(n log n)), y desaprovechamos que los alumnos ya estaban ordenados por cursos. Implementar
un algoritmo que reciba k filas (arreglos) de alumnos, cada una previamente ordenada de menor a mayor altura, y nos
devuelva un único arreglo con todos los alumnos ordeados por altura en tiempo menor a O(n log n). Indicar y justificar
la complejidad del algoritmo implementado.
*/

// k = cursos
// m = alumnos x curso

type Alumno struct{
	nombre string
	altura int
}

type Elem struct{
	valor Alumno
	idx int
	pos int
}

func cmp(e1, e2 Elem) int{
	return e1.valor.altura - e2.valor.altura
}

func kMerge(arr [][]Alumno, k, m int) []Alumno{
	res := make([]Alumno, 0, k * m)
	heap := CrearHeap[Elem](cmp)

	for i:=0; i<k; i++{
		heap.Encolar(Elem{valor: arr[i][0], idx: i, pos: 0})
	}

	for !heap.EstaVacia(){
		min := heap.Desencolar()
		res = append(res, min.valor)
		pos := min.pos + 1

		if pos < len(arr[min.idx]){
			heap.Encolar(Elem{valor: arr[min.idx][pos], idx: min.idx, pos: pos})
		}
	}
	return res
}

// Complejidad: O(n log k)
// n = k * m
// mejor qur n log n cuando k << n

/*
Las personas se modelan como vértices.
Las relaciones “falleció antes que” se modelan como aristas dirigidas.
La consistencia del conjunto de datos equivale a que el grafo no tenga ciclos y 
respete las relaciones de simultaneidad, lo que puede verificarse mediante un orden topológico.
*/

// TLDR; lo modelamos con el orden de quien fallecio antes que quien