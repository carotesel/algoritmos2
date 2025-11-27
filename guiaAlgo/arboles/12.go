/*(★★★★) 
♠ Implementar una primitiva para el AB que reciba dos arreglos (o listas) de cadenas. 
El primer arreglo corresponde al preorder de un árbol binario. 
El segundo al inorder del mismo árbol (ambos arreglos tienen los mismos elementos, sin repetidos). 

La función debe devolver un árbol binario que tenga dicho preorder e inorder. 
Indicar y justificar el orden de la primitiva (tener cuidado con este punto). 

Considerar que la estructura del árbol binario es:

    type Arbol struct {
        izq *Arbol
        der *Arbol
		clave int
    }*/

func Reconstruir(preorder, inorder []int) *Arbol {
	if preorder == nil{
		return nil
	}

	raiz := preorder[0]
	pos_raiz_inorder := 0

	for i:= 0; i<len(inorder); i++ { // puede ser O(n) en el peor caso
		if inorder[i] == raiz{
			pos_raiz_inorder = i
			break
		}
	}

	inorder_izq := inorder[: pos_raiz_inorder]
	inorder_der := inorder[pos_raiz_inorder+1:]

	tam_izq := len(inorder_izq)

	preorder_izq := preorder[1 : 1+tam_izq]
	preorder_der := preorder[tam_izq + 1::]

	raiz := &Arbol{clave: raiz}

	raiz.izq = Reconstruir(preorder_izq, inorder_izq)
	raiz.der = Reconstruir(preorder_der, inorder_der)

	return raiz

}

// ⭐ Complejidad

//Buscar la raíz en inorder → O(n)

//Llamar recursivamente n veces → O(n) veces

//Complejidad total: O(n²)

// ES CUADRATICO PORQUE ESE O(N) SE EJECUTA PARA CADA NODO PORQUE LA FC ES RECURSIVA, POR ESO ES CUADRATICO