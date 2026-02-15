"""5. Implementar un algoritmo que reciba un grafo no dirigido y determine si para cada par de vértices del grafo existe y es único el
camino entre ellos. El algoritmo implementado debe ser lineal en la cantidad de vértices y aristas. Justificar la complejidad del
algoritmo."""

def es_arbol(grafo):

    visitados = set()
    cola = Cola()

    v_ini = grafo.vertice_aleatorio()
    cola.Encolar((v_ini, None))  # (vertice, padre)

    while not cola.EstaVacia():
        v, padre = cola.Desencolar()

        if v in visitados:
            return False   # HAY CICLO

        visitados.add(v)

        for w in grafo.adyacentes(v):
            if w != padre:
                cola.Encolar((w, v))

    # verificar conexo
    return len(visitados) == grafo.cantidad_vertices()

"""3. Implementar un algoritmo que, dado un grafo no dirigido, conexo, y sin puentes (es decir, sin ninguna arista que al
quitarla formaría más de una componente conexa), determine una dirección para cada arista, para que el grafo dirigido
resultante sea fuertemente conexo (es decir, haya una única componente fuertemente conexa). Indicar y justificar la
complejidad del algoritmo."""

def completarDirig(grafo):
    grafo_D = Grafo(es_dirigido = True, vertices = grafo.obtener_vertices())
    aristas_visitadas = set()
    visitados = set()
    origen = grafo.vertice_random()
    dfs(grafo, grafo_D, origen, aristas_visitadas, visitados)
    return grafo_D

def dfs(grafo, grafo_D, origen, aristas_visitadas, visitados):
    visitados.add(origen)

    for w in grafo.adyacentes(origen):
        if (origen, w) in aristas_visitadas or (w, origen) in aristas_visitadas:
            continue

        aristas_visitadas.add((origen, w))
        aristas_visitadas.add((w, origen))
        grafo_D.agregarArista(origen, w)

        if w not in visitados:
            dfs(grafo, grafo_D, w, aristas_visitadas, visitados)

#O(V + E)
        
"NO DIR CONEXO -> DFS"

