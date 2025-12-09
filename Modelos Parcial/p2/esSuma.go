/*3. Implementar en Go una primitiva para el árbol binario func (ab *Arbol[int]) EsSuma() bool que determine si un
árbol es suma o no. Un árbol es suma cuando el valor de todo nodo padre es igual a la suma de sus hijos. Se considera
que si un nodo es inexistente su valor es 0. Si un nodo es hoja, entonces se define que es suma. Indicar y justificar la
complejidad de la primitiva. Al dorso del examen pueden ver tanto la estructura del árbol, como algunos ejemplos.

type Arbol[T any] struct {
dato T
izq *Arbol
der *Arbol
}
*/

func (ab *Arbol[int]) EsSuma() bool{
	ok, _ := esSumaAux(ab.dato)
	return ok
}

func esSumaAux(ab *Arbol[int]) (bool, int) {
	if ab == nil{
		return true, 0
	}

	// HOJAS: POR DEFINICION SON SUMA
	if ab.izq == nil && ab.der == nil{
		return true, ab.dato
	}

	okIzq, sumaIzq := esSumaAux(ab.izq)
    okDer, sumaDer := esSumaAux(ab.der)

	okNodo := ab.dato == sumaIzq + sumaDer

	return okIzq && okDer && okNodo, sumaIzq + sumaDer + ab.dato
}

// O(n)