import "math"
// 1)
// -----------------------------------------------------------------------------
// Implementar para el Hash Cerrado la primitiva Limpieza(), la cual se encarga
// de eliminar todos los borrados, asegurándose de dejar al Hash en un estado
// correcto.
//
// Pista: pensar en las búsquedas de los elementos que efectivamente se
// encuentran en el hash.
//
// Indicar y justificar la complejidad de la primitiva implementada.
// -----------------------------------------------------------------------------

const(
	VACIO estado = iota
	OCUPADO
	BORRADO
)

func[K comparable, V any] (hash *hashCerrado[K, V]) Limpieza(){
	if hash.cantidad == 0{
		hash.borrados = 0
		return
	}

	tabla_vieja := hash.tabla

	hash.tabla = CrearTabla[K,V](hash.tam)

	for _, celda := range tabla_vieja{
		if celda.estado == OCUPADO{
			hash.Guardar(celda.clave, celda.dato)
		}
	}

	hash.borrados = 0
}

// -----------------------------------------------------------------------------
// 3. Implementar en Go una primitiva para el árbol binario:
//
//     func (ab *Arbol[T]) EsCompleto() bool
//
// que determine si el árbol es completo.
//
// Indicar y justificar la complejidad de la primitiva.
//
// A fines del ejercicio, considerar que el árbol está definido como:
//
//     type Arbol[T any] struct {
//         dato T
//         izq  *Arbol[T]
//         der  *Arbol[T]
//     }
//
// -----------------------------------------------------------------------------

func (ab *Arbol[T]) Altura() int{
	if ab == nil{
		return 0
	}

	hIzq := ab.izq.Altura()
	hDer := ab.der.Altura()

	return 1 + max(hIzq, hDer)
}

func (ab *Arbol[T]) CantidadNodos() int{
	if ab == nil{
		return 0
	}

	nodosIzq := ab.izq.CantidadNodos()
	nodosDer := ab.der.CantidadNodos()

	return 1 + nodosIzq + nodosDer
}


func (ab *Arbol[T]) EsCompleto() bool{
	altura := ab.Altura()
	cantNodos := ab.CantidadNodos()
	return math.Pow(2, altura) - 1 == cantNodos
}

// - calculo altura
// - cuento nodos
// cantNodos == 2 ^ altura - 1