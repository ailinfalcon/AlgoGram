package TDAPost

import (
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
)

type post struct {
	id int
	contenido string
	publicador string
	cantLikes int
	likes TDADiccionario.DiccionarioOrdenado[string, string]
}

func CrearPost(usuario string, contenido string, cant int) Post {
	nuevoPost := new(post)
	nuevoPost.id = cant
	nuevoPost.publicador = usuario
	nuevoPost.contenido = contenido
	nuevoPost.cantLikes = 0
	nuevoPost.likes = TDADiccionario.CrearABB[string, string](strings.Compare)

	return nuevoPost
}

func (post *post) ObtenerId() int {
	return post.id
}

func (post *post) ObtenerPublicador() string {
	return post.publicador
}

func (post *post) ObtenerContenido() string {
	return post.contenido
}

func (post *post) ObtenerLikes() []int {
	return post.likes
}

func (post *post) VerUltimoPost() (int, string, string, int) {
	if !algogram.HayLoggeado() || algogram.usuarioLoggeado.feed.Cantidad() == 0 {
		fmt.Printf("Usuario no loggeado o no hay mas posts para ver\n")
		return 0, "", "", 0
	}

	postFeed := algogram.usuarioLoggeado.feed.Desencolar()

	return postFeed.post.id, postFeed.post.publicador.nombre, postFeed.post.contenido, postFeed.post.cantLikes // re feo jajaj
}

func (algogram *Algogram) LikearPost(id int) bool {
	if !algogram.HayLoggeado() || id >= algogram.posts.Largo() || id < 0 {
		fmt.Printf("Error: Usuario no loggeado o Post inexistente\n")
		return false
	}

	iter := algogram.posts.Iterador()

	for i := 0; i < id; i++ {
		iter.Siguiente()
	}

	postActual := iter.VerActual()

	if !postActual.likes.Pertenece(algogram.usuarioLoggeado.nombre) {
		postActual.likes.Guardar(algogram.usuarioLoggeado.nombre, algogram.usuarioLoggeado)
		postActual.cantLikes++
	}

	return true
}

func (algogram *Algogram) MostrarLikes(id int) ([]string, int) {
	var likes []string
	if id >= algogram.posts.Largo() || id < 0 {
		fmt.Printf("Error: Post inexistente o sin likes\n")
		return likes, 0
	}

	iter := algogram.posts.Iterador()

	for i := 0; i < id; i++ {
		iter.Siguiente()
	}
	postActual := iter.VerActual()
	if postActual.cantLikes == 0 {
		fmt.Printf("Error: Post inexistente o sin likes\n")
		return likes, 0
	}

	iterLikes := postActual.likes.Iterador()
	for iterLikes.HaySiguiente() {
		nombreUsuario, _ := iterLikes.VerActual()
		likes = append(likes, nombreUsuario)
		iterLikes.Siguiente()
	}

	return likes, postActual.cantLikes
}

