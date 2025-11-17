
def centralidad(grafo):
    central = {}
    for v in grafo:
        central[v] = 0
    
    for v in grafo:
        for w in grafo:
            padres = caminos_min(grafo, v, w)
            if padres[w] == None:
                continue

            actual = padres[w]
            
            while actual != v:
                central[actual]+= 1
                actual = padres[actual]
    
    return central

def caminos_min(grafo, origen, destino):

    padres = {origen: None}
    distancias = {}

    for v in grafo:
        distancias[v] = float("inf")
    
    distancias[origen] = 0

    heap_min = Heap()
    heap_min.Encolar((0, origen))

    while not heap_min.EstaVacia():
        v = heap_min.Desencolar()

        # va o no va?
        if v == destino:
            return padres
        
        for w in grafo.adyacentes(v):
            dist_w = distancias[v] + grafo.peso_arista(v, w)

            if dist_w < distancias[w]:
                distancias[w] = dist_w
                padres[w] = v
                heap_min.Encolar((dist_w, w))
    
    return padres


