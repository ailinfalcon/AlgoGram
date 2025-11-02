package algogram

import (
	// "tdas/diccionario"
	"tp2/errores"
)

type AlgoGram struct {
	//usuarios * diccionario.Diccionario[string, string]()
	usuarioLoggeado *Usuario
	hayLoggeado     bool
	posts *post // *lista.Lista[post] ?
}

type post struct {
	id int
	usuario *Usuario
	contenido string // (?
	// likes *heap.Heap[string]
	cantLikes int
}

func (algogram *AlgoGram) Login(nombre string) error {
	//if !algogram.usuarios.Pertenece(nombre) {
	//return errores.ErrorUsuarioInexistente{}
	//}

	if algogram.hayUsuarioLoggeado() {
		return errores.ErrorUsuarioLoggeado{}
	}

	algogram.loggearUsuario(nombre)

	//falta imprimir saludo
	return nil
}

func (algogram *AlgoGram) hayUsuarioLoggeado() bool {
	return algogram.hayLoggeado
}

func (algogram *AlgoGram) loggearUsuario(nombre string) {
	algogram.usuarioLoggeado.nombre = nombre
	algogram.hayLoggeado = true
}

func (algogram *AlgoGram) Logout() error {
	if !algogram.hayUsuarioLoggeado() {
		return errores.ErrorUsuarioNoLoggeado{}
	}

	algogram.desloggearUsuario()

	//falta imprimir saludo
	return nil
}

func (algogram *AlgoGram) desloggearUsuario() {
	//algogram.usuario = " "
	algogram.hayLoggeado = false
}

func (algogram *AlgoGram) MostrarLikes(id int) error {
	if algogram.posts.cantLikes == 0 { // || !heap.Pertenece(id)
		return errores.ErrorMostrarLikes{}
	}
	return nil
}
