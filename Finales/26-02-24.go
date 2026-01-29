/*
1. Un algoritmo iterativo sencillo para obtener la potencia de un número (b

n) tiene complejidad O(n). Tal como vieron en

el secundario (esperamos), sabemos que b
n =
b
n
2
2
. Utilizar esta propiedad para implementar un algoritmo que calcule

b
n, en tiempo O(log n). Justificar la complejidad del algoritmo implementado. Recordar tener cuidado con el caso que
n sea un valor impar.
*/

import (
    "fmt"
    "math"
)

func obtenerPotencia(b, n int) int{
	if n == 0{
		return 1
	}

	medio := n / 2

	res := pot(b, medio)

	res = res * res

	if n % 2 != 0{
		res = res * b
	}

	return res
}

/*
3. Realizar el seguimiento de aplicar CountingSort al siguiente conjunto de equipos de fútbol, ordenando por la cantidad
de descensos que tienen (entre paréntesis se indica la cantidad en cada caso). Implementar dicho algoritmo, e indicar y
justificar la complejidad del mismo.
Olimpo (4) - Boca (0) - Almagro (3) - Rosario Central (4) - Banfield (8) -
Sarmiento (2) - Defensa y Justicia (0) - Plantense (3) - River (1) - Independiente (1) -
Estudiantes LP (2) - Racing (1) - Tigre (8) - Velez (1) - Atlanta (4) - Gimnasia LP (5)
*/

type Equipo struct{
	nombre string
	desc int
}

func CountingSort(equipos []Equipo) []Equipo{
	arrFrecuencias := make([]int, 9)
	arrSumaAcumulada := make([]int, 9)

	for _, equipo := range equipos{
		arrFrecuencias[equipo.desc] += 1
	}

	arrSumaAcumulada[0] = 0

	for i:=1; i<9; i++{
		arrSumaAcumulada[i] = arrFrecuencias[i-1] + arrSumaAcumulada[i]
	}

	res := make([]Equipo, len(equipos))

	for i := 0; i < len(equipos); i++ {
		e := equipos[i]
		pos := arrSumaAcumulada[e.desc]
		res[pos] = e
		arrSumaAcumulada[e.desc]++
	}

	return res
}

/*
4. Implementar una función que dada una pila, determine si la misma se encuentra ordenada (es decir, se ingresaron los
elementos de menor a mayor). La pila debe quedar en el mismo estado al original al terminar la ejecución de la función.
Indicar y justificar la complejidad de la función.
*/

func pilaOrdenada(p PilaDinamica[int]) bool{

	if p.EstaVacia() {
        return true
    }

	pilaAux := CrearPilaDinamica[int]()

	tope := p.Desapilar()
	ordenada := true
	pilaAux.Apilar(tope)

	for !p.EstaVacia(){
		anterior := p.Desapilar()
		if anterior > tope{
			ordenada = false
		}
		pilaAux.Apilar(anterior)
		tope = anterior // siempre actualizo tope para comparar CONSECUTIVOS
	}

    for !pilaAux.EstaVacia() {
        p.Apilar(pilaAux.Desapilar())
    }
    
    return ordenada
}

// pila -> cola -> pila = INVIERTE 
// pila -> pila -> pila = MANTIENE ORDEN

/*
5. En una facultad contamos únicamente con 2 proyectores. Diferentes docentes de distintos cursos quieren utilizarlos
(dos cursos no pueden usar el mismo proyector si coindicen en horarios). Teniendo la información de los horarios de
cada curso, se pide definir si existe una forma de organizar la asignación para que todos tengan algún proyector (no
importa de momento cuál es la asignación, sólo si existe). Modelar este problema utilizando grafos, e implementar un
algoritmo que reciba un grafo de las características descriptas y resuelva el problema. Indicar y justificar la complejidad
del algoritmo implementado en función de las variables del problema.
*/

// El problema se resuelve con grafo bipartito

// Vertices = cursos 
// Hay arista si coinciden en horario

def es_bipartito(grafo):

	colores = {}

	for v in grafo:
		if v in colores:
			continue
		
		colores[v] = 0
		q = deque()
		q.append(v)

		while len(q) > 0:
			actual = q.popleft()

			for w in grafo.adyacentes(actual):
				if w in colores:
					if colores[w] == colores[actual]:
						return false
				else:
					colores[w] = 1 - colores[actual]
					q.append(w)
	return True

	