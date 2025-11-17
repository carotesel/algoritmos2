from grafo import Grafo
from collections import deque

# El grafo recibido tiene si o si pesos 1 o 2
def camino_minimo(grafo, origen):
   nuevo = convertir_no_pesado(grafo)
   padres = bfs(nuevo, origen)
   padres_limpios = limpiar(padres, grafo)
   return padres_limpios

def convertir_no_pesado(grafo):
    g_nuevo = Grafo(es_dirigido=False, vertices_init=[])
    procesadas = set()

    for v in grafo.obtener_vertices():
        g_nuevo.agregar_vertice(v)

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            # Evitar procesar dos veces la misma arista (A-B y B-A)
            if (w, v) in procesadas or (v, w) in procesadas:
                continue

            peso = grafo.peso_arista(v, w)

            if peso == 1:
                g_nuevo.agregar_arista(v, w, 1)
            else:  # peso == 2
                 aux = f"aux_{min(v, w)}_{max(v, w)}"
                 if aux not in g_nuevo.obtener_vertices():
                    g_nuevo.agregar_vertice(aux)
                    g_nuevo.agregar_arista(v, aux, 1)
                    g_nuevo.agregar_arista(aux, w, 1)

            # marcar como procesada
            procesadas.add((v, w))

    return g_nuevo

    
def bfs(nuevo, origen):

        cola = deque()
        padres = {}
        visitados = set()

        for v in nuevo.obtener_vertices():
            padres[v] = None
        
        cola.append(origen)
        visitados.add(origen)

        while cola:

            v = cola.popleft()

            for w in nuevo.adyacentes(v):

                if w not in visitados:

                    padres[w] = v
                    cola.append(w)
                    visitados.add(w)
        
        return padres


def limpiar(padres, grafo):

        dicc = {}

        for v in padres:

            if v in grafo.obtener_vertices():

                padre = padres[v]

                # Subir hasta encontrar un vértice original
                while padre is not None and padre not in grafo.obtener_vertices():
                    padre = padres[padre]

                dicc[v] = padre
        
        return dicc
