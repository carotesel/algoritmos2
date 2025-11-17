package altura

type ab struct {
    izq *ab
    der *ab
    dato int
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func (arbol *ab) Altura() int {
    if arbol == nil{
        return 0
    }

    izq := arbol.izq.Altura()
    der := arbol.der.Altura()

    return 1 + max(izq, der)
}

// Orden:
// T(n) = 2 T(n/2) + O(1)
// Log 2 (2) = 1 < C
// Compl: O(n^c) = O(n)
