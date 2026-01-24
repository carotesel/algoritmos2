"""4. Queremos resolver el problema de obtención de caminos mínimos en un grafo, pero con una variante. Nuestro grafo corresponde a la
red de calles de la ciudad. Como es bien sabido, el tránsito por las calles no es siempre el mismo, lo cual implica que el tiempo para
llegar a destino no es siempre el mismo. En nuestro grafo los pesos corresponden a cuándo (en tiempo) se llega de un vértice a otro a
partir de saber cuándo se comienza a transitar la arista, pero no es un valor constante sino una función (cada arista tiene su propia
función como peso), que dado el tiempo actual nos indica en qué tiempo pasamos por dicha arista hasta llegar al vértice destino (la
función es monótona creciente, así que no hay nunca ventaja en quedarse esperando en un lugar). Por ejemplo, si llego en tiempo t =
2 a un vértice v, veo la arista hacia w y aplico la función que se encuentra en la arista (v, w) con tiempo 2, y nos da como resultado
4, significa que en este caso llegaríamos a w en tiempo 4 (no que nos tomaría 4 para llegar, resultando en 6 en total). Implementar el
algoritmo de Dijkstra con las modificaciones necesarias para que funcione para un grafo de estas características, que permita obtener
el mínimo tiempo para llegar a todos los vértices desde un origen. Indicar y justificar la complejidad del algoritmo implementado
(considerar que las funciones de peso ejecutan en tiempo constante)."""

"""2. Implementar el algoritmo de Dijsktra para encontrar el camino mínimo de un vértice hacia los demás vértices del grafo, con la
siguiente modificación: si para llegar a un vértice tengo dos (o más) opciones de misma distancia mínima, elegir aquella cuyo primer
vértice en el camino sea de menor distancia (en caso que la distancia sea la misma, por la razón que fuere, entonces se puede elegir
cualquiera de aquellos). Indicar y justificar la complejidad del algoritmo implementado."""