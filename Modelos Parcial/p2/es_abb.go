/*✅ Ejercicio 3 — EsABB(cmp)

Consigna:
Implementar EsABB(raiz, cmp) que determine si un árbol binario cumple la propiedad de ABB 
usando una función de comparación cmp(a,b).

Debe verificarse que:
- todas las claves del subárbol izquierdo sean menores,
- todas las claves del subárbol derecho sean mayores.

La solución debe ser O(n).*/

// paso rango para verificar correctamente

func (ab *Arbol) EsABB (cmp func(a, b int) int) bool {
	return esABB(ab.raiz, nil, nil, cmp)
}

func esABB (nodo *nodo, min, max *int, cmp func(a, b int) int) bool {
	if nodo == nil{
		return true
	}

	if min != nil && cmp(nodo.clave, *min) <= 0{
		return false
	}

	if max != nil && cmp(nodo.clave, *max) > 0{
		return false
	}

	return esABB(nodo.izq, min, &nodo.clave, cmp) && esABB(nodo.der, &nodo.clave, max, cmp)
}
