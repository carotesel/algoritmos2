/*6.
(★★) Asumiendo que se tiene disponible una implementación completa del TDA Hash, se desea implementar 
una función que decida si dos Hash dados representan o no el mismo Diccionario. 

Considere para la solución que es de interés la mejor eficiencia temporal posible. 
Indique, para su solución, eficiencia en tiempo y espacio. 
Nota: Dos tablas de hash representan el mismo diccionario si tienen la misma cantidad de elementos; 
todas las claves del primero están en el segundo; todas las del segundo, en el primero; 
y los datos asociados a cada una de esas claves son iguales (se pueden comparar los valores con “==”). */ 

// 2 hash iguales:
// - misma cant de cosas
// - claves h1 -> claves h2
// - claves h2 -> claves h1
// - valores 1 == valores 2

// con verificar de 1 lado alcanza! (no hay orden xd)

func SonElMismo[K comparable, V comparable](dicc1, dicc2 Diccionario[K,V]) bool{

	if dicc1.Cantidad() != dicc2.Cantidad(){
		return false
	}

	iter1 := dicc1.Iterador()

	for iter1.HaySiguiente(){
		actual, valor := iter1.VerActual()

		if !dicc2.Pertenece(actual){
			return false
		} else if dicc2.Obtener(actual) != valor{
				return false
			}
		iter1.Siguiente()
	}
	return true
}
