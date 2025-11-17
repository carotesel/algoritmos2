from collections import deque

def ult_pudrirse(grafo, origen):

    cola = deque()
    cola.append(origen)
    dist = {origen: 0}
    
    visitados = set()
    visitados.add(origen)

    while cola:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                dist[w] += dist[v] + 1
                cola.append(w)
    
    max_dist = max(dist.values())

    return [v for v in dist if dist[v] == max_dist]
