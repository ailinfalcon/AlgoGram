package TDAPost

import (
	TDADiccionario "tdas/diccionario"
)

type Post interface {

	// Devuelve el id del post actual.
	ObtenerId() int

	// Devuelve el nombre del creador del post actual
	ObtenerPublicador() string

	// Devuelve el contenido del post actual
	ObtenerContenido() string

	// Devuelve un Diccionario Ordenado con los usuarios que le dieron like al post actual.
	ObtenerLikes() TDADiccionario.DiccionarioOrdenado[string, string]

	// Devuelve la cantidad de likes que tiene el post actual.
	ObtenerCantLikes() int

	//
	AgregarLike(string, string)
}
