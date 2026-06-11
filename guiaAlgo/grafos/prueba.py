def obtener_grado_g_no_dir(grafo):
    grado = {}

    for v in grafo:
        grado[v] = len(grafo.adyacentes(v))

    return grado

def obtener_grado_salida_g_dir(grafo):
    grado = {}

    for v in grafo:
        grado[v] = len(grafo.adyacentes(v))

    return grado

def obtener_grado_entrada_g_dir(grafo):
    grado = {}

    for v in grafo:
        grado[v] = 0
    
    for v in grafo:
        for w in grafo.adyacentes(v):
            grado[w] += 1

    return grado



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
    vertices = []

    dist[v] = 0

    for v in grafo:

        for w in grafo.adyacentes(v):

            dist[w] = dist[v] + 1
            if dist[w] == n:
                vertices.append(w)
    return w


"""función encontrar_ciclo(grafo):
    color[v] = BLANCO para todo v        ← nadie visitado

    para cada vértice v en el grafo:
        si color[v] == BLANCO:
            resultado = DFS(v)
            si resultado tiene ciclo:
                devolver ese ciclo

    devolver "no hay ciclo"

función DFS(v, padre, color, camino):
    color[v] = GRIS                      ← estoy visitando v ahora
    agregar v al camino

    para cada vecino w de v:
        si color[w] == GRIS:             ← ¡ciclo encontrado!
            devolver el ciclo desde w hasta v en el camino
        si color[w] == BLANCO:
            resultado = DFS(w, ...)
            si resultado tiene ciclo: devolver ciclo

    color[v] = NEGRO                     ← terminé con v
    sacar v del camino
    devolver "sin ciclo"""

# 0 = blanco
# 1 = gris
# 2 = negro

def encontrar_ciclo(grafo):
    colores = {}

    for v in grafo.obtener_vertices():
        colores[v] = 0      # todos blancos al principio

    for v in grafo.obtener_vertices():
        if colores[v] == 0:
            resultado = dfs(grafo, v, colores, [])
            if resultado is not None:
                return resultado

    return None

def dfs(grafo, v, colores, camino):

    colores[v] = 1 # lo marco gris
    camino.append(v) # lo agrego a camino actual

    for w in grafo.adyacentes(v):

        if colores[w] == 1: # hay un ady que es gris
            indice = camino.index(w)
            return camino[indice:]
        
        if colores[w] == 0: # es blanco
            res = dfs(grafo, w, colores, camino)
            if res is not None:
                return res
        
        # ya explore todo desde v
        colores[v] = 2
        camino.pop()
        return None
    

def bipartito(grafo):

    colores = {}
    cola = Cola()

    ini = grafo.obtener_vecrtices()[0]
    colores[ini] = 0
    cola.Encolar(ini)

    while not cola.EstaVacia():

        v = cola.Desencolar()

        for w in grafo.adyacentes(v):

            if w in colores:
                if colores[w] == colores[v]:
                    return False
            else:
                colores[w] = 1 - colores[v]
                cola.Encolar(w)

    return True
