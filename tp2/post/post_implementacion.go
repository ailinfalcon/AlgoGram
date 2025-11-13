package TDAPost

import (
	"strings"
	TDADiccionario "tdas/diccionario"
)

type post struct {
	id            int
	contenido     string
	publicador    string
	cantidadLikes int
	likes         TDADiccionario.DiccionarioOrdenado[string, string]
}

// Crear un nuevo post con el nombre de usuario, el contenido y el id recibido por parametro
func CrearPost(nombreUsuario string, contenido string, cant int) Post {
	nuevoPost := new(post)
	nuevoPost.id = cant
	nuevoPost.publicador = nombreUsuario
	nuevoPost.contenido = contenido
	nuevoPost.cantidadLikes = 0
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

func (post *post) ObtenerLikes() TDADiccionario.DiccionarioOrdenado[string, string] {
	return post.likes
}

func (post *post) ObtenerCantLikes() int {
	return post.cantidadLikes
}

func (post *post) AgregarLike(nombre string, usuarioLoggeado string) {
	if !post.likes.Pertenece(nombre) {
		post.likes.Guardar(nombre, usuarioLoggeado)
		post.cantidadLikes++
	}
}
