/*2. Un fanático del universo Quebrando Lo Malo (quien escribe este parcialito) desea saber cuándo fue 
la primera aparición de los personajes que aparecen tanto en Quebrando Lo Malo como en Mejor Llama a Saul 
(una precuela de la primera).

Para ello cuenta con dos diccionarios de tipo Diccionario[Personaje, []Episodio] (el valor es un arreglo de episodios
en los que aparece, el personaje, ordenado).

Implementar una función en Go:
func primeraAparicion(hashQuebrandoLoMalo, hashMejorLlamaASaul Diccionario[Personaje, []Episodio])
Diccionario[Personaje, []Episodio]

la cual, utilizando el iterador externo, devuelva un diccionario con la primera aparición de cada personaje que aparezca
en ambas series. El primer valor del array debe ser el episodio correspondiente a Mejor Llama a Saul, y el segundo
valor del array, el episodio correspondiente a Quebrando Lo Malo. Indicar y justificar la complejidad de la función
implementada.*/


func primeraAparicion(hashQuebrandoLoMalo, hashMejorLlamaASaul Diccionario[Personaje, []Episodio]) Diccionario[Personaje, []Episodio]{
	hash:= CrearDiccionario[Personaje, []Episodio]()

	iter := hashMejorLlamaASaul.Iterador()

	for iter.HaySiguiente() {
		personaje, episodios := iter.VerActual()

		if hashQuebrandoLoMalo.Pertenece(personaje){
			episodios_Quebrando_LoMalo := hashQuebrandoLoMalo.Obtener(personaje)
			arr_episodios := []Episodio{episodios[0], episodios_Quebrando_LoMalo[0]}
			hash.Guardar(personaje, arr_episodios)
		}
		iter.Siguiente()
	}

	return hash
}

// Complejidad: O(M)
// M = cant personajes better call saul