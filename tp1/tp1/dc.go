package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp1/calculadora"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		linea := scanner.Text()

		// ignorar líneas vacías
		if strings.TrimSpace(linea) == "" {
			continue
		}

		resultado, err := calculadora.EvaluarLinea(linea)
		if err != nil {
			fmt.Println("ERROR")
		} else {
			fmt.Println(resultado)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error leyendo entrada:", err)
	}
}
