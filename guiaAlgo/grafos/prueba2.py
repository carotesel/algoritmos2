class Grafo:
    def __init__(self, dirigido=False):
        self.dirigido = dirigido
        self.ady = {}

    def agregar_vertice(self, v):
        if v not in self.ady:
            self.ady[v] = []

    def agregar_arista(self, v, w):
        self.agregar_vertice(v)
        self.agregar_vertice(w)

        self.ady[v].append(w)

        if not self.dirigido:
            self.ady[w].append(v)

    def adyacentes(self, v):
        return self.ady[v]

    def __iter__(self):
        return iter(self.ady)

g = Grafo(dirigido=True)

g.agregar_arista("A", "B")
g.agregar_arista("A", "C")
g.agregar_arista("A", "D")

g.agregar_arista("B", "E")
g.agregar_arista("B", "F")

g.agregar_arista("C", "G")
g.agregar_arista("C", "H")

g.agregar_arista("D", "I")

g.agregar_arista("F", "G")


"""(★★★) Implementar un algoritmo que reciba un grafo dirigido, un vértice V y un número N, 
y devuelva una lista con todos los vértices que se encuentren a exactamente N aristas de distancia del vértice V. 
Indicar el tipo de recorrido utilizado y el orden del algoritmo. Justificar."""

def a_N_aristas(grafo, v, n):

    dist = {}
    visitados = set()
    cola = deque()
    res = []

    dist[v] = 0
    visitados.add(v)
    cola.append(v)

    while cola:

        x = cola.popleft()
        for w in grafo.adyacentes(x):
            if w not in visitados:
                visitados.add(w)
                dist[w] = dist[v] + 1

                if dist[w] == n:
                    res.append(w)

                cola.append(w)

    return res

"""(★★★) Un autor decidió escribir un libro con varias tramas que se puede leer de forma no lineal. Es decir, por ejemplo, después del capítulo 1 puede leer el 2 o el 73; pero la historia no tiene sentido si se abordan estos últimos antes que el 1.

Siendo un aficionado de la computación, el autor ahora necesita un orden para publicar su obra, y decidió modelar este problema como un grafo dirigido, en dónde los capítulos son los vértices y sus dependencias las aristas. Así existen, por ejemplo, las aristas (v1, v2) y (v1, v73).

Escribir un algoritmo que devuelva un orden en el que se puede leer la historia sin obviar ningún capítulo. Indicar la complejidad del algoritmo."""

def orden_topologico(grafo):

    grados_ent = grados_entrada(grafo)
    cola = Cola()
    orden = []

    for v in grafo:
        if grados_ent[v] == 0:
            cola.Encolar(v)
    
    while not cola.EstaVacia():
        x = cola.Desencolar()
        orden.append(x)

        for w in grafo.adyacentes(x):
                grados_ent[w] -= 1
                if grados_ent[w] == 0:
                    cola.Encolar(w)

    return res

"""(★★★★) Dado un número inicial X se pueden realizar dos tipos de operaciones sobre el número:
Multiplicar por 2
Restarle 1. 

Implementar un algoritmo que encuentra la menor cantidad de operaciones a realizar para convertir el número X en el número Y, 
con tan solo las operaciones mencionadas arriba (podemos aplicarlas la cantidad de veces que querramos)."""

# grafo se genera sobre la marcha
# BFS

def menor_operaciones(X, Y):

    if X == Y:
        return [X]
    
    cola = Cola()
    visitados = set()
    padre = {}
    distancia = {}

    cola.Encolar(X)
    visitados.add(X)
    padre[X] = None
    distancia[X] = 0

    while not cola.EstaVacia():

        x = cola.Desencolar()

        for w in [x * 2, x - 1]:

            if w not in visitados:

                visitados.add(w)
                cola.Encolar(w)
                padre[w] = x
                distancia[w] = distancia[x] + 1

                if w == Y:
                    return reconstruir_camino(padre, distancia, Y)
    return None

def reconstruir_camino(padre, distancia, Y):
    actual = Y
    camino = []
    while actual is not None:
        camino.append(actual)
        actual = padre[actual]
    return camino[::-1]


