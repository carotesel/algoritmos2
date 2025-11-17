"""Implementar un algoritmo que reciba un grafo no dirigido y determine la cant MINIMA de aristas que se le deberian agregar 
para que sea conexo. Si el grafo es conexo debe devolver cero."""

from collections import deque
import random


# solucion sencilla: cuento las componentes, si el grafo tiene 1 --> conexo
#si el grafo tiene mas, cant minima aristas = componentes - 1 

def cant_min_aristas_conexo(grafo):

    componentes = 0
    visitados = set()

    for v in grafo:
        bfs(grafo, v, visitados)
        componentes+= 1
    
    if componentes == 1:
        return 0
    
    return componentes - 1

def bfs(grafo, origen, visitados):

    visitados.add(origen)
    cola = deque()
    cola.append(origen)

    while cola:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                cola.append(w)

    
