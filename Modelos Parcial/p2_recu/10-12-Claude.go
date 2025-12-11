/*
1. Implementar un algoritmo que reciba un arreglo de n números, y un número k, y devuelva los k números dentro del
arreglo cuya suma sería la máxima (entre todos los posibles subconjuntos de k elementos de dicho arreglo). Indicar y
justificar la complejidad de la función implementada.
*/

func cmp(a, b int) int{
	return b - a
}

func sumaMax (arr []int, k int) []int{
	res := make([]int, 0, k)
	heap_max := CrearHeapArr(arr, cmp)

	for i:=0; i<k; i++{
		res = append(res, heap_max.Desencolar())
	}
	return res
}

// O(k log n)

/*
1. Implementar la primitiva Invertir() para el Heap, que invierta su forma de comparar los elementos (es decir, si era 
de máximos, ahora sea de mínimos), sin modificar las funciones y primitivas previamente implementadas (simplemente contener 
todo el cambio dentro de la primitiva). Indicar y justificar la complejidad del algoritmo. Indicar qué consecuencias podría 
tener esta forma de implementación si se invoca a la primitiva Invertir una cantidad k de veces, y cómo podría resolverse si 
se permitiera modificar otras funciones y/o primitivas (y/o la estructura del heap en sí).
*/

func (heap *heap[T]) Invertir(){
	for dato, i := range heap.datos{
		modificado := -dato
		heap.datos[i] = modificado
	}
	heapify(heap.datos, heap.cmp)
}

/*
2. Implementar una primitiva para el hash cerrado filtro(func(V) bool) que elimine del hash todas las claves del mismo que
tengan asociado un valor para el cuál la función devuelva false. 
Indicar y justificar la complejidad de la primitiva implementada.

Ejemplo: si mi diccionario es de cadenas a números como el de arriba, e invocamos a la primitiva con una función que devuelve 
true para los números pares y false para los impares, el diccionario debe quedar como el de abajo:
{"koala": 3, "rana": 2, "gato": 2, "perro": 5, "canguro": 4 }
↓
{"rana": 2, "gato": 2, "canguro": 4}
*/

func (dicc *hashCerrado[K, V]) Filtro(f func(V) bool){
	tabla := dicc.tabla
	nuevaTabla := make([]CeldaHash[K, V], hash.tam)
	dicc.tabla = nuevaTabla

	for _, celda := range tabla{
		if celda.estado == OCUPADO{
			if f(celda.valor){
				dicc.Guardar(celda.clave, celda.valor)
			}
		}
	}
}

// O(n)

/*
4. Implementar una función que reciba un arreglo A de n enteros y un número k,
   y devuelva un nuevo arreglo R tal que, para cada posición i de R, el valor
   sea la máxima suma obtenible de EXACTAMENTE k elementos dentro del rango
   A[0..i] (incluyendo a i). Si para una posición i no hay suficientes
   elementos (i < k-1), entonces R[i] debe valer -1.

   Ejemplo:
     A = [1, 5, 3, 4, 2, 8], k = 3
     Resultado esperado: R = [-1, -1, 9, 12, 12, 17]

   Requisito de eficiencia:
     La complejidad del algoritmo debe ser MEJOR que O(n * k).

   Indicar y justificar la complejidad del algoritmo implementado.
*/

func cmp(a, b int) int{
	return a - b
}

func maxSumaObtenibleK(a []int, k int) []int {
    heapMin := CrearHeap(cmp)  // heap de mínimos
    suma := 0
    r := make([]int, 0, len(a))

    for i, elem := range a {

        if i < k {
            heapMin.Encolar(elem)
            suma += elem

            if i < k-1 {
                r = append(r, -1)
            } else {          
                r = append(r, suma)
            }
            continue
        }

        // i >= k
        min := heapMin.VerMin()
        if elem > min {
            eliminado := heapMin.Desencolar()
            suma -= eliminado
            heapMin.Encolar(elem)
            suma += elem
        }

        r = append(r, suma)
    }

    return r
}

