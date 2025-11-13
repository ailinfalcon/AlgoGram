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

func CrearUsuario(nombreUsuario string, afinidadUsuario int) *usuario {
	return &usuario{
		nombre:   nombreUsuario,
		feed:     TDAHeap.CrearHeap[*postFeed](igualdadPostFeed),
		afinidad: afinidadUsuario,
	}
}

func (usuario *usuario) ObtenerAfinidad() int {
	return usuario.afinidad
}

func (usuario *usuario) ObtenerNombre() string {
	return usuario.nombre
}

func (usuario *usuario) AgregarPostFeed(post TDAPost.Post, afinidadPublicador int) {
	afinidad := math.Abs(float64(usuario.afinidad - afinidadPublicador))
	postFeed := crearPostFeed(post, int(afinidad))
	usuario.feed.Encolar(postFeed)
}

func (usuario *usuario) ObtenerPostFeed() TDAPost.Post {
	return usuario.feed.Desencolar().post
}

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
