def grados_entrada(g): #O(V+E)
    resultado={v:0 for v in g.obtener_vertices()}
    vertices=g.obtener_vertices()
    
    for vertice in vertices:
        adyacentes=g.adyacentes(vertice)
        for ady in adyacentes:
            resultado[ady]+=1
    
    return resultado

def grados_salida(g):#O(V)
    resultado={}
   
    for vertice in g.obtener_vertices():
        resultado[vertice]=len(g.adyacentes(vertice))
    
    return resultado

def amenazados(grafo):
    resultado=[]
    grados_ent=grados_entrada(grafo)
    grados_sal=grados_salida(grafo)
    
    for vertice in grafo.obtener_vertices():
        for ady in grafo.adyacentes(vertice):
            if grados_ent[ady]==1:
                if vertice not in resultado:
                    resultado.append(vertice)
            if grados_sal[vertice]==1:
                if ady not in resultado:
                    resultado.append(ady)
    
    return resultado