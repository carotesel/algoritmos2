/*(★★) Escribir una primitiva con la firma func (arbol *Arbol) Invertir() que invierta el árbol 
binario pasado por parámetro, de manera tal que los hijos izquierdos de cada nodo se conviertan en hijos derechos.

La estructura Arbol respeta la siguiente definición:

    type Arbol struct {
        izq *Arbol
        der *Arbol
    }

Indicar el orden de complejidad de la función implementada.*/

func (arbol *Arbol) Invertir(){
	if arbol == nil{
		return
	}

	arbol.izq, arbol.der = arbol.der, arbol.izq

	arbol.izq.Invertir()
	arbol.der.Invertir()

}

// Orden:
// T(n) = 2 T(n/2) + O(1)
// Log 2 (2) = 1 < C
// Compl: O(n^c) = O(n)