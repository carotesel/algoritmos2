/*
1. Implementar una primitiva para el ABB Predecesor(clave K) K que reciba una clave
(que puede estar en el árbol, o no) y devuelva la clave inmediatamente anterior a esta en el
recorrido inorder. Si no hay ninguna anterior, finalizar con un panic. Indicar y justificar
la complejidad de la primitiva.
*/

func (abb *Arbol[K, V]) Predecesor(clave K) K {
	if abb.raiz == nil {
		panic("Arbol vacío")
	}

	// Llamamos a la función recursiva auxiliar con candidato inicial vacío
	encontrado, res := predecesorRec(abb.raiz, abb.cmp, clave, nil, false)
	if !encontrado {
		panic("No hay predecesor")
	}
	return res
}

// ------------------------------------------------------------
// nodo: nodo actual del árbol
// cmp: función de comparación
// clave: clave buscada
// candidato: último valor < clave encontrado hasta ahora
// tieneCandidato: indica si candidato es válido
// ------------------------------------------------------------

func predecesorRec[K comparable, V any](
	nodo *nodoABB[K, V],
	cmp func(K, K) int,
	clave K,
	candidato K,
	tieneCandidato bool,
) (bool, K) {

	if nodo == nil {
		// terminó la búsqueda sin encontrar nodo igual a clave,
		// devolvemos el último candidato válido
		return tieneCandidato, candidato
	}

	// Caso exacto: nodo.clave == clave
	if cmp(clave, nodo.clave) == 0 {

		// Si tiene hijo izquierdo, el predecesor es el máximo del subárbol izquierdo
		if nodo.izq != nil {
			max := maximo(nodo.izq)
			return true, max
		}

		// Si no tiene hijo izquierdo, el predecesor es el último candidato registrado
		return tieneCandidato, candidato
	}

	// Si clave < nodo.clave → buscamos a la izquierda (este nodo no sirve como predecesor)
	if cmp(clave, nodo.clave) < 0 {
		return predecesorRec(nodo.izq, cmp, clave, candidato, tieneCandidato)
	}

	// Caso clave > nodo.clave:
	// nodo.clave es menor → podría ser predecesor → lo guardamos
	return predecesorRec(nodo.der, cmp, clave, nodo.clave, true)
}

// ------------------------------------------------------------
// Función auxiliar: encontrar el máximo de un subárbol
// ------------------------------------------------------------
func maximo[K comparable, V any](nodo *nodoABB[K, V]) K {
	for nodo.der != nil {
		nodo = nodo.der
	}
	return nodo.clave
}


/*
2. Implementar, para un Hash Cerrado, una primitiva Modificar(k int, aplicar
func(V) V) que dado un número k y una función a aplicar modifique el valor del
dato asociado a las primeras k claves para del hash que se encuentre, con el resultado
de aplicarle la función aplicar. En caso de que el Hash tenga menos de k elementos,
simplemente aplicar a todos los valores. Indicar y justificar la complejidad de la primitiva
implementada.
*/

func (dicc *hashCerrado[K, V]) Modificar(p int, aplicar func(V) V){
	
	if p <= 0{
		return
	}

	contador := 0

	for i:=0; i< len(dicc.tabla) && contador < p{
		celda := dicc.tabla[i]
		
		if celda.estado == OCUPADO{
			modificado := aplicar(celda.dato)
			celda.dato = modificado
			contador++
		}
	}
}


/*
3. Realizar un seguimiento sobre el TDA Heap (de máximos) que resulta de realizar las
siguientes operaciones. Se puede realizar el seguimiento directamente sobre la representación
en árbol del heap:
a. Crear un heap a partir del arreglo [6, 4, 2, 6, 5, 1, 0, 9].
b. Sobre el heap resultante del paso anterior, realizar: Encolar(7), Desencolar(),
Encolar(19),  Encolar(8), Desencolar(). DONE
*/