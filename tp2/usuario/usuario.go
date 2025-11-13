package usuario

import TDAPost "tp2/post"

type Usuario interface {

	//
	ObtenerAfinidad() int

	//
	ObtenerNombre() string

	//
	AgregarPostFeed(TDAPost.Post, int)

	//
	ObtenerPostFeed() TDAPost.Post

	//
	TienePostFeed() bool
}
