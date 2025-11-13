package usuario

import TDAPost "tp2/post"

type Usuario interface {
	/* Devuelve la afinidad del usuario */
	ObtenerAfinidad() int

	/* Devuelve el nombre del usuario */
	ObtenerNombre() string

	/* Agrega un nuevo post al feed del usuario */
	AgregarPostFeed(TDAPost.Post, int)

	/* Devuelve el siguiente post del feed del usuario actual */
	ObtenerPostFeed() TDAPost.Post

	/* Devuelve true si quedan post por ver en el feed, o false en caso contrario */
	TienePostFeed() bool
}
