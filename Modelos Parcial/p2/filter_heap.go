/*1. Implementar una primitiva Filter(f func(T) bool) para el Heap, que deje al heap únicamente con los elementos para
los que la función devuelva true. La primitiva debe funcionar en O(n), siendo n la cantidad de elementos inicialmente
en el heap. Por supuesto, luego de aplicar la operación, el heap debe quedar en un estado válido para poder seguir
operando. Justificar la complejidad de la primitiva implementada.*/

// NO BORRAR UNO POR UNO POR EL LOG N

func (h *ColaPrioridad[T]) Filter(f func(T) bool) {
	newDatos := make([]T, 0, h.cantidad)
	for _, x := range h.datos{ // O(n) en el peor caso
		if f(x) {
			newDatos = append(newDatos, x)
		}
	}
	
	h.datos = newDatos
	h.cantidad = len(newDatos)
	heapify(h.datos, h.cmp) // O(n)
}