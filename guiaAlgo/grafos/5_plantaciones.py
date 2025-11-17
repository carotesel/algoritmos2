"""1. En nuestra huerta tenemos unos rociadores de insecticidas automáticos. Cada rociador cuenta con la dosis apropiada para
cubrir hasta un máximo de 5 plantaciones a su alrededor. Es necesario averiguar si algún rociador tiene más de 5 plantaciones
alrededor ya que de tener una mayor cantidad la dosis sería insuficiente. Se tiene un grafo en donde los vértices son los
rociadores y plantas, es no pesado y dirigido (el origen de una arista es el rociador y el destino es una planta en su rango).
Implementar una función que reciba este grafo y devuelva, en caso que un rociador tenga más de 5 plantaciones alrededor,
el conjunto de plantaciones alrededor de dicho rociador (si hay más de un rociador que cumpla esta condición, devolver la
información correspondiente a cualquiera de estos). En caso contrario, devolver None. Indicar y justificar la complejidad de la
función."""

from collections import deque

def tiene_mas_de_cinco(grafo):

    for v in grafo.obtener_vertices():
        
        if len(grafo.adyacentes(v)) > 0:

            alcanzados = bfs(grafo, v, set())

            if len(alcanzados) > 5:
                return alcanzados
    return None

def bfs(grafo, v, visitados):

    visitados.add(v)
    res = set()
    cola = deque()

    cola.append(v)

    while cola:
        actual = cola.popleft()

        for w in grafo.adyacentes(actual):

            if w not in visitados:

                visitados.add(w)
                res.add(w)
                cola.append(w)
    return res




 