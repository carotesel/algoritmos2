/*
5.(★★) Implementar una función de orden O(n) que dado un arreglo de n números enteros,
devuelva true o false según si existe algún elemento que aparezca más de la mitad de las veces. 
Justificar el orden de la solución.*/

package masmitad

func masDeLaMitad(arr []int) bool{
	dicc_claves := CrearDiccionario[int, int]() // O(1)
	mitad := len(arr)/2 // O(1)

	for _, n range arr{ // O(n)
		if dicc_claves.Pertenece(n){
			cantidad := dicc_claves.Obtener(n)
			dicc_claves.Guardar(n, cantidad+1)
		} else{
			dicc_claves.GUardar(n, 1)
		}
	}

	// recorro el hash con un iterador EXTERNO (funcion -> barbara)

	iter := dicc_claves.Iterador()

	for iter.HaySiguiente(){ // O(n)
		_, valor := iter.VerActual()
		if valor > mitad{
			return true
		}
		iter.Siguiente()
	}

	return false
}

/* Complejidad:

O(n) + O(n) + 2.O(1) = 2 O(n) + 2 O(1) = O(n)