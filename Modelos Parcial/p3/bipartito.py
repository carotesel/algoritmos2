

def es_bipartito(grafo, origen, colores):
    colores={origen: 0}
    cola = Cola()
    cola.Encolar(origen)

    while not cola.EstaVacia():
        actual = cola.Desencolar()
        for w in grafo.adyacentes(actual):
            if w in colores:
                if colores[w] == colores[actual]:
                    return False
            else:
                colores[w] = 1 - colores[actual]
                cola.Encolar(w)
    
    return True

def bipartito(grafo):
    colores = {}

    for v in grafo:
        if v not in colores:
            if not es_bipartito(grafo, v, colores):
                return False
    return True