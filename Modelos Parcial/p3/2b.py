from collections import deque

def grados_entrada(grafo):

    entrada = {}

    for v in grafo:
        entrada[v] = 0

    for v in grafo:
        for w in grafo.adyacentes(v):
            entrada[w] += 1
    
    return entrada

def orden_topologico(grafo):

    g_entrada = grados_entrada(grafo)
    res = []
    cola = deque()

    for v in grafo.obtener_vertices():
        if g_entrada[v] == 0:
            cola.append(v)
    
    while cola:
        v = cola.popleft()
        res.append(v)

        for w in grafo.adyacentes(v):
            g_entrada[w] -= 1
            if g_entrada[w] == 0:
                cola.append(w)
    return res




