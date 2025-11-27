/*1. A lo largo de su trayectoria, la empresa ha tenido varias rotaciones de equipos. 
Próximamente habrá una nueva, y se busca que los nuevos equipos incluyan personas 
que hayan compartido equipos antes, para facilitar la transición.

Se dispone de una base de datos que registra, para N personas, en qué M equipos ha participado cada una. 
La base usa un hash donde la persona es la clave y el valor, una lista de equipos:

{
'Ana': ['Frontend-Platform', 'Growth-Squad'],
'Beto': ['Backend-Services', 'Frontend-Platform', 'Mobile-Core'],
'Carla': ['Mobile-Core'],
}

Realizar una función personasEnComun que reciba el hash, y el nombre de una persona,
y devuelva una lista con todas las personas que hayan trabajado en al menos uno de sus equipos listados. 

Indicar y justificar la complejidad del algoritmo implementado, expresada con las variables N y M del problema. 

Por ejemplo, si se consulta por ‘Beto’, la respuesta incluye a ‘Ana’ y ‘Carla’. Si se
pregunta por ‘Beto’, la respuesta incluye a ‘Beto’.*/

func personasEnComun (dicc hash[string, string], persona string) Lista[string]{
	lista := CrearListaEnlazada[string]()

	equipos_persona := dicc.Obtener(persona) // devuelve la lista de equipos

	hash := CrearDiccionario[string, bool]()

	for _, x := range equipos_persona{
		hash.Guardar(x, true)
	}

	iter := dicc.Iterador()

	for iter.HaySiguiente(){
		nombre, equiposPersona := iter.VerActual()

		// NO OLVIDAR ITERAR ARRAY DE EQUIPOS!!!!! NO ES UN STRING
		for _, eq := range equiposPersona{
			if equipos_persona.Pertenece(eq){
				lista.AgregarUltimo(nombre)
				break
			}
		}
		iter.Siguiente()
	}
	return lista
}