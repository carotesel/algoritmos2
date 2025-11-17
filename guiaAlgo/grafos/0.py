def obtener_aristas_dirigido(grafo):
    # devolver una lista de las aristas del grafo. no es lo mismo a-> b qaue b -> a
    aristas = []

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            peso = grafo.peso_arista(v, w)
            aristas.append((v, w, peso))

    return aristas

def obtener_aristas_no_dirigido(grafo):
    # devolver una lista de las aristas del grafo. a -> b == b -> a, duplicados! ver5ificar que w,v no este en visitados
    aristas = []
    visitados = set()
    
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if (w, v) not in visitados:
                peso = grafo.peso_arista(v, w)
                aristas.append((v, w, peso))
                visitados.add((v, w))
    
    return aristas