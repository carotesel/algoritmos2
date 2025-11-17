from collections import deque
from grafo import Grafo

"""Hacés el grafo no dirigido equivalente (como si todas las aristas fueran bidireccionales),
entonces cada conjunto de vértices conectados entre sí forma una componente débilmente conexa."""

def cantidad_componentes_debiles(grafo):
    
    # creo el grafo no dirigido equivalente
    
    vertices = grafo.obtener_vertices()
    g_no_dir = Grafo(es_dirigido = False, vertices_init = vertices)

    for v in vertices:
        for w in grafo.adyacentes(v):
            # evitar duplicados 
            if not g_no_dir.estan_unidos(v, w):
                g_no_dir.agregar_arista(v, w)
    
    visitados = set()
    componentes = 0

    # recorro el no dirigido y busco las componentes. 

    for v in g_no_dir.obtener_vertices():
        if v not in visitados:
            _bfs(g_no_dir, v, visitados)
            componentes += 1
    
    return componentes


def _bfs(g_no_dir, v, visitados):

    cola = deque()
    cola.append(v)
    visitados.add(v)

    while cola:

        actual = cola.popleft()

        for w in g_no_dir.adyacentes(actual):

            if w not in visitados:

                visitados.add(w)
                cola.append(w)





