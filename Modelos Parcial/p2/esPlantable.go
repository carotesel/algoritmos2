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

func (ab *Arbol) ArbolEsPlantable() bool{
	ok, _ := EsPlantable(ab)
	return ok
}

func esPlantable (ab *Arbol) (bool, int){
	if ab == nil{
		return true, 0
	}

	okIzq, hizq := esPlantable(ab.izq)
	okDer, hder := esPlantable(ab.der)

	altura := 1 + max(hizq, hder)

	condAltura := ab.dato <= altura

	condHijos := true

	if ab.izq != nil{
		condHijos = condHijos && ab.dato > ab.izq.dato
	}

	if ab.der != nil{
		condHijos = condHijos && ab.dato > ab.der.dato
	}

	return okIzq && okDer && condAltura && condHijos, altura
}

/*La función auxiliar visita cada nodo exactamente una vez. En cada visita hace una cantidad constante de operaciones: 
compara datos, combina resultados de los hijos y calcula la altura como el máximo entre dos valores.
No hay recorridos repetidos ni cálculos de altura adicionales.

Por lo tanto, el tiempo total es O(n), donde n es la cantidad de nodos del árbol. */