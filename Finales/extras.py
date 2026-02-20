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




# contar aristas g no dir

def contar_aristas(grafo):

    aristas = []
    aristas_visit = set()

    for v in grafo:
        for w in grafo.adyacentes(v):
            if (v, w) not in aristas_visit and (w, v) not in aristas_visit:
                aristas.append((v, w))
                aristas_visit.add((v, w))
                aristas_visit.add((w, v))
    return len(aristas)

def es_arbol(grafo):

    vert = grafo.obtener_vertices()

    if len(vert) == 0: 
        return True
    
    aristas = contar_aristas(grafo)

    if aristas != len(vert) - 1:
        return False
    
    visitados = set()
    origen = vert[0]

    dfs(grafo, origen, visitados)

    if len(visitados) != len(vert):
        return False
    return True

def dfs(grafo, origen, visitados):

    visitados.add(origen)

    for w in grafo.adyacentes(origen):
        if w not in visitados:
            dfs(grafo, w, visitados)

"""Ejercicio 51: Camino entre dos vértices (mejor con DFS)
Implementá una función que determine si existe algún camino entre dos vértices v y w. Complejidad: O(V+E)."""

def hay_camino(grafo, v, w):
    visitados = set()
    lo_hay = dfs(grafo, v, w, visitados)
    return lo_hay

def dfs(grafo, v, w, visitados):

    visitados.add(v)

    for x in grafo.adyacentes(v):
        if x == w:
            return True
        
        if x not in visitados:
            if dfs(grafo, x, w, visitados):
                return True
    return False

"""Ejercicio 52: Todos los caminos de v a w.
Implementá un algoritmo que encuentre todos los caminos simples (sin repetir vértices) de v a w. 
Complejidad: O(V! ) en el peor caso. Justificá."""

def todos_caminos(grafo, ini, fin):
    caminos = []
    camino = []
    visitados = set()
    dfs(grafo, ini, fin, visitados, camino, caminos)
    return caminos

def dfs(grafo, v, fin, visitados, camino, caminos):

    visitados.add(v)
    camino.append(v)

    for x in grafo.adyacentes(v):
        if x == fin:
            caminos.append(camino.copy())
        else:
           # sigo recorriendoooo el dfs
           for w in grafo.adyacentes(v):
            if w not in visitados:
                dfs(grafo, w, fin, visitados, camino, caminos)

    # BACKTRACKING -> sigo recorriendo y por eso saco el ultimo
    camino.pop()
    visitados.remove(v)

# Componentes conexas de grafo no dirigido (devolver cada componente no solo la cantidad)

def componentes(grafo):

    componentes = []
    visitados = set()
    
    for v in grafo:
        if v not in visitados:
            componente = []
            bfs(grafo, v, componente, visitados)
            componentes.append(componente)
        return componentes 
    
def  bfs(grafo, v, componente, visitados):

    cola = Cola()
    cola.Encolar(v)

    while cola:
        x = cola.Desencolar()
        componente.append(x)

        for w in grafo.adyacentes(x):
            if w not in visitados:
                visitados.add(w)
                cola.Encolar(w)






    