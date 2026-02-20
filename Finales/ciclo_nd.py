def dfs_ciclo(grafo, v, padres, visitados):

    visitados.add(v)

    for w in grafo.adyacentes(v):

        if w in visitados:
            if w != padres[v]:
                return reconstruir(v, w, padres)
        else:
            padres[w] = v
            ciclo = dfs_ciclo(grafo, w, padres, visitados)
            if ciclo is not None:
                return ciclo
        return None

def reconstruir(v, w, padres):
    x = w
    ciclo = []
    while x != v:
        ciclo.append(x)
        x = padres[x]
    ciclo.append(v)
    return ciclo.reverse()

def es_arbol(grafo):

    vert = grafo.obtener_vertices()

    if len(vert) == 0:
        return True
    
    padres = {}
    visitados = set()
    origen = vert[0]
    padres[origen] = None

    ciclo = dfs_ciclo(grafo, origen, padres, visitados)

    if ciclo is not None:
        return False
    
    if len(visitados) != len(vert):
        return False
    
    return True