def detectar_ciclo(grafo):
    visitados = set()
    padres = {}
    esta_rec = set()

    for v in grafo:
        if v not in visitados:
            ciclo = dfs(grafo, visitados, v, padres, esta_rec)
            if ciclo is not None:
                return ciclo
    return None

def dfs(grafo, visitados, v, padres, esta_rec):
    visitados.add(v)
    esta_rec.add(v)

    for w in grafo.adyacentes(v):
        if w in visitados:
            if w in esta_rec:
                return reconstruir_ciclo(padres, v, w)
        else:
            padres[w] = v
            ciclo = dfs(grafo, visitados, w, padres, esta_rec)
            if ciclo is not None:
                return ciclo
    return None

def reconstruir_ciclo(padres, v, w):
    ciclo = []
    x = w

    while x != v:
        ciclo.append(x)
        x = padres[x]
    ciclo.append(v)
    return ciclo.reverse()





############################

def hay_ciclo(grafo):
    visitados = set()
    padres = {}

    for v in grafo:
        if v not in visitados:
            ciclo = dfs(grafo, v, visitados, padres)
            if ciclo is not None:
                return ciclo
    return None

def dfs(grafo, v, visitados, padres):

    visitados.add(v)

    for w in grafo.adyacentes(v):
        if w in visitados:
            if w != padres[v]:
                return reconstruir_ciclo(padres, v, w)
        else:
            padres[w] = v
            ciclo = dfs(grafo, w, visitados, padres)
            if ciclo is not None:
                return ciclo
    return None

#######

# Ordeno aristas de menor a mayor, por cada arista si sus ady no son parte del mismo conj, la agrego al mst y 
# uno los conjuntos. 

def obtener_aristas(grafo):
    visitados = set()
    aristas = []

    for v in grafo:
        for w in grafo.ady(v):
            if w not in visitados:
                aristas.append((v, w, grafo.peso(v, w)))
        visitados.add(v)
    return aristas

def kruskal(grafo):

    conjuntos = UnionFind(grafo.obtener_vertices())
    aristas = sorted(obtener_aristas(grafo))
    arbol = Grafo(false, grafo.obtener_vertices())

    for a in aristas:
        v, w, peso = a
        if conjuntos.find(v) == conjuntos.find(w):
            continue
        arbol.agregar_arista(v, w, peso)
        conjuntos.union(v, w)
    return arbol

################

# MST Prim

# heap
# visitados 

def mst_prim(grafo):
    arbol = Grafo(false, grafo.obtener_vertices())
    origen = random.choice(grafo.obtener_vertices())
    heap_min = Heap()
    visitados = set()
    visitados.add(origen)

    for w in grafo.adyacentes(origen):
        heap_min.Encolar((grafo.peso(origen, w), origen, w))
    
    while heap_min:
        peso, v, w = heap_min.Desencolar()

        if w in visitados:
            continue
        arbol.agregar_arista(v, w, peso)
        visitados.add(w)

        for x in grafo.adyacentes(w):
            if x not in visitados:
               heap_min.Encolar((grafo.peso(w, x), w, x))
    return arbol
