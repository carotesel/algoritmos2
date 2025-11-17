def encontrar_ciclo(g):
  visitados = {}
  padre = {}
  for v in g.obtener_vertices():
    if v not in visitados:
      ciclo = dfs_ciclo(g, v, visitados, padre)
      if ciclo is not None:
        return ciclo
  return None

def dfs_ciclo(g, v, visitados, padre):
  visitados[v] = True
  for w in g.adyacentes(v):
    if w in visitados:
      # Si w fue visitado y es padre de v, entonces es la arista de donde
      # vengo (no es ciclo).
      # Si no es su padre, esta arista (v, w) cierra un ciclo que empieza
      # en w.
      if w != padre[v]:
        return reconstruir_ciclo(padre, v, w)
    else:
      padre[w] = v
      ciclo = dfs_ciclo(g, w, visitados, padre)
      if ciclo is not None:
        return ciclo

  # Si llegamos hasta acá es porque no encontramos ningún ciclo.
  return None

def reconstruir_ciclo(padre, inicio, fin):
  v = fin
  camino = []
  while v != inicio:
    camino.append(v)
    v = padre[v]
  camino.append(inicio)
  return camino[::-1]