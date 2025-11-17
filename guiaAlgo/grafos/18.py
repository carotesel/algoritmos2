
# como cada op peso = 1 -> bfs con cant ops

from collections import deque

def menor_cant_operaciones(x, y):

    cola = deque()
    visitados = set()

    cola.append((x, 0)) # encolo numero y cantidad operaciones
    visitados.add(x)

    while len(cola) > 0:

        v, ops = cola.popleft()

        if v == y:
            return ops
        
        # genero vertices
        sig1 = v - 1
        sig2 = v * 2

        # agrego a la cola a los que genere (actualizando las ops) asi si vuelve a entrar y sig == y, corta
        for siguiente in [sig1, sig2]:

            if siguiente not in visitados and siguiente > 0:
                cola.append((siguiente, ops + 1))
                visitados.add(siguiente)
                