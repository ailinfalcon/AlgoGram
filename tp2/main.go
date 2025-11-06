package algogram

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ruta := os.Args
	archivo, err := os.Open(ruta[1])

	if err != nil {
		fmt.Printf("Error al abrir el archivo")
		return
	}

	defer archivo.Close()

	algogram := Algogram(archivo)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		linea := s.Text()
		ProcesarComandos(algogram, linea)
	}
}
