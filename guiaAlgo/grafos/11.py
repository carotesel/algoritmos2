from collections import deque

def a_n_aristas(grafo, v, n):
   
   if n == 0:
       return [v]

   res = []

   grados = {v: 0}
   visitados = set([v])
   cola = deque([v])

   while cola:
       actual = cola.popleft()

       for w in grafo.adyacentes(actual):
           if w not in visitados:
               grados[w] = grados[actual] + 1

               if grados[w] == n:
                   res.append(w)

               cola.append(w)
               visitados.add(w)

   return res
