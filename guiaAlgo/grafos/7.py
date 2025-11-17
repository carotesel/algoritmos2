
def grafo_traspuesto(grafo):

    vertices = grafo.obtener_vertices()
    
    g_dirigido_res = Grafo(es_dirigido = True, vertices_init = vertices)

    for v in vertices:

        for w in grafo.adyacentes(v):

            g_dirigido_res.agregar_arista(w, v, 1)

    
    return g_dirigido_res


