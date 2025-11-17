
import random

def mst_prim(grafo):
    arbol = Grafo(es_dirigido = False, vertices = grafo.obtener_vertices())
    visitados = set()
    origen = random.choice(grafo.obtener_vertices())
    heap_min = Heap()
    visitados.add(origen)

    for w in grafo.adyacentes(origen):
        heap_min.Encolar((origen, w, grafo.peso(v, w)))
    
    while not heap_min.EstaVacia():
        v, w, peso = heap_min.Desencolar()

        if w in visitados:
            continue

        visitados.add(w)
        arbol.agregar_arista((v, w, peso))

        for x in grafo.adyacentes(w):
            if x not in visitados:
                heap_min.Encolar((w, x, grafo.peso(w, x)))
    
    return arbol



    

   
