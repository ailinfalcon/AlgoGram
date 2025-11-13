package usuario

import TDAPost "tp2/post"

type Usuario interface {
	// ObtenerAfinidad devuelve la afinidad del usuario.
	ObtenerAfinidad() int

	// ObtenerNombre devuelve el nombre del usuario .
	ObtenerNombre() string

	// AgregarPostFeed agrega un nuevo post al feed del usuario.
	AgregarPostFeed(TDAPost.Post, int)

	// ObtenerPostFeed devuelve el siguiente post del feed del usuario actual.
	ObtenerPostFeed() TDAPost.Post

	// TienePostFeed devuelve true si quedan post por ver en el feed, o false en caso contrario.
	TienePostFeed() bool
}
