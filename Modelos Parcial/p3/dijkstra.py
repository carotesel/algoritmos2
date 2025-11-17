
def dijkstra(grafo, origen, fin):
    padres = {origen:None}
    distancias = {}

    for v in grafo:
        grafo[v] = float("inf")
    
    distancias[origen] = 0
    heap_min = Heap()
    heap_min.Encolar((0, origen))

    while not heap_min.EstaVacia():
        _, v = heap_min.Desenacolar()

        if v == fin:
            return padres

        for w in grafo.adyacentes(v):
            dist_w = distancias[v] + grafo.peso(v, w)

            if dist_w < distancias[w]:
                padres[w] = v
                distancias[w] = dist_w
                heap_min.Encolar((dist_w, w))
    return padres



    