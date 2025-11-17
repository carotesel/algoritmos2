def puede_ser_no_dirigido(grafo):
    
    for v in grafo.obtener_vertices():

        for w in grafo.adyacentes(v):
            if not grafo.estan_unidos(w,v) or grafo.peso_arista(v, w) !=  grafo.peso_arista(w, v):
                return False
    
    return True