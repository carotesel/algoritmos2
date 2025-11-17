
def cant_caminos_min(grafo, origen, fin):
    padres = {origen: None}
    distancias = {}
    cant_caminos = {origen:1}

    for v in grafo:
        distancias[v] = float("inf")
    distancias[origen] = 0

    heap_min = Heap()
    heap_min.Encolar((0, origen))

    while not heap_min.EstaVacia():
        _, v = heap_min.Desencolar()
        for w in grafo.adyacentes(v):
            dist_w = distancias[v] + grafo.peso_arista(v, w)

            if dist_w < distancias[w]:
                distancias[w] = dist_w
                padres[w] = v
                cant_caminos[w] = cant_caminos[v]
                heap_min.Encolar((dist_w, w))
            
            elif dist_w == distancias[w]:
                cant_caminos[w] += cant_caminos[v]
    
    return cant_caminos[fin]



