package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	dic "tdas/diccionario"
	algogram "tp2/algogram"
)

const (
	_LOGIN              = "login"
	_LOGOUT             = "logout"
	_PUBLICAR           = "publicar"
	_VER_SIGUIENTE_FEED = "ver_siguiente_feed"
	_LIKEAR             = "likear_post"
	_MOSTRAR_LIKES      = "mostrar_likes"
)

// CargarUsuarios recibe un archivo del cual agrega el contenido de cada línea como clave de un diccionario,
// y como valor, el número de línea a la cual pertenece el contenido.
func CargarUsuarios(archivo *os.File) algogram.AlgoGram {
	s := bufio.NewScanner(archivo)
	cantUsuarios := 0

	usuarios := dic.CrearHash[string, int](func(a, b string) bool { return a == b })

	for s.Scan() {
		usuarios.Guardar(s.Text(), cantUsuarios)
		cantUsuarios++
	}

	algo := algogram.CrearAlgogram(usuarios)

	return algo
}

// ProcesarComandos separa el comando de los parametros recibidos, y ejecuta la función asociada al comando.
func ProcesarComandos(algogram algogram.AlgoGram, linea string) {
	comandos := guardarComandos()
	cmd, params, _ := strings.Cut(linea, " ")
	asignarComando(algogram, comandos, cmd, params)
}

// guardarComandos crea y devuelve un diccionario que relaciona el nombre de cada comando
// con la primitiva de AlgoGram que lo implementa.
func guardarComandos() dic.Diccionario[string, func(algogram.AlgoGram, string)] {
	comandos := dic.CrearHash[string, func(algogram.AlgoGram, string)](func(a, b string) bool { return a == b })

	comandos.Guardar(_LOGIN, func(algoG algogram.AlgoGram, parametro string) { ejecutarLogin(algoG, parametro) })
	comandos.Guardar(_LOGOUT, func(algoG algogram.AlgoGram, parametro string) { ejecutarLogout(algoG) })
	comandos.Guardar(_PUBLICAR, func(algoG algogram.AlgoGram, parametro string) { ejecutarPublicarPost(algoG, parametro) })
	comandos.Guardar(_VER_SIGUIENTE_FEED, func(algoG algogram.AlgoGram, parametro string) { ejecutarVerProximoPost(algoG) })
	comandos.Guardar(_LIKEAR, func(algoG algogram.AlgoGram, parametro string) { ejecutarLikearPost(algoG, parametro) })
	comandos.Guardar(_MOSTRAR_LIKES, func(algoG algogram.AlgoGram, parametro string) { ejecutarMostrarLikes(algoG, parametro) })

	return comandos
}

func ejecutarLogin(algogram algogram.AlgoGram, parametro string) {
	hayLoggeadoInicial := algogram.HayLoggeado()
	nombre := algogram.Login(parametro)
	if !hayLoggeadoInicial && parametro == nombre {
		fmt.Printf("Hola %v\n", nombre)
	}
}

func ejecutarLogout(algogram algogram.AlgoGram) {
	if algogram.Logout() {
		fmt.Println("Adios")
	}
}

func ejecutarPublicarPost(algogram algogram.AlgoGram, parametro string) {
	if algogram.PublicarPost(parametro) {
		fmt.Println("Post publicado")
	}
}

func ejecutarVerProximoPost(algogram algogram.AlgoGram) {
	post := algogram.VerProximoPost()
	if algogram.HayLoggeado() && post != nil {
		fmt.Printf(
			"Post ID %d\n%v dijo: %v\nLikes: %d\n",
			post.ObtenerId(), post.ObtenerPublicador(), post.ObtenerContenido(), post.ObtenerCantLikes(),
		)
	}
}

func ejecutarLikearPost(algogram algogram.AlgoGram, param string) {
	id, _ := esNumero(param)
	if algogram.LikearPost(id) {
		fmt.Println("Post likeado")
	}
}

func ejecutarMostrarLikes(algogram algogram.AlgoGram, parametro string) {
	id, _ := esNumero(parametro)
	likes, cantidad := algogram.MostrarLikes(id)

	if cantidad > 0 {
		fmt.Printf("El post tiene %d likes:\n", cantidad)
		imprimirUsuarios(likes)
	}
}

func imprimirUsuarios(likes []string) {
	for _, usuario := range likes {
		fmt.Printf("\t%v\n", usuario)
	}
}

// asignarComando recibe un comando, un parámetro y un diccionario de comandos con su primitiva de
// Algogram asociada. Si el comando recibido se encuentra en el diccionario, ejecuta la función
// correspondiente. Caso contrario, imprime por pantalla el mensaje "Comando inválido".
func asignarComando(algogram algogram.AlgoGram, comandos dic.Diccionario[string, func(algogram.AlgoGram, string)], comando, parametro string) {
	if !comandos.Pertenece(comando) {
		fmt.Printf("Comando inválido")
	} else {
		ejecutar := comandos.Obtener(comando)
		ejecutar(algogram, parametro)
	}
}

func esNumero(param string) (int, bool) {
	parametro, err := strconv.Atoi(param)
	if err != nil {
		return 0, false
	}
	return parametro, true
}
