/*(★★★) Definimos como quiebre en un árbol binario cuando ocurre que:
un hijo derecho tiene un solo hijo, y es el izquierdo
un hijo izquierdo tiene un solo hijo, y es el derecho
Implementar una primitiva para el árbol binario func (arbol Arbol) Quiebres() int 
que, dado un árbol binario, nos devuelva la cantidad de quiebres que tiene. 

La primitiva no debe modificar el árbol. La estructura del tipo Arbol es:

    type Arbol struct {
        izq *Arbol
        der *Arbol
    }
Indicar y justificar el orden de la primitiva, e indicar el tipo de recorrido implementado.*/

func (ab *Arbol) Quiebres() int{
	if ab == nil{
		return 0
	}

	izq := ab.izq.Quiebres()
	der := ab.der.Quiebres()

	quiebre := 0

	if ab.der != nil && ab.der.der == nil && ab.der.izq != nil{
		quiebre = 1
	}

	if ab.izq != nil && ab.izq.izq == nil && ab.izq.der != nil{
		quiebre = 1
	}

	return quiebre + izq + der
}

// complejidad O(n)