package usuario

import (
	"math"
	TDAHeap "tdas/cola_prioridad"
	TDAPost "tp2/post"
)

type usuario struct {
	nombre   string
	feed     TDAHeap.ColaPrioridad[*postFeed]
	afinidad int
}

type postFeed struct {
	post     TDAPost.Post
	afinidad int
}

// Crea un usuario con el nombre y afinidad recibido por parametro,
// y devuelve un puntero al usuario creado
func CrearUsuario(nombreUsuario string, afinidadUsuario int) *usuario {
	return &usuario{
		nombre:   nombreUsuario,
		feed:     TDAHeap.CrearHeap[*postFeed](igualdadPostFeed),
		afinidad: afinidadUsuario,
	}
}

// Devuelve la afinidad del usuario actual
func (usuario *usuario) ObtenerAfinidad() int {
	return usuario.afinidad
}

// Devuelve el nombre del usuario actual
func (usuario *usuario) ObtenerNombre() string {
	return usuario.nombre
}

// Agrega un nuevo post al feed del usuario actual.
func (usuario *usuario) AgregarPostFeed(post TDAPost.Post, afinidadPublicador int) {
	afinidad := math.Abs(float64(usuario.afinidad - afinidadPublicador))
	postFeed := crearPostFeed(post, int(afinidad))
	usuario.feed.Encolar(postFeed)
}

// Devuelve el siguiente post del feed del usuario actual
func (usuario *usuario) ObtenerPostFeed() TDAPost.Post {
	return usuario.feed.Desencolar().post
}

// Devuelve true si quedan post por ver en el feed, o false en caso contrario
func (usuario *usuario) TienePostFeed() bool {
	return usuario.feed.Cantidad() > 0
}

func crearPostFeed(post TDAPost.Post, afinidad int) *postFeed {
	nuevoPostFeed := new(postFeed)
	nuevoPostFeed.post = post
	nuevoPostFeed.afinidad = afinidad
	return nuevoPostFeed
}

func igualdadPostFeed(dato1, dato2 *postFeed) int {
	res := dato2.afinidad - dato1.afinidad

	if res == 0 {
		res = dato2.post.ObtenerId() - dato1.post.ObtenerId()
	}
	return res
}
