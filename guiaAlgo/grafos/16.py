def convertir_a_matriz(grafo):
    vertices = grafo.obtener_vertices()
    n = len(vertices)

    indices = {v: i for i, v in enumerate(vertices)}

    matriz = [[0] * n for _ in range(n)]

    for v in vertices:
        for w in grafo.adyacentes(v):
            i = indices[v]
            j = indices[w]
            peso = grafo.peso_arista(v, w)
            matriz[i][j] = peso
            matriz[j][i] = peso

    return matriz, vertices


##################

"""Diccionario de diccionarios
┌───────────────────────────┐
│ A: {B:10, C:5}           │
│ B: {A:10}                │
│ C: {A:5}                 │
└───────────────────────────┘
            ↓
Crear lista de vértices → [A, B, C]
            ↓
Mapear índices → A→0, B→1, C→2
            ↓
Matriz vacía 3x3 de ceros
            ↓
Rellenar con pesos usando peso_arista(v, w)
            ↓
┌───────────────────────────┐
│   A   B   C              │
│ A 0  10   5              │
│ B10   0   0              │
│ C 5   0   0              │
└───────────────────────────┘
            ↓
Return (matriz, [A, B, C])"""
