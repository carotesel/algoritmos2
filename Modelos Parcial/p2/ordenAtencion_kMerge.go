/*1. ¡El Dr. Casa necesita nuestra ayuda! Resulta que venia una caravana de k colectivos que estaban yendo por la Autopista
Buenos Aires - La Plata, y tuvieron un choque en cadena (nadie gravemente herido para nuestra suerte). En cada
colectivo viajaban h personas, y por razones de edad y patologías previas, aunque no hayan heridos graves, debemos
atender primero a ciertos pacientes que a otros. Por suerte para el Dr. Casa, su jefa, la Dra. Cudi, le ha asignado una
prioridad a cada paciente (siendo los pacientes con prioridad 1 aquellos con la mayor prioridad) y su compañero, el
Dr. Persecución, ha ordenado la fila de pacientes de cada colectivo por prioridad (considerar que cada fila es una cola
enlazada, donde los pacientes más prioritarios están al frente de la cola).
Teniendo k colas de h pacientes cada una, y una función ObtenerPrioridad(paciente) (que se ejecuta en O(1)),
implementar un algoritmo para que el Dr. Casa pueda saber el orden en el que deben ser atendidos los k · h pacientes.
Indicar y justificar la complejidad del algoritmo implementado.*/

type Item struct{
	persona string
	prioridad int
	indice int
}

func cmp(a, b Item) int{
	return a.prioridad - b.prioridad
}

func ordenAtencion(colas []ColaPrioridad[string]) Lista[string] {
	lista := CrearListaEnlazada[string]()
	heap := CrearHeap[Item](cmp)

	k := len(colas)

	// Inicializamos el heap con el primero de cada colectivo
	for i:=0; i < k; i++{ // O(k log k)
		if !colas[i].EstaVacia(){
			paciente := colas[i].Verprimero()
			prioridad := ObtenerPrioridad(paciente)
			heap.Encolar(Item{paciente, prioridad, i})
		}
	}

	for !heap.EstaVacia(){ // O(n)
		item := heap.Desencolar() // O(log k)
		lista.AgregarUltimo(item.persona)

		colas[item.indice].Desencolar()

		for !colas[item.indice].EstaVacia(){
			sig := colas[item.indice].Verprimero()
			p2 := ObtenerPrioridad(sig)
			heap.Encolar(Item{sig, p2, item.indice})
		}
	}
	return lista
}

// n = k * H
// O(k log k) + O(n log k) = O(n log k) = O(k * h log k)