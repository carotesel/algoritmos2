package suma

type ab struct {
    izq  *ab
    der  *ab
    dato int
}

func (arbol *ab) Suma() int {
    if arbol == nil {
        return 0
    }

    sumaIzq := arbol.izq.Suma()
    sumaDer := arbol.der.Suma()

    return arbol.dato + sumaIzq + sumaDer
}

// Complejidad: 2 T(n/2) + O(1) -> O(n)
