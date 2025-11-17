def comprobar_teorema(grafo):

    grado = {}
    cant_impar = 0
   
    for v in grafo.obtener_vertices():
        grado[v] = 0
        
        for w in grafo.adyacentes(v):
            grado[v]+= 1
    
    for item in grado:
        if grado[item] % 2 != 0:
            cant_impar+=1

    
    return cant_impar%2 == 0

