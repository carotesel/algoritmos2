package parciales

/*
1) ¡El Dr. Casa necesita nuestra ayuda! Resulta que venía una caravana de k colectivos
que estaban yendo por la Autopista Buenos Aires - La Plata, y tuvieron un choque en
cadena (nadie gravemente herido para nuestra suerte). En cada colectivo viajaban h
personas, y por razones de edad y patologías previas, aunque no hayan heridos graves,
debemos atender primero a ciertos pacientes que a otros. Por suerte para el Dr. Casa,
su jefa, la Dra. Cudi, le ha asignado una prioridad a cada paciente (siendo los
pacientes con prioridad 1 aquellos con la mayor prioridad) y su compañero, el
Dr. Persecución, ha ordenado la fila de pacientes de cada colectivo por prioridad
(considerar que cada fila es una cola enlazada, donde los pacientes más prioritarios
están al frente de la cola).

Teniendo k colas de h pacientes cada una, y una función ObtenerPrioridad(paciente)
(que se ejecuta en O(1)), implementar un algoritmo para que el Dr. Casa pueda saber
el orden en el que deben ser atendidos los k · h pacientes. Indicar y justificar la
complejidad del algoritmo implementado.
*/

// h = cant personas x colectivo
// k = cant colas

type Item struct{
    persona string
    prioridad int
    indice int
}

func cmp (i1, i2 Item) int{
    return i1.prioridad - i2.prioridad
}

func Ordenarprioridad(colas []ColaEnlazada[string]) Lista[string]{
    k := len(colas)
    lista := CrearListaEnlazada[string]()
    heap := CrearHeap([]Item, cmp)

    for i:=0; i < k; i++{
        paciente := colas[i].Verprimero()
        prioridad := ObtenerPrioridad(paciente)
        heap.Encolar(Item{persona: paciente, prioridad: prioridad, indice: i})
    }

    for !heap.EstaVacia(){
        actual := heap.Desencolar()
        lista.AgregarUltimo(actual.persona)
        colas[actual.indice].Desencolar()

        for !colas[actual.indice].EstaVacia(){
            paciente := colas[actual.indice].Verprimero()
            prioridad := ObtenerPrioridad(paciente)
            heap.Encolar(Item{persona: paciente, prioridad: prioridad, indice: actual.indice})
        }
    }

    return lista
}



/*
2) Un fanático del universo Quebrando Lo Malo (quien escribe este parcilito) desea saber
cuándo fue la primera aparición de los personajes que aparecen tanto en Quebrando Lo Malo
como en Mejor Llama a Saul (una precuela de la primera). Para ello cuenta con dos
diccionarios de tipo Diccionario[Personaje, []Episodio] (el valor es un arreglo de episodios
en los que aparece el personaje, ordenado).

Implementar una función en Go:

    func primeraAparicion(
        hashQuebrandoLoMalo Diccionario[Personaje, []Episodio],
        hashMejorLlamaASaul Diccionario[Personaje, []Episodio],
    ) Diccionario[Personaje, []Episodio]

La cual, utilizando el iterador externo, devuelva un diccionario con la primera aparición de
cada personaje que aparezca en ambas series. El primer valor del array debe ser el episodio
correspondiente a Mejor Llama a Saul, y el segundo valor del array, el episodio
correspondiente a Quebrando Lo Malo. Indicar y justificar la complejidad de la función
implementada.
*/

func primeraAparicion(hashQuebrandoLoMalo Diccionario[Personaje, []Episodio], hashMejorLlamaASaul Diccionario[Personaje, []Episodio],
    ) Diccionario[Personaje, []Episodio]{
        res := CrearDiccionario[Personaje, []Episodio]()

        iter := hashMejorLlamaASaul.iterador() 

        for iter.HaySiguiente(){ // O(n)
            personaje := iter.VerActual()
            episodios_BCS := hashMejorLlamaASaul.Obtener(personaje) // O(1)

            if hashQuebrandoLoMalo.Pertenece(personaje){
                episodios_BB := hashQuebrandoLoMalo.Obtener(personaje)
                res.Guardar(personaje, []Episodio{episodios_BCS[0], episodios_BB[0]})
            }
            iter.Siguiente()
        }
        return res
    }

    // N = cantidad BCS
    // M = cantidad BB

    // Complejidad: O(n)


/*3) Implementar en Go una primitiva para el árbol binario:

    func (ab *Arbol[int]) EsSuma() bool

que determine si un árbol es suma o no. Un árbol es suma cuando el valor de todo nodo padre
es igual a la suma de sus hijos. Se considera que si un nodo es inexistente su valor es 0.
Si un nodo es hoja, entonces se define que es suma. Indicar y justificar la complejidad
de la primitiva.
*/

 func (ab *Arbol[int]) EsSuma() bool{
    _, esSuma := esAbSuma(ab)
    return esSuma
 }

 func esAbSuma (ab *Arbol[int]) (int, bool){
    if ab == nil{
        return 0, true
    }

    // NO OLVIDAR NODOS HIJOS!!
    if ab.izq == nil && ab.der == nil{
        return 0, true
    }

    sumaIzq, esIzq := esAbSuma(ab.izq)
    sumaDer, esDer := esAbSuma(ab.der)

    return esIzq && esDer && ab.dato == (sumaIzq + sumaDer)
 }

 // Complejidad: O(n)
