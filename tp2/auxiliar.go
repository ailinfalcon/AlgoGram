package algogram

import (
	"bufio"
	"os"
	dic "tdas/diccionario"
	algogram "tp2/algogram"
)

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
