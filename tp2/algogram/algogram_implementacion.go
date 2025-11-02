package algogram

import (
	"tdas/diccionario"
	heap "tdas/cola_prioridad"
	"tdas/lista"
	"tp2/errores"
)

type Algogram struct {
	usuarios * diccionario.Diccionario[string, string]()
	usuarioLoggeado *Usuario
	hayLoggeado     bool
	posts *lista.Lista[post]
}

type post struct {
	id int
	usuario *Usuario
	contenido string // (?
	likes *heap.ColaPrioridad[string]
	cantLikes int
}

func CrearAlgogram(usuarios diccionario.Diccionario[string, string]) AlgoGram {
	algogram := new(Algogram)
	algogram.usuarioLoggeado = nil
	algogram.hayLoggeado = false
	algogram.usuarios = usuarios

	return algogram
}

func (algogram *AlgoGram) Login(nombre string) error {
	if !algogram.usuarios.Pertenece(nombre) {
		return errores.ErrorUsuarioInexistente{}
	}

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
	algogram.usuarioLoggeado = nil
	algogram.hayLoggeado = false
}

func (algogram *AlgoGram) MostrarLikes(id int) error {
	if algogram.posts.cantLikes == 0 { // || !heap.Pertenece(id)
		return errores.ErrorMostrarLikes{}
	}
	
	return nil
}
