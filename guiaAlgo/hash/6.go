package iguales

func SonIguales[K comparable, V comparable](d1, d2 Diccionario[K, V]) bool {
    if d1.Cantidad() != d2.Cantidad(){
        return false
    }

    // Si ambos vacíos, son iguales
    if d1.Cantidad() == 0 {
        return true
    }

    it_d1 := d1.Iterador()

    for it_d1.HaySiguiente(){
        c1, v1 := it_d1.VerActual()
        if !d2.Pertenece(c1) {
            return false
        } 
        if d2.Obtener(c1) != v1 {
            return false
        }
        it_d1.Siguiente()
    }

    return true
}