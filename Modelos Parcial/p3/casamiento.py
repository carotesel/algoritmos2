
# n = invitados
# se deben sentar con gente que se lleve bien entre ellos.

"""a. Podemos plantear un grafo no dirigido no pesado tal que vertices = personas, aristas = relacion entre ellas 
(positivas) dado que tenemos informacion acerca de quien se lleva bien con quien. Para la resolucion de este problema, 
nos es conveniente plantear el complemento de este grafo, es decir, las personas que no se llevan bien entre si, y a 
partir de el, intentar dividirlo en 2 grupos para que nadie que se lleve mal entre si se siente junto."""

def bipartito(grafo): #O(V+E)
    colores = {}
    for v in grafo:
        if v not in colores:
            if not _bipartito(grafo, v, colores):
                return False, colores
    return True, colores

def _bipartito(grafo, origen, colores): 
    colores[origen] = 0

    cola = Cola()
    cola.Encolar(origen)

    while not cola.Estavacia():
        actual = cola.Desencolar()

        for w in grafo.adyacentes(actual):
            if w not in colores:
                colores[w] = 1 - colores[actual]
            if colores[w] == colores[actual]:
                return False
    
    return True

def devolver_grupos(): 

    es_bip, colores = bipartito(grafo) #O(V+E)

    if not es_bip:
        return []
    
    res1 = []
    res2 = []

    for v in colores.items(): #O(V)
        if colores[v] == 0:
            res1.append(v)
        elif colores[v] == 1:
            res2.append(v)
    return res1, res2


#Complejidad: O(V+E) + O(V) = O(V+E) 