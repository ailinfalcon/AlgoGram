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

type parametro int

const (
	ninguno parametro = iota
	texto
	numero
)

type comando struct {
	cmd       string
	parametro parametro
	// aplicar func(string) error
}

/*
	var comandos = []string{
		_LOGIN, _LOGOUT, _PUBLICAR, _VER_SIGUIENTE_FEED, _LIKEAR, _MOSTRAR_LIKES,
	}
*/
var comandos = []comando{
	{_LOGIN, texto}, {_LOGOUT, ninguno}, {_PUBLICAR, texto}, {_VER_SIGUIENTE_FEED, ninguno},
	{_LIKEAR, numero}, {_MOSTRAR_LIKES, numero},
}

/*
	var comandos = []comando{
		{_LOGIN, texto, algogram.Login()},
		{_LOGOUT, ninguno, algogram.Logout()},
		{_PUBLICAR, texto, algogram.PublicarPost()},
		{_VER_SIGUIENTE_FEED, ninguno, algogram.VerProximoPost()},
		{_LIKEAR, numero, algogram.LikearPost()},
		{_MOSTRAR_LIKES, numero, algogram.MostrarLikes()},
	}
*/
func CargarUsuarios(archivo *os.File) algogram.AlgoGram {
	s := bufio.NewScanner(archivo)
	cantUsuarios := 0 //--> seria la afinidad

	usuarios := dic.CrearHash[string, int](func(a, b string) bool { return a == b })

	for s.Scan() {
		usuarios.Guardar(s.Text(), cantUsuarios)
		cantUsuarios++
	}

	algo := algogram.CrearAlgogram(usuarios)

	return algo
}

func ProcesarComandos(algogram algogram.AlgoGram, linea string) {
	/*
		token := strings.Fields(linea)
		cmd := token[0]
		var params string

		comando, esValido := buscarComando(cmd)
		if comando.parametro != ninguno {
			params = token[1]
		}
		if !esValido {

		}
	*/
	cmd, params, _ := strings.Cut(linea, " ") // corta la linea en el primer " "
	asignarComando(algogram, cmd, params)
}

// asignar comando con switch

func asignarComando(algogram algogram.AlgoGram, comando, parametro string) {

	switch comando {
	case _LOGIN:
		ejecutarLogin(algogram, parametro)
	case _LOGOUT:
		ejecutarLogout(algogram)
	case _PUBLICAR:
		ejecutarPublicarPost(algogram, parametro)
	case _VER_SIGUIENTE_FEED:
		ejecutarVerProximoPost(algogram)
	case _LIKEAR:
		ejecutarLikearPost(algogram, parametro)
	case _MOSTRAR_LIKES:
		ejecutarMostrarLikes(algogram, parametro)
	}
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
	if algogram.HayLoggeado() && post.ObtenerContenido() != "" {
		fmt.Printf(
			"Post ID %d\n%v dijo: %v\nLikes: %d\n",
			id, publicador, contenido, cantLikes,
		)
	}
}

func ejecutarLikearPost(algogram algogram.AlgoGram, param string) {
	num, _ := esNumero(param)
	if algogram.LikearPost(num) { // falta chequear que el post exista
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

// funcionaria si el struct comando fuese de tipo genérico
/*func asignarComando(algogram algogram.AlgoGram, comandoIngresado, parametro string) {
	for _, comando := range comandos {
		if comando.cmd == comandoIngresado {
			if comando.parametro == ninguno {
				comando.aplicar()
			}
			if comando.parametro == texto {
				comando.aplicar(parametro)
			}
			if comando.parametro == numero {
				num, _ := esNumero(comando.cmd)
				comando.aplicar(num)
			}
		}
	}
}
*/

func buscarComando(cmd string) (comando, bool) {
	for _, comando := range comandos {
		if comando.cmd == cmd {
			return comando, true
		}
	}

	return comando{}, false
}

func esNumero(param string) (int, bool) {
	parametro, err := strconv.Atoi(param)
	if err != nil {
		return 0, false
	}
	return parametro, true
}
