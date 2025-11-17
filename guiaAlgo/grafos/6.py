from collections import deque
import random

  # grafo no dirigido es arbol si
    # * aristas = vertices - 1 
    # * no hay ciclos
    # * es conexo

# contar aristas:

def cant_aristas(g):

    aristas = []
    visitados = set()

    for v in g.obtener_vertices():

        for w in g.adyacentes(v):

            if (w,v) not in visitados:
                aristas.append((v,w))
                visitados.add((v,w))
    return len(aristas)


# hay ciclo?

def tiene_ciclo(g):
    visitados = set()
    padre = {}

    for v in g.obtener_vertices():
        if v not in visitados:
            if dfs_ciclo(g, v, visitados, padre):
                return True
    return False


def dfs_ciclo(g, v, visitados, padre):
    visitados.add(v)
    for w in g.adyacentes(v):
        # Si el vecino ya fue visitado y no es el padre → ciclo
        if w in visitados and padre.get(v) != w:
            return True
        # Si no fue visitado, seguimos
        if w not in visitados:
            padre[w] = v
            if dfs_ciclo(g, w, visitados, padre):
                return True
    return False

# es conexo? 
def es_conexo(grafo):
    
    # recorro el grafo (bfs) y si coincide con el largo de vertices entonces ok

    vertices = grafo.obtener_vertices()

    if len(vertices) == 0:
        return True

    visitados = set()
    cola = deque()
    origen = random.choice(vertices)
    visitados.add(origen)
    cola.append(origen)

    while len(cola) > 0:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                cola.append(w)
                visitados.add(w)
    return len(visitados) == len(vertices)
    
    
def es_arbol(g):
    
    ciclo = tiene_ciclo(g)
    conexo = es_conexo(g)
    aristas = cant_aristas(g)
    cumple_aristas = aristas == len(g.obtener_vertices()) - 1

    return (not ciclo and conexo) or (not ciclo and cumple_aristas) or (conexo and cumple_aristas )