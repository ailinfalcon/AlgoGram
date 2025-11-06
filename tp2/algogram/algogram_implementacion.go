package TDAalgogram

import (
	"fmt"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	TDALista "tdas/lista"
	"tp2/errores"
)

type Algogram struct {
	usuarios        TDADiccionario.Diccionario[string, *usuario]
	usuarioLoggeado *usuario
	hayLoggeado     bool
	posts           TDALista.Lista[post]
}

type usuario struct {
	nombre   string
	feed     TDAHeap.ColaPrioridad[post]
	afinidad int
}

type post struct {
	id        int
	publicador   *usuario
	contenido string
	likes     TDADiccionario.DiccionarioOrdenado[string, *usuario]
	cantLikes int
}

func CrearAlgogram(us TDADiccionario.Diccionario[string, int]) AlgoGram {
	algogram := new(Algogram)
	algogram.usuarioLoggeado = nil
	algogram.hayLoggeado = false
	usuarios := TDADiccionario.CrearHash[string, *usuario](func(a, b string) bool { return a == b })

	iter := us.Iterador()

	for iter.HaySiguiente() {
		nombre, afinidad := iter.VerActual()
		usuarios.Guardar(nombre, crearUsuario(nombre, afinidad))

		iter.Siguiente()
	}

	algogram.usuarios = usuarios
	algogram.posts = nil

	return algogram
}

func crearUsuario(nombre string, afinidad int) *usuario {
	usuario := new(usuario)
	usuario.nombre = nombre
	usuario.afinidad = afinidad
	usuario.feed = TDAHeap.CrearHeap[post](igualdadPost) // Hay que hacer una funcion para que muestre por afinidad

	return usuario
}

func igualdadPost(post1, post2 post) int {
	return post2.id - post1.id
}

func (algogram *Algogram) Login(nombre string) error {
	if !algogram.usuarios.Pertenece(nombre) {
		return errores.ErrorUsuarioInexistente{}
	}

	if algogram.hayUsuarioLoggeado() {
		return errores.ErrorUsuarioLoggeado{}
	}

	algogram.loggearUsuario(nombre)

	fmt.Printf("Hola %s", nombre)
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

	fmt.Print("Adios")
	return nil
}

func (algogram *Algogram) desloggearUsuario() {
	algogram.usuarioLoggeado = nil
	algogram.hayLoggeado = false
}

func crearNuevoPost(u *usuario, contenido string, cant int) post {
	nuevoPost := new(post)
	nuevoPost.id = cant - 1
	nuevoPost.publicador = u
	nuevoPost.contenido = contenido
	nuevoPost.cantLikes = 0
	nuevoPost.likes = nil

	return *nuevoPost
}

func (algogram *Algogram) PublicarPost(contenido string) error {
	if !algogram.hayUsuarioLoggeado() {
		return errores.ErrorUsuarioNoLoggeado{}
	}

	post := crearNuevoPost(algogram.usuarioLoggeado, contenido, algogram.posts.Largo())
	algogram.posts.InsertarUltimo(post)

	fmt.Print("Post publicado")
	return nil
}

func (algogram *Algogram) VerProximoPost() error {
	if !algogram.hayUsuarioLoggeado() || algogram.usuarioLoggeado.feed.Cantidad() == 0{
		return errores.ErrorVerProximoPost{}
	}

	post := algogram.usuarioLoggeado.feed.Desencolar()
	id := post.id
	contenido := post.contenido
	publicador := post.publicador.nombre
	cantLikes := post.cantLikes

	fmt.Printf("%s dijo: %s\nLikes: %d\nId Post: %d", publicador, contenido, id, cantLikes)
	return nil
}

func (algogram *Algogram) LikearPost(id int) error {
	if !algogram.hayUsuarioLoggeado() || algogram.posts.Largo() >= id {
		return errores.ErrorLikearPost{}
	}

	iter := algogram.posts.Iterador()

	for i := 0; i <= id; i++ { // en el peor de los casos O(p)
		iter.Siguiente()
	}

	postActual := iter.VerActual()

	if !postActual.likes.Pertenece(algogram.usuarioLoggeado.nombre) {
		postActual.likes.Guardar(algogram.usuarioLoggeado.nombre, algogram.usuarioLoggeado) // O(log Up)
		postActual.cantLikes++
	}

	fmt.Print("Post likeado")
	return nil
}

func (algogram *Algogram) MostrarLikes(id int) error {
	if algogram.posts.Largo() >= id { // asumiendo que los posts estan en una lista
		return errores.ErrorMostrarLikes{}
	}

	iter := algogram.posts.Iterador()

	for i := 0; i <= id; i++ { // en el peor de los casos O(p)
		iter.Siguiente()
	}
	postActual := iter.VerActual()
	if postActual.cantLikes == 0 {
		return errores.ErrorMostrarLikes{}
	}
	fmt.Printf("El post tiene %d likes:", postActual.cantLikes)

	iterLikes := postActual.likes.Iterador()
	for iterLikes.HaySiguiente() { // O(Up)
		nombreUsuario, _ := iterLikes.VerActual()

		fmt.Print(nombreUsuario)

		iterLikes.Siguiente()
	}

	return nil
}