"""(★★★) 
♠
♠ Daniel está a punto de casarse y tiene un problema: gastó casi todo su dinero en la luna de miel. Contrató un salón para la fiesta donde sólo hay 2 mesas (muy, muy grandes, pero 2 en fin). Debe repartir a los 
n
n invitados entre las dos mesas, y su esposo le indicó una condición: en cada mesa debe sentarse gente que se lleve bien entre todos ellos. Daniel cuenta con la información de quién se lleva bien con quién, y necesita poder determinar si hay alguna forma de separar en dos grupos de gente donde en cada grupo todos se lleven bien entre sí.

a. Modelar este problema utilizando grafos, indicando claramente qué son los vértices y qué las aristas.

b. Implementar un algoritmo que reciba un grafo como el modelado en el punto (a) y devuelva ambos grupos de personas. Indicar y justificar la complejidad del algoritmo implementado.

IMPORTANTE: tener en cuenta que resolver el problema de forma directa puede ser difícil. Recomendamos plantearse el problema inverso: poder separar en dos grupos tal que en ningúno de los grupos haya un par que no se lleven bien."""

def es_bipartito_gral(grafo):

    colores = {}

    for v in grafo:
        if v not in colores:
            if not es_bipartito(grafo, v, colores):
                return None, None
    
    mesa1 = []
    mesa2 = []

    for c in colores:

        if colores[c] == 0:
            mesa1.append(c)
        else:
            mesa2.append(c)
    
    return mesa1, mesa2

def es_bipartito(grafo, v, colores):

    colores[v] = 0
    cola = Cola()
    cola.Encolar(v)

    while not cola.EstaVacia():

        x = cola.Desencolar()

        for w in grafo.adyacentes(x):

            if w not in colores:
                colores[w] = 1 - colores[x]
                cola.Encolar(w)
            else:
                if colores[w] == colores[x]:
                    return False
    return True 

"""(★★) Implementar un algoritmo que reciba un grafo no dirigido y determine la cantidad mínima de aristas que debería agregársele para que el grafo sea conexo. 
Obviamente, si el grafo ya es conexo el algoritmo debe devolver 0. Indicar y justificar la complejidad del algoritmo implementado."""

def cuantas_aristas_conexo(grafo):

    visitados = set()
    componentes = 0

    for v in grafo:
        if v not in visitados:
            dfs(grafo, v, visitados)
            componentes += 1
    
    return componentes - 1

def dfs(grafo, v, visitados):
    visitados.add(v)

    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs(grafo, w, visitados)

"""(★★★) La teoría de los 6 grados de separación dice que cualquiera en la Tierra puede estar conectado a cualquier otra persona del planeta a través de una cadena de 
conocidos que no tiene más de cinco intermediarios (conectando a ambas personas con solo seis enlaces). 

Suponiendo que se tiene un grafo G en el que cada vértice es una persona y cada arista conecta gente que se conoce (el grafo es no dirigido):

a. Implementar un algoritmo para comprobar si se cumple tal teoría para todo el conjunto de personas representadas en el grafo G. Indicar el orden del algoritmo.

"""

def se_cumple_teoria(grafo):

    for v in grafo:
        if not bfs(grafo, v):
            return False

    return True

def bfs(grafo, v):

    dist = {}

    for v in grafo:
        dist[v] = float("inf")
    
    cola = Cola()
    cola.Encolar(v)
    dist[v] = 0

    while cola:
        x = cola.Desencolar()

        for w in grafo.adyacentes(x):
            if dist[w] == float("inf"):
                dist[w] = dist[x] + 1
                cola.Encolar(w)
    
    for d in dist:
        if d != v and ((dist[v] > 6) or (dist[d] == float("inf"))) :
            return False
    return True

# Complejidad: O(V * (V + E))

"""(★★★) Matías está en Barcelona y quiere recorrer un museo. Su idea es hacer un recorrido bastante lógico: empezar en una sala (al azar), 
luego ir a una adyacente a ésta, luego a una adyancente a la segunda (si no fue visitada aún), y así hasta recorrer todas las salas. 
Cuando no tiene más salas adyacentes para visitar (porque ya fueron todas visitadas), simplemente vuelve por donde vino buscando otras salas adyacentes. 
Teniendo un grafo no dirigido, que representa el mapa del museo (donde los vértices son salas, 
y las aristas (v, w) indican que las salas v y w se encuentran conectadas), implementar un algoritmo que nos devuelva una lista con un recorrido posible 
de la idea de Matías para visitar las salas del museo. Indicar el recorrido utilizado y el orden del algoritmo. Justificar."""

