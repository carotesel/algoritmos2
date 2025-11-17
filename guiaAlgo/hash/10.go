package cantvalores

func (dicc *hashCerrado[K, V]) CantidadValoresDistintos() int {
    dic_valores := CrearHash[V, int]()

	for _, celda := range dicc.tabla{
		if celda.estado == OCUPADO{
			valorDicc := dicc.Obtener(celda.clave)

			if dic_valores.Pertenece(valorDicc){
				cant := dic_valores.Obtener(valorDicc)
				dic_valores.Guardar(valorDicc, cant+1)
			} else{
				dic_valores.Guardar(valorDicc, 1)
			}
		}
	}

	return dic_valores.Cantidad()	
}