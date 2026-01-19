/*4. Implementar un algoritmo que reciba dos cadenas (strings) y determine si son anagramas entre sí. Indicar y justificar
la complejidad del algortmo implementado.*/

func sonAnagramas(s1, s2 string) bool{
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