def recorrido_posible(grafo):
    res = []
    visitados = set()

    for v in grafo:
        if v not in visitados:
            dfs(grafo, v, visitados, res)
    
    return res

def dfs(grafo, v, visitados, res):

    visitados.add(v)
    res.append(v)

    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs(grafo, w, visitados, res)


"""3. Implementar un algoritmo que encuentre todas las componentes conexas de un grafo no dirigido. Indicar y justificar la
complejidad del algoritmo."""

def obtener_componentes(grafo):

    componentes = []
    visitados = set()

    for v in grafo:
        if v not in visitados:
            componente = []
            dfs(grafo, v, visitados, componente)
            componentes.append(componente)
    
    return componentes

def dfs(grafo, v, visitados, componente):
    componente.append(v)
    visitados.add(v)

    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs(grafo, w, visitados, componente)

# O(V + E)

"""2. Implementar un algoritmo que, dado un grafo pesado (con pesos positivos), un vértice v y otro w, determine el camino mínimo
de v a w dentro del grafo, con una modificación: en caso de encontrar más de un camino mínimo, que desempate por aquel de
menor cantidad de aristas. Considerar que, justamente, podrían haber varios caminos de una misma distancia, que a su vez
sean la mínima. Indicar y justificar la complejidad del algoritmo implementado.
Por ejemplo, en el grafo de arriba hay 2 caminos mínimos del vértice A al vértice D: A -> D, A -> E -> D, ambos de costo 3,
por lo que se debe elegir el primero por ser de menor cantidad de aristas.
"""

def dijkstra(grafo, origen, fin):
    padres = {origen: None}
    distancias = {}
    aristas = {}          # 🆕 cantidad de aristas del camino mínimo hasta cada vértice

    for v in grafo:
        distancias[v] = float("inf")
        aristas[v] = float("inf")   # 🆕
    
    distancias[origen] = 0
    aristas[origen] = 0             # 🆕
    heap_min = Heap()
    heap_min.Encolar((0, 0, origen))  # 🆕 (distancia, cant_aristas, vertice)

    while not heap_min.EstaVacia():
        dist_v, aris_v, v = heap_min.Desencolar()  # 🆕 desempaquetamos 3 valores

        # ignoramos si ya encontramos algo mejor para v
        if dist_v > distancias[v]:
            continue
        if dist_v == distancias[v] and aris_v > aristas[v]:  # 🆕
            continue

        if v == fin:
            return padres

        for w in grafo.adyacentes(v):
            dist_w = distancias[v] + grafo.peso(v, w)
            aris_w = aristas[v] + 1                          # 🆕

            if dist_w < distancias[w]:                       # caso 1: mejor distancia
                padres[w] = v
                distancias[w] = dist_w
                aristas[w] = aris_w                          # 🆕
                heap_min.Encolar((dist_w, aris_w, w))

            elif dist_w == distancias[w] and aris_w < aristas[w]:  # 🆕 caso 2: empate, menos aristas
                padres[w] = v
                aristas[w] = aris_w
                heap_min.Encolar((dist_w, aris_w, w))

    return padres



def dijkstra(grafo, ini, fin):

    padres = {}
    dist = {}
    aristas = {}

    for v in grafo:
        dist[v] = float("inf")
        aristas[v] = float("inf")
    
    padres[ini] = None
    dist[ini] = 0

    q = Heap()
    q.Encolar((0, 0, ini))

    while q:
        dis, arist, v = q.Desencolar()

        if dis > dist[v]:
            continue
        if dis == dist[v] and arist > aristas[v]:
            continue
        if v == fin:
            return padres, dist
        
        for w in grafo.adyacentes(v):
            nueva_dis = dis + grafo.peso(v, w)

            if nueva_dis < dist[w]:
                padres[w] = v
                dist[w] = nueva_dis
                aristas[w] = arist + 1
                q.Encolar((nueva_dis, aristas[w], w))
            
            elif nueva_dis == dist[w] and arist < aristas[w]:
                padres[w] = v
                aristas[w] = arist
                q.Encolar((nueva_dis, arist, w))


