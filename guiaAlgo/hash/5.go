package masmitad

func MasDeLaMitad(arr []int) bool {
    dicc := CrearHash[int, int]()
    mitad := len(arr) / 2

    for _, n := range arr{
        if dicc.Pertenece(n){
            cantidad := dicc.Obtener(n)
            dicc.Guardar(n, cantidad+1)
        } else{
            dicc.Guardar(n, 1)
        }
    }

    it := dicc.Iterador()

    for it.HaySiguiente(){
        _, v := it.VerActual()
        if v > mitad{
            return true
        }
        it.Siguiente()
    }
    return false
}