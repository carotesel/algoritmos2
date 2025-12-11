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
	paciente string
	prioridad int
	idx int
}

func ordenAtencion(colas []ColaEnlazada[string]) []string{

	k := len(colas)
	heap := CrearHeap[Item](cmp)
	res := make([]string, 0, k * len(colas[0]))

	// encolo los primeros de cada cola
	for i:=0; i<k; i++{
		paciente := colas[i].Verprimero()
		prioridad := ObtenerPrioridad(paciente)
		heap.Encolar(Item{paciente: paciente, prioridad: prioridad, idx: i})
	}

	// mientras el heap tenga cosas, encolo el mas chico, lo saco de la cola y encolo los q quedan en la cola
	for !heap.EstaVacia(){
		actual := heap.Desencolar()
		res = append(res, actual.paciente)
		colas[actual.idx].Desencolar()

		if !colas[actual.idx].EstaVacia(){
			paciente := colas[actual.idx].Verprimero()
			prioridad := ObtenerPrioridad(paciente)
			heap.Encolar(Item{paciente: paciente, prioridad: prioridad, idx: i})
		}
	}
	return res
}

// k merge de arr de arr de int. menor a mayor

type Item2 struct{
	nro int
	pos int
	idx int
}

func kMerge(arr [][]int) []int{
	k := len(arr)
	res := make([]int, 0, k * len(arr[0]))
	heap := CrearHeap[Item2](cmp)

	// encolo primeros
	for i:=0; i<k; i++{
		heap.Encolar(Item2{nro: arr[i][0], pos:0, idx: i})
	}

	for !heap.EstaVacia(){
		actual := heap.Desencolar()
		res = append(res, actual.nro)
		posicion := actual.pos + 1

		for posicion < len(arr[actual.idx]){
			heap.Encolar(Item2{nro: arr[actual.idx][posicion], pos:posicion, idx: actual.idx})
		}
	}
	return res
}

/*
Ejercicio 1
Implementar una primitiva para el Heap con la siguiente firma: func (heap *Heap[T]) KMenores(k int) []T
La primitiva debe devolver un arreglo con los k elementos menores del heap, ordenados de menor a mayor. 
El heap original no debe ser modificado. La primitiva debe funcionar en O(k log n), siendo n la cantidad de 
elementos en el heap. Indicar y justificar la complejidad de la primitiva implementada.
*/

func cmp(a, b T) int{
	return a.valor - b.valor
}

func (heap *Heap[T]) KMenores(k int) []T{
	heapAux:= CrearHeap[int](cmp)
	datos := heap.datos
	res := make([]T, 0)
	heapAux.Encolar(0) // raiz del heap

	for len(res) < k && !heapAux.EstaVacia(){
		min := heapAux.Desencolar()
		res = append(res, datos[min])

		izq := min * 2 + 1
		der := min * 2 + 2

		if izq < len(datos){
			heapAux.Encolar(izq)
		}

		if der < len(datos){
			heapAux.Encolar(der)
		}
	}

	return res
}

/*
Ejercicio 2
Implementar una primitiva para el árbol binario (no ABB) que determine si el mismo cumple propiedad de AVL. 
La estructura del árbol es: 
type arbol struct {
	izq  *arbol
	der  *arbol
	dato int
}
*/

func (ab *arbol) EsAVL(cmp) bool{
	esAvl, _ := ab.esAvlRec()
	return esAvl
}

func (ab *arbol) esAvlRec(bool,int){
	if ab == nil{
		return true, 0
	}

	esIzq, hIzq := ab.izq.esAvlRec
	esDer, hDer := ab.der.esAvlRec

	altura := max(hIzq, hder) + 1

	if !esIzq || !esDer {
		return false, 0
	}

	if math.abs(hIzq - hder) > 1{
		return false, 0
	} else{
		return true, altura
	}
}

/*
Ejercicio 3
Implementar un algoritmo que reciba un arreglo desordenado de n enteros y un número K y determinar en O(n) 
si existe un par de elementos en el arreglo que sumen exactamente K.
*/

/*
Ejercicio 4
3. Tenemos un arreglo de n números en el que cada elemento se encuentra a lo sumo k
posiciones de la que le correspondería si estuviera ordenado(2 ≤ k < n). Implementar una
función que reciba el arreglo y el valor de k y ordene el arreglo en O(n log k). Justificar la
complejidad del algoritmo implementado.
*/

func cmp (a, b int) int{
	return a - b
}

func casiOrdenado(arr []int, k int) []int{
	res := make([]ine, len(arr))
	heapMin := CrearHeap[int](cmp)
	index := 0 

	for i:=0; i<k; i++{ // O(log k)
		heapMin.Encolar(arr[i])
	}

	// O(n)
	for j:=k+1; j<len(arr); j++{
		res[index] = heapMin.Desencolar()
		index++
		heapMin.Encolar(arr[j])
	}

	for !heapMin.EstaVacia(){
		res[index] = heapMin.Desencolar()
		index++
	}
	return res
}


