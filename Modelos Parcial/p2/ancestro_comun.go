/*1. Implementar una primitiva para el ABB func (arbol *abb[K, V]) AncestroComun(clave1, clave2 K) K 
que reciba 2 claves y devuelva el último ancestro en común entre ambas claves. 

Dicho ancestro en común podría ser incluso alguna de estas claves. Si alguna clave pasada no se encuentra en el árbol, 
finalizar con panic. Indicar y justificar la complejidad de la primitiva implementada.

Mostramos ejemplos de resultados esperados de invocar la primitiva al árbol del dorso:
arbol.AncestroComun(1, 4) --> 2
arbol.AncestroComun(2, 4) --> 2
arbol.AncestroComun(9, 1) --> 5
*/

func (arbol *abb[K, V]) AncestroComun(clave1, clave2 K) K{

	if !arbol.Pertenece(clave1) || !arbol.Pertenece(clave2){
		panic("Una de las claves no pertenece al arbol")
	}

	if arbol == nil{
		panic("Árbol vacío")
	}

	if clave1 < arbol.clave && clave2 < arbol.clave {
		return arbol.izq.AncestroComun(clave1, clave2)
	}  
	
	if clave1 > arbol.clave && clave2 > arbol.clave{
		return arbol.der.AncestroComun(clave1, clave2)
	}
	return arbol.clave
}

//Complejidad: T(n/2) + O(1)
// A = 1, B = 2, C = 0
// log b (a) = 0 = C -> O(n^c log n) = O(log n) 
// En el peor caso puede ser o(n) si el arbol no esta balanceado