"""3. Implementar un algoritmo que reciba un grafo no dirigido y determine la cantidad máxima de aristas que se pueden agregar
al mismo de tal forma que no se reduzcan la cantidad de componentes conexas que hay en el mismo. Indicar y justificar la
complejidad del algoritmo implementado."""


def obtener_componentes(grafo):

    componentes = []
    visitados = set()
    aristas = []

    for v in grafo:
        if v not in visitados:
            componente = []
            aristas_comp = dfs(grafo, v, visitados, componente)
            componentes.append(componente)
            aristas.append(aristas_comp)
    
    return componentes, aristas

def dfs(grafo, v, visitados, componente):
    componente.append(v)
    visitados.add(v)
    aristas = 0

    for w in grafo.adyacentes(v):
        if w not in visitados:
            aristas += 1
            aristas += dfs(grafo, w, visitados, componente)
    return aristas

def resolver_problema(grafo):

    comps, aristas = obtener_componentes(grafo)
    agregables = 0

    for i in range(len(comps)):
        n = len(comps[i])
        posibles = n * (n-1)//2 - aristas[i]
        agregables += posibles
    return agregables



# ultimas pudrirse

def bfs(grafo, ini):

    dist = {}
    cola = Cola()
    maxDist = 0
    
    dist[ini] = 0
    cola.Encolar(ini)

    while not cola.EstaVacia():
        v = cola.Desencolar()

        for w in grafo.adyacentes(v):
            if w not in dist:
                dist[w] = dist[v] + 1
                if dist[w] > maxDist:
                    maxDist = dist[w]
                cola.Encolar(w)
    
    ultimas_en_pudrirse = []

    for d in dist:
        if dist[d] == maxDist:
            ultimas_en_pudrirse.append(d)
    
    return ultimas_en_pudrirse

#####################

def dijkstra(grafo, origen):

    padres = {}
    distancias = {}
    q = Heap()

    for v in grafo:
        distancias[v] = float("inf")
    
    padres[origen] = None
    distancias[origen] = 0
    q.Encolar(0, origen)

    while q:
        dist, v = q.Desencolar()

        for w in grafo.adyacentes(v):
            nueva = dist + grafo.peso(v, w)

            if nueva < dist[w]:
                distancias[w] = nueva
                padres[w] = v
                q.Encolar(nueva, w)
    
    return padres, distancias

def importantes(grafo, k):

    importancia = {}
    importantes = []

    for v in grafo: # O(V)
        importancia[v] = 0

    for v in grafo: # O(V * (V + E log V) * V)
        padres, dist = dijkstra(grafo, v) 
    
        for w in grafo:
            if w == v:
                continue          
            actual = padres[w]
            while actual != v:    
                importancia[actual] += 1
                actual = padres[actual]
    

    q = Heap()

    for i in importancia:
        q.Encolar(importancia[i], i)
    
    for i in range(k):
        imp, v = q.Desencolar()
        importantes.append(v)

    return importantes

# O(V ^ 2 + V.E log V) * V + O(K log V)
# O(V ^ 3 + V^2 E log V) + O(K log V)


"""1. En nuestra huerta tenemos unos rociadores de insecticidas automáticos. Cada rociador cuenta con la dosis apropiada para
cubrir hasta un máximo de 5 plantaciones a su alrededor. Es necesario averiguar si algún rociador tiene más de 5 plantaciones
alrededor ya que de tener una mayor cantidad la dosis sería insuficiente. Se tiene un grafo en donde los vértices son los
rociadores y plantas, es no pesado y dirigido (el origen de una arista es el rociador y el destino es una planta en su rango).
Implementar una función que reciba este grafo y devuelva, en caso que un rociador tenga más de 5 plantaciones alrededor,
el conjunto de plantaciones alrededor de dicho rociador (si hay más de un rociador que cumpla esta condición, devolver la
información correspondiente a cualquiera de estos). En caso contrario, devolver None. Indicar y justificar la complejidad de la
función."""

