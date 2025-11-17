from collections import deque

def grados_entrada(grafo):
    grados = {}

    for v in grafo.obtener_vertices():
        grados[v]=0
    
    for v in grafo.obtener_vertices():

        for w in grafo.adyacentes(v):
            grados[w] += 1
    
    return grados

def orden_topologico(grafo):

    grados_ent = grados_entrada(grafo)
    cola = deque()
    res = []

    for v in grafo.obtener_vertices():
        if grados_ent[v] == 0:
            cola.append(v)
    
    while cola:

        v = cola.popleft()
        res.append(v)

        for w in grafo.adyacentes(v):
            grados_ent[w]-= 1
            if grados_ent[w] == 0:
                cola.append(w)
    return res




