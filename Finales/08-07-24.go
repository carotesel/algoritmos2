/*
1. Dos cadenas X e Y son isomórficas si existe alguna transformación biyectiva de caracteres que permita obtener Y a partir
de X. Ejemplos: casa y bata son isomórficas, y la transformación es c → b, a → a, s → t. burro y pizza son isomórficas,
y la transformación es b → p, u → i, r → z, o → a. mesa y masa no son isomórficas, porque la transformación debe
ser biyectiva: no podemos incluir e → a y a → a. Escribir una función que reciba dos cadenas y determine si son
isomórficas. Indicar y justificar la complejidad de dicha función.
*/

// ISOMORFICAS: 2 HASHES ASI VEO SI UNO MAPEA AL OTRO Y VICEVERSA!

func isomorficas(cadena1, cadena2 string) bool{
	diccX := CrearHash[string, string]()
	diccY := CrearHash[string, string]()

	if len(cadena1) != len(cadena2){
		return false
	}

	for i:=0; i<len(cadena1); i++{
		x := cadena1[i]
		y := cadena2[i]

		if diccX.Pertenece(x){
			if diccX.Obtener(x) != y{
				return false
			}
		} else{
			diccX.Guardar(x, y)
		}

		if diccY.Pertenece(y){
			if diccY.Obtener(y) != x{
				return false
			}
		} else{
			diccX.Guardar(y, x)
		}
	}
	return true
}

/*
3. Indicar si las siguientes afirmaciones sobre grafos son verdaderas o falsas, y justificar detalladamente:
a. En un árbol, todo vértice de grado mayor o igual a 2 es un punto de articulación, y todos los de grado menor a 2
no lo son.
b. Si realizo un recorrido BFS y otro DFS partiendo desde el mismo vértice y los recorridos resultan iguales (los
diccionarios de padres resultantes son iguales), entonces el grafo debe ser un árbol.
c. Si en un grafo dirigido G calculo las componentes fuertemente conexas (CFC), y creo un nuevo grafo dirigido G0
en el que cada vértice representa a una CFC del grafo G, y cada arista de G0

(v, w) indica que un vértice de la

CFC v tiene una arista hacia un vértice de la CFC w, resulta que G0

es sí o sí un grafo dirigido y acíclico.
*/

// a. Es verdadero ya que los vertices de grado < 2 son hojas, por lo que si estos fueran removidos, no deja desconectado el arbol, por lo que no son puntos de articulacion. 

// b. Es verdadero ya que si bfs y dfs partiendo del mismo vertice devuelven el mismo dicc de padres,
// significa que no existen ciclos ni caminos alternativos en el grafo. Por lo tanto, el grafo es conexo y acíclico, es decir, un árbol.

// c. Verdadero, ya que al construir el grafo con v = cfc y aristas = hay arista entre 2 cfc,
// si o si el grafo no puede tener ciclos ya que los mismos estan contenidos en la cfc,
// por lo que si los hubiera, estaria contradiciendo la prop de cfc (seria una sola).
// Por lo tanto, 𝐺 es un grafo dirigido acíclico.

/*
4. Dado un arreglo de enteros ordenado de n elementos en el cual sus elementos van de 0 a M, con M  n, implementar
una función que determine en O(log n) si hay algún elemento que aparezca más de la mitad de la veces en el arreglo.
Justificar la complejidad del algoritmo implementado.
*/

func mitadVecesDYC(arr []int) bool{
	n := len(arr)

	if n == 0{
		return false
	}

	medio := arr[n/2]

	indicePrimer := buscarP(arr, medio, n)
	indiceUlt := buscarU(arr, medio, n)

	return indiceUlt - indicePrimer + 1 > n/2
}

func buscarP(arr []int, nro int, n int) int{
	ini := 0
	fin := n - 1
	res := -1

	if ini <= fin{
		medio := (ini + fin) / 2

		if arr[medio] == nro{
			res = medio
			fin = medio - 1
		} else if arr[medio] < nro{
			ini = medio + 1
		} else{
			fin = medio - 1
		}
	}

	return res
}

func buscarU(arr []int, nro int, n int) int{
	ini := 0
	fin := n - 1
	res := -1

	if ini <= fin{
		medio := (ini + fin) / 2

		if arr[medio] == nro{
			res = medio
			ini = medio + 1
		} else if arr[medio] < nro{
			ini = medio + 1
		} else{
			fin = medio - 1
		}
	}

	return res
}

// BUSCAR PRIMERO: INDICE + CHICO DONDE APARECE. SI LO ENCUENTRO SIGO REVISANDO A LA IZQ
// BUSCAR ULTIMO: INDICE + GRANDE DONDE APARECE. SI LO ENCUENTRO SIGO REVISANDO A LA DER

/*
5. Implementar una primitiva para la lista enlazada Modificar(modificador func(T) T) que modifique todos los datos
de la lista. Cada dato deberá a pasar ser el que resulta de aplicar la función modificador al dato que se encontraba
anteriormente en dicha posición. Por ejemplo, si la lista es [1, 2, 3] y la función func(elem int) int {return
elem * 2}, luego de ejecutar la primitiva con dicha función, la lista debe quedar como [2, 4, 6]. Indicar y justificar
la complejidad del algoritmo implementado.
*/

func (lista *listaEnlazada[T]) Modificar(modificador func(T) T){
	actual := lista.primero

	for actual != nil{
		actual.dato = modificador(actual.dato)
		actual = actual.siguiente
	}
}