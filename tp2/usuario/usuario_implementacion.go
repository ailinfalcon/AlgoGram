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
		nombre: nombreUsuario,
		feed: TDAHeap.CrearHeap[*postFeed](igualdadPostFeed),
		afinidad: afinidadUsuario,
	}
}

func (usuario *usuario) ObtenerAfinidad() int {
	return usuario.afinidad
}

func (usuario *usuario) ObtenerNombre() string {
	return usuario.nombre
}

func (usuario *usuario) AgregarPostFeed(post TDAPost.Post, afinidad int) {
	postFeed := crearPostFeed(post, afinidad)
	usuario.feed.Encolar(postFeed)
}

func (usuario *usuario) ObtenerPostFeed() TDAPost.Post {
	return usuario.feed.Desencolar().post
}

func crearPostFeed(post TDAPost.Post, afinidad int) *postFeed {
	nuevoPostFeed := new(postFeed)
	nuevoPostFeed.post = post
	nuevoPostFeed.afinidad = afinidad
	return nuevoPostFeed
}

func igualdadPostFeed(dato1, dato2 *postFeed) int {
	afinidad1 := math.Abs(float64(dato1.afinidad - dato1.post.ObtenerPublicador().ObtenerAfinidad()))
	afinidad2 := math.Abs(float64(dato2.afinidad - dato2.post.ObtenerPublicador().ObtenerAfinidad()))
	res := int(afinidad2 - afinidad1)

	if res == 0 {
		res = dato2.post.ObtenerId() - dato1.post.ObtenerId()
	}
	return res
}