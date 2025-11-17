
def centralidad(grafo):
    central = {}
    for v in grafo:
        central[v] = 0
    
    for v in grafo:
        for w in grafo:
            padres = caminos_minimos(v, w, grafo)
            if padres[w] is None:
                continue
            actual = padres[w]
            while actual != v:
                central[actual]+= 1
                actual = padres[actual]
    return central

def caminos_minimos(origen, fin, grafo):

    padres = {origen:None}
    distancias = {}
    heap_min = Heap()
    heap_min.Encolar((0, origen))

    for v in grafo:
        distancias[v] = float("inf")
    distancias[origen] = 0

    while not heap_min.EstaVacia():
        _, v = heap_min.Desencolar()
        if v == fin:
            return padres
        for w in grafo.adyacentes(v):
            dist_w = distancias[v] + grafo.peso_arista(v, w)
            if dist_w < distancias[w]:
                distancias[w] = dist_w
                padres[w] = v
                heap_min.Encolar((dist_w, w))
    return padres

