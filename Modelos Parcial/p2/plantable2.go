/*
3. Implementar en Go una primitiva para Árbol Binario func (ab *Arbol[int]) ArbolEsPlantable() bool que determine si un
árbol es plantable, o no. Para que lo sea, todo nodo debe cumplir: el dato del nodo debe ser mayor al dato de sus hijos 
(si los tiene), y además, el dato del nodo no puede superar la altura de dicho nodo. 
Implementar la primitiva en O(n), y justificar su complejidad.
*/

func (ab *Arbol[int]) ArbolEsPlantable() bool{
	ok, _ := esPlant(ab)
}

func esPlant (ab *Arbol[int]) (bool, int){
	if ab == nil{
		return true, 0
	}

	okIzq, hIzq := esPlant(ab.izq)
	okDer, hDer := esPlant(ab.der)

	hNodo := max(hIzq, hDer) + 1

	condHijos := true

	if ab.izq != nil{
		condHijos = condHijos && ab.dato > ab.izq.dato
	}

	if ab.der != nil{
		condHijos = condHijos && ab.dato > ab.der.dato
	}

	return okIzq && okDer && ab.dato <= hNodo && condHijos, hNodo
}
