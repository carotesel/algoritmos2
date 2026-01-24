
def dijkstra(grafo, origen, fin):
    padres = {origen:None}
    distancias = {}

    for v in grafo:
        distancias[v] = float("inf")
    
    distancias[origen] = 0
    heap_min = Heap()
    heap_min.Encolar((0, origen))

    while not heap_min.EstaVacia():
        dist_v, v = heap_min.Desenacolar()

        if dist_v > distancias[v]:
            continue

        if v == fin:
            return padres

        for w in grafo.adyacentes(v):
            dist_w = distancias[v] + grafo.peso(v, w)

            if dist_w < distancias[w]:
                padres[w] = v
                distancias[w] = dist_w
                heap_min.Encolar((dist_w, w))
    return padres

def reconstruir_camino(padres, fin):
    camino = []

    while fin is not None:
        camino.append(fin)
        fin = padres[fin]
    
    return camino[::-1]



    