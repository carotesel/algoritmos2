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



/*3) Implementar en Go una primitiva para el árbol binario:

    func (ab *Arbol[int]) EsSuma() bool

que determine si un árbol es suma o no. Un árbol es suma cuando el valor de todo nodo padre
es igual a la suma de sus hijos. Se considera que si un nodo es inexistente su valor es 0.
Si un nodo es hoja, entonces se define que es suma. Indicar y justificar la complejidad
de la primitiva.
*/

