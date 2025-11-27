/*
2. En nuestro juego de rol táctico, Final Fantasy Algorithms, el orden de ataque se decide por el atributo iniciativa: 
los de mayor iniciativa atacan primero. En esta batalla, un hechizo global está a punto de activarse, por lo que solo 
quedan T turnos en los que se pueda realizar un ataque. 

Se desea saber, de los N personajes, cuáles lograrán atacar antes de que se active el hechizo Final Fantástico
Algorítmico. Se tiene una lista con los personajes que participarán en el combate como una estructura de formato 
(nombre string, iniciativa int):

[ ('Ma-Go Lang', 95), ('Bárbara', 75), ('Cléri-Go Lang', 60), ('Arquera de bugs', 90) ]

Se pide realizar una función determinarOrdenDeAtaque que reciba la lista de combatientes, y la cantidad T turnos 
de turnos restantes. La función debe devolver una lista con los nombres de los personajes que logran actuar en esa ventana 
de tiempo, ordenados por turno en el que actúan. 

Indicar y justificar la complejidad del algoritmo implementado, expresada con las variables N y T del problema.*/

func cmp(p1, p2 PersPuntaje) int{
	return p2.puntaje - p1.puntaje
}

type PersPuntaje struct{
	nombre string
	puntaje int
}

// COMO ES LISTA DE TUPLAS, CONVERTIR A ARRAY A MANO

func determinarOrdenDeAtaque (lista ListaEnlazada[PersPuntaje], T int) ListaEnlazada[string] {
	
	// Convertir la lista a arreglo
    arr := make([]PersPuntaje, 0, lista.Largo())
    
	it := lista.Iterador()
    for it.HaySiguiente() {
        arr = append(arr, it.VerActual())
        it.Siguiente()
    }
	
	heap := CrearHeapArr(lista, arr) // heap maximos. O(n)
	res := CrearListaEnlazada[string]()

	if lista.Largo() == 0{
		return res // lista vacia
	}

	for i:= 0; i < T && !heap.EstaVacia(); i++{ // O(T)
		max := heap.Desencolar() // O(log n)
		res.AgregarUltimo(max.nombre)
	}

	return res
}

// Complejidad: O(n + T log n)