package TDAPost

import (
	TDADiccionario "tdas/diccionario"
)

type Post interface {

	// ObtenerId devuelve el id del post actual.
	ObtenerId() int

	// ObtenerPublicador devuelve el nombre del creador del post actual
	ObtenerPublicador() string

	// ObtenerContenido devuelve el contenido del post actual
	ObtenerContenido() string

	// ObtenerLikes devuelve un Diccionario Ordenado con los usuarios que le dieron like al post actual.
	ObtenerLikes() TDADiccionario.DiccionarioOrdenado[string, string]

	// ObtenerCantLikes devuelve la cantidad de likes que tiene el post actual.
	ObtenerCantLikes() int

	// AgregarLike agrega el like del usuario loggeado a los likes del post actual
	AgregarLike(string)
}
