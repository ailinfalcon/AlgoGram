package algogram

import (
	"bufio"
	"os"
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

var comandos = []string{
	_LOGIN, _LOGOUT, _PUBLICAR, _VER_SIGUIENTE_FEED, _LIKEAR, _MOSTRAR_LIKES,
}

func Algogram(archivo *os.File) error {
	s := bufio.NewScanner(archivo)
	cantUsuarios := 0 //--> seria la afinidad

	usuarios := dic.CrearHash[string, int](func(a, b string) bool { return a == b })

	for s.Scan() {
		usuarios.Guardar(s.Text(), cantUsuarios)
		cantUsuarios++
	}

	algogram.CrearAlgogram(usuarios)

	return nil
}

func ProcesarComandos(linea string) error {
	token := strings.Fields(linea)
	cmd, param := token[0], token[1]

	if esComandoValido(cmd) {

	}

	return nil
}

func esComandoValido(cmd string) bool {
	for _, comando := range comandos {
		if comando == cmd {
			return true
		}
	}

	return false
}
