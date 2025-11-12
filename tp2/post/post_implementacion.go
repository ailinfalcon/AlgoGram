package TDAPost

import (
	"strings"
	TDADiccionario "tdas/diccionario"
	TDAUsuario "tp2/usuario"
)

type post struct {
	id            int
	contenido     string
	publicador    TDAUsuario.Usuario
	cantidadLikes int
	likes         TDADiccionario.DiccionarioOrdenado[string, TDAUsuario.Usuario]
}

func CrearPost(usuario TDAUsuario.Usuario, contenido string, cant int) Post {
	nuevoPost := new(post)
	nuevoPost.id = cant
	nuevoPost.publicador = usuario
	nuevoPost.contenido = contenido
	nuevoPost.cantidadLikes = 0
	nuevoPost.likes = TDADiccionario.CrearABB[string, TDAUsuario.Usuario](strings.Compare)

	return nuevoPost
}

func (post *post) ObtenerId() int {
	return post.id
}

func (post *post) ObtenerPublicador() TDAUsuario.Usuario {
	return post.publicador
}

func (post *post) ObtenerContenido() string {
	return post.contenido
}

func (post *post) ObtenerLikes() TDADiccionario.DiccionarioOrdenado[string, TDAUsuario.Usuario] {
	return post.likes
}

func (post *post) ObtenerCantLikes() int {
	return post.cantidadLikes
}

func (post *post) GuardarLike(nombre string, usuarioLoggeado TDAUsuario.Usuario) {
	if !post.likes.Pertenece(nombre) {
		post.likes.Guardar(nombre, usuarioLoggeado)
		post.cantidadLikes++
	}
}
