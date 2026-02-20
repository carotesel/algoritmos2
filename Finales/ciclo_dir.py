def dfs_ciclo_dir(grafo, v, padres, visitados, esta_rec):

    visitados.add(v)
    esta_rec.add(v)

    for w in grafo.adyacentes(v):

        if w in visitados:
            if w in esta_rec:
                return reconstruir_ciclo(v, w, padres)
        else:
            padres[w] = v
            ciclo = dfs_ciclo_dir(grafo, w, padres, visitados, esta_rec)
            if ciclo is not None:
                return ciclo
        return None
    
def hay_ciclo_dir(grafo):

    visitados = set()
    padres = {}
    esta_rec = set()

    for v in grafo:
        if v not in visitados:
            ciclo = dfs_ciclo_dir(grafo, v, padres, visitados, esta_rec)
            if ciclo is not None:
                return ciclo
        return None
    
def reconstruir_ciclo(v, w, padres):
    x = w
    ciclo = []
    while x != v:
        ciclo.append(x)
        x = padres[x]
    ciclo.append(v)
    return ciclo.reverse()