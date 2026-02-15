"""Implementar un algoritmo que determine si un grafo no dirigido tiene ciclos.
Complejidad lineal."""

"""Implementar un algoritmo que determine si un grafo no dirigido tiene ciclos.
Complejidad lineal."""

def hay_ciclo_ND(grafo):
    padres = {}
    visitados = set()

    for origen in grafo:
        if origen in visitados:
            continue
        visitados.add(origen)
        padres[origen] = None
        cola = Cola()
        cola.Encolar(origen)

        while cola:
            v = cola.Desencolar()

            for w in grafo.adyacentes(v):

                if w not in visitados:
                    visitados.add(w)
                    padres[w] = v
                    cola.Encolar(w)
                
                if padres[v] != w:
                    return True
    return False

"""2
Implementar un algoritmo que determine si un grafo dirigido tiene ciclos.
Complejidad lineal."""

def hay_ciclo_dir(grafo):
    visitados = set()

    for v in grafo:
        if v in visitados:
            continue
        cola = Cola()
        cola.Encolar(v)
        visitados.add(v)

        for w in grafo.adyacentes(v):

            if w not in visitados:
                visitados.add(w)
                cola.Encolar(w)
            else:
                return True
    return False
