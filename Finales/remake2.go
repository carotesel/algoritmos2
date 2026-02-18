/*
3. Implementar un algoritmo que, dado un grafo no dirigido, conexo, y sin puentes (es decir, sin ninguna arista que al
quitarla formaría más de una componente conexa), determine una dirección para cada arista, para que el grafo dirigido
resultante sea fuertemente conexo (es decir, haya una única componente fuertemente conexa). Indicar y justificar la
complejidad del algoritmo.
*/

def determinar_direcc(grafo):
    visitados = set()
    origen = grafo.vertice_aleatorio()
    padres = {origen: None}
    aristas_visitadas = set()

    g_dir = Grafo(es_dirigido=True, vertices=grafo.obtener_vertices())

    _dfs(grafo, origen, visitados, padres, g_dir, aristas_visitadas)
    return g_dir


def _dfs(grafo, v, visitados, padres, g_dir, aristas_visitadas):
    visitados.add(v)

    for w in grafo.adyacentes(v):

        # no procesar dos veces la misma arista
        if (v, w) in aristas_visitadas or (w, v) in aristas_visitadas:
            continue

        aristas_visitadas.add(arista)

        if w not in visitados:
            # arista de árbol
            padres[w] = v
            g_dir.agregarArista(v, w)
            _dfs(grafo, w, visitados, padres, g_dir, aristas_visitadas)
        elif w != padres[v]:
            # arista de retorno (cierra ciclo)
            g_dir.agregarArista(v, w)



/*
4. Tenemos un arreglo de n elementos en el que cada elemento se encuentra a lo sumo k posiciones de la que le correspondería
si el arreglo estuviera ordenado (2 ≤ k ≤ n). Implementar un algoritmo de ordenamiento que funcione en O(n log k).
*/

func cmp(a, b int) int{
	return a - b
}

func k_pos(arr []int, k int) []int{
	n := len(arr)
	res := make([]int, n, 0)
	if k >= n{
		k = n -1
	}
	heap := CrearHeapArr[int](cmp, arr[:k+1])
	actual := k+1

	for actual < n{
		res = append(res, heap.Desencolar())
		heap.Encolar(arr[actual])
		actual++
	}

	// Vaciar el heap
	for !heap.EstaVacia() {
		res = append(res, heap.Desencolar())
	}

	return res
}

/*1. Implementar un algoritmo que reciba un arreglo de n enteros (con n ≥ 3) en el que todos sus elementos son iguales
salvo 1, y determine (utilizando división y conquista) cual es dicho elemento no repetido. Indicar y justificar la
complejidad del algoritmo implementado.*/

func distintaa(arr []int) int{
    return laDis(arr, 0, len(arr)-1, arr[0])
}

func laDis(arr []int, ini, fin, repet int) int{
	if ini == fin{
		return arr[ini]
	}

	medio := (ini + fin)/2

	if arr[medio] != repet{
		return arr[medio]
	}
	if arr[medio-1] != repet{
		return arr[medio-1]
	}
	if arr[medio+1] != repet{
		return arr[medio+1]
	}

	cantIzq := medio - ini
	cantDer := fin - medio

	if cantIzq % 2 == 1{ // impar = algo raro
		return laDis(arr, ini, medio-1, repet)
	} else{
		return laDis(arr, medio+1, fin, repet)
	}
}

// O(n) si o si porque no hay forma de asegurar que este a la izq o a la der sin recorrerlos.

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

/*
1. Implementar:

    ClavesEnRangoHastaNivel(abb, ini, fin, M) → Lista
    
    que devuelva todas las claves del ABB que están en el rango `[ini, fin]` y 
	no superen el nivel M de profundidad (raíz está en nivel 0).
    
    La lista debe quedar ordenada.
    
    Indicar complejidad según h y M.
*/

func (abb *abb[K, V]) ClavesRangoNivel(M int, ini, fin K) Lista[V]{
	lista:= CrearListaEnlazada[V]()
	clavesRango(abb.raiz, abb.cmp, ini, fin, nivel, 0, lista)
	return lista
}

func clavesRango(nodo *nodoABB[K, V], cmp func(K, K) int, ini, fin K, nivel, nActual int, lista ListaEnlazada[V]){
	if nodo == nil{
		return 
	}

	if nActual > nivel{
		return
	}

	// si nodo < ini -> der (no arranca aun)
	if cmp(nodo.clave, iniciativa) < 0{
		return clavesRango(nodo.der, cmp, ini, fin, nivel, nActual + 1, lista)
	}

	// si nodo > fin -> izq (me pase)
	if cmp(nodo.clave, fin) > 0{
		return clavesRango(nodo.izq, cmp, ini, fin, nivel, nActual + 1, lista)
	}

	lista.AgregarUltimo(nodo.clave)

	clavesRango(nodo.der, cmp, ini, fin, nivel, nActual + 1, lista)
	clavesRango(nodo.izq, cmp, ini, fin, nivel, nActual + 1, lista)
}

