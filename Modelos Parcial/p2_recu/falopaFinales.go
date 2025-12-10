/*
04/07/24
2. Implementar una primitiva para el heap func (heap *heap[T]) DiferenciaSimetrica(otro *heap[T]) ColaPrioridad[T], 
que reciba otro Heap y cree un nuevo Heap con los elementos del primero que no se ecuentren en el segundo, y viceversa 
(es decir, la diferencia simétrica entre ambos). La función de comparación del nuevo heap debe ser la del primer heap. 
Indicar y justificar la complejidad del algoritmo implementado.
*/

// 08/07/24
/*
1. Dos cadenas X e Y son isomórficas si existe alguna transformación biyectiva de caracteres que permita obtener Y a partir
de X. Ejemplos: casa y bata son isomórficas, y la transformación es c → b, a → a, s → t. burro y pizza son isomórficas,
y la transformación es b → p, u → i, r → z, o → a. mesa y masa no son isomórficas, porque la transformación debe
ser biyectiva: no podemos incluir e → a y a → a. Escribir una función que reciba dos cadenas y determine si son
isomórficas. Indicar y justificar la complejidad de dicha función.
*/

/*
2. Dado un arreglo de enteros ordenado de n elementos en el cual sus elementos van de 0 a M, con M n, implementar
una función que determine en O(log n) si hay algún elemento que aparezca más de la mitad de la veces en el arreglo.
Justificar la complejidad del algoritmo implementado.
*/

/*
3. Implementar una primitiva para la lista enlazada Modificar(modificador func(T) T) que modifique todos los datos
de la lista. Cada dato deberá a pasar ser el que resulta de aplicar la función modificador al dato que se encontraba
anteriormente en dicha posición. Por ejemplo, si la lista es [1, 2, 3] y la función func(elem int) int {return
elem * 2}, luego de ejecutar la primitiva con dicha función, la lista debe quedar como [2, 4, 6]. Indicar y justificar
la complejidad del algoritmo implementado.
*/

/*
15/07/24

1. Implementar en Go una primitiva para un árbol binario izquierdista, que reciba la cantidad de nodos que tiene, y
devuelva el dato del elemento más a abajo y a la derecha del árbol. En los árboles de las figuras mostradas, se debe
devolver en ambos casos 4. Para que el ejercicio se pueda considerar como aprobable, debe resolverse en no más que
O(n), sin contar con otros errores. Para que se considere completamente bien, debe ejecutar en O(log n). Justificar la
complejidad del algoritmo implementado. A fines del ejercicio, considerar que la estructura del árbol binario es:

type ab[T any] struct {
izquierda *ab[T]
derecha *ab[T]
dato T
}*/

/*
2. Implementar una primitiva para el árbol binario EsABB(func(T, T) int) bool que reciba una función de comparación y determine
si el árbol cumple con la propiedad de ABB para dicha función de comparación. Indicar y justificar la complejidad del algoritmo
implementado.
A fines del ejercicio, considerar que la estructura del árbol es la indicada en el dorso a este examen.

type ab[T any] struct {
izquierda *ab[T]
derecha *ab[T]
dato T
}
*/

/*
### 19/12/24

1. Implementar en Go una primitiva que reciba un árbol binario que representa un heap (árbol binario izquierdista, que
cumple la propiedad de heap), y devuelva la representación en arreglo del heap. La firma de la primitiva debe ser
RepresentacionArreglo() []T. Indicar y justificar la complejidad de la primitiva. La estructura del árbol binario es:
type ab[T any] struct {
izquierda *ab[T]
derecha *ab[T]
dato T
}

### 17/02/25

1. Responder las siguientes preguntas sobre árboles, detalladamente.
a. Suponer que implementamos una primitiva que invierte todo hijo derecho con el hijo izquierdo, similar a lo
practicado en clase, pero se realiza sobre un ABB. ¿Podemos decir que ahora el ABB se comporta con la propiedad
de ABB invertida?
b. Suponer que queremos hacer un ejercicio típico de árboles, de los sencillos, como calcular la cantidad de hojas de
un árbol. En general, lo implementaríamos de forma recursiva, por ejemplo, con un preorder. ¿Cómo lo harías,
siguiendo la misma lógica de recorrido, para realizarlo de forma iterativa? ¿Utilizarías alguna otra estructura de
datos vista en la materia? ¿cómo quedaría la complejidad?
c. Si un árbol cumple propiedad de ABB, podemos recorrer los elementos en orden en tiempo lineal. Si un árbol
cumple la propiedad de heap (de mínimos), ¿se puede hacer lo mismos?

### 14/07/25

1. Implementar la primitiva Invertir() para el Heap, que invierta su forma de comparar los elementos (es decir, si era de máximos,
ahora sea de mínimos), sin modificar las funciones y primitivas previamente implementadas (simplemente contener todo el cambio
dentro de la primitiva). Indicar y justificar la complejidad del algoritmo. Indicar qué consecuencias podría tener esta forma de
implementación si se invoca a la primitiva Invertir una cantidad k de veces, y cómo podría resolverse si se permitiera modificar
otras funciones y/o primitivas (y/o la estructura del heap en sí). DONE

2. Implementar la primitiva Interseccion(otro *abb[K, V]) Lista[K] para el ABB que nos devuelva una lista ordenada con la
intersección entre el árbol y el recibido por parámetro, que estén ocupando el mismo lugar en el árbol. Indicar y justificar la
complejidad del algoritmo implementado. En el ejemplo a continuación, la intersección sería [4, 10, 18, 20].
    
    

### 21/07/25

1. Implementar un algoritmo que reciba un arreglo de n números, y un número k, y devuelva los k números dentro del
arreglo cuya suma sería la máxima (entre todos los posibles subconjuntos de k elementos de dicho arreglo). Indicar y
justificar la complejidad de la función implementada.
DONE

### 11/08/25

1. Explicar detalladamente cómo modificarías la implementación del ABB para poder tener una primitiva Maximo y Minimo que nos
devuelva las claves máximas y mínimas, y que se realice en tiempo constante.

2. Implementar una primitiva para el hash cerrado filtro(func(V) bool) que elimine del hash todas las claves del mismo que
tengan asociado un valor para el cuál la función devuelva false. Indicar y justificar la complejidad de la primitiva implementada.
Ejemplo: si mi diccionario es de cadenas a números como el de arriba, e invocamos a la primitiva con una función que devuelve true
para los números pares y false para los impares, el diccionario debe quedar como el de abajo:
{"koala": 3, "rana": 2, "gato": 2, "perro": 5, "canguro": 4 }. DONE

↓

{"rana": 2, "gato": 2, "canguro": 4}
*/