/* 
2. Implementar una primitiva eliminarColisiones(clave K) []K para el hash,
que elimine del hash todas las claves que colisionen con la clave pasada
por parámetro en el estado actual (sin eliminar dicha clave del
diccionario, si se encuentra) y devuelva dichas claves.

Implementar tanto para el hash abierto como para el hash cerrado. Si no
se implementa para alguno, el ejercicio no estará aprobable.

Indicar y justificar la complejidad de la primitiva para ambos casos.
*/

// HASH ABIERTO

func (dicc *hashAbierto[K, V]) eliminarColisiones(clave K) []K{
	indice := hash(clase) % dicc.tam
	tabla := dicc.tabla[indice] // TABLA ES UNA LISTA

	res := make([]K, 0)

	iter := tabla.Iterador()

	for iter.HaySiguiente(){
		actual := iter.VerActual()
		if actual != clave{
			res = append(res, actual.clave)
			dicc.cantidad--
			iter.Borrar()
		}
		iter.Siguiente()
	}

	dicc.tabla[indice] = nuevaTabla

	if (dicc.cantidad / len(dicc.tabla)) < FACTOR_MINIMO{
		nuevoT := len(dicc.tabla) / 2

		if nuevoT < FACTOR_MINIMO{
			nuevoT = FACTOR_MINIMO
		}
		dicc.redimensionar(nuevoT)
	}

	if (dicc.cantidad / len(dicc.tabla)) > FACTOR_MAXIMO{
		nuevoT := len(dicc.tabla) * 2
		dicc.redimensionar(nuevoT)
	}
	return res
}

// HASH CERRADO

func (dicc *hashCerrado[K, V]) eliminarColisiones(clave K) []K{
	cHashed := hash(clave) % dicc.tam
	tabla := dicc.tabla
	
	dicc.tabla = make([]CeldaHashCerrado[K, V], dicc.tam)	
	dicc.borrados = 0
	dicc.cantidad = 0

	res := make([]K, 0)

	for i := range tabla{
		celda := &tabla[i]
		if celda.estado != OCUPADO{
			continue
		}
		if clave == celda.clave{
			dicc.Guardar(celda.clave, celda.dato)
		} else if clave != celda.clave && cHashed == hash(celda.clave){
			res = append(res, celda.clave)
			dicc.cantidad--
		} else{
			dicc.Guardar(celda.clave, celda.dato)
		}
	}
	return res
}

/*
5. Implementar una primitiva del ABB que, dado un valor entero M, una clave inicial "inicio"
   y una clave final "fin", devuelva una lista con todos los datos cuyas claves estén
   dentro del rango [inicio, fin] y que además se encuentren dentro de los primeros M niveles
   del árbol (considerando a la raíz en el nivel 1).

   Ejemplo:
     Si se tiene el siguiente ABB:

              10
             /  \
            5    15
           / \   / \
          3   8 12 20
             /     \
            7       14

     y se llama con M = 3, inicio = 5 y fin = 15,
     el resultado debe contener los datos de las claves 10, 5, 8, 12, 15 (en cualquier orden).

   Se debe indicar y justificar la complejidad temporal del algoritmo implementado.
*/

func (ab *abb[K, V]) ClavesRangoNivel(M int, ini, fin K) Lista[V]{
	lista := CrearListaEnlazada[V]()
	ab.raiz.clavesRangoN(M, 1, ini, fin, lista, ab.cmp)
	return lista
}

func (nodo *nodoABB[K, V]) clavesRangoN(M, nActual int, ini, fin K, lista Lista[V], cmp func(K, K) int){
	if nodo == nil{
		return
	}

	if nActual > M{
		return
	}

	if cmp(nodo.clave, ini) < 0{
		return nodo.der.clavesRangoN(M, nActual+1, ini, fin, lista, ab.cmp)
	}

	if cmp(nodo.clave, fin) > 0{
		return nodo.izq.clavesRangoN(M, nActual+1, ini, fin, lista, ab.cmp)
	}

	lista.AgregarUltimo(nodo.dato)

	nodo.izq.clavesRangoN(M, nActual+1, ini, fin, lista, ab.cmp)
	nodo.der.clavesRangoN(M, nActual+1, ini, fin, lista, ab.cmp)
}

/* (FALOPA)
2. Implementar la primitiva Interseccion(otro *abb[K, V]) Lista[K] para el ABB que nos devuelva una lista ordenada con la
intersección entre el árbol y el recibido por parámetro, que estén ocupando el mismo lugar en el árbol. 
Indicar y justificar la complejidad del algoritmo implementado. 

En el ejemplo a continuación, la intersección sería [4, 10, 18, 20].
*/

