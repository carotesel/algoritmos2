from collections import deque
import random

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
    return len(visitados) == len(grafo)

#En el peor caso, el algoritmo recorre todos los vértices del grafo (O(V))
#y, para cada vértice, inspecciona todas las posibles conexiones (adyacentes) en la matriz,
#lo que cuesta O(V) por vértice.
#Por lo tanto, la complejidad total del recorrido BFS sobre una matriz de adyacencia es:

#T(V) = O(V × V) = O(V²)
    
    
