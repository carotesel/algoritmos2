# O(V + E)

def grados(g):
    # devolver un diccionario string -> int
    res = {}

    for v in g.obtener_vertices():
        res[v] = len(g.adyacentes(v))

    return res

# O(V) + O(V + E) → O(V + E)
def grados_entrada(g):
    # devolver un diccionario string -> int
    res = {}

    for v in g.obtener_vertices():
        res[v] = 0

    for v in g.obtener_vertices():
        for w in g.adyacentes(v):
            res[w]+=1
    return res

# O(V + E)
def grados_salida(g):
    res = {}
    # devolver un diccionario string -> int

    for v in g.obtener_vertices():
        res[v] = len(g.adyacentes(v))

    return res