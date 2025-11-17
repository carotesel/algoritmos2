"""Implementar la fc grafo_potencia(grafo, k) que reciba un grafo no dirigido y un entero k, y nos devuelva un nuevo grafo,
que tenga los mismos vertices que el grafo recibido, y cuyas aristas (v, w) indican que v esta a lo sumo k aristas de w 
en el grafo original.
"""

def grafo_potencia(grafo, k):

    nuevo_g = Grafo(es_dirigido=False, vertices=grafo.obtener_vertices())

    for v in grafo:
        dist = bfs_dist(grafo, v, k)
    
        for w, d in dist.items():
            if d <= k and w != v:
                nuevo_g.agregar_arista(v, w)

    return nuevo_g

def bfs_dist(grafo, origen, k):
    dist = {origen:0}
    cola = cola()
    cola.Encolar(origen)

    while cola:
        v = cola.Desencolar()

        if dist[v] == k: # no expando en adyacentes pero sigo revisando los otros vertices!!
            continue

        for w in grafo.adyacentes(v):
            if w not in dist:
                dist[w] = dist[v] + 1
                cola.Encolar(w)
    
    return dist






