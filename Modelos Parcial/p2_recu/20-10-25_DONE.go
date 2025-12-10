/*
1. A lo largo de su trayectoria, la empresa ha tenido varias rotaciones de equipos. Próximamente habrá una nueva, y se busca 
que los nuevos equipos incluyan personas que hayan compartido equipos antes, para facilitar la transición.
Se dispone de una base de datos que registra, para N personas, en qué M equipos ha participado cada una. La base usa un hash 
donde la persona es la clave y el valor, una lista de equipos:

{
'Ana': ['Frontend-Platform', 'Growth-Squad'],
'Beto': ['Backend-Services', 'Frontend-Platform', 'Mobile-Core'],
'Carla': ['Mobile-Core'],
}

Realizar una función personasEnComun que reciba el hash, y el nombre de una persona , y devuelva una lista con todas las 
personas que hayan trabajado en al menos uno de sus equipos listados. Indicar y justificar la complejidad del algoritmo 
implementado, expresada con las variables N y M del problema. Por ejemplo, si se consulta por ‘Beto’, la respuesta incluye a 
‘Ana’ y ‘Carla’. Si se pregunta por ‘Beto’, la respuesta incluye a ‘Beto’.
*/

func personasEnComun (dicc Diccionario[string, []string], persona string) Lista[string]{

	if !dicc.Pertenece(persona){
		return CrearListaEnlazada[string]()
	}

	lista := CrearListaEnlazada[string]()
	equipos := dicc.Obtener(persona) // O(1)

	diccEquipos := CrearDiccionario[string, bool]() 

	for _, e := range equipos{ // O(M)
		diccEquipos.Guardar(e, true)
	}

	iter := dicc.Iterador()

	for iter.HaySiguiente(){ // O(N)
		pActual, pEquipos := iter.VerActual()

		for _, e := range pEquipos{
			if diccEquipos.Pertenece(e){ // O(1)
				lista.AgregarUltimo(pActual) // O(1)
				break // NO OLVIDAR BREAK
			}
		}
		iter.Siguiente()
	}
	return lista
}

// O(M) + O(N) = O(M + N)

/*
2. En nuestro juego de rol táctico, Final Fantasy Algorithms, el orden de ataque se decide por el atributo iniciativa: los de 
mayor iniciativa atacan primero. En esta batalla, un hechizo global está a punto de activarse, por lo que solo quedan T turnos 
en los que se pueda realizar un ataque. Se desea saber, de los N personajes, cuáles lograrán atacar antes de que se active el 
hechizo Final Fantástico Algorítmico. Se tiene una lista con los personajes que participarán en el combate como una estructura 
de formato (nombre string, iniciativa int):

[ ('Ma-Go Lang', 95), ('Bárbara', 75), ('Cléri-Go Lang', 60), ('Arquera de bugs', 90) ]
Se pide realizar una función determinarOrdenDeAtaque que reciba la lista de combatientes, y la cantidad T turnos de turnos 
restantes. La función debe devolver una lista con los nombres de los personajes que logran actuar en esa ventana de tiempo, 
ordenados por turno en el que actúan. Indicar y justificar la complejidad del algoritmo implementado, 
expresada con las variables N y T del problema.
*/

type Combatiente struct{
	nombre string
	iniciativa int
}

func cmp(c1, c2 Combatiente) int{
	return c2.iniciativa - c1.iniciativa
}

func determinarOrdenDeAtaque (lista Lista[Combatiente], T int) Lista[string]{
	heap_max := CrearHeap(cmp)
	res := CrearListaEnlazada[string]()

	iter := lista.Iterador()

	for iter.HaySiguiente(){ // O(n)
		actual := iter.VerActual()
		heap_max.Encolar(actual)
		iter.Siguiente()
	}

	if T > heap_max.Cantidad(){
		T = heap_max.Cantidad()
	}

	for i:=0; i<T; i++{ // O(T log n)
		combat := heap_max.Desencolar()
		lista.AgregarUltimo(combat.nombre)
	}
	return res
}

/*
3. Implementar en Go una primitiva para Árbol Binario func (ab *Arbol[int]) ArbolEsPlantable() bool que determine si un
árbol es plantable, o no. Para que lo sea, todo nodo debe cumplir: el dato del nodo debe ser mayor al dato de sus hijos 
(si los tiene), y además, el dato del nodo no puede superar la altura de dicho nodo. 
Implementar la primitiva en O(n), y justificar su complejidad.
A fines del ejercicio considerar la estructura del árbol como la definida al dorso.
*/

func (ab *Arbol[int]) ArbolEsPlantable() bool{
	esPlant, _ := esPlantable(ab)
	return esPlant
}

func esPlantable (ab *Arbol[int]) (bool, int){
	if ab == nil{
		return true, 0
	}

	esIzq, hIzq := esPlantable(ab.izq)
	esDer, hDer := esPlantable(ab.der)

	hNodo := max(hIzq, hDer) + 1

	condHijos := true

	if ab.izq != nil{
		condHijos = condHijos && ab.izq.dato < ab.dato
	}
	
	if ab.der != nil{
		condHijos = condHijos && ab.der.dato < ab.dato
	}
	

	return ab.dato <= hNodo && condHijos && esIzq && esDer, hNodo
}
