/*3. Implementar en Go una primitiva para Árbol Binario func (ab *Arbol[int]) ArbolEsPlantable() bool 
que determine si un árbol es plantable, o no. Para que lo sea, todo nodo debe cumplir: el dato del nodo 
debe ser mayor al dato de sus hijos (si los tiene), y además, el dato del nodo no puede superar la altura 
de dicho nodo. Implementar la primitiva en O(n), y justificar su complejidad.

A fines del ejercicio considerar la estructura del árbol como la definida al dorso.

type Arbol struct {
dato int
izq *Arbol
der *Arbol
}*/

func (ab *Arbol[int]) ArbolEsPlantable() bool{
	ok, _ := esPlantableAux(ab.dato)
	return ok
}

func esPlantableAux (ab *Arbol[int]) (bool, int){
	if ab == nil{
		return true, 0
	}

	okIzq, sumaIzq := esPlantableAux(ab.izq)
    okDer, sumaDer := esPlantableAux(ab.der)

	altura := max(sumaIzq, sumaDer) + 1

	condAltura := ab.dato >= altura

	condHijos := true

	if ab.izq != nil{
		condHijos = condHijos && ab.dato > ab.izq.dato
	}

	if ab.der != nil{
		condHijos = condHijos && ab.dato > ab.der.dato
	}

	return okIzq && okDer && condAltura && condHijos, altura
}
