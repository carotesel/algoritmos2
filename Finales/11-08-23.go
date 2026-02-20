func main(arr []int) int {

	var mayoritario int

	if arr[0] == arr[1] || arr[0] == arr[2] {
		mayoritario = arr[0]
	} else {
		mayoritario = arr[1]
	}

	return buscarDiferente(arr, mayoritario, 0, len(arr)-1)
}

func buscarDiferente(arr []int, mayoritario int, izquierda, derecha int) int {
    // Caso base: cuando izquierda == derecha, hemos encontrado el elemento
    if izquierda == derecha {
        return izquierda
    }
    
    medio := (izquierda + derecha) / 2
    
    // Si el elemento del medio es diferente al mayoritario, ¡es el que buscamos!
    if arr[medio] != mayoritario {
        return medio
    }
    
    // Si el del medio es mayoritario, el diferente está en la otra mitad
    // Pero ¿en cuál? Depende de si el primer elemento es mayoritario o no
    
    if arr[izquierda] != mayoritario {
        // El diferente está en la mitad izquierda
        return buscarDiferente(arr, mayoritario, izquierda, medio-1)
    } else {
        // El diferente está en la mitad derecha
        return buscarDiferente(arr, mayoritario, medio+1, derecha)
    }
}

/**/
def plan_minimo_cuatrimestres(grafo):

    grados = grados_entrada(grafo)
    cola = Cola()

    # materias sin correlativas
    for v in grafo.obtener_vertices():
        if grados[v] == 0:
            cola.Encolar(v)

    resultado = []

    while not cola.EstaVacia():

        siguiente_cola = Cola()
        nivel_actual = []

        while not cola.EstaVacia():
            v = cola.Desencolar()
            nivel_actual.append(v)

            for w in grafo.adyacentes(v):
                grados[w] -= 1
                if grados[w] == 0:
                    siguiente_cola.Encolar(w)

        resultado.append(nivel_actual)
        cola = siguiente_cola

    return resultado


/*4. Implementar un algoritmo que reciba dos cadenas (strings) y determine si son anagramas entre sí. Indicar y justificar
la complejidad del algortmo implementado.*/

ffunc sonAnagramas(s1, s2 string) bool{
	if len(s1) != len(s2){
		return false
	}

	dicc := CrearDiccionario[string, int]()

	for _, letra := range s1{
		if dicc.Pertenece(letra){
			dicc.Guardar(letra, dicc.Obtener(letra)+1)
		} else {
			dicc.Guardar(letra, 1)
		}
	}

	for _, letra := range s2{
		if !dicc.Pertenece(letra){
			return false
		}
		cantidad := dicc.Obtener(letra) - 1 
		if cantidad == 0{
			dicc.Borrar(letra)
		} else {
			dicc.Guardar(letra, cantidad)
		}
	}

	return dicc.Cantidad() == 0
}

// n = len(s1)
// m = len(s2)
// Complejidad: O(n)+O(m)=O(n+m)
	
/*5. Implementar una primitiva de árbol binario de búsqueda que devuelva un diccionario en el cual las claves sean los
niveles (int) y los datos sean listas de todos las claves del ABB que se encuentran en dicho nivel. Indicar y justificar la
complejidad del algoritmo implementado.*/

type Elem[K comparable, V any] struct {
	nodo  *nodoABB[K, V]
	nivel int
}

func (abb *abb[K, V]) Niveles() Diccionario[int, Lista[K]] {

	dicc := CrearHash[int, Lista[K]]()
	if abb.raiz == nil {
		return dicc
	}

	cola := CrearColaEnlazada[Elem[K,V]]()
	cola.Encolar(Elem[K,V]{nodo: abb.raiz, nivel: 0})

	for !cola.EstaVacia() {

		v := cola.Desencolar()

		if !dicc.Pertenece(v.nivel){
			lista := CrearListaEnlazada[K]()
			dicc.Guardar(v.nivel, lista)
		} else{
			lista := dicc.Obtener(v.nivel)
		}

		lista.AgregarUltimo(v.nodo.clave)
		

		if v.nodo.izq != nil {
			cola.Encolar(Elem[K,V]{v.nodo.izq, v.nivel+1})
		}

		if v.nodo.der != nil {
			cola.Encolar(Elem[K,V]{v.nodo.der, v.nivel+1})
		}
	}

	return dicc
}
