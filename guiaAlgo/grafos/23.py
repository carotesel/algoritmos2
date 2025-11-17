from collections import deque

def diametro(grafo):
    
    max_dist_global = 0

    for v in grafo.obtener_vertices():
        distancias = bfs(v, grafo)
        max_dist = max(distancias.values())

        if max_dist > max_dist_global:
            max_dist_global = max_dist
    
    return max_dist_global

def bfs(v, grafo):

    cola = deque()
    visitados = set()
    distancias = {}

    cola.append(v)
    distancias[v] = 0
    visitados.add(v)
    
    while cola:
        ver = cola.popleft()

        for w in grafo.adyacentes(ver):

            if w not in visitados:
                distancias[w] = distancias[ver] + 1
                visitados.add(w)
                cola.append(w)
    
    return distancias
