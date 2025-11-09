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
	token := strings.Fields(linea)
	cmd, params := token[0], token[1]

	if !esComandoValido(cmd) {

	}

	asignarComando(algogram, cmd, params)
}

// asignar comando con switch

func asignarComando(algogram algogram.AlgoGram, comando, parametro string) {

	switch comando {
	case _LOGIN:
		algogram.Login(parametro)
	case _LOGOUT:
		algogram.Logout()
	case _PUBLICAR:
		algogram.PublicarPost(parametro)
	case _VER_SIGUIENTE_FEED:
		ejecutarPublicarPost(algogram)
	case _LIKEAR:
		num, _ := esNumero(parametro)
		algogram.LikearPost(num)
	case _MOSTRAR_LIKES:
		num, _ := esNumero(parametro)
		likes, cantidad := algogram.MostrarLikes(num)
		fmt.Printf("El post tiene %d likes", cantidad)
		fmt.Println(likes)
	}
}

func ejecutarPublicarPost(algogram algogram.AlgoGram) {
	res := algogram.VerProximoPost()
	fmt.Printf(
		"Post ID %d\n%v dijo: %v\nLikes: %d\n",
		res.Id, res.Publicador, res.Contenido, res.CantLikes,
	)
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

func esComandoValido(cmd string) bool {
	for _, comando := range comandos {
		if comando.cmd == cmd {
			return true
		}
	}

	return false
}

func esNumero(param string) (int, bool) {
	parametro, err := strconv.Atoi(param)
	if err != nil {
		return 0, false
	}
	return parametro, true
}
