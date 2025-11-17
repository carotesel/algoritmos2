
# no dirigido
def obtener_ciclo(grafo):
    visitados = set()
    padres = {}

    for v in grafo:
        if v not in visitados:
            visitados.add(v)
            ciclo = _dfs(grafo, v, visitados, padres)

            if ciclo is not None:
                return ciclo
    return None

def _dfs(grafo, origen, visitados, padres):

    visitados.add(origen)
    padres[origen] = None

    for w in grafo.adyacentes(origen):
        if w not in visitados:
            padres[w] = origen
            camino = _dfs(grafo, w, visitados, padres)
            if camino is not None:
                return camino
        else:
            if w != padres[origen]:
                return reconstruir_ciclo(padres, origen, w)
    return None

def reconstruir_ciclo(padres, origen, destino):

    res = []
    actual = destino

    while actual != origen:
        res.append(actual)
        actual = padres[actual]
    res.append(origen)

    return res[::-1]

# dirigido
def encontrar_ciclo(g):
    visitados = set()
    esta_recursion = set()
    padre = {}
  
    for v in g.obtener_vertices():
        if v not in visitados:
            ciclo = dfs_ciclo(g, v, visitados, esta_recursion, padre)
            if ciclo is not None:
                return ciclo
    return None


def dfs_ciclo(g, v, visitados, esta_recursion, padre):
    visitados.add(v)
    esta_recursion.add(v)

    for w in g.adyacentes(v):
        if w not in visitados:
            padre[w] = v
            ciclo = dfs_ciclo(g, w, visitados, esta_recursion, padre)
            if ciclo is not None:
                return ciclo
        elif w in esta_recursion:
            return reconstruir_ciclo(padre, w, v)
    
    return None