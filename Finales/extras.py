"""Dado un grafo dirigido y aciclico, determinar si admite un unico orden topologico. 
Justificar la complejidad."""

def ordenUnico(grafo):

    grados_entrada = grados_entrada(grafo)

    procesados = 0

    cola = cola()

    for v in grafo:
        if grados_entrada[v] == 0:
            cola.Encolar(v)
        
    while cola:
        if len(cola) > 1:
            return False
        x = cola.Desencolar()
        procesados += 1
        
        for w in grafo.adyacentes(x):
            grados_entrada[w] -= 1
            if grados_entrada[w] == 0:
                cola.Encolar(w)
        
    return procesados == len(grafo.obtener_vertices())


# Componentes conexas de grafo no dirigido (devolver cada componente no solo la cantidad)

def componentes_conexas(grafo):
    componentes = []
    visitados = set()

    for v in grafo:
        if v not in visitados:
            componente = []
            bfs(grafo, v, visitados, componente)
            componentes.append(componente)
    return componentes # se puede devolver len tmb

def bfs(grafo, v, visitados, componente):
    visitados.add(v)
    cola = Cola()
    cola.Encolar(v)

    while cola:
        x = cola.Desencolar()
        componente.append(x)

        for w in grafo.adyacentes(x):
            if w not in visitados:
                visitados.add(w)
                cola.encolar(w)

# Dado un grafo pesado con pesos positivos y dos vertices v y w, determinar la cantidad de 
# caminos minimos entre v y w. Justificar la complejidad.

def cantidad_caminos_min(grafo, origen, fin):
    cantidad_caminos = {}
    padres = {}
    padres[origen] = None
    distancias = {}

    for v in grafo:
        distancias[v] = float("inf")
        cantidad_caminos[v] = 0
    
    distancias[origen] = 0
    cantidad_caminos[origen] = 1
    heap_min = Heap()
    heap_min.Encolar((0, origen))
    visitados = set() # IMPORTANTE ASI NO PROCESO CAMINOS 2 VECES

    while heap_min:
        dist_v, v = heap_min.Desencolar()

        if v in visitados:
            continue

        visitados.add(v)

        for w in grafo.adyacentes(v):
            nueva_dist = dist_v + grafo.peso_arista(v, w)

            if nueva_dist < distancias[w]:
                cantidad_caminos[w] = cantidad_caminos[v] # todos los caminos a v se extienden a w
                distancias[w] = nueva_dist
                padres[w] = v
                heap_min.Encolar((nueva_dist, w))
            
            # Encontramos OTRO camino de igual longitud mínima
            elif nueva_dist == distancias[w]:
                cantidad_caminos[w] += cantidad_caminos[v]  # Sumamos los caminos alternativos
    
    return cantidad_caminos[fin]

# O(V + E log V)


"""3. Implementar un algoritmo que, dado un grafo no dirigido, conexo, y sin puentes (es decir, sin ninguna arista que al
quitarla formaría más de una componente conexa), determine una dirección para cada arista, para que el grafo dirigido
resultante sea fuertemente conexo (es decir, haya una única componente fuertemente conexa). Indicar y justificar la
complejidad del algoritmo.
*/"""

def definir_direc(grafo):
    visitados = set()
    aristas_vis = set()
    origen = grafo.vertice_aleatorio()
    g_dir = Grafo(true, grafo.obtener_vertices())
    dfs(grafo, g_dir, aristas_vis, visitados, origen)
    return g_dir

def dfs(grafo, g_dir, aristas_vis, visitados, origen):
    visitados.add(origen)

    for w in grafo.adyacentes(origen):

        if (origen, w) in aristas_vis or (w, origen) in aristas_vis:
            continue

        aristas_vis.add((origen, w))
        aristas_vis.add((w, origen))

        if w not in visitados:
            g_dir.agregar_arista(origen, w)
            dfs(grafo, g_dir, aristas_vis, visitados, w)
        
        else: # w visitado oero arista no -> arista de retorno
            g_dir.agregar_arista(origen, w)




