/*Se tiene un árbol binario que representa la fase eliminatoria del mundial. En cada nodo guarda el nombre del país, así
como la cantidad de goles que convirtió en dicha fase (incluyendo la tanda de penales, si fuera necesario). El padre
del nodo debe si o si tener al hijo que ganó (tuvo mayor cantidad de goles). 

Implementar una primitiva para el árbol donde solamente están los nombres de los equipos en las hojas (no en los internos), 
y deje el árbol completado con los ganadores en cada fase. Se puede asumir que el árbol es o bien completo, 
o que al menos todos los nodos internos tienen exactamente 2 hijos. 

La cantidad de goles en la raíz no es relevante. La estructura del árbol es:

type Arbol struct {
pais string
goles int
izq *Arbol
der *Arbol
}

Tomando el ejempo del dorso, si invocamos para el árbol de la izquierda, debe quedar como el de la derecha.*/

// RENOMBRO EL PAIS CON EL MAS GRANDE DE SUS HIJOS XD

func (ab *Arbol) CompletarNodos(){

	if ab == nil{
		return
	}

	if ab.izq == nil || ab.der == nil{
		return
	}

	ab.izq.CompletarNodos()
	ab.der.CompletarNodos()

	if (ab.izq != nil && ab.der != nil) && ab.izq.goles > ab.der.goles{
		ab.pais = ab.izq.pais
	} else {
		ab.pais = ab.der.pais
	}
}
