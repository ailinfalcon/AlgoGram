package algogram

import (
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	"tp2/errores"
)

type Algogram struct {
	usuarios        TDADiccionario.Diccionario[string, *usuario]
	usuarioLoggeado *usuario
	hayLoggeado     bool
	posts           TDAHeap.ColaPrioridad[post]
}

type usuario struct {
	nombre   string
	feed     TDAHeap.ColaPrioridad[post]
	afinidad int
}

type post struct {
	id        int
	usuario   *usuario
	contenido string
	likes     TDADiccionario.DiccionarioOrdenado[string, *usuario]
	cantLikes int
}

func CrearAlgogram(usuarios TDADiccionario.Diccionario[string, *usuario]) AlgoGram {
	algogram := new(Algogram)
	algogram.usuarioLoggeado = nil
	algogram.hayLoggeado = false
	algogram.usuarios = usuarios
	algogram.posts = nil

	return algogram
}

func (algogram *Algogram) Login(nombre string) error {
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

func (algogram *Algogram) hayUsuarioLoggeado() bool {
	return algogram.hayLoggeado
}

func (algogram *Algogram) loggearUsuario(nombre string) {
	usuario := algogram.usuarios.Obtener(nombre)
	algogram.usuarioLoggeado = usuario
	algogram.hayLoggeado = true
}

func (algogram *Algogram) Logout() error {
	if !algogram.hayUsuarioLoggeado() {
		return errores.ErrorUsuarioNoLoggeado{}
	}

	algogram.desloggearUsuario()

	//falta imprimir saludo
	return nil
}

func (algogram *Algogram) desloggearUsuario() {
	algogram.usuarioLoggeado = nil
	algogram.hayLoggeado = false
}

func crearNuevoPost(u *usuario, contenido string) post {
	nuevoPost := new(post)
	//nuevoPost.id = //calcular con la cantidad de post que hay en total
	nuevoPost.usuario = u
	nuevoPost.contenido = contenido
	nuevoPost.cantLikes = 0
	nuevoPost.likes = nil

	return *nuevoPost
}

func (algogram *Algogram) PublicarPost(contenido string) error {
	if !algogram.hayUsuarioLoggeado() {
		return errores.ErrorUsuarioNoLoggeado{}
	}

	//post := crearNuevoPost(algogram.usuarioLoggeado, contenido)
	//algogram.posts.Encolar

	//falta imprimir "Post Publicado"
	return nil
}

func (algogram *Algogram) LikearPost(id int) error

func (algogram *Algogram) MostrarLikes(id int) error {
	/*if algogram.posts.cantLikes == 0 { // || !heap.Pertenece(id)
		return errores.ErrorMostrarLikes{}
	}*/

	return nil
}
