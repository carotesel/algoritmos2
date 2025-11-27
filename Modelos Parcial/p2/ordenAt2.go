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

// k = colectivos
// h = personas x colectivo
// ObtenerPrioridad(paciente)

type Item struct{
	nombre string
	prioridad int
	idx int
}

func cmp(i, i2 Item) int{
	return i.prioridad - i2.prioridad
}

func ordenAtencion(colas []ColaEnlazada[string]) []string{
	heap := CrearHeap([]Item, cmp)
	k := len(colas)
	res := make([]string, k * len(colas[0]))

	for i:=0; i < k; i++{
		paciente := colas[i].Verprimero()
		prioridad := ObtenerPrioridad(colas[i][0])
		heap.Encolar(Item{paciente, prioridad, i})
	}

	for !heap.EstaVacia(){
		min := heap.Desencolar()
		res = append(res, min.nombre)

		// saco el q acabo de guardar
		colas[min.idx].Desencolar()

		for !colas[min.idx].EstaVacia(){
			paciente := colas[min.idx].Verprimero()
			prioridad := ObtenerPrioridad(paciente)
			heap.Encolar(Item{paciente, prioridad, min.idx})
		}

	}
	return res
}