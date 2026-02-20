"""Implementar un algoritmo que determine si un grafo no dirigido tiene ciclos.
Complejidad lineal."""

def hay_ciclo_ND(grafo):
    padres = {}
    visitados = set()

    for origen in grafo:
        if origen in visitados:
            continue
        visitados.add(origen)
        padres[origen] = None
        cola = Cola()
        cola.Encolar(origen)

        while cola:
            v = cola.Desencolar()

            for w in grafo.adyacentes(v):

                if w not in visitados:
                    visitados.add(w)
                    padres[w] = v
                    cola.Encolar(w)
                
                if padres[v] != w:
                    return True
    return False

"""2
Implementar un algoritmo que determine si un grafo dirigido tiene ciclos.
Complejidad lineal."""



#################### hay ciclo NO DIRIGIDO?

def hay_ciclo(grafo):

    padres = {}
    visitados = set()

    for v in grafo:
        if v not in visitados:
           ciclo = dfs_ciclo(grafo, v, padres, visitados)
           if ciclo is not None:
               return ciclo
    return None

def dfs_ciclo(grafo, v, padres, visitados):

    visitados.add(v)

    for w in grafo.adyacentes(v):
        if w in visitados:
            if w != padres[v]:
                return reconstruir_ciclo(padres, v, w)
        else:
            padres[w] = v
            ciclo = dfs_ciclo(grafo, w, padres, visitados)
            if ciclo is not None:
                return ciclo
         
    return None

def reconstruir_ciclo(padres, v, w):

    ciclo = []

    while w != None:
        ciclo.append(w)
        w = padres[w]
    ciclo.reverse()
    return ciclo



# grafo no dirigido es arbol?

def dfs_ciclo(grafo, v, padres, visitados):

    visitados.add(v)

    for w in grafo.adyacentes(v):

        if w in visitados:
            if w != padres[v]:
                return reconstruir_ciclo(v, w, padres)
        else:
            padres[w] = v
            ciclo = dfs_ciclo(grafo, w, padres, visitados)
            if ciclo is not None:
                return ciclo
    return None

def reconstruir_ciclo(v, w, padres):
    ciclo = []

    while w != None:
        ciclo.append(w)
        w = padres[w]
    ciclo.reverse()
    return ciclo

# E = V - 1
# No ciclo - 
# Es conexo

def es_arbol(grafo):

    vertices = list(grafo.obtener_vertices())
    
    if len(vertices) == 0:
        return True

    visitados = set()
    origen = vertices[0]
    padres = {origen: None}
    
    # si tiene ciclo
    if dfs_ciclo(grafo, origen, padres, visitados) is not None:
        return False
    
    if len(visitados) != len(vertices):
        return False
    
    return True
    




