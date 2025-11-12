package TDAalgogram

import (
	"fmt"
	"math"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	TDALista "tdas/lista"
	TDAPost "tp2/post"
)

type Algogram struct {
	usuarios        TDADiccionario.Diccionario[string, *usuario]
	usuarioLoggeado *usuario
	posts           TDALista.Lista[TDAPost.Post]
}

type usuario struct {
	nombre   string
	feed     TDAHeap.ColaPrioridad[*postFeed]
	afinidad int
}

type postFeed struct {
	post     TDAPost.Post
	afinidad int
}

func CrearAlgogram(us TDADiccionario.Diccionario[string, int]) *Algogram {
	usuarios := cargarUsuarios(us)

	return &Algogram{
		usuarioLoggeado: nil,
		usuarios:        usuarios,
		posts:           TDALista.CrearListaEnlazada[TDAPost.Post](),
	}
}

func (algogram *Algogram) HayLoggeado() bool {
	return algogram.usuarioLoggeado != nil
}

func (algogram *Algogram) Login(nombre string) string {
	if !algogram.usuarios.Pertenece(nombre) {
		fmt.Printf("Error: usuario no existente\n")
		return ""
	}

	if algogram.HayLoggeado() {
		fmt.Printf("Error: Ya habia un usuario loggeado\n")
		return ""
	}

	algogram.loggearUsuario(nombre)
	return nombre
}

func (algogram *Algogram) Logout() bool {
	if !algogram.HayLoggeado() {
		fmt.Printf("Error: no habia usuario loggeado\n")
		return false
	}

	algogram.desloggearUsuario()

	return true
}

func (algogram *Algogram) PublicarPost(contenido string) bool {
	if !algogram.HayLoggeado() {
		fmt.Printf("Error: no habia usuario loggeado\n")
		return false
	}

	post := TDAPost.CrearPost(algogram.usuarioLoggeado.nombre, contenido, algogram.posts.Largo())
	algogram.posts.InsertarUltimo(post)

	iter := algogram.usuarios.Iterador()
	for iter.HaySiguiente() {
		nombre, usuario := iter.VerActual()
		if nombre != algogram.usuarioLoggeado.nombre {
			postFeed := crearNuevoPostFeed(post, usuario.afinidad)
			usuario.feed.Encolar(postFeed)
		}
		iter.Siguiente()
	}

	return true
}

func (algogram *Algogram) VerProximoPost() TDAPost.Post {
	if !algogram.HayLoggeado() || algogram.usuarioLoggeado.feed.Cantidad() == 0 {
		fmt.Printf("Usuario no loggeado o no hay mas posts para ver\n")
		return nil
	}

	postFeed := algogram.usuarioLoggeado.feed.Desencolar()

	// return postFeed.post.id, postFeed.post.publicador.nombre, postFeed.post.contenido, postFeed.post.cantLikes
	return postFeed.post
}

func (algogram *Algogram) loggearUsuario(nombre string) {
	usuario := algogram.usuarios.Obtener(nombre)
	algogram.usuarioLoggeado = usuario
}

func (algogram *Algogram) desloggearUsuario() {
	algogram.usuarioLoggeado = nil
}

func cargarUsuarios(us TDADiccionario.Diccionario[string, int]) TDADiccionario.Diccionario[string, *usuario] {
	usuarios := TDADiccionario.CrearHash[string, *usuario](func(a, b string) bool { return a == b })

	iter := us.Iterador()

	for iter.HaySiguiente() {
		nombre, afinidad := iter.VerActual()
		usuarios.Guardar(nombre, crearUsuario(nombre, afinidad))

		iter.Siguiente()
	}

	return usuarios
}

func crearUsuario(nombre string, afinidad int) *usuario {
	usuario := new(usuario)
	usuario.nombre = nombre
	usuario.afinidad = afinidad
	usuario.feed = TDAHeap.CrearHeap[*postFeed](igualdadPostFeed)

	return usuario
}

func crearNuevoPostFeed(post TDAPost.Post, afinidad int) *postFeed {
	nuevoPostFeed := new(postFeed)
	nuevoPostFeed.post = post
	nuevoPostFeed.afinidad = afinidad
	return nuevoPostFeed
}

func igualdadPostFeed(dato1, dato2 *postFeed) int {
	afinidad1 := math.Abs(float64(dato1.afinidad - dato1.post.publicador.afinidad))
	afinidad2 := math.Abs(float64(dato2.afinidad - dato2.post.publicador.afinidad))
	res := int(afinidad2 - afinidad1)

	if res == 0 {
		res = dato2.post.ObtenerId() - dato1.post.ObtenerId()
	}
	return res
}