def mas_de_5_plantaciones(grafo):


    for v in grafo: # O(V + E)
     if len(grafo.adyacentes(v)) > 5:
        return list(grafo.adyacentes(v))
    return None
    

#######

def bfs(grafo, ini, fin):

    distancias = {}
    padres = {}
    cola = Cola()

    distancias[ini] = 0
    padres[ini] = None
    cola.Encolar(ini)

    while cola:
        v = cola.Desencolar()

        if v == fin:
            return padres, distancias
        
        for w in grafo.adyacentes(v):
            if w not in distancias:
                padres[w] = v
                distancias[w] = distancias[v] + 1
                cola.Encolar(w)
    return padres, distancias

def reconstruir_camino(fin, padres):
    camino = []
    actual = fin

    while actual is not None:
        camino.append(actual)
        actual = padres[actual]
    
    camino.reverse()
    return camino

"""2. La exitosa empresa AlgoConnect está desarrollando dos aplicaciones:
AlgoFriends: Es una red social en donde los usuarios pueden subir imágenes y entablar relaciones. Cualquier usuario
puede conectarse con cualquier otro.
AlgoBuy: Es una plataforma de ventas, en donde cada usuario se registra como comprador o vendedor. El único motivo
de conexión posible entre usuarios es hacer consultas sobre un producto, es por esto que los compradores sólo pueden
interactuar con vendedores, y viceversa. Cada aplicación cuenta con su propio grafo, en ambos casos los grafos son no
dirigidos, no pesados, en donde los vértices son los usuarios y las aristas representan que dos usuarios se han conectado.
Lamentablemente hubo un error en el software que borró las etiquetas de estos grafos, por lo que debemos re-etiquetarlos para
saber si cada grafo corresponde a AlgoFriends o a AlgoBuy. Implementar una función que dado dos grafos nos devuelva
dos Strings que correspondan a la asignación de cada uno de los grafos. Debe devolver "AlgoFriends", "AlgoBuy" si el
primero correspondería al de AlgoFriends y el segundo al de AlgoBuy, "AlgoBuy", "AlgoFriends", si es lo contrario, o
directamente None, None si dado cómo son los grafos, no es posible determinar cuál aplicación es cada uno. Indicar y justificar
la complejidad de la función."""

def es_bipartito_gral(g):
    colores = {}

    for v in g:
        if v not in colores:
            if not es_bipartito(g, colores, v):
                return False
    return True

def es_bipartito(grafo, colores, v):
    colores[v] = 0
    cola = Cola()
    cola.Encolar(v)

    while cola:
        x = cola.Desencolar()

        for w in grafo.adyacentes(x):
            if w not in colores:
                colores[w] = 1 - colores[x]
                cola.Encolar(w)
            else:
                if colores[w] == colores[x]:
                    return False
    return True

def cual_es_cual(g1, g2):

    es_bip1 = es_bipartito_gral(g1)
    es_bip2 = es_bipartito_gral(g2)

    if not es_bip1 and not es_bip2:
        return None, None
    if es_bip1 and es_bip2:
        return None, None
    if es_bip1:
        return "Algobuy", "Algofriends"
    else:
        return "Algofriends", "Algobuy"
    

"""1. Implementar un algoritmo que reciba un grafo dirigido y acíclico y determine si dicho grafo admite un único orden
topológico. Indicar y justificar la complejidad de la función. Pista: pensar qué condición puede darse para que exista
más de un posible orden topológico."""

def es_unico_top(grafo):

    grados_ent = grados_entrada(grafo)
    cola = Cola()
    orden = []

    for v in grafo:
        if grados_ent[v] == 0:
            cola.Encolar(v)
    
    while cola:
        v = cola.Desencolar()
        if not cola.EstaVacia():
            return False
        orden.append(v)

        for w in grafo.adyacentes(v):
            grados_ent[w] -= 1
            if grados_ent[w] == 0:
                cola.Encolar(w)
    return True

def grados_ent(grafo):
    grados = {}
    
    for v in grafo:
        grados[v] = 0
    
    for v in grafo:
        for w in grafo.adyacentes(v):
            grados[w] += 1
    return grados


# orden topologico