func (ab *abb[K, V]) Interseccion(otro *abb[K, V]) Lista[K] {
	lista := CrearListaEnlazada[K]()
	ab.raiz.interseccion(otro.raiz, lista)
	return lista
}

func (nodo *nodoABB[K, V]) interseccion(otro *nodoABB[K, V], lista Lista[K]){
	if nodo == nil || otro == nil{
		return
	}

	nodo.izq.interseccion(otro.izq, lista)

	if nodo.clave == otro.clave{
		lista.AgregarUltimo(nodo.clave)
	}

	nodo.der.interseccion(otro.der, lista)
}

/*
1. Implementar una primitiva para el ABB Predecesor(clave K) K que reciba una clave
(que puede estar en el árbol, o no) y devuelva la clave inmediatamente anterior a esta en el
recorrido inorder. Si no hay ninguna anterior, finalizar con un panic. Indicar y justificar
la complejidad de la primitiva.
*/

func (ab *abb[K,V]) Predecesor(clave K) K {
    nodo := buscar(ab.raiz, clave)
    if nodo == nil {
        panic("clave no existe")
    }

    // Caso sencillo
    if nodo.izq != nil {
        return max(nodo.izq)
    }

    // Caso sin hijo izquierdo → buscar último menor desde raíz
    return ultimoMenor(ab.raiz, clave)
}

func buscar(n *nodoABB[K,V], clave K) *nodoABB[K,V] {
    if n == nil {
        return nil
    }
    if clave < n.clave {
        return buscar(n.izq, clave)
    }
    if clave > n.clave {
        return buscar(n.der, clave)
    }
    return n
}

func max(n *nodoABB[K,V]) K {
    for n.der != nil {
        n = n.der
    }
    return n.clave
}

func ultimoMenor(n *nodoABB[K,V], clave K) K {
    var candidato *nodoABB[K,V]

    for n != nil {
        if clave > n.clave {
            candidato = n
            n = n.der
        } else {
            n = n.izq
        }
    }

    if candidato == nil {
        panic("no tiene predecesor")
    }
    return candidato.clave
}

/*
1. Implementar en Go una primitiva que reciba un árbol binario que representa un heap (árbol binario izquierdista, que
cumple la propiedad de heap), y devuelva la representación en arreglo del heap. La firma de la primitiva debe ser
RepresentacionArreglo() []T. Indicar y justificar la complejidad de la primitiva. La estructura del árbol binario es:

type ab[T any] struct {
izquierda *ab[T]
derecha *ab[T]
dato T
}
*/

func (ab *ab[T]) RepresentacionArreglo() []T{

	if ab == nil{
		return nil
	}

	res := make([]T, 0)
	res = append(res, ab.dato)

	res = append(res, ab.izq.RepresentacionArreglo())
	res = append(res, ab.der.RepresentacionArreglo())
	return res
}

// O(n)

/*1. Implementar en Go una primitiva para un árbol binario izquierdista, que reciba la cantidad de nodos que tiene, y
devuelva el dato del elemento más a abajo y a la derecha del árbol. En los árboles de las figuras mostradas, se debe
devolver en ambos casos 4. Para que el ejercicio se pueda considerar como aprobable, debe resolverse en no más que
O(n), sin contar con otros errores. Para que se considere completamente bien, debe ejecutar en O(log n). Justificar la
complejidad del algoritmo implementado. A fines del ejercicio, considerar que la estructura del árbol binario es:

type ab[T any] struct {
izquierda *ab[T]
derecha *ab[T]
dato T
}*/

// es un bfs porque es x nivel xd
func (ab *ab[T]) MasAbajoDer(nodos int) T{
	if ab == nil{
		return
	}

	actual := ab
	cola := CrearColaEnlazada[T]()
	cola.Encolar(actual)

	for !cola.EstaVacia(){
		actual = cola.Desencolar()
		if actual.izq != nil{
			cola.Encolar(actual.izq)
		}
		if actual.der != nil{
			cola.Encolar(actual.der)
		}
	}

	return actual.dato
}


