package algogram

import (
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

	Algogram(archivo)
}
