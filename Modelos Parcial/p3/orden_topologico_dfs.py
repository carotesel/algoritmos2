
def orden_top_dfs(grafo):
    visitados = set()
    pila = Pila()

    for v in grafo:
        if v not in visitados:
            visitados.add(v)
            _dfs(grafo, v, visitados, pila)
    return pila_a_lista(pila)

def _dfs(grafo, origen, visitados, pila):

    for w in grafo.adyacentes(origen):
        if w not in visitados:
            visitados.add(w)
            _dfs(grafo, w, visitados, pila)
    pila.Apilar(origen)

def pila_a_lista(pila):
    res = []

    while not pila.EstaVacia():
        actual = pila.Desapilar()
        res.append(actual)
    
    return res[::-1]