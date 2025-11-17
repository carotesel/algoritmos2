package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	temporal := *x
	*x = *y
	*y = temporal
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	} else {
		maximo := vector[0]
		maxIndex := 0

		for i := 1; i < len(vector); i++ {
			if vector[i] > maximo {
				maximo = vector[i]
				maxIndex = i
			}
		}
		return maxIndex
	}
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo;
// 0 si son iguales;
// o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	i := 0

	for i < len(vector1) && i < len(vector2) {
		if vector1[i] != vector2[i] {
			if vector1[i] < vector2[i] {
				return -1
			} else if vector1[i] > vector2[i] {
				return 1
			}
		}
		i++
	}
	if len(vector1) < len(vector2) {
		return -1
	} else if len(vector2) < len(vector1) {
		return 1
	}
	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección. Lo hago con max en vez de min, el max queda ordenado al final.
func Seleccion(vector []int) {
	n := len(vector)

	if n < 2 {
		return
	} else {
		for i := n; i > 1; i-- {
			posMaximo := Maximo(vector[:i])

			if posMaximo != i-1 {
				Swap(&vector[posMaximo], &vector[i-1])
			}
		}
	}

}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).

func sumaRecursiva(vector []int, n int) int {
	// caso base: n = 1
	// otro caso: vector[n-1] + vector[n-2]

	if n == 1 {
		return vector[n-1]
	} else {
		return vector[n-1] + sumaRecursiva(vector, n-1)
	}

}

func Suma(vector []int) int {
	n := len(vector)

	if n < 1 {
		return 0
	} else {
		suma := sumaRecursiva(vector, n)
		return suma
	}
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).

func EsPalindromoRecursiva(cadena string, n int) bool {
	// caso base: 0 o 1 caracteres => palíndromo
	// caso no base: si cadena [0] != cadea [n-1] => no palindromo
	if n <= 1 {
		return true
	}
	if cadena[0] != cadena[n-1] {
		return false
	}
	// llamamos sacando primera y última letra, y reduciendo n en 2
	return EsPalindromoRecursiva(cadena[1:n-1], n-2)
}

func EsCadenaCapicua(cadena string) bool {
	return EsPalindromoRecursiva(cadena, len(cadena))
}
