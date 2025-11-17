package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDADiccionario "tdas/diccionario"
	algogram "tp2/algogramSistema"
)

type comandosAlgogram func(string, algogram.AlgoGram) string

func crearHashComandos() TDADiccionario.Diccionario[string, comandosAlgogram] {
	cmp := func(a, b string) bool { return a == b }
	hash := TDADiccionario.CrearHash[string, comandosAlgogram](cmp)

	comandosString := []string{"login", "logout", "publicar", "ver_siguiente_feed", "likear_post", "mostrar_likes"}
	comandosFunc := []comandosAlgogram{
		func(argumento string, red algogram.AlgoGram) string {
			return red.Login(argumento)
		},
		func(argumento string, red algogram.AlgoGram) string {
			return red.Logout()
		},
		func(argumento string, red algogram.AlgoGram) string {
			return red.Publicar(argumento)
		},
		func(argumento string, red algogram.AlgoGram) string {
			return red.VerSiguienteFeed()
		},
		func(argumento string, red algogram.AlgoGram) string {
			argumentoInt, err := strconv.Atoi(argumento)
			if err != nil {
				return "ERROR"
			}
			return red.LikearPost(argumentoInt)
		},
		func(argumento string, red algogram.AlgoGram) string {
			argumentoInt, err := strconv.Atoi(argumento)
			if err != nil {
				return "ERROR"
			}
			return red.MostrarLikes(argumentoInt)
		},
	}

	for i := 0; i < len(comandosString); i++ {
		hash.Guardar(comandosString[i], comandosFunc[i])
	}

	return hash
}

func leerUsuarios(ruta string) ([]string, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return []string{}, err
	}
	defer archivo.Close()

	usuarios := make([]string, 0)

	linea := bufio.NewScanner(archivo)

	for linea.Scan() {
		usuario := linea.Text()
		usuario = strings.Trim(usuario, "\n")
		if usuario != "" {
			usuarios = append(usuarios, usuario)
		}
	}
	return usuarios, nil
}

func crearComando(ingreso string) (string, string) {
	linea := strings.Fields(ingreso)
	comando := linea[0]
	argumento := ""

	if len(linea) >= 2 {
		argumentoSeparado := linea[1:]
		argumento = strings.Join(argumentoSeparado, " ")
	}
	return comando, argumento
}

func main() {
	rutaUsuarios := os.Args[1]

	usuarios, err := leerUsuarios(rutaUsuarios)

	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", "ERROR")
	}

	algoGram := algogram.CrearAlgoGram(usuarios)
	hashComandos := crearHashComandos()

	ingreso := bufio.NewScanner(os.Stdin)
	for ingreso.Scan() {
		comando, argumento := crearComando(ingreso.Text())

		resultado := "ERROR"

		if hashComandos.Pertenece(comando) {
			comandoAEjecutar := hashComandos.Obtener(comando)

			resultado = comandoAEjecutar(argumento, algoGram)
		}

		fmt.Fprintf(os.Stdout, "%s\n", resultado)
	}
}
