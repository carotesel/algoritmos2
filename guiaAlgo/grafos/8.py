#hago bfs de distancias
from collections import deque

def seis_grados(grafo):

    vertices = grafo.obtener_vertices()

    for v in vertices:
        if not bfs_dist(v, vertices, grafo):
            return False
    
    return True
    

def bfs_dist(v_ini, vertices, grafo): # calculo las distancias de 1 vertice inicial a todos

    distancias = {v_ini:0}
    visitados = set()
    visitados.add(v_ini)
    cola = deque()
    cola.append(v_ini)

    while len(cola) > 0:
        v = cola.popleft()

        if distancias[v] >= 6:
            continue

        for w in grafo.adyacentes(v):
            if w not in visitados:
                distancias[w] = distancias[v] + 1

                if distancias[w] > 6:
                    return False

                visitados.add(w)
                cola.append(w)

    if len(distancias) < len(vertices):
        return False
    else:
        return True


    
#Si el grafo mantiene fija la cantidad de vértices pero las aristas pueden variar, una matriz de adyacencia es conveniente porque permite agregar o eliminar relaciones en tiempo constante y no requiere reestructurar su tamaño.
#Además, verificar si dos personas están conectadas es inmediato (O(1)).

#Sin embargo, la matriz ocupa siempre O(V²) espacio, por lo que resulta ineficiente si el grafo es disperso. Además, para recorrer los vecinos de una persona es necesario revisar toda la fila, lo que cuesta O(V).

#En resumen, la matriz es una buena opción si el grafo es denso o de tamaño moderado, pero poco eficiente en espacio para grafos grandes con pocas conexiones.