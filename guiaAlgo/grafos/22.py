"""(★★★) Implementar una función que reciba un grafo no dirigido y 
determine si el mismo no tiene ciclos de una cantidad impar de vértices. 
Indicar y justificar la complejidad de la función.
"""

# equivale a decir: es bipartito?

from collections import deque

def sin_ciclos_impares(grafo):

    colores = {}

    for v in grafo.obtener_vertices():

        if v in colores:
            continue 

        q = deque()

        q.append(v)
        colores[v] = 0

        while q:

            ver = q.popleft()

            for w in grafo.adyacentes(ver):

                if w in colores:

                    if colores[ver] == colores[w]:
                        return False
                else:

                    colores[w] = 1 - colores[ver]
                    q.append(w)
            
    return True