/*
1. Se tiene una cadena que contiene () y ningún otro caracter (considerar que un único caracter es de tipo rune). Un
ejercicio típico es dada una cadena averiguar si está balanceada (es decir, todos los símbolos de apertura se cierran, y
además respetan el orden en el que se abrieron. Ejemplos balanceados: "()()()", o "(())()". No balanceados: "(()",
o ")(".
Teniendo en cuenta esto, se tiene una cadena que se asegura que en caso de borrar algunos paréntesis la cadena será
balanceada, se pide implementar una función func cantBorradosBalanceada(cadena string) int que dada una
cadena de este tipo, devuelva la cantidad mínima de paréntesis que se deben borrar para que la cadena esté balanceada.
Indicar y justificar la complejidad del algoritmo.
Ejemplos:
cadena: '()' -> 0
cadena: ')(' -> 2
cadena: '(()' -> 1
cadena: ')(()' -> 2
*/

func cantBorradosBalanceada(cadena string) int{
	p := CrearPilaDinamica[string]()
	cantBorrar := 0

	for _, c := range cadena{
		if c == '('{
			p.Apilar(c)
		} else{
			if p.EstaVacia(){
				cantBorrar++
			} else{
				p.Desapilar()
			}
		}
	}

	for !p.EstaVacia(){
		cantBorrar++
		p.Desapilar()
	}

	return cantBorrar
}

// complejidad: o(n)

/*
// 1.er parcialito – 26/09/2022

// 1. Implementar una funcion `balancedado(texto string) bool`
//    que retorne si `texto` esta balanceado o no.
//    `texto` solo puede contener los siguientes caracteres:
//    { }, [ ], ( ), < >.
//
//    Indicar y justificar la complejidad de la funcion implementada.
//
//    Un texto esta balanceado si cada agrupador abre y cierra en un
//    orden correcto. Por ejemplo:
//
//    balancedado("{{{[()]}}}") => true
//    balancedado("{[}")        => false
//    balancedado("{()}]")      => false
//    balancedado("{[()()]}")   => true
//    balancedado("(){}([])")   => true
*/

func balancedado(texto string) bool{
	p := CrearPilaDinamica[string]()

	for _, c := range cadena{
		if c == '('|| c == '[' || c == '{' || c == '<'{
			p.Apilar(c)
		} else{

			if p.EstaVacia() {
                return false  // Cierre sin apertura
            }

			elem := p.Desapilar()

			if c == ')' && elem != '('{
				return false
			} else if c == '}' && elem != '{'{
				return false
			} else if c == ']' && elem != '['{
				return false
			} else if c == '>' && elem != '<'{
				return false
			}
		}
	}
	return p.EstaVacia() // la pila debe quedar vacia!
}

/*1. Implementar una función suma_total(arreglo []float) float que, por división y conquista, devuelva la suma
de todos los elementos. Indicar y justificar adecuadamente la complejidad de la función implementada.*/

func suma_total(arreglo []float) float{
	return suma(arreglo, 0, len(arr)-1)
}

func suma(arr []float, ini, fin int) float{
	if ini > fin{
		return 0 // arr vacio
	}

	if ini == fin{
		return arr[ini] // arr con 1 elemento
	}

	medio := (ini + fin)/2

	izq := suma(arr, ini, medio)
	der := suma(arr, medio+1, fin)

	return izq + der
}

// O(n)



/*2. Implementar una función que reciba un arreglo A de n enteros y un número k y devuelva un nuevo arreglo en el que
para cada posición i de dicho arreglo, contenga el resultado de la multiplicación de los primeros k máximos del arreglo A
entre las posición [0;i] (incluyendo a i). Las primeras k − 1 posiciones del arreglo a devolver deben tener como valor -1.
Por ejemplo, para el arreglo [1, 5, 3, 4, 2, 8] y k = 3, el resultado debe ser [-1, -1, 15, 60, 60, 160]. Indicar
y justificar la complejidad del algoritmo implementado.
*/

func multiplicar(a []int, k int) []int{
	n := len(a)
	res := make([]int, n)
	heap_min := CrearHeap[int](cmp)
	acumulado := 1

	for i:=0; i<n; i++{

		x := a[i]

		if heap_min.Largo() < k{ // menos de k elementos -> encolo y multiplico de una
			heap_min.Encolar(x)
			acumulado = acumulado * x
		
		} else if x > heap_min.VerMin(){ // mas de k elementos y nuevo elem es mayor que el minimo -> desencolo y divido asi lo desacumulo pero desp encolo el nuevo porque es medio un heap de max en realidad
			sacado := heap_min.Desencolar()
			acumulado = acumulado / sacado
			heap_min.Encolar(x)
			acumulado = acumulado * x
		}

		if i < k{
			res[i] = -1
		} else{
			res[i] = acumulado
		}
	}
	return res
}