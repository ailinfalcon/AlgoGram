package TDAPost

import (
	TDADiccionario "tdas/diccionario"
)

type Post interface {

	//
	ObtenerId() int

	//
	ObtenerPublicador() string

	//
	ObtenerContenido() string

	//
	ObtenerLikes() TDADiccionario.DiccionarioOrdenado[string, string]

	//
	ObtenerCantLikes() int

	//
	AgregarLike(string, string)
}
