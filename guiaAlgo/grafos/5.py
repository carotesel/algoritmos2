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


def reconstruir_ciclo(padre, inicio, fin):
    v = fin
    camino = [v]
  
    while v != inicio:
        if v not in padre:   # si llegamos a un nodo sin padre, cortamos
            break
        v = padre[v]
        camino.append(v)
    
    camino.reverse()

    # 🔧 Si el primer y último son iguales (A,B,A), borrá el último
    if len(camino) > 1 and camino[0] == camino[-1]:
        camino.pop()

    return camino
