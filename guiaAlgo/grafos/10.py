from collections import deque

def es_bipartito(grafo):
    
    colores = {}

    for v_inicial in grafo.obtener_vertices():
        if v_inicial in colores:
            continue  # ya visitado en otra componente

        q = deque()

        colores[v_inicial] = 0
        q.append(v_inicial)

        while len(q) > 0:
            v = q.popleft()
            for w in grafo.adyacentes(v):
                if w in colores:
                    if colores[v] == colores[w]:
                        return False
                else:
                    colores[w] = 1 - colores[v] # opuesto
                    q.append(w) # encolo para procesar
    return True
               
# usa recorrido bfs