def topologico(grafo):

    orden = []
    g_entrada = grados_ent(grafo)
    cola = Cola()

    for v in grafo:
        if g_entrada[v] == 0:
            cola.Encolar(v)
    
    while cola:
        v = cola.Desencolar()
        orden.append(v)

        for w in grafo.adyacentes(v):
            g_entrada[w] -= 1
            if g_entrada[w] == 0:
                cola.Encolar(w)
    return orden

"""1. Implementar un algoritmo que, dado un grafo pesado (con pesos positivos), un vértice v y otro w, determine la cantidad de
caminos mínimos que hay de v a w dentro del grafo. Considerar que, justamente, podrían haber varios caminos de una misma
distancia, que a su vez sean la mínima. Indicar y justificar la complejidad del algoritmo implementado.
Por ejemplo, en el grafo de abajo hay 3 caminos mínimos del vértice A al vértice E: A -> E, A -> B -> F -> D -> E, A -> H
-> F -> D -> E, todos de costo 8."""


def dijkstra(grafo, ini, fin):

    padres = {}
    distancias = {}
    cant_caminos = {}
    q = heap()

    for v in grafo:
        distancias[v] = float("inf")
        cant_caminos[v] = 0
    
    padres[ini] = None
    distancias[ini] = 0
    cant_caminos[ini] = 1
    q.Encolar(0, v)

    while q:
        dist, v = q.Desencolar()

        if dist > distancias[v]:
            continue

        for w in grafo.adyacentes(v):
            nueva = dist + grafo.peso(v, w)

            if nueva < distancias[w]:
                distancias[w] = nueva
                padres[w] = v
                cant_caminos[w] = cant_caminos[v]
                q.Encolar(nueva, w)
            elif nueva == distancias[w]:
                cant_caminos[w] += cant_caminos[v]
    return cant_caminos[fin]

    
# MST

def kruskal(grafo):

    arbol = Grafo(es_dirigido=False, vertices=grafo.obtener_vertices())
    aristas = sorted(obtener_aristas(grafo))
    conjuntos = UnionFind(grafo.obtener_vertices())

    for a in aristas:
        v, w, peso = a

        if conjuntos.find(v) == conjuntos.find(w):
            continue

        arbol.agregar_arista(v, w, peso)
        conjuntos.union(v, w)
    return arbol


"""Ejercicio — Dijkstra con desempate
Implementar el algoritmo de Dijkstra para encontrar el camino mínimo desde un vértice origen hacia todos los demás vértices del grafo, con la siguiente modificación:
Si para llegar a un vértice existen dos o más opciones de misma distancia mínima, elegir aquella cuyo primer vértice en el camino desde el origen sea de menor distancia (es decir, el vértice que viene inmediatamente después del origen en el camino). En caso de que esa distancia también sea la misma, se puede elegir cualquiera de esas opciones.
Indicar y justificar la complejidad del algoritmo implementado."""

def dijkstra_modif(grafo, origen):

    padres = {}
    distancias = {}
    primer_vert = {}

    for v in grafo:
        distancias[v] = float("inf")
    
    padres[origen] = None
    distancias[origen] = 0
    primer_vert[origen] = None

    q = Heap()
    q.Encolar(0, origen)

    while not q.EstaVacia():

        dist, v = q.Desencolar()

        if dist > dist[v]:
            continue

        for w in grafo.adyacentes(v):

            nueva = dist + grafo.peso(v, w)

            if v == origen:
                primer_v = w
            else:
                primer_v = primer_vert[v]
            
            if nueva < dist[w]:
                dist[w] = nueva
                padres[w] = v
                primer_vert[w] = primer_v
                q.Encolar(nueva, w)
            elif nueva == dist[w]:
                if dist[primer_v] < dist[primer_v[w]]:
                    padres[w] = v
                    primer_vert[w] = primer_v
                    q.Encolar(nueva, w)


def vertice_corte(grafo, s, t):
    visitados = set()
    res = []

    for v in grafo:
        if v not in visitados:
            dfs(grafo, visitados, v, set(), res)
    
    return res

def dfs(grafo, visitados, v, este_dfs, res):

    este_dfs.add(v)

    for w in grafo.adyacentes(v):
        if w in este_dfs:
            res.append(w)
        elif w not in visitados:
            dfs(grafo, visitados, w, este_dfs, res)
    visitados.add(v)

