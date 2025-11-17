#(★★★) Matías está en Barcelona y quiere recorrer un museo. 
# Su idea es hacer un recorrido bastante lógico: empezar en una sala (al azar), luego ir a una adyacente a ésta, 
# luego a una adyancente a la segunda (si no fue visitada aún), y así hasta recorrer todas las salas. 
# Cuando no tiene más salas adyacentes para visitar (porque ya fueron todas visitadas),
# simplemente vuelve por donde vino buscando otras salas adyacentes. 
# 
# Teniendo un grafo no dirigido, que representa el mapa del museo (donde los vértices son salas, 
# y las aristas (v, w) indican que las salas v y w se encuentran conectadas), 
# implementar un algoritmo que nos devuelva una lista con un recorrido posible de la idea de Matías 
# para visitar las salas del museo. Indicar el recorrido utilizado y el orden del algoritmo. Justificar.


# CUANDO NO HAY MAS VUELVE --> DFS

import random

def recorrer_museo(grafo):

    salas = grafo.obtener_vertices()

    sala_inicial = random.choice(salas)

    visitadas = set()

    recorrido = []

    _dfs(grafo, sala_inicial, visitadas, recorrido)

    return recorrido


def _dfs(grafo, sala_inicial, visitadas, recorrido):

    recorrido.append(sala_inicial)
    visitadas.add(sala_inicial)

    for w in grafo.adyacentes(sala_inicial):

        if w not in visitadas:

            _dfs(grafo, w, visitadas, recorrido)



#Complejidad:
# O (V+E) DFS lol

#Se utiliza un recorrido en profundidad (DFS) porque Matías recorre una sala, luego se adentra en una adyacente, y solo retrocede cuando no quedan más por visitar.

#Cada vértice (sala) y cada arista (conexión) se procesan una sola vez, por lo que la complejidad temporal es O(V + E) y la espacial es O(V) por las estructuras auxiliares.




    



