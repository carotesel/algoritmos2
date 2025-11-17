package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

// leerEnterosDesdeArchivo abre el archivo indicado, lee línea por línea,
// convierte cada línea a int y devuelve un slice con los números.
func leerEnterosDesdeArchivo(ruta string) ([]int, error) {
	// Abrimos el archivo
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo %s: %v", ruta, err)
	}
	defer archivo.Close() //  cerrar el archivo al salir de la función

	var numeros []int
	scanner := bufio.NewScanner(archivo)

	// Leer línea por línea
	for scanner.Scan() {
		linea := scanner.Text()
		numero, err := strconv.Atoi(linea)
		if err != nil {
			// Si falla la conversión, devolvemos el error
			return nil, fmt.Errorf("error al convertir '%s' a entero en %s: %v", linea, ruta, err)
		}
		numeros = append(numeros, numero)
	}

	// Verificamos si hubo errores durante la lectura
	err = scanner.Err()

	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo %s: %v", ruta, err)
	}

	return numeros, nil
}

func mostrarItemsArray(vector []int) {
	for _, element := range vector {
		fmt.Println(element)
	}
}

func main() {
	// Leemos el primer archivo
	arrayFile1, err := leerEnterosDesdeArchivo("archivo1.in")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Leemos el segundo archivo
	arrayFile2, err := leerEnterosDesdeArchivo("archivo2.in")
	if err != nil {
		fmt.Println(err)
		return
	}

	// comparar arreglo mas grande

	comparacion := ejercicios.Comparar(arrayFile1, arrayFile2)

	// resultado comparacion:
	// si es 0 -> vec 1 == vec 2
	// si es 1 -> vec 1 > vec 2
	// si es -1 -> vec 1 < vec 2

	// Elegimos el arreglo mayor (si son iguales, usamos el primero)
	var mayor []int
	if comparacion >= 0 {
		mayor = arrayFile1
	} else {
		mayor = arrayFile2
	}

	// Ordenamos y mostramos
	ejercicios.Seleccion(mayor)
	mostrarItemsArray(mayor)
}
