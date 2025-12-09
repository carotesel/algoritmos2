/*
1. Implementar una primitiva para el ABB Predecesor(clave K) K que reciba una clave
(que puede estar en el árbol, o no) y devuelva la clave inmediatamente anterior a esta en el
recorrido inorder. Si no hay ninguna anterior, finalizar con un panic. Indicar y justificar
la complejidad de la primitiva.
*/


/* 
2. Implementar una primitiva eliminarColisiones(clave K) []K para el hash,
que elimine del hash todas las claves que colisionen con la clave pasada
por parámetro en el estado actual (sin eliminar dicha clave del
diccionario, si se encuentra) y devuelva dichas claves.

Implementar tanto para el hash abierto como para el hash cerrado. Si no
se implementa para alguno, el ejercicio no estará aprobable.

Indicar y justificar la complejidad de la primitiva para ambos casos.
*/

/*
3. ¡El Dr. Casa necesita nuestra ayuda! Resulta que venía una caravana de k colectivos
que estaban yendo por la Autopista Buenos Aires - La Plata, y tuvieron un choque en
cadena (nadie gravemente herido para nuestra suerte). En cada colectivo viajaban h
personas, y por razones de edad y patologías previas, aunque no hayan heridos graves,
debemos atender primero a ciertos pacientes que a otros. Por suerte para el Dr. Casa,
su jefa, la Dra. Cudi, le ha asignado una prioridad a cada paciente (siendo los
pacientes con prioridad 1 aquellos con la mayor prioridad) y su compañero, el
Dr. Persecución, ha ordenado la fila de pacientes de cada colectivo por prioridad
(considerar que cada fila es una cola enlazada, donde los pacientes más prioritarios
están al frente de la cola).

Teniendo k colas de h pacientes cada una, y una función ObtenerPrioridad(paciente)
(que se ejecuta en O(1)), implementar un algoritmo para que el Dr. Casa pueda saber
el orden en el que deben ser atendidos los k · h pacientes. Indicar y justificar la
complejidad del algoritmo implementado.
*/

// h = cant personas x colectivo
// k = cant colas

/*
4. Implementar una función que reciba un arreglo A de n enteros y un número k,
   y devuelva un nuevo arreglo R tal que, para cada posición i de R, el valor
   sea la máxima suma obtenible de EXACTAMENTE k elementos dentro del rango
   A[0..i] (incluyendo a i). Si para una posición i no hay suficientes
   elementos (i < k-1), entonces R[i] debe valer -1.

   Ejemplo:
     A = [1, 5, 3, 4, 2, 8], k = 3
     Resultado esperado: R = [-1, -1, 9, 12, 12, 17]

   Requisito de eficiencia:
     La complejidad del algoritmo debe ser MEJOR que O(n * k).

   Indicar y justificar la complejidad del algoritmo implementado.
*/

/*
5. Implementar una primitiva del ABB que, dado un valor entero M, una clave inicial "inicio"
   y una clave final "fin", devuelva una lista con todos los datos cuyas claves estén
   dentro del rango [inicio, fin] y que además se encuentren dentro de los primeros M niveles
   del árbol (considerando a la raíz en el nivel 1).

   Ejemplo:
     Si se tiene el siguiente ABB:

              10
             /  \
            5    15
           / \   / \
          3   8 12 20
             /     \
            7       14

     y se llama con M = 3, inicio = 5 y fin = 15,
     el resultado debe contener los datos de las claves 10, 5, 8, 12, 15 (en cualquier orden).

   Se debe indicar y justificar la complejidad temporal del algoritmo implementado.
*/